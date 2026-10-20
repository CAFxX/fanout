package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32; F2 uint16 }
type GenS2 struct { F0 uint; F1 uint32 }
type GenS3 struct { F0 uint; F1 int16 }
type GenS4 struct { F0 bool; F1 int32; F2 bool }
type GenS5 struct { F0 int; F1 uint64; F2 int16 }
type GenS6 struct { F0 string; F1 int8 }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 uint16; F1 string }
type GenS9 struct { F0 uint8; F1 string; F2 string }
type GenS10 struct { F0 uint8; F1 int32; F2 uint }
func genFn0() bool {
    switch 117 {
    case -57:
        var v7 string = "foo bar"
        _ = v7
    case 325:
        v8 := []int{-420, 127, 4294967295, -203}
        _ = v8
    default:
        switch "" {
        case "hello":
            v9 := GenS1{F0: 0, F1: -1, F2: 42}
            _ = v9
        }
    }
    for fi10 := 0; fi10 < 11; fi10++ {
        for fi11 := 0; fi11 < 13; fi11++ {
            var v12 uint8 = 42
            _ = fi11
            _ = v12
        }
        switch "" {
        case "":
            var v13 uint16 = 255
            _ = v13
        case "foo bar":
            fmt.Println(fi10)
        case "a":
            if true {
                var v14 uint16 = 32767
                _ = v14
            }
        }
        _ = fi10
    }
    v15 := GenS1{F0: -2147483648, F1: -2147483648, F2: 94}
    if false {
        for fi16 := 0; fi16 < 12; fi16++ {
            v17 := GenS2{F0: 4294967295, F1: 42}
            var v18 int32 = 1
            _ = fi16
            _ = v17
            _ = v18
        }
    }
    _ = v15
    return false
}
func genFn1() int {
    var v19 uint8 = 205
    fmt.Println(v19)
    for fi20 := 0; fi20 < 15; fi20++ {
        v19 += v19
        _ = fi20
    }
    var v21 float64 = 0.0
    for fi22 := 0; fi22 < 11; fi22++ {
        fmt.Println(v19)
        if (42 < 127) {
            for fi23 := 0; fi23 < 20; fi23++ {
                v24 := []int{fi23}
                var v25 string = ("\n\t" + "")
                _ = fi23
                _ = v24
                _ = v25
            }
            for fi26 := 0; fi26 < 2; fi26++ {
                var v27 int64 = -845
                v27 = v27
                _ = fi26
                _ = v27
            }
            fmt.Println(fi22, fi22)
        } else {
            switch ("" + "\n\t") {
            case "":
                for fi28 := 0; fi28 < 7; fi28++ {
                    var v29 float64 = -2.25
                    v19 += 22
                    var v30 string = "\n\t"
                    fmt.Println(fi22, v29)
                    _ = fi28
                    _ = v29
                    _ = v30
                }
            case "foo bar":
                for fi31 := 0; fi31 < 5; fi31++ {
                    v32 := []int{32767, (^(fi22)), (fi31 << ((65535) & 63)), fi31}
                    _ = fi31
                    _ = v32
                }
            default:
                if ((2147483647 != 4294967295) && (42 < 100)) {
                    switch ((("a" + "hello") + "") + "") {
                    case "hello":
                        v21 -= ((v21 * float64(float32(v21))) - v21)
                    case "a":
                        if false {
                            v33 := []int{}
                            _ = v33
                        } else {
                            var v34 int32 = 221
                            for fi35 := 0; fi35 < 10; fi35++ {
                                fmt.Println(v34, fi22, fi22)
                                _ = fi35
                            }
                            v36 := []int{(fi22 % ((fi22) | 1)), (-(fi22)), fi22, -9223372036854775808}
                            _ = v34
                            _ = v36
                        }
                    case "foo bar":
                        if (!(1 == 1)) {
                            v19 += (^(v19))
                            v19 = (^(v19))
                        }
                    }
                }
            }
            if (65535 != 2147483647) {
                if ((("a" != "\n\t") || true) && (float32(v21) <= ((float32((((float64(1e10) + float64(1e10)) + ((((v21 + (v21 - (float64(1.5) / float64(((0.0 / 1.5) - 0.1))))) / v21) - -2.25) - 0.0)) / (v21 * (1e10 * v21)))) / 0.1) + 3.14159))) {
                    if (false || false) {
                        var v37 int = fi22
                        for fi38 := 0; fi38 < 6; fi38++ {
                            switch (("\n\t" + "hello") + "") {
                            case "":
                                var v39 int32 = -2147483648
                                _ = v39
                            case "\n\t":
                                v40 := GenS2{F0: 281, F1: uint32(v19)}
                                _ = v40
                            default:
                                switch (((("hello" + "foo bar") + "hello") + (("a" + "") + ("\n\t" + ((("a" + ("hello" + "hello")) + ((("hello" + "") + ("\n\t" + "a")) + (("hello" + "a") + "\n\t"))) + "foo bar")))) + ("\n\t" + ("hello" + ((("foo bar" + "") + "a") + "a")))) {
                                case "foo bar":
                                    v19 += (v19 + v19)
                                case "\n\t":
                                    var v41 uint32 = 2147483648
                                    _ = v41
                                case "":
                                    v42 := GenS1{F0: -499, F1: 42, F2: 65535}
                                    _ = v42
                                default:
                                    if (0 != -507) {
                                        v19 = v19
                                    }
                                }
                            }
                            _ = fi38
                        }
                        switch "foo bar" {
                        case "foo bar":
                            v37 += (fi22 % ((v37) | 1))
                        case "hello":
                            var v43 float64 = 3.14159
                            _ = v43
                        case "\n\t":
                            v44 := []int{v37, fi22}
                            _ = v44
                        default:
                            var v45 float32 = 1.5
                            _ = v45
                        }
                        _ = v37
                    } else {
                        if (2 < 2147483648) {
                            v46 := []int{(fi22 % ((2147483647) | 1)), (fi22 & -2147483648)}
                            var v47 int64 = -646
                            var v48 int = 0
                            _ = v46
                            _ = v47
                            _ = v48
                        } else {
                            v49 := GenS1{F0: 2, F1: -664, F2: 904}
                            switch "" {
                            case "foo bar":
                                var v50 int64 = 938
                                _ = v50
                            case "\n\t":
                                if (!(111 > -38)) {
                                    v19 = v19
                                } else {
                                    switch "\n\t" {
                                    case "hello":
                                        var v51 uint8 = (v19 % (((v19 >> ((651) & 63))) | 1))
                                        _ = v51
                                    default:
                                        v52 := GenS2{F0: 759, F1: 0}
                                        _ = v52
                                    }
                                    v19 -= v19
                                }
                            default:
                                if true {
                                    for fi53 := 0; fi53 < 19; fi53++ {
                                        switch 32767 {
                                        case -710:
                                            var v54 bool = (!(true && (fi53 > (fi22 % (((fi53 << ((2) & 63))) | 1)))))
                                            _ = v54
                                        default:
                                            v55 := GenS3{F0: 2, F1: -14}
                                            _ = v55
                                        }
                                        _ = fi53
                                    }
                                    v56 := []int{fi22}
                                    _ = v56
                                } else {
                                    switch (fi22 & 127) {
                                    case 4294967295:
                                        for fi57 := 0; fi57 < 13; fi57++ {
                                            v19 = 136
                                            _ = fi57
                                        }
                                    case -1:
                                        if (((!(-1 != -328)) || true) || (false || false)) {
                                            v58 := []int{(fi22 ^ fi22), fi22, (fi22 ^ fi22)}
                                            _ = v58
                                        }
                                    default:
                                        var v59 uint8 = (v19 / (((v19 & v19)) | 1))
                                        _ = v59
                                    }
                                    switch "a" {
                                    case "":
                                        if (!true) {
                                            fmt.Println(v19, fi22, v19)
                                            v60 := []int{(fi22 - fi22), fi22, (fi22 % ((fi22) | 1)), -9223372036854775808}
                                            _ = v60
                                        }
                                    case "foo bar":
                                        if false {
                                            switch "foo bar" {
                                            case "":
                                                v21 = (v21 * 1e10)
                                            case "\n\t":
                                                switch "a" {
                                                case "foo bar":
                                                    v21 -= 1e10
                                                }
                                            default:
                                                v21 = 0.1
                                            }
                                            if (!true) {
                                                v61 := []int{(fi22 & (fi22 * 32767)), -436}
                                                _ = v61
                                            }
                                        } else {
                                            v21 += v21
                                            v62 := GenS4{F0: false, F1: 1, F2: (true && true)}
                                            v63 := GenS1{F0: -965, F1: 42, F2: 2}
                                            _ = v62
                                            _ = v63
                                        }
                                    default:
                                        v64 := GenS5{F0: 9223372036854775807, F1: 691, F2: 434}
                                        _ = v64
                                    }
                                    if ((245 >= fi22) || false) {
                                        v19 = 42
                                    }
                                }
                            }
                            _ = v49
                        }
                        for fi65 := 0; fi65 < 18; fi65++ {
                            switch (fi65 + (fi65 << ((352) & 63))) {
                            case -531:
                                switch (((("hello" + (("a" + "\n\t") + "a")) + "hello") + "\n\t") + ("a" + "a")) {
                                case "foo bar":
                                    v19 = v19
                                case "":
                                    if false {
                                        var v66 int8 = -128
                                        if true {
                                            v67 := GenS4{F0: true, F1: -1, F2: (false || true)}
                                            if (291 > 255) {
                                                if ((^(v19)) <= 255) {
                                                    v68 := GenS6{F0: ("a" + ""), F1: v66}
                                                    _ = v68
                                                }
                                                v19 = (v19 / ((v19) | 1))
                                                var v69 int32 = -1
                                                _ = v69
                                            }
                                            _ = v67
                                        }
                                        v19 = (v19 << ((50) & 63))
                                        _ = v66
                                    }
                                }
                            }
                            v21 = v21
                            if false {
                                v70 := GenS7{F0: -955}
                                switch ("foo bar" + "hello") {
                                case "\n\t":
                                    v21 = 3.14159
                                }
                                switch 32768 {
                                case -373:
                                    fmt.Println(fi22)
                                    v71 := GenS4{F0: (v19 <= v19), F1: 813, F2: true}
                                    _ = v71
                                case -9223372036854775808:
                                    if ("foo bar" <= "") {
                                        var v72 uint8 = 93
                                        v21 += v21
                                        v72 -= (v72 | 127)
                                        _ = v72
                                    } else {
                                        v19 = (v19 << ((65535) & 63))
                                    }
                                case -462:
                                    var v73 bool = ((false && false) && ((true || false) && false))
                                    _ = v73
                                }
                                _ = v70
                            } else {
                                switch (fi22 ^ -846) {
                                case -9223372036854775808:
                                    v74 := GenS8{F0: 65535, F1: ("a" + "a")}
                                    _ = v74
                                case 9223372036854775807:
                                    var v75 uint64 = uint64(v19)
                                    _ = v75
                                case 32768:
                                    v76 := []int{(fi22 / (((fi22 << ((719) & 63))) | 1)), (fi22 ^ (-(fi22))), -9223372036854775808}
                                    _ = v76
                                default:
                                    v19 -= 222
                                }
                                var v77 uint32 = 2
                                fmt.Println(v21, v77)
                                _ = v77
                            }
                            _ = fi65
                        }
                    }
                } else {
                    v19 -= (v19 * 128)
                    if (!(true && true)) {
                        var v78 int16 = 1000
                        for fi79 := 0; fi79 < 20; fi79++ {
                            if ((!true) && false) {
                                switch ("hello" + "\n\t") {
                                case "":
                                    var v80 string = "foo bar"
                                    _ = v80
                                case "foo bar":
                                    v19 += v19
                                }
                            }
                            _ = fi79
                        }
                        _ = v78
                    }
                    var v81 uint = 4294967295
                    _ = v81
                }
                for fi82 := 0; fi82 < 14; fi82++ {
                    v83 := []int{fi22, (fi22 & (fi22 & (fi82 / ((fi22) | 1))))}
                    fmt.Println(fi82, fi22)
                    _ = fi82
                    _ = v83
                }
                switch -9223372036854775808 {
                case 9223372036854775807:
                    for fi84 := 0; fi84 < 2; fi84++ {
                        v21 += (1.5 * v21)
                        _ = fi84
                    }
                case -665:
                    fmt.Println(v21, v19, v19)
                }
            } else {
                v85 := []int{-2147483648}
                switch fi22 {
                case 128:
                    if (1 < 518) {
                        fmt.Println(v21)
                    } else {
                        _ = len(v85)
                        v19 += 214
                    }
                case 0:
                    v19 = (v19 << ((42) & 63))
                case 625:
                    var v86 string = ""
                    _ = v86
                default:
                    switch fi22 {
                    case -9223372036854775808:
                        v21 = v21
                    default:
                        switch "\n\t" {
                        case "a":
                            for fi87 := 0; fi87 < 18; fi87++ {
                                v88 := GenS6{F0: "a", F1: 0}
                                var v89 uint16 = 2
                                _ = fi87
                                _ = v88
                                _ = v89
                            }
                        }
                    }
                }
                _ = v85
            }
        }
        v90 := GenS5{F0: 32768, F1: 18446744073709551615, F2: -898}
        _ = fi22
        _ = v90
    }
    _ = v19
    _ = v21
    return 127
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[string]bool)
    v3 := make(map[uint32]bool)
    var v4 uint32 = 460
    var v5 uint32 = v4
    v6 := make(map[uint32]bool)
    var v91 bool = true
    v92 := GenS9{F0: 120, F1: ((((("" + "a") + "a") + "\n\t") + "") + ""), F2: ""}
    fmt.Println(v91)
    var v93 int16 = -3
    for fi94 := 0; fi94 < 3; fi94++ {
        delete(v6, (v4 * 19))
        _ = fi94
    }
    for fi95 := 0; fi95 < 2; fi95++ {
        v91 = false
        v91 = v91
        if (4294967295 != 159) {
            fmt.Println(v91, v93, v91)
            v5 += 2147483647
            v4 = v5
        } else {
            if false {
                if (!(v91 && true)) {
                    switch ("a" + "") {
                    case "foo bar":
                        v96 := []int{(fi95 - (fi95 / ((fi95) | 1))), (fi95 + -313), (fi95 | (fi95 % ((-9223372036854775808) | 1)))}
                        _ = v96
                    case "hello":
                        var v97 uint64 = 2
                        _ = v97
                    default:
                        for fi98 := 0; fi98 < 7; fi98++ {
                            var v99 float32 = 0.0
                            v100 := []int{genFn1(), fi95}
                            var v101 int16 = v93
                            _ = fi98
                            _ = v100
                            _ = v101
                            _ = v99
                        }
                    }
                    v93 = (v93 % (((-(v93))) | 1))
                }
            } else {
                v102 := GenS3{F0: 65535, F1: -943}
                _ = v102
            }
            v103 := GenS10{F0: 143, F1: 42, F2: 4294967295}
            _ = v103
        }
        _ = fi95
    }
    v1[v5] = "hello"
    v4 = (v5 * (v5 ^ 65535))
    switch 128 {
    case 42:
        _ = len(v6)
    default:
        _ = len(v1)
    }
    v91 = (v91 && ((113 < 42) || (664 >= 2147483647)))
    switch -1 {
    case -956:
        if v91 {
            if false {
                fmt.Println(v5, v4, v93)
                for fi104 := 0; fi104 < 8; fi104++ {
                    var v105 int8 = -8
                    for fi106 := 0; fi106 < 18; fi106++ {
                        v5 -= (^(v4))
                        var v107 int16 = v93
                        v3[v5] = true
                        v93 += (v93 ^ 1000)
                        if v91 {
                            v91 = (!v91)
                            v91 = ((fi106 ^ fi104) >= 309)
                            v105 = 127
                        }
                        v93 = v93
                        _ = fi106
                        _ = v107
                    }
                    var v108 int = fi104
                    _ = fi104
                    _ = v105
                    _ = v108
                }
            }
        }
    default:
        var v109 uint32 = v5
        _ = v109
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v91:", v91)
    fmt.Println("v93:", v93)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v91
    _ = v92
    _ = v93
}
