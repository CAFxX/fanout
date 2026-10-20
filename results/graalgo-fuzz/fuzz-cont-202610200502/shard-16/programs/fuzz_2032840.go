package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint8; F1 string }
type GenS4 struct { F0 int; F1 int; F2 string }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint32; F1 int32; F2 int16 }
type GenS7 struct { F0 int16; F1 uint64; F2 int32 }
type GenS8 struct { F0 uint8; F1 int32; F2 int32 }
type GenS9 struct { F0 uint32; F1 uint }
type GenS10 struct { F0 bool; F1 int }
type GenS11 struct { F0 uint64; F1 int16 }
func genFn0(p0 uint16) uint16 {
    v7 := GenS1{F0: 65535}
    v8 := GenS1{F0: 255}
    v9 := GenS1{F0: 42}
    for fi10 := 0; fi10 < 5; fi10++ {
        v11 := []int{42, fi10, (fi10 ^ (^(fi10))), fi10}
        _ = fi10
        _ = v11
    }
    for fi12 := 0; fi12 < 3; fi12++ {
        var v13 uint32 = 65535
        v13 -= v13
        _ = fi12
        _ = v13
    }
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return (p0 ^ (p0 >> ((255) & 63)))
}
func genFn1(p0 bool) bool {
    if p0 {
        var v14 uint = 0
        var v15 uint32 = 42
        var v16 int8 = 0
        _ = v14
        _ = v15
        _ = v16
    }
    if false {
        switch "\n\t" {
        case "\n\t":
            for fi17 := 0; fi17 < 3; fi17++ {
                v18 := GenS2{F0: 927}
                fmt.Println(p0)
                switch 42 {
                case 130:
                    v19 := GenS3{F0: 11, F1: "foo bar"}
                    _ = v19
                case 32767:
                    if false {
                        p0 = (p0 && (!p0))
                        p0 = true
                    }
                }
                _ = fi17
                _ = v18
            }
        case "a":
            p0 = (0 <= 4294967295)
        default:
            v20 := []int{128, 128, -2147483648, -9223372036854775808}
            _ = v20
        }
        var v21 int16 = 32767
        fmt.Println(v21)
        _ = v21
    }
    for fi22 := 0; fi22 < 19; fi22++ {
        for fi23 := 0; fi23 < 3; fi23++ {
            v24 := []int{(^(fi23)), (fi23 | fi22)}
            p0 = (p0 && p0)
            if (p0 || (-2.25 > 0.1)) {
                var v25 float64 = 0.0
                var v26 uint32 = 2147483648
                _ = len(v24)
                var v27 int32 = 1
                _ = v25
                _ = v26
                _ = v27
            }
            _ = fi23
            _ = v24
        }
        for fi28 := 0; fi28 < 12; fi28++ {
            p0 = (108 > 1)
            if (p0 && ((fi22 * fi22) >= (^(fi28)))) {
                switch "a" {
                case "hello":
                    switch "a" {
                    case "hello":
                        v29 := []int{fi22}
                        _ = v29
                    case "foo bar":
                        if (213 < 18446744073709551615) {
                            switch "" {
                            case "a":
                                var v30 uint = 2
                                _ = v30
                            }
                        }
                    case "":
                        switch (^(fi22)) {
                        case 128:
                            switch 4294967295 {
                            case 55:
                                if p0 {
                                    v31 := []int{fi28, (fi28 % ((128) | 1)), fi22}
                                    if len(v31) > 0 {
                                        v31[0] = (^(fi22))
                                    }
                                    switch "hello" {
                                    case "foo bar":
                                        if ((1.5 <= 1.5) || p0) {
                                            _ = len(v31)
                                            var v32 uint8 = 0
                                            switch fi28 {
                                            case -9223372036854775808:
                                                if len(v31) > 0 {
                                                    v31[0] = 2147483647
                                                }
                                                p0 = ((((false || p0) || false) && p0) || ((p0 && true) || (42 >= 42)))
                                            case 32768:
                                                v31 = append(v31, fi22)
                                                if (2147483648 >= 65535) {
                                                    if ((fi28 + -398) != fi22) {
                                                        fmt.Println(p0)
                                                        v33 := GenS2{F0: 197}
                                                        fmt.Println(fi22, p0, fi28)
                                                        _ = v33
                                                    } else {
                                                        v34 := GenS4{F0: fi22, F1: (fi28 % ((fi28) | 1)), F2: "\n\t"}
                                                        p0 = ((1.5 < -2.25) || (((v32 >> ((1) & 63)) < 1) || (2 > 42)))
                                                        _ = v34
                                                    }
                                                    var v35 string = "hello"
                                                    _ = v35
                                                }
                                                v36 := GenS3{F0: 127, F1: ("foo bar" + "")}
                                                var v37 int = int(v32)
                                                _ = v36
                                                _ = v37
                                            case 127:
                                                var v38 bool = p0
                                                _ = v38
                                            }
                                            _ = v32
                                        } else {
                                            p0 = false
                                            v39 := GenS5{F0: 120}
                                            if len(v31) >= 2 {
                                                v31 = v31[1:len(v31)]
                                            }
                                            p0 = p0
                                            _ = v39
                                        }
                                    }
                                    _ = v31
                                } else {
                                    switch "a" {
                                    case "a":
                                        switch -459 {
                                        case -1:
                                            var v40 float64 = 1.5
                                            _ = v40
                                        case 89:
                                            var v41 uint32 = 2147483647
                                            _ = v41
                                        case 42:
                                            if (("hello" + ((("a" + "foo bar") + "hello") + ("\n\t" + "a"))) == "hello") {
                                                v42 := []int{(fi22 - (fi28 / ((fi28) | 1))), (fi22 % ((fi28) | 1)), (fi28 + fi28)}
                                                if len(v42) > 0 {
                                                    v42[0] = -9223372036854775808
                                                }
                                                _ = v42
                                            }
                                        }
                                    case "hello":
                                        v43 := []int{-469, fi28, (fi22 / (((fi22 + (fi22 >> ((0) & 63)))) | 1)), 496}
                                        _ = v43
                                    case "":
                                        p0 = ((42 >= 212) || p0)
                                    }
                                }
                            default:
                                v44 := GenS6{F0: 629, F1: -2147483648, F2: 719}
                                _ = v44
                            }
                        case 4294967295:
                            var v45 int = (fi22 * (fi22 + fi28))
                            _ = v45
                        case 9223372036854775807:
                            fmt.Println(p0, fi22)
                        default:
                            v46 := GenS7{F0: 42, F1: 42, F2: 2147483647}
                            _ = v46
                        }
                    }
                case "":
                    v47 := GenS8{F0: 1, F1: -2147483648, F2: 42}
                    _ = v47
                case "a":
                    var v48 int8 = -1
                    _ = v48
                default:
                    fmt.Println(fi22)
                }
                p0 = p0
                switch fi28 {
                case 9223372036854775807:
                    v49 := []int{}
                    _ = v49
                default:
                    v50 := GenS6{F0: 0, F1: 0, F2: -1}
                    _ = v50
                }
            }
            var v51 int32 = 2147483647
            _ = fi28
            _ = v51
        }
        _ = fi22
    }
    p0 = (65535 > 4294967295)
    fmt.Println(p0)
    _ = p0
    return p0
}

func main() {
    var v1 float32 = 3.14159
    v2 := make(map[uint32]string)
    var v3 int8 = 0
    var v4 int32 = -2147483648
    var v5 int16 = -710
    v6 := make(map[string]int)
    for fi52 := 0; fi52 < 17; fi52++ {
        var v53 uint32 = 4294967295
        v5 -= int16(v53)
        fmt.Println(v3, v5)
        _ = fi52
        _ = v53
    }
    v3 = (v3 | (v3 / (((v3 ^ (v3 << ((0) & 63)))) | 1)))
    var v54 uint64 = 18446744073709551615
    var v55 uint8 = 4
    v55 = v55
    for fi56 := 0; fi56 < 19; fi56++ {
        v54 = (v54 + 4294967295)
        var v57 float64 = 1e10
        _ = fi56
        _ = v57
    }
    fmt.Println(v4, v55)
    if (v1 < v1) {
        v58 := []int{32767}
        if true {
            if (42 > v54) {
                var v59 int64 = 566
                for fi60 := 0; fi60 < 16; fi60++ {
                    v5 += v5
                    for fi61 := 0; fi61 < 11; fi61++ {
                        fmt.Println(v4, v1)
                        _ = fi61
                    }
                    v4 -= -515
                    _ = fi60
                }
                for fi62 := 0; fi62 < 1; fi62++ {
                    v4 -= 736
                    _ = fi62
                }
                _ = v59
            }
        } else {
            var v63 uint32 = 0
            _ = v63
        }
        _ = v58
    } else {
        v64 := GenS1{F0: 65535}
        _ = v64
    }
    switch int(v4) {
    case -258:
        var v65 uint16 = 437
        _ = v65
    case 4294967295:
        v66 := GenS3{F0: v55, F1: "a"}
        _ = v66
    case 32768:
        v67 := GenS9{F0: 897, F1: 42}
        _ = v67
    default:
        if (!(!(!false))) {
            v68 := GenS9{F0: 2147483648, F1: 255}
            if (140 != (v55 & 2)) {
                v54 = 2
                switch ("\n\t" + "") {
                case "foo bar":
                    var v69 float64 = 3.14159
                    _ = v69
                case "a":
                    for fi70 := 0; fi70 < 17; fi70++ {
                        v55 += v55
                        _ = fi70
                    }
                }
            } else {
                fmt.Println(v5, v3, v4)
                if false {
                    v71 := GenS2{F0: 229}
                    _ = v71
                } else {
                    switch "foo bar" {
                    case "foo bar":
                        if true {
                            var v72 bool = true
                            v54 = (^(v54))
                            if true {
                                v4 = 1
                                for fi73 := 0; fi73 < 9; fi73++ {
                                    for fi74 := 0; fi74 < 1; fi74++ {
                                        switch "\n\t" {
                                        case "":
                                            v75 := GenS7{F0: (v5 / ((32767) | 1)), F1: v54, F2: 42}
                                            _ = v75
                                        case "a":
                                            v76 := GenS6{F0: 0, F1: (v4 << ((4294967295) & 63)), F2: v5}
                                            _ = v76
                                        case "\n\t":
                                            if (v55 <= v55) {
                                                v77 := GenS10{F0: ((v55 >> ((255) & 63)) > v55), F1: 32767}
                                                _ = v77
                                            }
                                        }
                                        var v78 string = "hello"
                                        var v79 int32 = (v4 % (((v4 << ((578) & 63))) | 1))
                                        fmt.Println(fi74, fi73)
                                        _ = fi74
                                        _ = v78
                                        _ = v79
                                    }
                                    _ = fi73
                                }
                            } else {
                                fmt.Println(v5, v3)
                                for fi80 := 0; fi80 < 16; fi80++ {
                                    var v81 int8 = (v3 + 32)
                                    v3 -= v81
                                    _ = fi80
                                    _ = v81
                                }
                            }
                            _ = v72
                        } else {
                            v1 -= (float32((1e10 + (float64(v1) / ((float64(((v1 - 0.1) / v1)) - 1e10) - ((1e10 * 0.0) * 1e10))))) - 0.0)
                            var v82 float32 = (v1 / -2.25)
                            if ((!(0.1 == ((((v1 - v82) * (v82 - v1)) / float32(0.1)) * (3.14159 / v1)))) || (!false)) {
                                v83 := []int{-2147483648, 9223372036854775807, 128}
                                v84 := GenS1{F0: 255}
                                var v85 uint32 = 2147483648
                                _ = v83
                                _ = v84
                                _ = v85
                            } else {
                                v86 := v2[2147483647]
                                fmt.Println(v82, v55, v5)
                                _ = v86
                            }
                            _ = v82
                        }
                    case "":
                        if false {
                            v4 = v4
                            v5 = v5
                        }
                    }
                    v5 -= v5
                }
                var v87 uint32 = 1
                _ = v87
            }
            for fi88 := 0; fi88 < 5; fi88++ {
                fmt.Println(v1)
                _ = fi88
            }
            _ = v68
        }
    }
    var v89 string = "\n\t"
    _ = len(v6)
    v89 = (((("" + ("foo bar" + "\n\t")) + (((v89 + ("\n\t" + v89)) + (v89 + (v89 + ""))) + v89)) + "foo bar") + v89)
    v90 := []int{9223372036854775807, 584}
    if ((v5 % (((^(v5))) | 1)) > (v5 | (v5 % ((v5) | 1)))) {
        v91 := v6[((v89 + (v89 + "hello")) + "hello")]
        _ = v91
    }
    v54 -= v54
    for fi92 := 0; fi92 < 5; fi92++ {
        v93 := GenS11{F0: v54, F1: int16(v55)}
        if false {
            var v94 int8 = int8(v5)
            _ = v94
        }
        var v95 int32 = v4
        _ = fi92
        _ = v93
        _ = v95
    }
    v55 -= 128
    for fi96 := 0; fi96 < 17; fi96++ {
        switch fi96 {
        case 42:
            var v97 float32 = v1
            _ = v97
        case 127:
            var v98 int8 = (v3 + -94)
            _ = v98
        default:
            switch v89 {
            case "a":
                v5 = 393
            case "\n\t":
                fmt.Println(v89, v3)
            default:
                if (((^(v54)) < 4294967295) && true) {
                    if false {
                        var v99 uint16 = 449
                        for fi100 := 0; fi100 < 13; fi100++ {
                            v101 := GenS6{F0: uint32(v54), F1: (v4 ^ (v4 - (v4 | int32(v99)))), F2: (-(v5))}
                            v2[2] = v89
                            if ((v54 * uint64(v55)) > v54) {
                                v54 = 278
                                if (fi100 > fi96) {
                                    switch (v89 + v89) {
                                    case "hello":
                                        if len(v90) > 0 {
                                            v90[0] = fi100
                                        }
                                    }
                                    var v102 int16 = v5
                                    switch (fi96 << ((14) & 63)) {
                                    case 0:
                                        if ("\n\t" != (("\n\t" + (v89 + v89)) + "")) {
                                            fmt.Println(v1, v54)
                                            var v103 uint = 681
                                            _ = v103
                                        } else {
                                            if len(v90) >= 2 {
                                                v90 = v90[1:len(v90)]
                                            }
                                            var v104 int16 = v5
                                            v89 = "foo bar"
                                            _ = v104
                                        }
                                    default:
                                        var v105 uint8 = (v55 >> ((363) & 63))
                                        _ = v105
                                    }
                                    _ = v102
                                } else {
                                    v90 = append(v90, (fi100 - (fi100 ^ fi96)))
                                    v106 := GenS5{F0: (v55 - v55)}
                                    _ = v106
                                }
                            } else {
                                fmt.Println(fi100, v54)
                                v55 = (v55 & 227)
                                switch (("hello" + ("a" + v89)) + "\n\t") {
                                case "a":
                                    v3 = v3
                                }
                                v55 = (v55 >> ((0) & 63))
                                var v107 float64 = (3.14159 + (float64((float32(0.1) + 3.14159)) + 0.0))
                                _ = v107
                            }
                            _ = fi100
                            _ = v101
                        }
                        var v108 uint16 = v99
                        _ = v108
                        _ = v99
                    } else {
                        if (float64((v1 * 3.14159)) < 1e10) {
                            if len(v90) > 0 {
                                v90[0] = (fi96 >> ((1) & 63))
                            }
                        } else {
                            var v109 int32 = v4
                            switch (((v89 + v89) + ("\n\t" + ((v89 + (v89 + v89)) + v89))) + ("hello" + (v89 + v89))) {
                            case "a":
                                v89 = "a"
                            default:
                                v110 := GenS5{F0: (v55 % (((^(v55))) | 1))}
                                _ = v110
                            }
                            _ = v109
                        }
                    }
                    for fi111 := 0; fi111 < 2; fi111++ {
                        if false {
                            var v112 float64 = 3.14159
                            fmt.Println(v54, v1)
                            _ = v112
                        }
                        _ = fi111
                    }
                }
            }
        }
        _ = fi96
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v54:", v54)
    fmt.Println("v55:", v55)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v89:", v89)
    fmt.Println("slicelen v90:", len(v90), v90)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v54
    _ = v55
    _ = v6
    _ = v89
    _ = v90
}
