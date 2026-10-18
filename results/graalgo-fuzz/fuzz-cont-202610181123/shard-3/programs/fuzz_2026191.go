package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint32; F1 bool; F2 int8 }
type GenS3 struct { F0 uint16; F1 uint; F2 int16 }
type GenS4 struct { F0 uint16; F1 int8; F2 int64 }
type GenS5 struct { F0 int8; F1 uint64; F2 int8 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 uint64; F1 uint16; F2 uint }
type GenS9 struct { F0 string }
type GenS10 struct { F0 int16; F1 uint8; F2 string }
func genFn0(p0 uint) uint16 {
    p0 -= (p0 / ((p0) | 1))
    v5 := GenS1{F0: 1}
    var v6 uint8 = 18
    _ = p0
    _ = v5
    _ = v6
    return 32767
}
func genFn1(p0 string) string {
    fmt.Println(p0)
    switch (((p0 + p0) + "foo bar") + ((p0 + "\n\t") + "foo bar")) {
    case "a":
        v7 := []int{-495, -2147483648}
        _ = v7
    default:
        for fi8 := 0; fi8 < 20; fi8++ {
            var v9 uint16 = 277
            if false {
                v9 -= 28
            } else {
                v9 = (v9 % ((v9) | 1))
            }
            _ = fi8
            _ = v9
        }
    }
    p0 = p0
    p0 = (p0 + (p0 + p0))
    v10 := []int{0, 9223372036854775807, 9223372036854775807}
    _ = p0
    _ = v10
    return p0
}

func main() {
    var v1 uint8 = 127
    v2 := make(map[string]uint32)
    var v3 uint64 = 0
    var v4 uint8 = 243
    fmt.Println(v3, v4, v3)
    v1 = (v1 / (((v4 % ((v1) | 1))) | 1))
    var v11 uint32 = 132
    var v12 int32 = int32(v3)
    for fi13 := 0; fi13 < 5; fi13++ {
        if false {
            for fi14 := 0; fi14 < 7; fi14++ {
                var v15 uint32 = v11
                _ = fi14
                _ = v15
            }
        }
        v4 += v1
        for fi16 := 0; fi16 < 18; fi16++ {
            switch "" {
            case "foo bar":
                fmt.Println(v1, v12)
            case "a":
                v3 += (^(v3))
                var v17 int16 = 1
                v18 := GenS2{F0: (^(v11)), F1: true, F2: 0}
                _ = v17
                _ = v18
            default:
                fmt.Println(fi16, v12, fi16)
            }
            _ = len(v2)
            _ = fi16
        }
        _ = fi13
    }
    for fi19 := 0; fi19 < 18; fi19++ {
        v11 -= 682
        v11 -= (v11 + (^(v11)))
        _ = fi19
    }
    var v20 int32 = (v12 << ((907) & 63))
    var v21 int8 = -88
    v20 = 42
    for fi22 := 0; fi22 < 7; fi22++ {
        var v23 uint64 = v3
        switch "\n\t" {
        case "a":
            if false {
                if (v1 != v4) {
                    fmt.Println(v11, v4, v23)
                    for fi24 := 0; fi24 < 15; fi24++ {
                        v3 = (^(v23))
                        if (float32(1.5) <= (((0.0 + 1e10) - ((1.5 + ((float32(3.14159) / 3.14159) / (1e10 * 3.14159))) - 3.14159)) * 1.5)) {
                            if (!false) {
                                v3 = v3
                                var v25 string = ""
                                if ((v3 | v23) == (^(v23))) {
                                    if false {
                                        fmt.Println(v21)
                                        v26 := []int{-333, fi22, fi24}
                                        _ = v26
                                    }
                                } else {
                                    var v27 uint16 = 0
                                    _ = v27
                                }
                                _ = v25
                            } else {
                                v28 := GenS3{F0: genFn0(42), F1: 4294967295, F2: 1}
                                _ = v28
                            }
                            switch ((("\n\t" + "hello") + "") + "hello") {
                            case "a":
                                if true {
                                    v29 := GenS4{F0: genFn0(0), F1: (v21 & 42), F2: 346}
                                    _ = v29
                                }
                            case "foo bar":
                                var v30 float64 = -2.25
                                _ = v30
                            case "":
                                if true {
                                    v31 := GenS5{F0: (v21 ^ (v21 >> ((255) & 63))), F1: (v23 * uint64(v4)), F2: -35}
                                    switch 1 {
                                    case 625:
                                        var v32 uint8 = 255
                                        _ = v32
                                    case 557:
                                        v3 -= v3
                                        v3 += 1
                                        if (true && ("a" == "\n\t")) {
                                            var v33 uint64 = (v23 ^ 770)
                                            var v34 bool = (v11 > v11)
                                            _ = v33
                                            _ = v34
                                        } else {
                                            switch (fi22 | fi24) {
                                            case 524:
                                                var v35 uint64 = 483
                                                _ = v35
                                            case -9223372036854775808:
                                                if false {
                                                    v1 += (v1 % ((114) | 1))
                                                } else {
                                                    v11 -= 4294967295
                                                }
                                            case -105:
                                                switch (-(fi22)) {
                                                case -483:
                                                    if true {
                                                        v36 := GenS2{F0: (v11 ^ 2147483648), F1: false, F2: (-(v21))}
                                                        switch "\n\t" {
                                                        case "\n\t":
                                                            v2["\n\t"] = v11
                                                        case "hello":
                                                            v4 -= v1
                                                        default:
                                                            if false {
                                                                switch ("hello" + "a") {
                                                                case "hello":
                                                                    fmt.Println(v11, v20, v4)
                                                                }
                                                                if false {
                                                                    _ = len(v2)
                                                                    var v37 uint8 = v1
                                                                    _ = v37
                                                                }
                                                            } else {
                                                                v23 = 889
                                                                v3 += (v23 - v23)
                                                                fmt.Println(fi24, fi22)
                                                            }
                                                        }
                                                        if (42 != 9223372036854775807) {
                                                            v38 := v2["a"]
                                                            var v39 uint64 = 620
                                                            _ = v38
                                                            _ = v39
                                                        } else {
                                                            switch ("" + ("\n\t" + "\n\t")) {
                                                            case "a":
                                                                if true {
                                                                    v12 -= v20
                                                                    var v40 int8 = (^(v21))
                                                                    if false {
                                                                        if false {
                                                                            var v41 uint64 = (v3 & (^(v3)))
                                                                            if (255 > 934) {
                                                                                var v42 uint8 = uint8(v3)
                                                                                _ = len(v2)
                                                                                _ = v42
                                                                            } else {
                                                                                v43 := GenS6{F0: v1}
                                                                                _ = v43
                                                                            }
                                                                            _ = v41
                                                                        } else {
                                                                            var v44 int = fi22
                                                                            var v45 uint8 = v4
                                                                            v46 := GenS7{F0: 42}
                                                                            _ = v44
                                                                            _ = v45
                                                                            _ = v46
                                                                        }
                                                                    } else {
                                                                        fmt.Println(v11)
                                                                    }
                                                                    if ((v3 | (v3 + v23)) >= (v23 >> ((42) & 63))) {
                                                                        fmt.Println(v1, fi24)
                                                                        fmt.Println(fi24, v11, v40)
                                                                        delete(v2, "")
                                                                    }
                                                                    if true {
                                                                        var v47 int16 = 1
                                                                        v48 := GenS5{F0: v40, F1: (v3 - (v23 << ((1) & 63))), F2: v21}
                                                                        v20 = (v12 | v12)
                                                                        _ = v47
                                                                        _ = v48
                                                                    }
                                                                    _ = v40
                                                                } else {
                                                                    fmt.Println(v11, v21, v11)
                                                                    v1 = (v4 ^ 255)
                                                                    var v49 uint32 = (v11 / ((0) | 1))
                                                                    _ = v49
                                                                }
                                                            default:
                                                                switch "foo bar" {
                                                                case "foo bar":
                                                                    var v50 int8 = v21
                                                                    _ = v50
                                                                case "\n\t":
                                                                    switch (fi22 - 128) {
                                                                    case 416:
                                                                        v1 = (v4 - (v1 / (((v1 << ((906) & 63))) | 1)))
                                                                    default:
                                                                        v51 := GenS5{F0: (v21 % ((v21) | 1)), F1: v23, F2: v21}
                                                                        _ = v51
                                                                    }
                                                                case "hello":
                                                                    var v52 string = "foo bar"
                                                                    _ = v52
                                                                }
                                                            }
                                                            switch "a" {
                                                            case "hello":
                                                                var v53 uint64 = v3
                                                                _ = v53
                                                            default:
                                                                fmt.Println(v12)
                                                            }
                                                            switch "a" {
                                                            case "":
                                                                var v54 int32 = 357
                                                                _ = v54
                                                            }
                                                        }
                                                        fmt.Println(v11, v4)
                                                        v55 := []int{fi24, fi24}
                                                        _ = v36
                                                        _ = v55
                                                    } else {
                                                        var v56 int16 = 1000
                                                        v57 := GenS8{F0: 178, F1: 1, F2: 1}
                                                        v58 := GenS3{F0: 96, F1: 585, F2: (v56 * -781)}
                                                        v1 -= (v4 % (((^(v4))) | 1))
                                                        v59 := GenS6{F0: v1}
                                                        _ = v56
                                                        _ = v57
                                                        _ = v58
                                                        _ = v59
                                                    }
                                                default:
                                                    v11 += 626
                                                }
                                            }
                                            if false {
                                                var v60 float64 = ((((3.14159 - 0.1) / -2.25) - 0.0) - -2.25)
                                                v61 := []int{fi24, fi22, fi24}
                                                fmt.Println(fi24)
                                                _ = v60
                                                _ = v61
                                            }
                                        }
                                    default:
                                        switch ("a" + "a") {
                                        case "hello":
                                            if (42 == 32767) {
                                                delete(v2, (("hello" + "a") + (((("foo bar" + "") + "hello") + ("foo bar" + "\n\t")) + "")))
                                                if false {
                                                    var v62 int32 = (v12 + v12)
                                                    _ = v62
                                                } else {
                                                    v63 := []int{(fi22 ^ 374), fi22, fi22}
                                                    switch fi22 {
                                                    case -1:
                                                        fmt.Println(v23, v20)
                                                    case 128:
                                                        v21 = (v21 & v21)
                                                    }
                                                    _ = v63
                                                }
                                            }
                                        case "a":
                                            v4 = (v1 ^ v1)
                                        }
                                    }
                                    v23 = 1
                                    _ = v31
                                } else {
                                    switch ("foo bar" + (("hello" + "") + "hello")) {
                                    case "\n\t":
                                        switch "foo bar" {
                                        case "a":
                                            switch 9223372036854775807 {
                                            case 2147483647:
                                                v1 = v4
                                            case 127:
                                                if (true || ((-856 >= -32768) || false)) {
                                                    v23 += v3
                                                }
                                            case 42:
                                                v12 = (v12 * (v12 - (v20 << ((183) & 63))))
                                            default:
                                                v1 = 2
                                            }
                                        default:
                                            switch fi24 {
                                            case 75:
                                                var v64 uint64 = v23
                                                _ = v64
                                            default:
                                                v65 := []int{fi22}
                                                _ = v65
                                            }
                                        }
                                    default:
                                        v4 += 255
                                    }
                                    fmt.Println(v11, v4, v12)
                                }
                            }
                            if false {
                                if (false || (2 > 4294967295)) {
                                    fmt.Println(v23, v4)
                                    var v66 uint64 = v3
                                    _ = v66
                                }
                                if (!true) {
                                    v23 = v23
                                    var v67 float32 = 3.14159
                                    if (((false || false) || (true && false)) && (65535 != 509)) {
                                        v11 -= (v11 << ((42) & 63))
                                    }
                                    _ = v67
                                }
                            }
                        }
                        if false {
                            v1 = (v4 - 255)
                            var v68 int64 = 0
                            v3 -= (v23 | v23)
                            _ = v68
                        } else {
                            var v69 int8 = 126
                            v12 += (v12 + v20)
                            if true {
                                if (v23 != v3) {
                                    if (false || ((!true) || false)) {
                                        v11 = (v11 ^ 2147483648)
                                    } else {
                                        v69 += v21
                                        v21 = v21
                                        v70 := GenS9{F0: "hello"}
                                        _ = v70
                                    }
                                    fmt.Println(v11)
                                } else {
                                    switch (fi22 - (-(fi24))) {
                                    case 1:
                                        if false {
                                            v71 := GenS7{F0: 42}
                                            _ = v71
                                        } else {
                                            _ = len(v2)
                                        }
                                    case 42:
                                        v3 = 2
                                        var v72 int = fi22
                                        if (!(((v1 * v1) > (v1 * (v4 % ((127) | 1)))) || (0.0 < (1.5 / ((1.5 - ((0.0 + 0.0) * (-2.25 * 3.14159))) / 3.14159))))) {
                                            v12 = int32(v3)
                                            v69 = v69
                                        } else {
                                            var v73 uint32 = v11
                                            _ = v73
                                        }
                                        _ = v72
                                    case 127:
                                        switch "\n\t" {
                                        case "hello":
                                            fmt.Println(v3, v11)
                                        case "":
                                            var v74 int8 = v69
                                            _ = v74
                                        case "\n\t":
                                            switch fi22 {
                                            case -1:
                                                if (false && true) {
                                                    v20 -= (^(v12))
                                                    v75 := []int{32767, 128}
                                                    v1 = (v4 & (v1 << ((383) & 63)))
                                                    _ = v75
                                                } else {
                                                    v76 := GenS9{F0: "foo bar"}
                                                    v21 = (v69 % ((v21) | 1))
                                                    _ = v76
                                                }
                                            default:
                                                var v77 string = "hello"
                                                _ = v77
                                            }
                                        }
                                    }
                                    v2["foo bar"] = v11
                                    var v78 int64 = int64(v21)
                                    v79 := GenS6{F0: v4}
                                    v80 := []int{fi22}
                                    _ = v78
                                    _ = v79
                                    _ = v80
                                }
                                var v81 uint16 = uint16(v11)
                                _ = v81
                            }
                            _ = v69
                        }
                        _ = fi24
                    }
                } else {
                    var v82 int64 = -9223372036854775808
                    var v83 float32 = 1.5
                    _ = v82
                    _ = v83
                }
                for fi84 := 0; fi84 < 11; fi84++ {
                    v23 -= v3
                    _ = fi84
                }
            }
        }
        _ = fi22
        _ = v23
    }
    if (!true) {
        switch "\n\t" {
        case "a":
            var v85 int16 = 1000
            _ = v85
        }
    } else {
        fmt.Println(v20, v12)
        fmt.Println(v12)
    }
    fmt.Println(v3, v4)
    var v86 float64 = 3.14159
    for fi87 := 0; fi87 < 19; fi87++ {
        for fi88 := 0; fi88 < 4; fi88++ {
            switch -2147483648 {
            case -1:
                if (float64(-2.25) != ((v86 / (float64((1e10 - (3.14159 - 0.0))) + v86)) / 1.5)) {
                    var v89 int16 = 1000
                    v89 = -996
                    if (!false) {
                        v90 := GenS10{F0: (v89 >> ((2) & 63)), F1: v1, F2: "\n\t"}
                        _ = v90
                    }
                    if (255 >= v1) {
                        if (!("\n\t" >= ("\n\t" + "foo bar"))) {
                            var v91 bool = (uint8(v3) < (v4 >> ((uint(v3)) & 63)))
                            _ = v91
                        } else {
                            v92 := v2[("a" + "\n\t")]
                            var v93 bool = false
                            _ = v92
                            _ = v93
                        }
                        v1 -= (v4 / ((v4) | 1))
                        if false {
                            v94 := GenS3{F0: 65535, F1: 4294967295, F2: (v89 * v89)}
                            _ = v94
                        }
                    } else {
                        var v95 int8 = (v21 ^ v21)
                        var v96 uint16 = 42
                        _ = v95
                        _ = v96
                    }
                    _ = v89
                } else {
                    if (("foo bar" + (("" + "foo bar") + ("a" + ("foo bar" + "a")))) > ((("hello" + "hello") + "a") + ("" + ""))) {
                        v1 -= 1
                    } else {
                        var v97 uint16 = 2
                        switch (fi88 - fi87) {
                        case -2147483648:
                            v4 += (v1 << ((1) & 63))
                        case -135:
                            delete(v2, ("hello" + "a"))
                        }
                        _ = v97
                    }
                    v21 = 1
                }
            }
            var v98 bool = (v86 >= (((v86 + (0.0 * v86)) * ((v86 - v86) - (v86 * v86))) + v86))
            _ = fi88
            _ = v98
        }
        _ = fi87
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v86:", v86)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v86
}
