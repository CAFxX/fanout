package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32; F2 uint }
type GenS2 struct { F0 uint32; F1 int64 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 bool; F1 uint8; F2 uint32 }
type GenS5 struct { F0 int64; F1 int32; F2 uint }
type GenS6 struct { F0 uint16; F1 int32 }
type GenS7 struct { F0 uint16 }
func genFn0() uint32 {
    v4 := GenS1{F0: 0, F1: -1, F2: 300}
    var v5 string = ("" + ("a" + "hello"))
    var v6 uint8 = 255
    if false {
        var v7 int32 = 17
        for fi8 := 0; fi8 < 5; fi8++ {
            for fi9 := 0; fi9 < 2; fi9++ {
                v10 := []int{4294967295}
                var v11 int32 = v7
                v6 += (v6 & v6)
                _ = fi9
                _ = v10
                _ = v11
            }
            v7 = v7
            _ = fi8
        }
        switch "hello" {
        case "foo bar":
            for fi12 := 0; fi12 < 16; fi12++ {
                v13 := GenS1{F0: 0, F1: -1, F2: 287}
                v7 += v7
                v14 := []int{fi12, (fi12 ^ fi12), 490, fi12}
                _ = fi12
                _ = v13
                _ = v14
            }
        case "\n\t":
            switch ("hello" + ((((v5 + v5) + v5) + ("a" + ((v5 + "foo bar") + ""))) + v5)) {
            case "a":
                var v15 int = 684
                _ = v15
            }
        case "hello":
            v16 := []int{128}
            _ = v16
        }
        _ = v7
    }
    var v17 string = v5
    _ = v17
    _ = v4
    _ = v5
    _ = v6
    return 0
}
func genFn1(p0 bool) int {
    if (false || (-447 <= 2147483647)) {
        for fi18 := 0; fi18 < 1; fi18++ {
            v19 := GenS2{F0: 928, F1: -9223372036854775808}
            fmt.Println(p0, fi18)
            p0 = p0
            _ = fi18
            _ = v19
        }
        fmt.Println(p0)
    } else {
        if p0 {
            var v20 float64 = ((3.14159 - 3.14159) + (1e10 + 1.5))
            _ = v20
        } else {
            if (-38 == 850) {
                v21 := []int{-728, 915}
                p0 = (((0.0 * -2.25) > ((3.14159 - (0.1 * 3.14159)) * (0.1 / -2.25))) || false)
                _ = v21
            } else {
                v22 := GenS1{F0: 215, F1: 42, F2: 65535}
                _ = v22
            }
            for fi23 := 0; fi23 < 3; fi23++ {
                switch "" {
                case "\n\t":
                    if p0 {
                        var v24 int = fi23
                        p0 = p0
                        _ = v24
                    }
                default:
                    var v25 bool = p0
                    _ = v25
                }
                _ = fi23
            }
        }
    }
    p0 = p0
    for fi26 := 0; fi26 < 17; fi26++ {
        v27 := []int{(fi26 & fi26), fi26, (fi26 << ((4294967295) & 63))}
        var v28 int8 = 42
        if (fi26 > (fi26 | 656)) {
            p0 = (!p0)
            p0 = ((!((!p0) && ((!(2 == 2)) && p0))) && (p0 || p0))
            switch ("foo bar" + "\n\t") {
            case "foo bar":
                for fi29 := 0; fi29 < 19; fi29++ {
                    switch "foo bar" {
                    case "a":
                        v30 := GenS3{F0: "hello"}
                        fmt.Println(p0, p0)
                        switch (fi29 % (((fi26 << ((0) & 63))) | 1)) {
                        case 127:
                            if ((("\n\t" + ("foo bar" + "hello")) + ("" + (("" + ((("\n\t" + ("" + "\n\t")) + "") + ("hello" + "\n\t"))) + (("foo bar" + ("" + ("" + "\n\t"))) + "foo bar")))) < "hello") {
                                var v31 float64 = 1e10
                                if (!p0) {
                                    fmt.Println(fi26, v31, fi26)
                                    v32 := GenS1{F0: 2, F1: -2147483648, F2: 0}
                                    _ = v32
                                }
                                if (-361 < 1) {
                                    v33 := GenS4{F0: false, F1: 0, F2: 2}
                                    p0 = (float32(v31) <= 0.1)
                                    p0 = false
                                    _ = v33
                                } else {
                                    var v34 int64 = 1
                                    v28 -= (v28 >> ((65535) & 63))
                                    if p0 {
                                        v35 := GenS3{F0: "foo bar"}
                                        v34 -= v34
                                        var v36 uint8 = 127
                                        _ = v35
                                        _ = v36
                                    }
                                    _ = v34
                                }
                                _ = v31
                            } else {
                                var v37 int16 = -160
                                var v38 int32 = 42
                                var v39 bool = (int32(v37) == -2147483648)
                                _ = v37
                                _ = v38
                                _ = v39
                            }
                        case 42:
                            p0 = (42 < 688)
                        case 128:
                            var v40 float64 = 1.5
                            _ = v40
                        default:
                            fmt.Println(p0, p0, fi26)
                        }
                        _ = v30
                    default:
                        _ = len(v27)
                        var v41 float64 = 1.5
                        if (426 > 9223372036854775807) {
                            var v42 bool = ((0.1 == float32(((1e10 + v41) / v41))) && (float64(0.1) >= v41))
                            v43 := GenS4{F0: (!(v42 && p0)), F1: 167, F2: 788}
                            _ = v42
                            _ = v43
                        }
                        _ = v41
                    }
                    if (818 > 255) {
                        v44 := GenS2{F0: 0, F1: 0}
                        _ = v44
                    } else {
                        switch (-(fi26)) {
                        case -1:
                            if (92 <= 42) {
                                p0 = (!((4294967295 == 4294967295) || (!(1 >= 360))))
                                v27 = append(v27, 9223372036854775807)
                            } else {
                                v45 := GenS1{F0: 778, F1: 2147483647, F2: 255}
                                switch "a" {
                                case "a":
                                    v28 += v28
                                case "":
                                    v46 := GenS1{F0: 1, F1: 42, F2: 65535}
                                    _ = v46
                                }
                                _ = v45
                            }
                        case 0:
                            fmt.Println(p0)
                        }
                        v47 := GenS2{F0: 2, F1: 1}
                        _ = v47
                    }
                    var v48 float64 = (0.1 + 3.14159)
                    _ = fi29
                    _ = v48
                }
            case "hello":
                for fi49 := 0; fi49 < 8; fi49++ {
                    var v50 float32 = (float32(3.14159) / 1.0)
                    v27 = append(v27, fi26)
                    _ = fi49
                    _ = v50
                }
            }
        } else {
            fmt.Println(p0)
            var v51 bool = p0
            v27 = append(v27, 42)
            _ = v51
        }
        _ = fi26
        _ = v27
        _ = v28
    }
    _ = p0
    return -9223372036854775808
}

func main() {
    v1 := make(map[string]int)
    v2 := make(map[int]uint32)
    var v3 uint8 = 54
    v3 -= (^(v3))
    v3 += v3
    if true {
        switch genFn1(true) {
        case 1:
            switch "" {
            case "a":
                var v52 int = genFn1((!false))
                _ = v52
            default:
                if (525 <= 1) {
                    fmt.Println(v3)
                }
            }
        }
    }
    var v53 int = genFn1((1 <= 2147483647))
    v3 = 2
    if false {
        switch v53 {
        case 127:
            delete(v2, (v53 - 2147483647))
        case 235:
            var v54 uint16 = 42
            _ = v54
        }
    }
    var v55 uint32 = 2
    if (true || false) {
        if (!(uint(v3) >= 4294967295)) {
            for fi56 := 0; fi56 < 16; fi56++ {
                for fi57 := 0; fi57 < 12; fi57++ {
                    var v58 int32 = 1
                    var v59 int = genFn1((!((v55 % ((genFn0()) | 1)) >= (v55 / ((771) | 1)))))
                    _ = fi57
                    _ = v58
                    _ = v59
                }
                var v60 uint8 = (^(v3))
                var v61 float64 = (float64(-2.25) - 0.1)
                _ = fi56
                _ = v60
                _ = v61
            }
            v55 += v55
        } else {
            if false {
                for fi62 := 0; fi62 < 10; fi62++ {
                    var v63 float64 = 0.0
                    if (!false) {
                        _ = len(v1)
                        if false {
                            for fi64 := 0; fi64 < 5; fi64++ {
                                v63 += ((3.14159 + v63) / 1.5)
                                switch "foo bar" {
                                case "foo bar":
                                    v65 := GenS3{F0: "a"}
                                    _ = v65
                                case "hello":
                                    var v66 int16 = 42
                                    _ = v66
                                default:
                                    v55 += (v55 / (((v55 - genFn0())) | 1))
                                }
                                _ = fi64
                            }
                            if ((629 <= 0) && (!false)) {
                                switch "" {
                                case "\n\t":
                                    var v67 int = (v53 | (v53 / (((fi62 / (((v53 / (((fi62 / ((v53) | 1))) | 1))) | 1))) | 1)))
                                    _ = v67
                                case "a":
                                    if true {
                                        v68 := GenS1{F0: 600, F1: 2147483647, F2: 255}
                                        v69 := GenS2{F0: v55, F1: 42}
                                        _ = v68
                                        _ = v69
                                    } else {
                                        if false {
                                            v70 := GenS5{F0: 921, F1: 42, F2: 1}
                                            v53 = v53
                                            _ = v70
                                        }
                                    }
                                }
                                fmt.Println(v53, v63, v3)
                                fmt.Println(v63, v53)
                            } else {
                                for fi71 := 0; fi71 < 12; fi71++ {
                                    var v72 uint = 18446744073709551615
                                    _ = fi71
                                    _ = v72
                                }
                            }
                        } else {
                            if true {
                                v63 += 1.5
                                v53 += (v53 % ((fi62) | 1))
                            }
                            v3 = 127
                        }
                        for fi73 := 0; fi73 < 4; fi73++ {
                            if true {
                                if ((((1.5 * v63) - (3.14159 / (v63 * float64(1.5)))) + v63) <= -2.25) {
                                    var v74 int16 = -32768
                                    var v75 float32 = (float32(v63) * float32(0.1))
                                    _ = v74
                                    _ = v75
                                }
                            }
                            v76 := GenS6{F0: 32767, F1: -2147483648}
                            _ = fi73
                            _ = v76
                        }
                    } else {
                        fmt.Println(v3, v3, fi62)
                    }
                    _ = fi62
                    _ = v63
                }
                var v77 int64 = -9223372036854775808
                _ = v77
            }
            if (false && (false && false)) {
                if true {
                    v53 = -1
                    v3 = (v3 >> ((65535) & 63))
                }
            } else {
                v53 -= v53
            }
            var v78 uint8 = 0
            _ = v78
        }
        fmt.Println(v53)
    } else {
        v79 := GenS3{F0: "a"}
        _ = v79
    }
    v2[32768] = (v55 | v55)
    if false {
        for fi80 := 0; fi80 < 3; fi80++ {
            fmt.Println(v53)
            v81 := []int{(v53 | int(v55)), (fi80 % (((v53 - genFn1(true))) | 1))}
            var v82 uint16 = 65535
            _ = fi80
            _ = v81
            _ = v82
        }
    } else {
        v55 = v55
    }
    if true {
        var v83 float64 = 3.14159
        switch v53 {
        case 9223372036854775807:
            switch ("a" + "a") {
            case "a":
                v84 := GenS2{F0: genFn0(), F1: 0}
                _ = v84
            case "":
                v3 += v3
            default:
                if (0 >= -128) {
                    v85 := GenS3{F0: "\n\t"}
                    var v86 bool = (-687 == -127)
                    fmt.Println(v83, v86, v53)
                    _ = v85
                    _ = v86
                }
            }
        }
        if (!("hello" < "foo bar")) {
            v3 += (v3 ^ 2)
            v87 := GenS4{F0: true, F1: v3, F2: 928}
            _ = v87
        } else {
            if (true || false) {
                v88 := GenS2{F0: genFn0(), F1: -1}
                _ = v88
            } else {
                switch (v53 ^ (v53 * (v53 + v53))) {
                case 9223372036854775807:
                    if (-1 == 915) {
                        v89 := []int{(v53 / ((v53) | 1)), v53}
                        var v90 float64 = v83
                        _ = v89
                        _ = v90
                    }
                case -2147483648:
                    if (1 < -701) {
                        for fi91 := 0; fi91 < 6; fi91++ {
                            for fi92 := 0; fi92 < 18; fi92++ {
                                v55 = (v55 ^ (^(v55)))
                                _ = fi92
                            }
                            v83 += float64(float32(float64(3.14159)))
                            var v93 int64 = 42
                            _ = fi91
                            _ = v93
                        }
                    } else {
                        v94 := []int{}
                        v95 := GenS3{F0: "a"}
                        delete(v2, genFn1(true))
                        _ = v94
                        _ = v95
                    }
                case 187:
                    v96 := []int{}
                    _ = v96
                default:
                    for fi97 := 0; fi97 < 18; fi97++ {
                        v98 := v1["a"]
                        for fi99 := 0; fi99 < 15; fi99++ {
                            var v100 int32 = 1
                            v101 := GenS4{F0: (v100 <= (v100 % ((v100) | 1))), F1: v3, F2: (v55 % ((v55) | 1))}
                            v3 = 1
                            _ = fi99
                            _ = v100
                            _ = v101
                        }
                        _ = fi97
                        _ = v98
                    }
                }
                var v102 float32 = float32(-2.25)
                _ = v102
            }
        }
        _ = v83
    } else {
        var v103 int64 = -9223372036854775808
        _ = v103
    }
    for fi104 := 0; fi104 < 2; fi104++ {
        for fi105 := 0; fi105 < 1; fi105++ {
            if (4294967295 >= 2) {
                v106 := GenS6{F0: uint16(v3), F1: 2147483647}
                switch fi105 {
                case 2147483647:
                    fmt.Println(fi104, fi104)
                    v55 += (v55 * genFn0())
                default:
                    var v107 int = (fi104 >> ((893) & 63))
                    _ = v107
                }
                v55 -= v55
                _ = v106
            }
            if (-2147483648 >= 1) {
                var v108 string = ("foo bar" + "a")
                v3 += 42
                v108 = v108
                _ = v108
            } else {
                v3 += v3
                fmt.Println(v55, fi105, fi105)
                if (false || false) {
                    _ = len(v1)
                    v109 := v1["foo bar"]
                    _ = v109
                }
            }
            if false {
                v110 := []int{(v53 << ((310) & 63))}
                _ = v110
            }
            _ = fi105
        }
        v55 += v55
        switch genFn1(true) {
        case -229:
            if ((v3 + v3) <= (v3 ^ v3)) {
                var v111 uint = 255
                var v112 uint8 = v3
                _ = v111
                _ = v112
            } else {
                v3 = (v3 % ((v3) | 1))
                fmt.Println(v55)
            }
        case 9223372036854775807:
            _ = len(v2)
        }
        _ = fi104
    }
    v1["a"] = (v53 / ((0) | 1))
    v3 -= v3
    switch (("" + ("" + ("" + ""))) + (("" + ("" + (("a" + "a") + (("hello" + "") + ("\n\t" + "a"))))) + "foo bar")) {
    case "hello":
        var v113 float32 = 1e10
        _ = v113
    case "a":
        fmt.Println(v3, v55, v3)
    }
    switch (v53 & (v53 + (v53 + 2147483647))) {
    case 9223372036854775807:
        switch genFn1(true) {
        case -2147483648:
            if (-710 <= -166) {
                var v114 int32 = -2147483648
                v115 := GenS7{F0: 65535}
                _ = v114
                _ = v115
            } else {
                for fi116 := 0; fi116 < 7; fi116++ {
                    var v117 int16 = 32767
                    _ = fi116
                    _ = v117
                }
                v118 := []int{(v53 << ((65535) & 63)), genFn1(true)}
                _ = v118
            }
        }
    }
    if (false && false) {
        var v119 uint = 255
        _ = v119
    }
    for fi120 := 0; fi120 < 11; fi120++ {
        _ = len(v1)
        v55 += (v55 << ((42) & 63))
        _ = fi120
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v53:", v53)
    fmt.Println("v55:", v55)
    _ = v1
    _ = v2
    _ = v3
    _ = v53
    _ = v55
}
