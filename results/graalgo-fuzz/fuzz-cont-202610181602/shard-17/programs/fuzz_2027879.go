package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint64; F1 string; F2 int }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int8; F1 string; F2 int }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 int64; F1 uint8 }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 uint32; F1 uint16; F2 int8 }

func main() {
    var v1 int16 = 42
    v2 := make(map[uint32]int)
    v3 := make(map[uint32]uint32)
    for fi4 := 0; fi4 < 14; fi4++ {
        for fi5 := 0; fi5 < 1; fi5++ {
            if false {
                switch (fi4 | -162) {
                case 42:
                    switch "hello" {
                    case "foo bar":
                        var v6 int8 = -1
                        _ = v6
                    case "":
                        var v7 uint16 = 0
                        _ = v7
                    default:
                        v8 := []int{2147483647}
                        _ = v8
                    }
                case -1:
                    v9 := GenS1{F0: 0}
                    _ = v9
                case 128:
                    if (-426 != -1) {
                        if (1 > 965) {
                            if (!(-2147483648 > -2147483648)) {
                                v1 -= v1
                                v1 = (^(v1))
                            }
                            v1 = v1
                            switch (fi5 % (((fi5 << ((2) & 63))) | 1)) {
                            case 4294967295:
                                v10 := []int{(fi5 >> ((4294967295) & 63)), -913, int(v1)}
                                _ = v10
                            case 42:
                                delete(v3, 2147483647)
                            default:
                                switch "" {
                                case "foo bar":
                                    var v11 string = "hello"
                                    _ = v11
                                case "a":
                                    v1 = 1
                                default:
                                    fmt.Println(v1, fi4)
                                }
                            }
                        } else {
                            if true {
                                var v12 uint8 = 98
                                _ = v12
                            } else {
                                if ((2147483647 >= fi5) && false) {
                                    var v13 int16 = 1
                                    _ = v13
                                }
                            }
                            if ((v1 >> ((42) & 63)) > v1) {
                                var v14 int32 = -24
                                _ = v14
                            }
                            v1 = 834
                            if (329 > 255) {
                                v15 := []int{32768, 1, (fi5 + (^(fi4)))}
                                _ = v15
                            } else {
                                switch (("\n\t" + "") + ((("hello" + ("a" + "a")) + ("hello" + ("hello" + "hello"))) + "foo bar")) {
                                case "":
                                    var v16 int32 = 42
                                    _ = v16
                                case "hello":
                                    switch ("a" + "") {
                                    case "hello":
                                        fmt.Println(fi5, v1)
                                    case "a":
                                        fmt.Println(fi4)
                                    case "\n\t":
                                        v1 = 1000
                                    default:
                                        var v17 uint16 = 65535
                                        _ = v17
                                    }
                                case "foo bar":
                                    v18 := []int{-996, (fi5 / ((fi4) | 1)), fi5, -559}
                                    _ = v18
                                default:
                                    fmt.Println(v1, fi5)
                                }
                            }
                            if (false || ((65535 >= 2) || (!(0 != 755)))) {
                                if true {
                                    var v19 float64 = float64(float32(-2.25))
                                    switch 0 {
                                    case -899:
                                        var v20 string = "hello"
                                        _ = v20
                                    default:
                                        var v21 int64 = 660
                                        _ = v21
                                    }
                                    v1 += (^(v1))
                                    var v22 float32 = ((((float32(v19) - 1e10) / 3.14159) * -2.25) + -2.25)
                                    _ = v19
                                    _ = v22
                                } else {
                                    v23 := []int{2147483647, (fi4 - -2147483648), fi5}
                                    _ = v23
                                }
                                var v24 int64 = 1
                                var v25 uint32 = 4294967295
                                var v26 uint32 = (v25 >> ((2) & 63))
                                _ = v24
                                _ = v25
                                _ = v26
                            } else {
                                _ = len(v3)
                                var v27 bool = true
                                _ = v27
                            }
                            var v28 float64 = 1.5
                            _ = v28
                        }
                        v29 := GenS1{F0: 1}
                        _ = v29
                    }
                    v30 := v3[2]
                    var v31 bool = false
                    _ = v30
                    _ = v31
                }
                var v32 float64 = (-2.25 + -2.25)
                _ = v32
            }
            fmt.Println(fi4)
            _ = fi5
        }
        switch fi4 {
        case -1:
            v1 = v1
        case -566:
            v33 := GenS2{F0: 18446744073709551615, F1: "a", F2: (-(fi4))}
            _ = v33
        default:
            switch "foo bar" {
            case "":
                v34 := GenS3{F0: 2}
                _ = v34
            case "\n\t":
                v35 := GenS4{F0: 100, F1: ("\n\t" + ""), F2: fi4}
                _ = v35
            default:
                if false {
                    v36 := GenS5{F0: -1}
                    for fi37 := 0; fi37 < 13; fi37++ {
                        if (-128 < 1) {
                            v38 := GenS3{F0: 415}
                            var v39 int32 = -1
                            var v40 bool = (3.14159 != 3.14159)
                            _ = v38
                            _ = v39
                            _ = v40
                        } else {
                            if false {
                                switch ("\n\t" + ((("a" + "") + (("a" + ("hello" + "\n\t")) + "foo bar")) + "a")) {
                                case "hello":
                                    v41 := v2[65535]
                                    _ = v41
                                case "foo bar":
                                    var v42 uint64 = 837
                                    _ = v42
                                }
                                fmt.Println(fi37, fi37, v1)
                            } else {
                                var v43 uint32 = 42
                                _ = v43
                            }
                        }
                        fmt.Println(v1)
                        fmt.Println(v1)
                        fmt.Println(fi37)
                        _ = fi37
                    }
                    var v44 uint8 = 0
                    _ = v36
                    _ = v44
                }
            }
        }
        var v45 bool = (!true)
        _ = fi4
        _ = v45
    }
    v46 := GenS6{F0: 95, F1: 120}
    var v47 int8 = 42
    delete(v3, 457)
    var v48 string = ("\n\t" + "foo bar")
    v47 += v47
    v47 = v47
    var v49 uint16 = 42
    if (((-(v47)) >= (v47 % ((v47) | 1))) || false) {
        if true {
            var v50 bool = true
            for fi51 := 0; fi51 < 15; fi51++ {
                for fi52 := 0; fi52 < 6; fi52++ {
                    switch "\n\t" {
                    case "a":
                        v53 := GenS1{F0: 18446744073709551615}
                        var v54 uint8 = 1
                        fmt.Println(v47, fi51)
                        _ = v53
                        _ = v54
                    }
                    delete(v3, 42)
                    _ = fi52
                }
                switch (("foo bar" + v48) + ("\n\t" + ((("\n\t" + v48) + v48) + ("a" + ("\n\t" + (((v48 + ((v48 + v48) + "\n\t")) + (v48 + v48)) + v48)))))) {
                case "":
                    v55 := GenS7{F0: false}
                    _ = v55
                case "hello":
                    var v56 int64 = int64(v49)
                    _ = v56
                default:
                    v50 = false
                }
                _ = fi51
            }
            for fi57 := 0; fi57 < 20; fi57++ {
                var v58 float32 = -2.25
                v48 = v48
                _ = fi57
                _ = v58
            }
            _ = v50
        } else {
            for fi59 := 0; fi59 < 7; fi59++ {
                var v60 bool = (0 != 132)
                v3[42] = 930
                _ = fi59
                _ = v60
            }
            v1 = (v1 * (-(v1)))
            if true {
                fmt.Println(v1)
            }
        }
        for fi61 := 0; fi61 < 13; fi61++ {
            var v62 int8 = v47
            _ = fi61
            _ = v62
        }
    } else {
        var v63 uint8 = 2
        var v64 int8 = (v47 * 33)
        var v65 string = (v48 + (v48 + ("hello" + "foo bar")))
        _ = v63
        _ = v64
        _ = v65
    }
    v47 += v47
    if true {
        switch 42 {
        case -2147483648:
            v66 := GenS8{F0: (v49 >> ((569) & 63))}
            _ = v66
        case -9223372036854775808:
            v67 := GenS5{F0: -1}
            _ = v67
        default:
            v68 := v2[1]
            _ = v68
        }
        v69 := []int{9223372036854775807, 127, 32767, -617}
        switch v48 {
        case "hello":
            v70 := GenS3{F0: v49}
            _ = v70
        }
        _ = v69
    }
    if (95 <= 1) {
        if (!false) {
            switch v48 {
            case "a":
                v48 = v48
            case "hello":
                v1 = (v1 ^ v1)
            case "foo bar":
                for fi71 := 0; fi71 < 7; fi71++ {
                    v47 += v47
                    for fi72 := 0; fi72 < 19; fi72++ {
                        v73 := GenS5{F0: 0}
                        _ = fi72
                        _ = v73
                    }
                    if true {
                        v74 := GenS2{F0: 18446744073709551615, F1: "hello", F2: fi71}
                        fmt.Println(fi71, v47)
                        if (!true) {
                            if false {
                                var v75 uint64 = 0
                                var v76 uint64 = (^(v75))
                                var v77 uint8 = 255
                                _ = v75
                                _ = v76
                                _ = v77
                            } else {
                                switch "\n\t" {
                                case "foo bar":
                                    switch fi71 {
                                    case 633:
                                        if false {
                                            fmt.Println(v48, v1, v49)
                                            if true {
                                                v78 := []int{(fi71 + (fi71 << ((0) & 63))), (fi71 + 32768), -9223372036854775808, fi71}
                                                _ = v78
                                            } else {
                                                v79 := GenS6{F0: -142, F1: 255}
                                                var v80 float64 = ((((0.0 - (float64(1e10) - ((((-2.25 / 3.14159) * -2.25) * -2.25) - 1e10))) * (((0.1 + (((0.0 - 3.14159) - 1.5) + 0.1)) * (float64(0.1) + 1.5)) * (3.14159 * -2.25))) - 0.1) / ((1e10 / (1e10 + (0.0 / (0.0 + float64(float32(((float64(1.5) + 1e10) / (-2.25 + 3.14159)))))))) * 0.1))
                                                _ = v79
                                                _ = v80
                                            }
                                        }
                                    default:
                                        switch (fi71 + (fi71 ^ (fi71 << ((4294967295) & 63)))) {
                                        case 1:
                                            fmt.Println(v1, v1)
                                        case -1:
                                            v49 = 2
                                        }
                                    }
                                case "a":
                                    v81 := GenS2{F0: 18446744073709551615, F1: v48, F2: (-(fi71))}
                                    _ = v81
                                case "hello":
                                    v49 = 1
                                }
                            }
                            v82 := GenS2{F0: 4294967295, F1: (v48 + v48), F2: fi71}
                            var v83 int16 = (v1 % (((v1 * (v1 & 0))) | 1))
                            _ = v82
                            _ = v83
                        } else {
                            var v84 int8 = v47
                            for fi85 := 0; fi85 < 12; fi85++ {
                                if false {
                                    switch (v48 + (v48 + v48)) {
                                    case "a":
                                        v86 := GenS8{F0: (v49 / ((v49) | 1))}
                                        _ = v86
                                    case "\n\t":
                                        v49 += v49
                                    case "hello":
                                        switch "foo bar" {
                                        case "foo bar":
                                            var v87 int8 = (v47 / ((1) | 1))
                                            _ = v87
                                        case "\n\t":
                                            fmt.Println(fi71, v47)
                                        }
                                    }
                                    if true {
                                        v2[42] = fi71
                                        var v88 uint = 65535
                                        delete(v2, 361)
                                        _ = v88
                                    }
                                    v89 := []int{}
                                    _ = v89
                                }
                                _ = fi85
                            }
                            var v90 uint16 = 944
                            _ = v84
                            _ = v90
                        }
                        _ = v74
                    }
                    _ = fi71
                }
            default:
                v91 := GenS9{F0: 952, F1: v49, F2: -46}
                _ = v91
            }
        }
    } else {
        var v92 bool = (2147483647 >= -1)
        fmt.Println(v1, v47)
        switch 32768 {
        case 200:
            var v93 int64 = 1
            _ = v93
        case 32767:
            switch "foo bar" {
            case "a":
                var v94 float64 = 1e10
                _ = v94
            case "\n\t":
                if (!false) {
                    if v92 {
                        v95 := []int{4294967295, 164, 4294967295, 672}
                        _ = v95
                    } else {
                        if false {
                            if ((v49 % ((v49) | 1)) == (^(v49))) {
                                var v96 int = -9223372036854775808
                                v48 = ("" + "")
                                _ = v96
                            } else {
                                var v97 float64 = 1e10
                                _ = v97
                            }
                        }
                    }
                    var v98 uint64 = 0
                    _ = v98
                } else {
                    switch 853 {
                    case 127:
                        if (366 <= 2147483648) {
                            switch -9223372036854775808 {
                            case 0:
                                var v99 uint16 = v49
                                _ = v99
                            case -635:
                                var v100 int8 = v47
                                _ = v100
                            case 226:
                                v49 += v49
                            }
                            v101 := GenS1{F0: 176}
                            v1 = v1
                            _ = v101
                        }
                    case 341:
                        var v102 uint32 = 42
                        _ = v102
                    case 4294967295:
                        for fi103 := 0; fi103 < 18; fi103++ {
                            var v104 uint32 = 42
                            switch "hello" {
                            case "a":
                                v105 := GenS9{F0: (v104 * (^(v104))), F1: v49, F2: (v47 / (((^(v47))) | 1))}
                                _ = v105
                            case "":
                                var v106 int = (fi103 & fi103)
                                _ = v106
                            }
                            v107 := GenS3{F0: uint16(v104)}
                            _ = fi103
                            _ = v104
                            _ = v107
                        }
                    }
                }
            case "":
                if (9223372036854775807 > 9223372036854775807) {
                    v108 := GenS2{F0: 3, F1: ((v48 + (v48 + "hello")) + ("" + "")), F2: 689}
                    v1 = -768
                    _ = v108
                }
            }
        case 0:
            v49 += (^(v49))
        }
        _ = v92
    }
    v109 := GenS5{F0: -1}
    v48 = v48
    fmt.Println(v48, v48)
    var v110 bool = (9223372036854775807 != 9223372036854775807)
    var v111 int32 = -2147483648
    fmt.Println("v1:", v1)
    fmt.Println("v110:", v110)
    fmt.Println("v111:", v111)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    _ = v1
    _ = v109
    _ = v110
    _ = v111
    _ = v2
    _ = v3
    _ = v46
    _ = v47
    _ = v48
    _ = v49
}
