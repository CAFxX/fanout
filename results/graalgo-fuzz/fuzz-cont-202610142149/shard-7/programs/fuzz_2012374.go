package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32; F2 string }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int64; F1 int32; F2 uint8 }
type GenS4 struct { F0 int8; F1 uint }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int16; F1 string }
type GenS7 struct { F0 int16; F1 int32 }
type GenS8 struct { F0 uint32 }
type GenS9 struct { F0 string; F1 uint32 }
type GenS10 struct { F0 uint16; F1 int64; F2 uint8 }

func main() {
    v1 := make(map[int]uint32)
    var v2 uint = 70
    v3 := make(map[int]int)
    var v4 int8 = 0
    var v5 int64 = 42
    v4 = (^(v4))
    var v6 float64 = (0.0 / 0.1)
    for fi7 := 0; fi7 < 7; fi7++ {
        fmt.Println(fi7, v6, v5)
        for fi8 := 0; fi8 < 8; fi8++ {
            v9 := []int{580, fi8, (^(fi8)), fi8}
            v5 -= v5
            v4 -= v4
            _ = fi8
            _ = v9
        }
        _ = fi7
    }
    if false {
        v10 := GenS1{F0: 65535, F1: 942, F2: "hello"}
        _ = len(v1)
        for fi11 := 0; fi11 < 17; fi11++ {
            delete(v1, 127)
            _ = fi11
        }
        _ = v10
    } else {
        if (-1 >= 133) {
            switch (("a" + "a") + "\n\t") {
            case "\n\t":
                if false {
                    var v12 uint8 = 1
                    for fi13 := 0; fi13 < 20; fi13++ {
                        switch fi13 {
                        case 2147483647:
                            v4 = v4
                        case 32768:
                            for fi14 := 0; fi14 < 15; fi14++ {
                                switch fi13 {
                                case -1:
                                    v4 -= 1
                                }
                                fmt.Println(v12, fi13, v5)
                                if (-1 != (v5 - v5)) {
                                    var v15 uint16 = 638
                                    var v16 int16 = int16(v12)
                                    var v17 uint8 = (v12 >> ((1) & 63))
                                    v18 := GenS1{F0: (v15 - v15), F1: 588, F2: ""}
                                    switch "\n\t" {
                                    case "a":
                                        var v19 string = ((((("" + "\n\t") + "foo bar") + "foo bar") + ("foo bar" + "\n\t")) + "")
                                        _ = v19
                                    case "\n\t":
                                        var v20 bool = false
                                        _ = v20
                                    default:
                                        var v21 uint = 670
                                        _ = v21
                                    }
                                    _ = v15
                                    _ = v16
                                    _ = v17
                                    _ = v18
                                }
                                _ = fi14
                            }
                        }
                        _ = fi13
                    }
                    fmt.Println(v12)
                    _ = v12
                }
            case "":
                switch (("foo bar" + "a") + "\n\t") {
                case "\n\t":
                    for fi22 := 0; fi22 < 16; fi22++ {
                        var v23 int = -1
                        if false {
                            fmt.Println(v6, v23, v23)
                            for fi24 := 0; fi24 < 8; fi24++ {
                                fmt.Println(v6, fi24, v2)
                                _ = fi24
                            }
                        } else {
                            switch "\n\t" {
                            case "hello":
                                switch (^(fi22)) {
                                case 121:
                                    if (0 >= 2) {
                                        fmt.Println(v23)
                                    }
                                case 1:
                                    v2 = v2
                                case 4294967295:
                                    var v25 int16 = 1000
                                    _ = v25
                                default:
                                    v26 := GenS1{F0: 42, F1: uint32(v2), F2: ("hello" + "\n\t")}
                                    _ = v26
                                }
                            default:
                                for fi27 := 0; fi27 < 19; fi27++ {
                                    fmt.Println(v6, v23, v2)
                                    v28 := GenS1{F0: 286, F1: 754, F2: ""}
                                    _ = fi27
                                    _ = v28
                                }
                            }
                            for fi29 := 0; fi29 < 17; fi29++ {
                                v30 := []int{v23, fi29}
                                v31 := GenS2{F0: int16(v5)}
                                fmt.Println(fi22, v6)
                                _ = fi29
                                _ = v30
                                _ = v31
                            }
                        }
                        if (42 > 255) {
                            var v32 uint64 = 42
                            for fi33 := 0; fi33 < 12; fi33++ {
                                v34 := GenS1{F0: 32767, F1: 0, F2: ""}
                                var v35 uint16 = uint16(v2)
                                if (uint32(v32) <= 2147483648) {
                                    v2 = v2
                                } else {
                                    v6 -= (((v6 / (v6 / v6)) - (float64((1.5 + 3.14159)) + float64(((((float32(-2.25) / 1.5) + 3.14159) * 3.14159) / 1e10)))) / 3.14159)
                                    var v36 uint32 = 0
                                    var v37 uint = 200
                                    _ = v36
                                    _ = v37
                                }
                                _ = fi33
                                _ = v34
                                _ = v35
                            }
                            _ = v32
                        } else {
                            switch "hello" {
                            case "\n\t":
                                if (1 == 586) {
                                    for fi38 := 0; fi38 < 14; fi38++ {
                                        v39 := GenS1{F0: 42, F1: 212, F2: "\n\t"}
                                        if (737 < 2) {
                                            v40 := GenS3{F0: (v5 - (v5 - (v5 / ((v5) | 1)))), F1: 690, F2: 128}
                                            _ = v40
                                        }
                                        if (float64(((3.14159 / ((((-2.25 + 1.5) / 1.5) - 0.0) * 1e10)) / 0.1)) == 1e10) {
                                            delete(v1, (fi38 * 1))
                                            if (4294967295 <= 42) {
                                                if (!false) {
                                                    v41 := []int{fi22, 32767, (v23 / (((v23 / ((fi22) | 1))) | 1))}
                                                    v4 = -128
                                                    var v42 uint = 245
                                                    _ = v41
                                                    _ = v42
                                                }
                                                switch fi22 {
                                                case 127:
                                                    var v43 int8 = int8(v5)
                                                    _ = v43
                                                case 4294967295:
                                                    if true {
                                                        if true {
                                                            fmt.Println(v2)
                                                            _ = len(v3)
                                                        } else {
                                                            var v44 uint32 = 65535
                                                            _ = v44
                                                        }
                                                        var v45 uint64 = 2
                                                        v6 -= (1e10 + (3.14159 / float64(-2.25)))
                                                        _ = v45
                                                    }
                                                }
                                                _ = len(v3)
                                            } else {
                                                v46 := GenS3{F0: (v5 + v5), F1: -681, F2: 154}
                                                var v47 float64 = ((float64(-2.25) * v6) - float64(0.1))
                                                var v48 int64 = v5
                                                v49 := GenS1{F0: 470, F1: 65535, F2: ""}
                                                v50 := []int{(v23 >> ((v2) & 63)), v23, 2147483647, 1}
                                                _ = v46
                                                _ = v47
                                                _ = v48
                                                _ = v49
                                                _ = v50
                                            }
                                            switch "hello" {
                                            case "\n\t":
                                                if true {
                                                    v51 := GenS4{F0: v4, F1: v2}
                                                    if true {
                                                        if ((!false) || true) {
                                                            v4 -= (v4 << ((598) & 63))
                                                            var v52 string = "hello"
                                                            var v53 uint32 = 65535
                                                            v23 -= fi38
                                                            v54 := GenS4{F0: v4, F1: 18446744073709551615}
                                                            _ = v52
                                                            _ = v53
                                                            _ = v54
                                                        }
                                                        var v55 string = ("" + "hello")
                                                        v23 = (fi22 / (((v23 * fi22)) | 1))
                                                        v56 := GenS5{F0: -750}
                                                        _ = v55
                                                        _ = v56
                                                    }
                                                    _ = v51
                                                }
                                            case "":
                                                v57 := GenS6{F0: -131, F1: ("" + "\n\t")}
                                                _ = v57
                                            default:
                                                var v58 float32 = (((3.14159 * 0.1) / (0.1 + 1.5)) - (((1e10 - -2.25) / ((0.1 * 1.5) * (1.5 + 0.1))) * 1e10))
                                                _ = v58
                                            }
                                            if (255 == 127) {
                                                delete(v1, fi38)
                                                fmt.Println(v4, v2)
                                            } else {
                                                v59 := GenS1{F0: 32767, F1: uint32(v2), F2: "foo bar"}
                                                fmt.Println(v6, v23)
                                                delete(v1, v23)
                                                _ = v59
                                            }
                                        } else {
                                            if ((65535 != 42) || true) {
                                                fmt.Println(v4, fi38)
                                            } else {
                                                v2 = v2
                                            }
                                        }
                                        _ = fi38
                                        _ = v39
                                    }
                                    for fi60 := 0; fi60 < 9; fi60++ {
                                        v61 := GenS7{F0: -124, F1: 1}
                                        var v62 int16 = 344
                                        _ = fi60
                                        _ = v61
                                        _ = v62
                                    }
                                } else {
                                    v2 = v2
                                    v4 -= (v4 + (-(v4)))
                                }
                            case "a":
                                switch "" {
                                case "hello":
                                    switch "a" {
                                    case "hello":
                                        if ("a" <= "hello") {
                                            for fi63 := 0; fi63 < 15; fi63++ {
                                                var v64 float32 = 1.5
                                                switch "hello" {
                                                case "foo bar":
                                                    v1[(^(v23))] = 1
                                                case "a":
                                                    v65 := GenS2{F0: 733}
                                                    _ = v65
                                                default:
                                                    v3[(^(fi22))] = (fi22 / ((391) | 1))
                                                }
                                                if true {
                                                    fmt.Println(v6, v5)
                                                    var v66 uint = (v2 & v2)
                                                    switch 32768 {
                                                    case -1:
                                                        fmt.Println(v66, v5, v5)
                                                    case 0:
                                                        if false {
                                                            switch "\n\t" {
                                                            case "foo bar":
                                                                v67 := GenS6{F0: 32767, F1: "\n\t"}
                                                                _ = v67
                                                            }
                                                            v66 = 349
                                                        }
                                                    case 127:
                                                        switch (v23 % (((fi22 >> ((v66) & 63))) | 1)) {
                                                        case -2147483648:
                                                            if ((v5 ^ (v5 / ((9223372036854775807) | 1))) >= v5) {
                                                                v68 := []int{fi63, fi22}
                                                                v69 := GenS1{F0: 65535, F1: 65535, F2: ""}
                                                                _ = v68
                                                                _ = v69
                                                            } else {
                                                                if true {
                                                                    v70 := GenS8{F0: 4294967295}
                                                                    v5 = (v5 + v5)
                                                                    v23 = -9223372036854775808
                                                                    _ = v70
                                                                }
                                                                delete(v3, fi63)
                                                            }
                                                        case 2147483647:
                                                            if true {
                                                                v6 += v6
                                                                v3[fi63] = v23
                                                            } else {
                                                                var v71 int8 = v4
                                                                _ = v71
                                                            }
                                                        }
                                                    }
                                                    _ = v66
                                                }
                                                _ = fi63
                                                _ = v64
                                            }
                                            var v72 bool = ((1000 >= 42) && true)
                                            _ = v72
                                        }
                                    case "\n\t":
                                        var v73 int16 = -32768
                                        _ = v73
                                    default:
                                        delete(v1, (fi22 >> ((v2) & 63)))
                                    }
                                default:
                                    if true {
                                        for fi74 := 0; fi74 < 11; fi74++ {
                                            var v75 uint16 = 798
                                            var v76 string = ("foo bar" + "\n\t")
                                            _ = fi74
                                            _ = v75
                                            _ = v76
                                        }
                                        var v77 int16 = -32768
                                        switch v23 {
                                        case 2147483647:
                                            v78 := []int{(v23 * int(v5)), 32768}
                                            _ = v78
                                        }
                                        _ = v77
                                    }
                                }
                            default:
                                if (false && (true || false)) {
                                    v79 := GenS9{F0: ("\n\t" + (("hello" + (("a" + "\n\t") + ("" + "a"))) + "")), F1: 65535}
                                    v3[(fi22 / (((fi22 % (((fi22 % ((-1) | 1))) | 1))) | 1))] = v23
                                    if false {
                                        var v80 float64 = ((v6 - 0.1) / 1.5)
                                        switch ("a" + ("foo bar" + ("hello" + ("a" + ((("hello" + "a") + ("foo bar" + ("" + ""))) + "foo bar"))))) {
                                        case "\n\t":
                                            v81 := GenS2{F0: 1000}
                                            _ = v81
                                        default:
                                            if (("" + "foo bar") > "a") {
                                                switch (v23 << ((18446744073709551615) & 63)) {
                                                case -1:
                                                    fmt.Println(v23, v23)
                                                case -2147483648:
                                                    v5 -= v5
                                                case 42:
                                                    var v82 uint64 = 356
                                                    _ = v82
                                                }
                                                v83 := GenS7{F0: 1000, F1: 451}
                                                _ = v83
                                            }
                                        }
                                        v6 = v6
                                        _ = v80
                                    } else {
                                        var v84 uint16 = 2
                                        var v85 uint16 = (v84 / ((v84) | 1))
                                        _ = v84
                                        _ = v85
                                    }
                                    _ = v79
                                }
                            }
                            var v86 float32 = (0.1 + (1e10 - 1e10))
                            var v87 float64 = -2.25
                            _ = v86
                            _ = v87
                        }
                        _ = fi22
                        _ = v23
                    }
                case "hello":
                    for fi88 := 0; fi88 < 13; fi88++ {
                        delete(v3, fi88)
                        for fi89 := 0; fi89 < 13; fi89++ {
                            v90 := []int{(fi89 << ((405) & 63)), (fi89 / ((int(v4)) | 1)), (fi88 - -553)}
                            _ = fi89
                            _ = v90
                        }
                        v4 += v4
                        _ = fi88
                    }
                default:
                    v5 += v5
                }
            case "foo bar":
                var v91 int = -120
                _ = v91
            default:
                v92 := GenS4{F0: (v4 + v4), F1: v2}
                _ = v92
            }
            for fi93 := 0; fi93 < 20; fi93++ {
                v94 := v3[fi93]
                if false {
                    fmt.Println(v6, v2, v4)
                    if (811 > 2147483647) {
                        var v95 bool = false
                        _ = v95
                    }
                    switch "a" {
                    case "\n\t":
                        var v96 float32 = (((-2.25 + 0.0) - 0.0) + 1e10)
                        _ = v96
                    }
                }
                for fi97 := 0; fi97 < 10; fi97++ {
                    switch "hello" {
                    case "foo bar":
                        switch 418 {
                        case 2147483647:
                            v98 := GenS10{F0: 614, F1: (-(v5)), F2: 42}
                            _ = v98
                        case 32767:
                            v94 = v94
                        case -761:
                            v6 += ((1e10 + v6) - (v6 - 0.1))
                        }
                    case "hello":
                        v99 := v3[fi97]
                        _ = v99
                    }
                    v100 := GenS1{F0: 255, F1: 4294967295, F2: ("foo bar" + "")}
                    _ = fi97
                    _ = v100
                }
                _ = fi93
                _ = v94
            }
        }
        v101 := []int{}
        v2 = (v2 & v2)
        _ = v101
    }
    v102 := []int{4294967295, -365, 42}
    if (!false) {
        for fi103 := 0; fi103 < 2; fi103++ {
            v6 = v6
            v5 -= 42
            _ = fi103
        }
        v5 += 1
    }
    var v104 bool = (-1 >= -2147483648)
    v105 := GenS6{F0: -1, F1: "\n\t"}
    v4 += v4
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v102:", len(v102), v102)
    fmt.Println("v104:", v104)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v102
    _ = v104
    _ = v105
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
