package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 bool; F1 int; F2 uint8 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 bool; F1 int; F2 int32 }
type GenS6 struct { F0 int16; F1 uint }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 uint; F1 int; F2 int16 }
type GenS9 struct { F0 uint; F1 bool; F2 uint8 }
type GenS10 struct { F0 int32; F1 uint }
type GenS11 struct { F0 string; F1 uint64; F2 int }
type GenS12 struct { F0 uint8 }
type GenS13 struct { F0 uint64 }
func genFn0(p0 int, p1 uint, p2 uint16) uint {
    v6 := []int{(p0 & (^(p0))), 9223372036854775807}
    for fi7 := 0; fi7 < 19; fi7++ {
        for fi8 := 0; fi8 < 12; fi8++ {
            _ = len(v6)
            p2 = (p2 + p2)
            if len(v6) > 0 {
                v6[0] = p0
            }
            var v9 int64 = 929
            _ = fi8
            _ = v9
        }
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        _ = fi7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return p1
}

func main() {
    v1 := make(map[string]uint32)
    var v2 uint = 0
    v3 := make(map[int]string)
    var v4 int8 = 0
    v5 := make(map[int]bool)
    v10 := GenS1{F0: 18446744073709551615, F1: 566}
    if (("foo bar" + "\n\t") <= "foo bar") {
        var v11 uint64 = 2
        v11 += 42
        for fi12 := 0; fi12 < 2; fi12++ {
            v13 := GenS1{F0: v11, F1: 997}
            _ = fi12
            _ = v13
        }
        _ = v11
    } else {
        switch 4294967295 {
        case 0:
            v4 = v4
        default:
            v14 := GenS2{F0: (4294967295 > 1)}
            _ = v14
        }
        switch -2147483648 {
        case 32767:
            v4 = v4
        default:
            v15 := GenS3{F0: false, F1: -696, F2: 0}
            _ = v15
        }
    }
    fmt.Println(v2, v2)
    var v16 float32 = -2.25
    fmt.Println(v16, v4)
    var v17 uint64 = 733
    for fi18 := 0; fi18 < 13; fi18++ {
        for fi19 := 0; fi19 < 13; fi19++ {
            if (v2 == 2) {
                if false {
                    v4 += 0
                    if ((fi19 - fi18) == (fi19 & fi19)) {
                        fmt.Println(fi19, v16)
                        switch fi19 {
                        case -404:
                            v20 := []int{fi18}
                            _ = v20
                        default:
                            if false {
                                var v21 string = "\n\t"
                                _ = v21
                            } else {
                                v4 = v4
                                var v22 int = fi18
                                var v23 uint32 = 750
                                v2 = 924
                                _ = v22
                                _ = v23
                            }
                        }
                        var v24 uint8 = 25
                        _ = v24
                    } else {
                        switch "a" {
                        case "hello":
                            switch "a" {
                            case "\n\t":
                                v25 := []int{(fi19 & (^(fi18)))}
                                if len(v25) > 0 {
                                    v25[0] = 4294967295
                                }
                                _ = v25
                            default:
                                var v26 uint8 = uint8(v17)
                                _ = v26
                            }
                        case "\n\t":
                            v27 := GenS4{F0: -9223372036854775808}
                            _ = v27
                        }
                        fmt.Println(fi19, v16, fi18)
                        var v28 int8 = v4
                        _ = v28
                    }
                    if true {
                        v2 -= v2
                        var v29 uint8 = 128
                        _ = v29
                    } else {
                        v30 := []int{(fi19 * -205), -2147483648, -2147483648, fi18}
                        delete(v5, (fi18 - fi19))
                        var v31 float64 = (1e10 + 3.14159)
                        _ = v30
                        _ = v31
                    }
                    if true {
                        if (3.14159 >= 1.5) {
                            v32 := GenS5{F0: ((!("a" == "hello")) && false), F1: (fi19 * fi18), F2: 42}
                            if false {
                                var v33 int32 = 2147483647
                                _ = v33
                            }
                            var v34 int = 123
                            if false {
                                fmt.Println(fi18, v34)
                            } else {
                                v35 := []int{(fi19 >> ((v2) & 63)), fi19}
                                v16 = v16
                                v36 := GenS6{F0: 91, F1: (v2 << ((genFn0(fi19, v2, 255)) & 63))}
                                _ = v35
                                _ = v36
                            }
                            _ = v32
                            _ = v34
                        } else {
                            v4 -= v4
                        }
                    } else {
                        switch ("hello" + ("hello" + "a")) {
                        case "hello":
                            if false {
                                v4 = v4
                                v37 := []int{32767, 42}
                                _ = v37
                            }
                        case "a":
                            if (280 <= 4294967295) {
                                switch "hello" {
                                case "\n\t":
                                    fmt.Println(v17, fi19)
                                case "":
                                    var v38 float64 = 1e10
                                    _ = v38
                                case "foo bar":
                                    v16 = (1.5 + v16)
                                default:
                                    var v39 int32 = 2147483647
                                    v3[fi18] = ("" + ("hello" + "hello"))
                                    v40 := []int{fi18}
                                    _ = v39
                                    _ = v40
                                }
                                v4 += v4
                                _ = len(v3)
                                _ = len(v1)
                            }
                        default:
                            v17 = (v17 >> ((v2) & 63))
                            v41 := GenS7{F0: 65535}
                            switch fi18 {
                            case -1:
                                v42 := []int{(fi18 + (fi19 & 1))}
                                _ = v42
                            case -2147483648:
                                switch "hello" {
                                case "hello":
                                    v2 -= 4294967295
                                    v43 := GenS3{F0: (2147483647 != 428), F1: (fi19 + fi19), F2: 76}
                                    switch (-(fi19)) {
                                    case 42:
                                        if (!(false || false)) {
                                            switch "\n\t" {
                                            case "hello":
                                                switch (fi19 ^ (-(fi19))) {
                                                case -747:
                                                    var v44 uint16 = 42
                                                    _ = v44
                                                default:
                                                    fmt.Println(v2)
                                                    switch (fi18 << (((v2 * v2)) & 63)) {
                                                    case 4294967295:
                                                        switch 32768 {
                                                        case -9223372036854775808:
                                                            var v45 uint16 = 255
                                                            v16 += float32(3.14159)
                                                            _ = v45
                                                        case 9223372036854775807:
                                                            var v46 float32 = (3.14159 + v16)
                                                            _ = v46
                                                        case -674:
                                                            fmt.Println(fi18)
                                                        }
                                                    }
                                                }
                                            case "a":
                                                v2 = 65535
                                            case "\n\t":
                                                v47 := GenS8{F0: v2, F1: fi19, F2: -675}
                                                _ = v47
                                            default:
                                                if (true || false) {
                                                    var v48 float64 = (0.0 * -2.25)
                                                    if (1000 >= 884) {
                                                        if true {
                                                            var v49 uint16 = 255
                                                            _ = v49
                                                        }
                                                    } else {
                                                        var v50 int64 = 58
                                                        var v51 uint32 = uint32(v17)
                                                        v52 := GenS9{F0: 4294967295, F1: false, F2: 128}
                                                        _ = v50
                                                        _ = v51
                                                        _ = v52
                                                    }
                                                    _ = v48
                                                } else {
                                                    v53 := v3[fi18]
                                                    switch (v53 + "a") {
                                                    case "foo bar":
                                                        v3[fi19] = v53
                                                    case "hello":
                                                        v54 := GenS1{F0: (v17 / ((v17) | 1)), F1: 42}
                                                        _ = v54
                                                    }
                                                    _ = v53
                                                }
                                            }
                                            v17 += v17
                                        }
                                    }
                                    v4 += -1
                                    _ = v43
                                case "":
                                    v55 := v5[4294967295]
                                    _ = v55
                                case "\n\t":
                                    v17 = 464
                                default:
                                    v56 := GenS1{F0: (v17 >> ((42) & 63)), F1: 1}
                                    _ = v56
                                }
                            }
                            _ = v41
                        }
                    }
                }
                v57 := GenS7{F0: 42}
                var v58 int64 = int64(v17)
                _ = v57
                _ = v58
            } else {
                v59 := GenS2{F0: true}
                switch (("\n\t" + "foo bar") + ("" + "a")) {
                case "\n\t":
                    var v60 float64 = 0.1
                    _ = v60
                case "a":
                    v17 -= (v17 | 18446744073709551615)
                }
                _ = v59
            }
            var v61 uint8 = 42
            v62 := GenS4{F0: -52}
            _ = fi19
            _ = v61
            _ = v62
        }
        var v63 int = (fi18 ^ fi18)
        for fi64 := 0; fi64 < 10; fi64++ {
            if (float64((((1e10 * float32(1e10)) / v16) + 3.14159)) < -2.25) {
                _ = len(v3)
                if true {
                    v65 := []int{(fi64 * (v63 / ((2147483647) | 1)))}
                    _ = v65
                } else {
                    var v66 uint16 = 692
                    _ = v66
                }
                v2 -= genFn0(2147483647, v2, 255)
            }
            var v67 uint32 = 1
            if ((v4 / ((v4) | 1)) == v4) {
                v68 := GenS8{F0: (^(v2)), F1: (fi64 + fi18), F2: -32768}
                var v69 int32 = -1
                _ = v68
                _ = v69
            } else {
                if ((^(v67)) >= v67) {
                    v2 = v2
                    fmt.Println(v67, v2)
                    fmt.Println(v16, fi18, v4)
                    v17 += (v17 >> (((v2 & (v2 % ((42) | 1)))) & 63))
                }
            }
            _ = fi64
            _ = v67
        }
        _ = fi18
        _ = v63
    }
    fmt.Println(v16, v16, v17)
    for fi70 := 0; fi70 < 1; fi70++ {
        v71 := GenS10{F0: 0, F1: 42}
        var v72 int16 = -933
        if false {
            if false {
                v1[("a" + "a")] = 1
                if false {
                    if ((0.0 == v16) || (!false)) {
                        fmt.Println(v16, v17, v72)
                    }
                    switch fi70 {
                    case -95:
                        var v73 uint16 = 42
                        _ = v73
                    }
                }
                switch "foo bar" {
                case "\n\t":
                    if true {
                        if (!(v72 >= -1)) {
                            v74 := GenS3{F0: (1 != 1), F1: (^(fi70)), F2: 0}
                            _ = v74
                        }
                        if (((!true) && true) && (255 != v2)) {
                            for fi75 := 0; fi75 < 10; fi75++ {
                                v76 := GenS11{F0: "\n\t", F1: (v17 << ((v2) & 63)), F2: -509}
                                _ = len(v1)
                                _ = fi75
                                _ = v76
                            }
                        } else {
                            var v77 float32 = (v16 + 0.1)
                            v16 = 1.5
                            v78 := []int{(fi70 / ((9223372036854775807) | 1)), (fi70 * 1), (fi70 - fi70), -1}
                            _ = v77
                            _ = v78
                        }
                    } else {
                        v79 := GenS5{F0: (1 < 65535), F1: (fi70 + (fi70 << ((genFn0(-518, (^(v2)), 65535)) & 63))), F2: -1}
                        for fi80 := 0; fi80 < 9; fi80++ {
                            v17 -= 2
                            _ = fi80
                        }
                        if true {
                            var v81 string = ("foo bar" + (("\n\t" + "hello") + ("foo bar" + ("hello" + ""))))
                            _ = v81
                        }
                        _ = v79
                    }
                }
            } else {
                switch "hello" {
                case "hello":
                    fmt.Println(v2)
                case "":
                    v16 = v16
                }
            }
            var v82 int32 = 2147483647
            if (!false) {
                v83 := []int{-941, 231, 51}
                if (2 == 2147483647) {
                    var v84 bool = (793 <= 255)
                    var v85 uint32 = 2147483647
                    _ = v84
                    _ = v85
                } else {
                    fmt.Println(v72, v17, fi70)
                }
                _ = v83
            } else {
                switch ("a" + ("foo bar" + "")) {
                case "a":
                    switch "\n\t" {
                    case "foo bar":
                        var v86 int8 = v4
                        _ = v86
                    }
                case "foo bar":
                    for fi87 := 0; fi87 < 19; fi87++ {
                        v17 -= v17
                        var v88 string = ("" + "hello")
                        var v89 float64 = (3.14159 - 0.1)
                        _ = fi87
                        _ = v88
                        _ = v89
                    }
                case "hello":
                    switch (fi70 << ((65535) & 63)) {
                    case 822:
                        if true {
                            for fi90 := 0; fi90 < 12; fi90++ {
                                v91 := []int{}
                                var v92 string = "hello"
                                if (995 == -871) {
                                    if true {
                                        v92 = ""
                                    } else {
                                        var v93 int8 = (v4 & -128)
                                        _ = v93
                                    }
                                    v94 := GenS12{F0: 1}
                                    _ = v94
                                }
                                _ = fi90
                                _ = v91
                                _ = v92
                            }
                        }
                    case 9223372036854775807:
                        switch (("hello" + "") + "") {
                        case "":
                            var v95 int64 = -9223372036854775808
                            _ = v95
                        default:
                            switch "a" {
                            case "hello":
                                v96 := GenS3{F0: ((false || true) || (uint8(v17) > 47)), F1: -2147483648, F2: 42}
                                _ = v96
                            }
                        }
                    case 0:
                        v17 -= 1
                    }
                default:
                    var v97 bool = (false && ((false || false) || ((v72 & int16(fi70)) != (v72 * 1))))
                    _ = v97
                }
                var v98 int32 = int32(v72)
                _ = v98
            }
            _ = v82
        } else {
            var v99 int16 = v72
            var v100 int16 = (-(v72))
            v101 := []int{806, (fi70 & int(v100)), 10}
            _ = v100
            _ = v101
            _ = v99
        }
        _ = fi70
        _ = v71
        _ = v72
    }
    for fi102 := 0; fi102 < 9; fi102++ {
        var v103 bool = (((42 == -32768) || ((true || ((v16 + v16) > 0.1)) && (2147483648 < 2147483647))) || ((!false) || false))
        var v104 int64 = 649
        v4 = v4
        _ = fi102
        _ = v103
        _ = v104
    }
    var v105 uint8 = 60
    v106 := []int{-9223372036854775808, 32768}
    v17 -= (^(v17))
    _ = len(v3)
    for fi107 := 0; fi107 < 16; fi107++ {
        if (1 > 1) {
            var v108 float32 = float32(1.5)
            v109 := GenS11{F0: (("\n\t" + ("\n\t" + "")) + "a"), F1: v17, F2: (fi107 << (((v2 << (((v2 % (((v2 << ((4294967295) & 63))) | 1))) & 63))) & 63))}
            var v110 uint16 = 42
            _ = v108
            _ = v109
            _ = v110
        } else {
            v111 := GenS8{F0: (v2 % ((v2) | 1)), F1: fi107, F2: -93}
            _ = v111
        }
        for fi112 := 0; fi112 < 3; fi112++ {
            v113 := GenS13{F0: uint64(v105)}
            _ = fi112
            _ = v113
        }
        v105 += (v105 / ((v105) | 1))
        _ = fi107
    }
    fmt.Println(v17, v17)
    if true {
        for fi114 := 0; fi114 < 10; fi114++ {
            fmt.Println(v17)
            _ = fi114
        }
    } else {
        fmt.Println(v105, v17)
        if len(v106) > 0 {
            v106[0] = 32768
        }
        v115 := v1[("\n\t" + ("hello" + "hello"))]
        _ = v115
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v105:", v105)
    fmt.Println("slicelen v106:", len(v106), v106)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v105
    _ = v106
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
