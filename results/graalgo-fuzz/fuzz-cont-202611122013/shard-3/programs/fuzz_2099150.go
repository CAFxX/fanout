package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8; F2 int }
type GenS2 struct { F0 string; F1 int64; F2 string }
type GenS3 struct { F0 int; F1 int32 }
type GenS4 struct { F0 int32; F1 uint8; F2 uint64 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 uint8; F1 uint32 }
type GenS9 struct { F0 uint32; F1 int32 }
type GenS10 struct { F0 int16 }
type GenS11 struct { F0 int32; F1 uint16; F2 int8 }
func genFn0() int8 {
    var v5 uint32 = 0
    if (v5 == v5) {
        if false {
            var v6 int = 960
            _ = v6
        } else {
            v7 := []int{347, -801}
            for fi8 := 0; fi8 < 18; fi8++ {
                var v9 string = "hello"
                _ = fi8
                _ = v9
            }
            if false {
                if (127 <= 90) {
                    var v10 int8 = -32
                    v11 := GenS1{F0: 86, F1: 1, F2: 423}
                    _ = v10
                    _ = v11
                } else {
                    var v12 bool = ((!((false || false) || (-1 <= 32767))) || (!(0 == -82)))
                    if len(v7) > 0 {
                        v7[0] = 13
                    }
                    _ = v12
                }
                var v13 uint = 154
                _ = v13
            }
            _ = v7
        }
        if (!false) {
            switch "hello" {
            case "foo bar":
                switch ("" + "") {
                case "foo bar":
                    v14 := GenS1{F0: 605, F1: 255, F2: 913}
                    _ = v14
                case "a":
                    var v15 int32 = -1
                    _ = v15
                case "\n\t":
                    for fi16 := 0; fi16 < 18; fi16++ {
                        if (3.14159 == (3.14159 + -2.25)) {
                            v17 := []int{(-(fi16))}
                            _ = v17
                        } else {
                            if false {
                                var v18 uint64 = 4294967295
                                _ = v18
                            } else {
                                switch fi16 {
                                case -7:
                                    v19 := []int{(fi16 * fi16), 127, (fi16 << ((42) & 63)), 780}
                                    _ = v19
                                }
                                var v20 uint32 = v5
                                _ = v20
                            }
                            for fi21 := 0; fi21 < 2; fi21++ {
                                if (((((-2.25 != float32((0.0 / 1.0))) && (!((false && true) && true))) && (((false && false) || (!false)) || false)) || false) && (0 <= 128)) {
                                    if (false && (4294967295 >= 65535)) {
                                        v22 := []int{fi16, 1, (fi21 % (((fi21 + (fi21 & fi16))) | 1)), int(v5)}
                                        _ = v22
                                    }
                                    v5 = (v5 * 884)
                                } else {
                                    v23 := []int{}
                                    v23 = append(v23, (fi21 * (fi21 >> ((470) & 63))))
                                    _ = v23
                                }
                                _ = fi21
                            }
                            v5 -= 4294967295
                        }
                        _ = fi16
                    }
                }
            default:
                v24 := GenS2{F0: "\n\t", F1: 42, F2: ("hello" + "a")}
                _ = v24
            }
        } else {
            var v25 bool = (0.1 <= 1.5)
            fmt.Println(v25, v5)
            _ = v25
        }
        fmt.Println(v5)
    }
    var v26 int = 127
    for fi27 := 0; fi27 < 3; fi27++ {
        v5 = v5
        v28 := GenS2{F0: ("a" + ""), F1: 42, F2: "\n\t"}
        _ = fi27
        _ = v28
    }
    v29 := GenS2{F0: "hello", F1: 806, F2: ""}
    _ = v26
    _ = v29
    _ = v5
    return -83
}
func genFn1(p0 int64) uint {
    v30 := GenS1{F0: 42, F1: 42, F2: 834}
    if false {
        fmt.Println(p0)
        var v31 string = ("hello" + (("foo bar" + "hello") + ("a" + "")))
        v32 := []int{}
        _ = v31
        _ = v32
    }
    _ = p0
    _ = v30
    return 884
}

func main() {
    var v1 int8 = -128
    var v2 uint8 = 149
    var v3 int16 = 1000
    var v4 int8 = (^(v1))
    v33 := GenS1{F0: genFn1(-712), F1: (v2 ^ 42), F2: int(v1)}
    v34 := []int{-2147483648, 127, -427, 32767}
    v1 = genFn0()
    switch "hello" {
    case "hello":
        if (!(1.5 <= (0.0 + 1e10))) {
            for fi35 := 0; fi35 < 12; fi35++ {
                v36 := GenS3{F0: (fi35 >> ((genFn1(-9223372036854775808)) & 63)), F1: -1}
                _ = fi35
                _ = v36
            }
            v1 -= (^(v4))
            if ((v3 ^ (v3 - v3)) == (v3 - (-(v3)))) {
                var v37 int16 = v3
                _ = v37
            } else {
                v38 := GenS2{F0: ((("hello" + "") + "") + ((("\n\t" + ("" + "")) + "\n\t") + "foo bar")), F1: 0, F2: "foo bar"}
                var v39 int64 = -305
                _ = v38
                _ = v39
            }
        }
    case "":
        v40 := GenS3{F0: 32767, F1: -143}
        _ = v40
    default:
        var v41 uint8 = v2
        _ = v41
    }
    switch "a" {
    case "hello":
        if (42 >= 1) {
            for fi42 := 0; fi42 < 12; fi42++ {
                v34 = append(v34, fi42)
                _ = fi42
            }
            for fi43 := 0; fi43 < 4; fi43++ {
                fmt.Println(v3, v2, v2)
                var v44 float64 = (1.5 * (1e10 - 0.1))
                switch (-(fi43)) {
                case 42:
                    v45 := GenS4{F0: 2147483647, F1: 127, F2: 1}
                    _ = v45
                default:
                    v1 = genFn0()
                }
                _ = fi43
                _ = v44
            }
            if false {
                if ("a" <= (("hello" + "\n\t") + "foo bar")) {
                    var v46 float32 = 1.5
                    _ = v46
                }
                var v47 float64 = float64(1e10)
                switch "foo bar" {
                case "\n\t":
                    if ((v2 | v2) <= v2) {
                        fmt.Println(v4, v4)
                        if len(v34) >= 2 {
                            v34 = v34[1:len(v34)]
                        }
                    }
                case "a":
                    v48 := GenS5{F0: v2}
                    _ = v48
                }
                _ = v47
            }
        } else {
            v4 -= (^(v1))
            if len(v34) >= 2 {
                v34 = v34[1:len(v34)]
            }
        }
    case "":
        var v49 float64 = (3.14159 * 1.5)
        _ = v49
    case "\n\t":
        var v50 int = 2147483647
        _ = v50
    default:
        v2 = v2
    }
    fmt.Println(v3)
    for fi51 := 0; fi51 < 10; fi51++ {
        fmt.Println(v2, fi51, v2)
        for fi52 := 0; fi52 < 12; fi52++ {
            if len(v34) >= 2 {
                v34 = v34[1:len(v34)]
            }
            _ = fi52
        }
        _ = fi51
    }
    if false {
        if true {
            var v53 uint = 964
            _ = len(v34)
            if len(v34) > 0 {
                v34[0] = -667
            }
            _ = v53
        }
    }
    var v54 uint16 = 65535
    switch (("hello" + "") + "hello") {
    case "\n\t":
        if len(v34) >= 2 {
            v34 = v34[1:len(v34)]
        }
    default:
        var v55 bool = true
        _ = v55
    }
    switch "foo bar" {
    case "\n\t":
        v56 := GenS6{F0: 1}
        _ = v56
    case "hello":
        if true {
            for fi57 := 0; fi57 < 8; fi57++ {
                fmt.Println(v54, v3, v3)
                v54 += v54
                _ = fi57
            }
            v58 := GenS3{F0: -221, F1: 42}
            _ = v58
        } else {
            for fi59 := 0; fi59 < 15; fi59++ {
                var v60 uint = genFn1(-9223372036854775808)
                v61 := GenS2{F0: ("" + "\n\t"), F1: -9223372036854775808, F2: "hello"}
                _ = fi59
                _ = v60
                _ = v61
            }
            if ((42 == 202) && false) {
                var v62 int32 = 632
                v54 = v54
                switch 1 {
                case 42:
                    switch (("hello" + "foo bar") + "hello") {
                    case "a":
                        if true {
                            if false {
                                switch (("hello" + "a") + (("" + "foo bar") + "a")) {
                                case "":
                                    v54 += v54
                                default:
                                    v63 := GenS4{F0: (-(v62)), F1: 137, F2: 316}
                                    _ = v63
                                }
                                var v64 int32 = v62
                                fmt.Println(v1, v64)
                                _ = v64
                            }
                            v2 = (v2 | v2)
                        }
                    case "":
                        var v65 int8 = v1
                        _ = v65
                    case "foo bar":
                        var v66 int64 = 42
                        _ = v66
                    default:
                        v67 := GenS2{F0: "a", F1: 466, F2: ("hello" + "\n\t")}
                        _ = v67
                    }
                }
                _ = v62
            } else {
                var v68 int64 = 1
                for fi69 := 0; fi69 < 20; fi69++ {
                    v34 = append(v34, int(v3))
                    _ = fi69
                }
                var v70 bool = true
                _ = v68
                _ = v70
            }
        }
    case "a":
        fmt.Println(v3)
    }
    if len(v34) >= 2 {
        v34 = v34[1:len(v34)]
    }
    for fi71 := 0; fi71 < 13; fi71++ {
        fmt.Println(v2, fi71)
        if true {
            switch (fi71 * fi71) {
            case -956:
                if (265 == 18446744073709551615) {
                    if true {
                        switch fi71 {
                        case 337:
                            v4 = v1
                        default:
                            if (v2 == (^(v2))) {
                                switch "foo bar" {
                                case "":
                                    switch fi71 {
                                    case 4294967295:
                                        fmt.Println(v4, v1)
                                    case -617:
                                        var v72 uint64 = 4294967295
                                        _ = v72
                                    case -185:
                                        v4 = -15
                                    default:
                                        v73 := GenS7{F0: (v1 | v4)}
                                        _ = v73
                                    }
                                case "a":
                                    var v74 int16 = v3
                                    _ = v74
                                }
                                if false {
                                    v75 := GenS5{F0: v2}
                                    for fi76 := 0; fi76 < 14; fi76++ {
                                        if false {
                                            v77 := GenS3{F0: fi71, F1: 525}
                                            _ = v77
                                        }
                                        v1 = (^(v1))
                                        _ = fi76
                                    }
                                    _ = v75
                                } else {
                                    var v78 int = (fi71 >> ((genFn1(9223372036854775807)) & 63))
                                    _ = v78
                                }
                                var v79 int8 = (v1 << ((521) & 63))
                                _ = v79
                            } else {
                                var v80 string = "foo bar"
                                v81 := GenS3{F0: fi71, F1: int32(v3)}
                                var v82 int32 = -420
                                _ = v80
                                _ = v81
                                _ = v82
                            }
                        }
                        v54 = v54
                    }
                    v83 := GenS2{F0: (("a" + ("foo bar" + "\n\t")) + ("foo bar" + "a")), F1: 42, F2: ("a" + (((((((("\n\t" + ("" + ("foo bar" + (("a" + "foo bar") + "a")))) + "hello") + (("foo bar" + "\n\t") + "a")) + "foo bar") + "hello") + (((("a" + ("hello" + "a")) + "") + ((("\n\t" + "") + ("" + ("foo bar" + ""))) + ("a" + ("\n\t" + ("\n\t" + (("foo bar" + "\n\t") + "foo bar")))))) + "hello")) + ("\n\t" + (("" + "foo bar") + ("hello" + (("a" + "foo bar") + "\n\t"))))) + "a"))}
                    _ = v83
                } else {
                    v34 = append(v34, (-(fi71)))
                    v84 := GenS2{F0: ("foo bar" + ""), F1: 620, F2: "a"}
                    v85 := GenS7{F0: v4}
                    _ = v84
                    _ = v85
                }
            case -777:
                for fi86 := 0; fi86 < 20; fi86++ {
                    v87 := GenS6{F0: -885}
                    if (true || true) {
                        var v88 int64 = 485
                        v54 = v54
                        switch 0 {
                        case 4294967295:
                            if len(v34) > 0 {
                                v34[0] = 2147483647
                            }
                        case -672:
                            v88 -= (^(v88))
                            var v89 int = fi71
                            var v90 uint8 = v2
                            _ = v89
                            _ = v90
                        default:
                            v34 = append(v34, fi71)
                        }
                        _ = v88
                    }
                    _ = fi86
                    _ = v87
                }
            case 4294967295:
                v91 := GenS8{F0: (^(v2)), F1: 2}
                _ = v91
            }
            v92 := GenS1{F0: 2, F1: v2, F2: fi71}
            _ = v92
        }
        fmt.Println(v2, v1)
        _ = fi71
    }
    if true {
        v3 -= (^(v3))
    }
    var v93 float32 = float32(0.1)
    var v94 float32 = 0.0
    if false {
        for fi95 := 0; fi95 < 20; fi95++ {
            v34 = append(v34, int(v4))
            v96 := GenS9{F0: 507, F1: -231}
            _ = fi95
            _ = v96
        }
        v97 := GenS6{F0: -625}
        _ = v97
    } else {
        switch -9223372036854775808 {
        case -2147483648:
            if false {
                if false {
                    var v98 int32 = 2147483647
                    v99 := GenS10{F0: v3}
                    _ = v98
                    _ = v99
                }
                v100 := GenS9{F0: uint32(v54), F1: int32(v2)}
                _ = v100
            }
        case 32768:
            v93 = v94
        default:
            if (v4 == genFn0()) {
                switch ("hello" + "\n\t") {
                case "\n\t":
                    v101 := GenS8{F0: v2, F1: uint32(v2)}
                    _ = v101
                case "a":
                    for fi102 := 0; fi102 < 3; fi102++ {
                        var v103 uint = 0
                        _ = fi102
                        _ = v103
                    }
                case "hello":
                    for fi104 := 0; fi104 < 15; fi104++ {
                        var v105 float32 = 3.14159
                        _ = fi104
                        _ = v105
                    }
                default:
                    fmt.Println(v54, v2)
                }
                switch -950 {
                case 32768:
                    for fi106 := 0; fi106 < 5; fi106++ {
                        v1 = (-(v4))
                        for fi107 := 0; fi107 < 10; fi107++ {
                            v54 -= (v54 >> ((255) & 63))
                            _ = fi107
                        }
                        _ = fi106
                    }
                }
                v108 := GenS11{F0: 750, F1: (v54 % ((uint16(v2)) | 1)), F2: int8(v3)}
                _ = v108
            } else {
                v1 = (v1 & (v1 % (((v4 >> ((genFn1(9223372036854775807)) & 63))) | 1)))
            }
        }
        if false {
            var v109 uint32 = 2
            _ = v109
        } else {
            _ = len(v34)
        }
    }
    v1 += 100
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("v4:", v4)
    fmt.Println("v54:", v54)
    fmt.Println("v93:", v93)
    fmt.Println("v94:", v94)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v54
    _ = v93
    _ = v94
}
