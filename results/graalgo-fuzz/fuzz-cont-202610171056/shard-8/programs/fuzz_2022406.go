package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16; F2 string }
type GenS2 struct { F0 uint; F1 uint }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int32; F1 uint8 }
type GenS5 struct { F0 uint; F1 uint; F2 uint16 }
type GenS6 struct { F0 uint8; F1 uint64 }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 int }

func main() {
    var v1 int64 = 972
    var v2 float32 = ((1e10 / 1.0) + float32(1e10))
    var v3 int64 = 1
    v4 := make(map[uint32]bool)
    v5 := make(map[int]int)
    v6 := make(map[int]bool)
    v7 := GenS1{F0: 0, F1: 42, F2: ("hello" + (("foo bar" + "foo bar") + ""))}
    if true {
        var v8 int16 = 421
        _ = v8
    }
    var v9 float64 = 0.1
    var v10 int64 = (v1 - v3)
    switch 4294967295 {
    case 32768:
        if true {
            v11 := GenS1{F0: 271, F1: -32768, F2: "a"}
            _ = v11
        } else {
            if false {
                if false {
                    fmt.Println(v3, v2)
                    for fi12 := 0; fi12 < 19; fi12++ {
                        delete(v4, 2147483647)
                        var v13 float32 = (v2 / 1.0)
                        v10 = 42
                        _ = fi12
                        _ = v13
                    }
                }
                if false {
                    if (!false) {
                        for fi14 := 0; fi14 < 12; fi14++ {
                            for fi15 := 0; fi15 < 17; fi15++ {
                                v2 = v2
                                _ = fi15
                            }
                            v16 := GenS1{F0: 36, F1: -8, F2: (("\n\t" + "") + "hello")}
                            for fi17 := 0; fi17 < 18; fi17++ {
                                v18 := []int{}
                                v19 := GenS1{F0: 2, F1: 0, F2: "\n\t"}
                                _ = fi17
                                _ = v18
                                _ = v19
                            }
                            _ = fi14
                            _ = v16
                        }
                        v3 = v1
                        var v20 float64 = v9
                        _ = v20
                    }
                    for fi21 := 0; fi21 < 14; fi21++ {
                        for fi22 := 0; fi22 < 8; fi22++ {
                            v23 := GenS1{F0: 42, F1: 32767, F2: ((("a" + "a") + "") + "")}
                            v10 = (v3 / ((9223372036854775807) | 1))
                            var v24 uint64 = 18446744073709551615
                            fmt.Println(fi22, v1)
                            switch "" {
                            case "":
                                v2 -= v2
                            case "hello":
                                v9 = (0.1 - float64(((((v2 + (v2 + float32((v9 / v9)))) - 3.14159) * v2) / v2)))
                            }
                            v10 = v3
                            v10 = v10
                            _ = fi22
                            _ = v23
                            _ = v24
                        }
                        if (775 > 2147483648) {
                            fmt.Println(v9, v1, v2)
                            v25 := []int{fi21, -76, (-(fi21))}
                            _ = v25
                        } else {
                            var v26 string = "foo bar"
                            switch ("hello" + v26) {
                            case "foo bar":
                                v1 = 0
                            case "":
                                if (float64((v2 * (v2 - float32(v9)))) > 0.0) {
                                    fmt.Println(v1, v10, v26)
                                    v27 := []int{0, -280, 9223372036854775807}
                                    _ = v27
                                }
                            case "\n\t":
                                v1 = 187
                            }
                            for fi28 := 0; fi28 < 15; fi28++ {
                                fmt.Println(v3, v10)
                                switch fi21 {
                                case 128:
                                    v10 += (v10 | v10)
                                case 1:
                                    v2 = 1.5
                                }
                                v26 = v26
                                _ = fi28
                            }
                            _ = v26
                        }
                        v3 -= 0
                        _ = fi21
                    }
                } else {
                    for fi29 := 0; fi29 < 12; fi29++ {
                        switch -266 {
                        case 0:
                            var v30 int32 = -381
                            _ = v30
                        case 128:
                            v31 := []int{fi29}
                            _ = v31
                        default:
                            for fi32 := 0; fi32 < 14; fi32++ {
                                if true {
                                    fmt.Println(fi32, v9)
                                    v9 += 1e10
                                    var v33 int = fi32
                                    fmt.Println(fi32, v1)
                                    var v34 int16 = -32768
                                    _ = v33
                                    _ = v34
                                }
                                _ = fi32
                            }
                        }
                        v35 := []int{fi29, (fi29 * (fi29 | (fi29 * fi29))), -9223372036854775808}
                        if len(v35) >= 2 {
                            v35 = v35[1:len(v35)]
                        }
                        _ = fi29
                        _ = v35
                    }
                    v2 = v2
                    for fi36 := 0; fi36 < 6; fi36++ {
                        delete(v6, fi36)
                        if false {
                            var v37 int8 = 55
                            _ = v37
                        }
                        _ = fi36
                    }
                }
                v38 := GenS1{F0: 255, F1: 42, F2: "a"}
                _ = v38
            } else {
                v39 := GenS1{F0: 999, F1: -32768, F2: ((("a" + ("hello" + "foo bar")) + (("\n\t" + "\n\t") + "foo bar")) + "hello")}
                if true {
                    for fi40 := 0; fi40 < 10; fi40++ {
                        fmt.Println(v1, fi40, v2)
                        switch "\n\t" {
                        case "":
                            var v41 uint64 = 18446744073709551615
                            _ = v41
                        case "foo bar":
                            v2 = (v2 * 3.14159)
                        case "a":
                            v42 := GenS2{F0: 18446744073709551615, F1: 928}
                            _ = v42
                        }
                        var v43 int32 = -594
                        _ = fi40
                        _ = v43
                    }
                    for fi44 := 0; fi44 < 17; fi44++ {
                        if (65535 != 65535) {
                            v45 := []int{fi44, (fi44 + -2147483648), fi44, (fi44 + fi44)}
                            if (v2 == (-2.25 - v2)) {
                                if false {
                                    var v46 uint32 = 203
                                    if (!(!false)) {
                                        for fi47 := 0; fi47 < 14; fi47++ {
                                            v48 := GenS2{F0: 0, F1: 592}
                                            if (1000 == 32767) {
                                                if true {
                                                    var v49 string = "\n\t"
                                                    _ = v49
                                                } else {
                                                    fmt.Println(v9, v10, v1)
                                                    switch (fi44 + (fi47 << ((18446744073709551615) & 63))) {
                                                    case -465:
                                                        var v50 int8 = 119
                                                        _ = v50
                                                    }
                                                    if (true || (false && false)) {
                                                        fmt.Println(v3, v9, v3)
                                                    } else {
                                                        v46 = (v46 - v46)
                                                        v51 := GenS3{F0: 1}
                                                        _ = v51
                                                    }
                                                    v1 = (v10 - (v10 ^ (v10 << ((255) & 63))))
                                                    if (241 != 42) {
                                                        var v52 uint8 = 128
                                                        switch "foo bar" {
                                                        case "":
                                                            var v53 float32 = (0.1 / v2)
                                                            _ = v53
                                                        default:
                                                            v10 += v1
                                                        }
                                                        _ = v52
                                                    }
                                                    switch "\n\t" {
                                                    case "foo bar":
                                                        v2 += (v2 - (1e10 - ((v2 + ((v2 * float32(v9)) + 1.5)) / v2)))
                                                    case "":
                                                        if (true && true) {
                                                            var v54 bool = true
                                                            fmt.Println(fi47, v10, v54)
                                                            switch (fi44 >> ((1) & 63)) {
                                                            case 42:
                                                                v45 = append(v45, (fi44 ^ 4294967295))
                                                            }
                                                            _ = v54
                                                        }
                                                    case "a":
                                                        if (true || (v46 == v46)) {
                                                            v55 := GenS4{F0: 1, F1: 1}
                                                            v56 := GenS5{F0: 2, F1: 248, F2: 941}
                                                            v10 -= 797
                                                            _ = v55
                                                            _ = v56
                                                        }
                                                    default:
                                                        if true {
                                                            var v57 string = "a"
                                                            switch (v57 + v57) {
                                                            case "a":
                                                                if (!false) {
                                                                    var v58 int64 = (v10 + v3)
                                                                    _ = v58
                                                                }
                                                            default:
                                                                var v59 float64 = v9
                                                                _ = v59
                                                            }
                                                            _ = v57
                                                        }
                                                        v1 -= 120
                                                        v3 = v3
                                                    }
                                                    switch 9223372036854775807 {
                                                    case 647:
                                                        var v60 float32 = 0.1
                                                        v61 := GenS6{F0: 1, F1: 616}
                                                        delete(v6, (fi47 | 2147483647))
                                                        _ = v60
                                                        _ = v61
                                                    case 2147483647:
                                                        v2 = 0.0
                                                    case -246:
                                                        var v62 uint32 = (v46 << ((18446744073709551615) & 63))
                                                        _ = v62
                                                    }
                                                }
                                                switch ("" + ("" + "\n\t")) {
                                                case "hello":
                                                    var v63 int64 = (v10 * (v3 | v3))
                                                    _ = v63
                                                case "\n\t":
                                                    if (("foo bar" == ("" + ("hello" + "\n\t"))) && true) {
                                                        var v64 string = "hello"
                                                        switch (fi44 - int(v46)) {
                                                        case -1:
                                                            v1 -= (v3 ^ (v3 - v3))
                                                        case 128:
                                                            var v65 int = fi47
                                                            _ = v65
                                                        case 127:
                                                            switch ("foo bar" + ((v64 + ("" + "a")) + v64)) {
                                                            case "foo bar":
                                                                v9 += v9
                                                            case "hello":
                                                                switch 32768 {
                                                                case -221:
                                                                    v10 = 0
                                                                default:
                                                                    v66 := GenS7{F0: -128}
                                                                    _ = v66
                                                                }
                                                                var v67 int32 = -207
                                                                _ = v67
                                                            default:
                                                                v68 := GenS2{F0: 2, F1: 65535}
                                                                _ = v68
                                                            }
                                                        default:
                                                            var v69 float64 = 0.0
                                                            _ = v69
                                                        }
                                                        fmt.Println(v64)
                                                        _ = v64
                                                    }
                                                case "a":
                                                    v70 := GenS3{F0: 1}
                                                    _ = v70
                                                }
                                                v46 = v46
                                            } else {
                                                var v71 bool = false
                                                v2 = ((3.14159 / (1e10 / v2)) * v2)
                                                _ = v71
                                            }
                                            _ = fi47
                                            _ = v48
                                        }
                                        if len(v45) > 0 {
                                            v45[0] = fi44
                                        }
                                        var v72 int64 = -9223372036854775808
                                        _ = v72
                                    }
                                    _ = v46
                                } else {
                                    v73 := GenS7{F0: 0}
                                    fmt.Println(v2, v10, v2)
                                    var v74 int32 = -1
                                    _ = v73
                                    _ = v74
                                }
                                for fi75 := 0; fi75 < 20; fi75++ {
                                    fmt.Println(v10, v9)
                                    _ = fi75
                                }
                                fmt.Println(fi44, v9, fi44)
                            }
                            _ = v45
                        } else {
                            v1 = (v3 % (((v10 >> ((4294967295) & 63))) | 1))
                            var v76 bool = true
                            _ = v76
                        }
                        _ = fi44
                    }
                }
                _ = v39
            }
            v77 := GenS6{F0: 128, F1: 2}
            for fi78 := 0; fi78 < 13; fi78++ {
                for fi79 := 0; fi79 < 8; fi79++ {
                    v80 := GenS2{F0: 255, F1: 1}
                    _ = fi79
                    _ = v80
                }
                _ = fi78
            }
            _ = v77
        }
    }
    v81 := GenS5{F0: 150, F1: 4294967295, F2: 511}
    for fi82 := 0; fi82 < 8; fi82++ {
        v1 = v1
        v83 := []int{}
        _ = fi82
        _ = v83
    }
    if (96 < 486) {
        var v84 int = 0
        if ((-244 >= (v3 | (v1 >> ((50) & 63)))) && (!(965 >= 1))) {
            var v85 int = (v84 >> ((18446744073709551615) & 63))
            _ = v85
        }
        if true {
            if true {
                fmt.Println(v9, v1, v3)
                v86 := GenS6{F0: 127, F1: 0}
                _ = v86
            } else {
                v87 := []int{v84}
                if (true && false) {
                    fmt.Println(v1)
                    v3 -= v1
                }
                _ = v87
            }
            var v88 string = (("hello" + "foo bar") + "foo bar")
            _ = v88
        }
        _ = v84
    } else {
        v89 := GenS6{F0: 42, F1: 18446744073709551615}
        for fi90 := 0; fi90 < 3; fi90++ {
            for fi91 := 0; fi91 < 17; fi91++ {
                if (!true) {
                    v92 := v5[(fi91 * -153)]
                    if (false || false) {
                        var v93 int8 = -1
                        _ = v93
                    } else {
                        var v94 bool = false
                        var v95 int32 = 546
                        fmt.Println(v94, v1)
                        v10 = (-(v10))
                        v96 := []int{fi91}
                        _ = v94
                        _ = v95
                        _ = v96
                    }
                    _ = v92
                } else {
                    if (184 != 2) {
                        v97 := GenS3{F0: 0}
                        var v98 uint64 = 701
                        v2 += v2
                        _ = v97
                        _ = v98
                    }
                    var v99 int16 = 42
                    v1 = v3
                    _ = v99
                }
                var v100 uint8 = 96
                _ = fi91
                _ = v100
            }
            if (v2 != v2) {
                if true {
                    var v101 float32 = (0.1 * (v2 / v2))
                    _ = v101
                } else {
                    v102 := v4[522]
                    _ = v102
                }
                v103 := GenS8{F0: 630}
                if (-2147483648 == 42) {
                    var v104 uint16 = 633
                    _ = v104
                } else {
                    if (813 >= (fi90 >> ((0) & 63))) {
                        if ((-1 > -128) && (((v9 != 3.14159) || false) && (-9223372036854775808 > (v1 - (-(v1)))))) {
                            if (94 >= 699) {
                                v3 = -9223372036854775808
                                v105 := []int{}
                                v9 += ((float64(v2) / v9) + (0.1 / v9))
                                _ = v105
                            }
                            switch (fi90 - fi90) {
                            case -9223372036854775808:
                                for fi106 := 0; fi106 < 10; fi106++ {
                                    var v107 uint8 = 2
                                    v108 := GenS6{F0: (v107 << ((1) & 63)), F1: 751}
                                    v3 = -1
                                    _ = fi106
                                    _ = v107
                                    _ = v108
                                }
                            }
                        } else {
                            if false {
                                v1 = v3
                                if true {
                                    var v109 uint16 = 326
                                    var v110 uint16 = 2
                                    fmt.Println(v2, v109, fi90)
                                    _ = v109
                                    _ = v110
                                }
                            }
                        }
                    } else {
                        var v111 int16 = 361
                        for fi112 := 0; fi112 < 17; fi112++ {
                            if (0 >= 1) {
                                fmt.Println(v3)
                                var v113 string = (("foo bar" + ("a" + ("" + "hello"))) + ("\n\t" + "a"))
                                var v114 string = "\n\t"
                                var v115 int8 = 74
                                var v116 int8 = (v115 + v115)
                                _ = v113
                                _ = v114
                                _ = v115
                                _ = v116
                            }
                            _ = fi112
                        }
                        fmt.Println(v10, v10)
                        _ = v111
                    }
                }
                _ = v103
            } else {
                for fi117 := 0; fi117 < 6; fi117++ {
                    fmt.Println(fi90)
                    fmt.Println(fi117, fi90)
                    _ = fi117
                }
                v4[65535] = true
            }
            var v118 string = "hello"
            _ = fi90
            _ = v118
        }
        _ = v89
    }
    var v119 int = 42
    for fi120 := 0; fi120 < 17; fi120++ {
        var v121 int16 = -227
        fmt.Println(v10, v10)
        _ = fi120
        _ = v121
    }
    v2 = (float32((v9 + v9)) * float32(v9))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v119:", v119)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v119
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v81
    _ = v9
}
