package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool; F2 int16 }
type GenS2 struct { F0 uint16; F1 bool }
type GenS3 struct { F0 int32; F1 uint64 }
type GenS4 struct { F0 int32; F1 uint64 }
type GenS5 struct { F0 int; F1 uint; F2 uint }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 bool; F1 int32 }
type GenS8 struct { F0 int32; F1 uint64; F2 int16 }
func genFn0(p0 bool) int {
    fmt.Println(p0)
    p0 = (1.5 < (0.1 + (1.5 / 1.0)))
    v6 := GenS1{F0: 1, F1: p0, F2: 32767}
    if true {
        fmt.Println(p0)
        for fi7 := 0; fi7 < 13; fi7++ {
            p0 = p0
            var v8 float64 = -2.25
            switch (("a" + ("\n\t" + ("\n\t" + "\n\t"))) + "foo bar") {
            case "a":
                if (1 <= 42) {
                    fmt.Println(fi7)
                    var v9 uint64 = 2
                    switch "foo bar" {
                    case "hello":
                        var v10 string = ("a" + "\n\t")
                        _ = v10
                    }
                    _ = v9
                } else {
                    v11 := []int{-654, 32768, fi7, fi7}
                    _ = v11
                }
            default:
                switch (fi7 * (fi7 ^ (-(fi7)))) {
                case -9223372036854775808:
                    var v12 int8 = -1
                    _ = v12
                case 9223372036854775807:
                    for fi13 := 0; fi13 < 15; fi13++ {
                        fmt.Println(fi13, fi7, fi7)
                        _ = fi13
                    }
                case 32767:
                    var v14 uint16 = 813
                    _ = v14
                default:
                    fmt.Println(fi7, fi7)
                }
            }
            _ = fi7
            _ = v8
        }
    }
    var v15 float64 = float64(1.5)
    _ = p0
    _ = v15
    _ = v6
    return 1
}
func genFn1(p0 uint16, p1 bool, p2 int16) uint {
    v16 := GenS1{F0: 952, F1: (0 > 18446744073709551615), F2: (p2 << ((491) & 63))}
    for fi17 := 0; fi17 < 20; fi17++ {
        for fi18 := 0; fi18 < 13; fi18++ {
            var v19 int16 = p2
            v20 := GenS2{F0: p0, F1: p1}
            if p1 {
                var v21 uint8 = 2
                fmt.Println(p2, p1, v19)
                p2 -= p2
                p2 = p2
                _ = v21
            }
            _ = fi18
            _ = v19
            _ = v20
        }
        _ = fi17
    }
    var v22 uint64 = 0
    fmt.Println(p1, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    _ = v22
    return 65535
}

func main() {
    var v1 float64 = (0.1 - 1e10)
    var v2 float32 = (-2.25 - -2.25)
    v3 := make(map[int]uint32)
    v4 := make(map[int]uint32)
    v5 := []uint32{844, 0, 0, 0}
    v23 := GenS3{F0: -1, F1: 4294967295}
    if ((255 < 255) || (!(0 == 1))) {
        for fi24 := 0; fi24 < 17; fi24++ {
            if (!true) {
                var v25 int8 = 1
                _ = v25
            } else {
                v26 := GenS3{F0: -2147483648, F1: 499}
                for fi27 := 0; fi27 < 17; fi27++ {
                    v1 = float64(v2)
                    if false {
                        v2 = v2
                    }
                    var v28 uint32 = 2147483648
                    _ = fi27
                    _ = v28
                }
                switch ("foo bar" + "hello") {
                case "a":
                    for fi29 := 0; fi29 < 7; fi29++ {
                        v30 := GenS1{F0: -930, F1: (-30 < 127), F2: 878}
                        fmt.Println(fi24, v2)
                        var v31 uint16 = 32767
                        _ = fi29
                        _ = v30
                        _ = v31
                    }
                case "\n\t":
                    var v32 bool = (!(42 != 0))
                    _ = v32
                case "hello":
                    for fi33 := 0; fi33 < 2; fi33++ {
                        fmt.Println(fi33)
                        fmt.Println(fi24)
                        _ = fi33
                    }
                default:
                    if false {
                        for fi34 := 0; fi34 < 20; fi34++ {
                            v35 := GenS4{F0: 2147483647, F1: 0}
                            switch genFn0(false) {
                            case 2147483647:
                                if true {
                                    switch "foo bar" {
                                    case "hello":
                                        var v36 uint = 1
                                        _ = v36
                                    case "\n\t":
                                        var v37 uint32 = 4294967295
                                        v2 += (v2 * (((v2 + 0.1) * 1.5) + 1e10))
                                        _ = v37
                                    default:
                                        fmt.Println(v2, fi24)
                                    }
                                }
                            case -450:
                                if len(v5) > 0 {
                                    v5[0] = 1
                                }
                                var v38 float32 = float32(0.1)
                                v38 = 0.0
                                _ = v38
                            case 9:
                                var v39 int64 = 1
                                _ = v39
                            default:
                                switch genFn0(true) {
                                case -12:
                                    v40 := GenS5{F0: fi24, F1: 65535, F2: genFn1(42, false, 107)}
                                    _ = v40
                                case 926:
                                    v2 = 0.0
                                    v41 := GenS1{F0: -802, F1: true, F2: 539}
                                    if (0 < 0) {
                                        v2 += 0.0
                                    } else {
                                        v4[127] = 2147483647
                                        var v42 int8 = 100
                                        v5 = append(v5, 2147483647)
                                        var v43 int32 = 1
                                        var v44 uint16 = 794
                                        _ = v42
                                        _ = v43
                                        _ = v44
                                    }
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                    fmt.Println(fi34)
                                    switch fi34 {
                                    case 0:
                                        switch genFn0(true) {
                                        case -184:
                                            switch 46 {
                                            case -879:
                                                switch fi34 {
                                                case 127:
                                                    var v45 uint64 = 18446744073709551615
                                                    _ = v45
                                                case 4294967295:
                                                    v5 = append(v5, 42)
                                                case 0:
                                                    v2 = float32(3.14159)
                                                default:
                                                    v3[32768] = 2147483647
                                                }
                                            default:
                                                var v46 uint32 = 277
                                                _ = v46
                                            }
                                        case -951:
                                            if false {
                                                v47 := GenS4{F0: 813, F1: 625}
                                                v2 = v2
                                                switch 157 {
                                                case 2147483647:
                                                    if true {
                                                        var v48 int = (fi34 * -988)
                                                        var v49 int8 = 0
                                                        _ = v48
                                                        _ = v49
                                                    }
                                                }
                                                v1 -= (v1 - v1)
                                                _ = v47
                                            } else {
                                                var v50 float32 = ((1.5 / v2) * -2.25)
                                                _ = v50
                                            }
                                        case -9223372036854775808:
                                            _ = len(v5)
                                        default:
                                            v2 = v2
                                        }
                                        switch fi24 {
                                        case 1:
                                            var v51 uint64 = 4294967295
                                            _ = v51
                                        case -612:
                                            v1 = -2.25
                                        }
                                        if (("" + "\n\t") == "a") {
                                            v1 += (float64(3.14159) / v1)
                                            switch "\n\t" {
                                            case "a":
                                                var v52 uint32 = 0
                                                _ = v52
                                            default:
                                                var v53 int8 = 1
                                                _ = v53
                                            }
                                            switch (fi24 ^ (fi24 % ((fi24) | 1))) {
                                            case 0:
                                                if (((((0.1 - 0.1) + -2.25) <= 1.5) && true) || ((65535 > 771) || ((((381 <= -1) || false) || (!true)) && (42 < 0)))) {
                                                    var v54 int32 = 764
                                                    _ = v54
                                                }
                                            }
                                        } else {
                                            delete(v3, 32767)
                                            var v55 int32 = -2147483648
                                            if (65535 == 2) {
                                                var v56 uint64 = 0
                                                _ = v56
                                            } else {
                                                fmt.Println(fi34, v1, v1)
                                            }
                                            switch fi34 {
                                            case 990:
                                                fmt.Println(v1, v2)
                                            }
                                            _ = v55
                                        }
                                        var v57 int = (-(fi24))
                                        _ = v57
                                    case 4294967295:
                                        fmt.Println(fi24, fi24, v2)
                                    default:
                                        if ("a" > "") {
                                            var v58 float64 = 0.1
                                            v58 -= v1
                                            var v59 uint8 = 55
                                            _ = v58
                                            _ = v59
                                        }
                                        switch genFn0(true) {
                                        case -2147483648:
                                            switch fi34 {
                                            case -970:
                                                switch (-(fi24)) {
                                                case 9223372036854775807:
                                                    var v60 int8 = 18
                                                    _ = v60
                                                }
                                            case 32768:
                                                var v61 uint32 = 1
                                                var v62 uint64 = 0
                                                var v63 float32 = -2.25
                                                v64 := GenS4{F0: 2147483647, F1: (v62 & (v62 << ((65535) & 63)))}
                                                v65 := GenS3{F0: 1, F1: v62}
                                                _ = v61
                                                _ = v62
                                                _ = v63
                                                _ = v64
                                                _ = v65
                                            case 2147483647:
                                                var v66 uint32 = 317
                                                _ = v66
                                            default:
                                                v2 += 0.1
                                            }
                                        case 0:
                                            v1 = 0.0
                                        default:
                                            v2 += v2
                                        }
                                        v1 = 1.5
                                        var v67 float32 = v2
                                        _ = v67
                                    }
                                    _ = v41
                                case 128:
                                    v1 -= (v1 - (3.14159 + (v1 + v1)))
                                }
                            }
                            v68 := GenS4{F0: 0, F1: 0}
                            _ = fi34
                            _ = v35
                            _ = v68
                        }
                    }
                }
                _ = v26
            }
            switch 9223372036854775807 {
            case 423:
                fmt.Println(v2)
            case 32768:
                delete(v4, genFn0((false || ((255 <= 2) || false))))
            default:
                for fi69 := 0; fi69 < 6; fi69++ {
                    v5 = append(v5, 562)
                    var v70 uint = 255
                    _ = fi69
                    _ = v70
                }
            }
            _ = fi24
        }
        switch "a" {
        case "":
            v71 := GenS4{F0: -505, F1: 18446744073709551615}
            _ = v71
        case "a":
            v72 := GenS6{F0: -128}
            _ = v72
        case "hello":
            v1 = (0.1 / v1)
        default:
            _ = len(v5)
        }
        fmt.Println(v1, v1)
    }
    v1 += float64((((v2 * (-2.25 + 1e10)) - float32(v1)) - v2))
    for fi73 := 0; fi73 < 1; fi73++ {
        fmt.Println(v1, v2)
        _ = fi73
    }
    v2 = v2
    switch 706 {
    case 920:
        var v74 int8 = -1
        _ = v74
    case 2147483647:
        for fi75 := 0; fi75 < 6; fi75++ {
            fmt.Println(v1, v2)
            if ((fi75 << ((65535) & 63)) >= (fi75 & fi75)) {
                for fi76 := 0; fi76 < 5; fi76++ {
                    switch fi76 {
                    case -2147483648:
                        var v77 uint64 = 0
                        _ = v77
                    case 22:
                        fmt.Println(v1, v2)
                    case 9223372036854775807:
                        if true {
                            v1 = float64(0.0)
                        }
                    }
                    v1 -= float64((((v2 * v2) + 1.5) + float32(0.0)))
                    _ = fi76
                }
                if ("a" >= ("hello" + "")) {
                    v78 := GenS1{F0: 32767, F1: true, F2: 0}
                    v1 -= v1
                    v1 += (3.14159 - v1)
                    _ = v78
                }
            } else {
                var v79 string = "a"
                switch v79 {
                case "hello":
                    v80 := GenS6{F0: -128}
                    _ = v80
                case "a":
                    var v81 int16 = -32768
                    _ = v81
                case "foo bar":
                    v1 = -2.25
                }
                switch -2147483648 {
                case 472:
                    if (((v2 - 0.0) + (0.1 / (v2 + v2))) > v2) {
                        v5 = append(v5, 4294967295)
                        var v82 uint64 = 242
                        _ = v82
                    } else {
                        if (fi75 <= fi75) {
                            for fi83 := 0; fi83 < 5; fi83++ {
                                v84 := GenS1{F0: -32768, F1: true, F2: 792}
                                var v85 uint32 = 40
                                _ = fi83
                                _ = v84
                                _ = v85
                            }
                            v86 := GenS7{F0: (1 != 18446744073709551615), F1: -1}
                            _ = v86
                        } else {
                            var v87 string = "\n\t"
                            _ = v87
                        }
                        var v88 int32 = 2147483647
                        _ = v88
                    }
                case 831:
                    var v89 uint16 = 976
                    _ = v89
                default:
                    for fi90 := 0; fi90 < 4; fi90++ {
                        var v91 float32 = ((float32(v1) + (((v2 - v2) - 0.0) * float32((((float64(float32((1e10 - ((v1 - 0.0) * 0.1)))) * float64(0.0)) * -2.25) * v1)))) + v2)
                        _ = fi90
                        _ = v91
                    }
                }
                _ = v79
            }
            _ = fi75
        }
    default:
        if false {
            v2 -= float32(1e10)
            _ = len(v5)
        } else {
            var v92 string = (("\n\t" + "hello") + ("" + ("a" + (("foo bar" + "") + ((("" + ("foo bar" + ("a" + "hello"))) + ("\n\t" + "hello")) + "a")))))
            _ = len(v3)
            _ = v92
        }
    }
    v93 := GenS8{F0: -1, F1: 965, F2: -32768}
    var v94 float32 = v2
    if false {
        fmt.Println(v2)
        v2 -= v94
    }
    var v95 uint64 = 166
    v96 := GenS5{F0: 127, F1: 65535, F2: 0}
    var v97 uint = 65535
    delete(v3, 547)
    v98 := GenS4{F0: 2147483647, F1: v95}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v94:", v94)
    fmt.Println("v95:", v95)
    fmt.Println("v97:", v97)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v93
    _ = v94
    _ = v95
    _ = v96
    _ = v97
    _ = v98
}
