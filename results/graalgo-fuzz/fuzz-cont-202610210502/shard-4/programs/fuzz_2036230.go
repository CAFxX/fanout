package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int16; F1 int32; F2 uint16 }
type GenS4 struct { F0 uint; F1 uint32; F2 int8 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 uint16; F1 uint8; F2 uint64 }
type GenS7 struct { F0 int64; F1 int16 }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 int8 }
type GenS10 struct { F0 string; F1 int8; F2 int16 }
type GenS11 struct { F0 int }
type GenS12 struct { F0 bool; F1 uint32 }
type GenS13 struct { F0 uint16; F1 uint; F2 uint32 }
type GenS14 struct { F0 int16; F1 int8 }
type GenS15 struct { F0 int8; F1 uint16; F2 string }
type GenS16 struct { F0 int32; F1 uint64; F2 int8 }
type GenS17 struct { F0 int64; F1 uint }
type GenS18 struct { F0 int; F1 bool; F2 int16 }
type GenS19 struct { F0 uint16; F1 int }
func genFn0() int32 {
    for fi4 := 0; fi4 < 14; fi4++ {
        fmt.Println(fi4)
        _ = fi4
    }
    if true {
        v5 := []int{127, -2147483648, -956, -27}
        _ = v5
    }
    switch -9223372036854775808 {
    case 50:
        v6 := []int{551}
        _ = v6
    }
    var v7 float32 = 3.14159
    _ = v7
    return -2147483648
}
func genFn1(p0 string, p1 uint8, p2 uint8) int64 {
    p2 = p2
    if true {
        if (65535 > 1) {
            if true {
                switch 32768 {
                case 128:
                    var v8 string = p0
                    _ = v8
                case -60:
                    for fi9 := 0; fi9 < 2; fi9++ {
                        fmt.Println(p1)
                        for fi10 := 0; fi10 < 5; fi10++ {
                            v11 := []int{761, (fi9 & 2147483647)}
                            _ = fi10
                            _ = v11
                        }
                        _ = fi9
                    }
                }
                for fi12 := 0; fi12 < 17; fi12++ {
                    for fi13 := 0; fi13 < 18; fi13++ {
                        var v14 string = "foo bar"
                        _ = fi13
                        _ = v14
                    }
                    p0 = (p0 + p0)
                    switch p0 {
                    case "foo bar":
                        p1 += p1
                    case "":
                        p2 = p2
                    case "a":
                        if true {
                            if false {
                                for fi15 := 0; fi15 < 4; fi15++ {
                                    v16 := []int{-459}
                                    v16 = append(v16, fi15)
                                    if false {
                                        v17 := GenS1{F0: (377 > -348)}
                                        _ = v17
                                    }
                                    _ = fi15
                                    _ = v16
                                }
                            } else {
                                if (128 < (fi12 % ((2147483647) | 1))) {
                                    if (0 > 42) {
                                        var v18 string = p0
                                        _ = v18
                                    } else {
                                        var v19 uint = 0
                                        var v20 int16 = 42
                                        v21 := []int{}
                                        _ = v19
                                        _ = v20
                                        _ = v21
                                    }
                                } else {
                                    var v22 int64 = 0
                                    p2 = (^(p1))
                                    _ = v22
                                }
                                for fi23 := 0; fi23 < 19; fi23++ {
                                    v24 := GenS1{F0: (((1 < 1) && true) && false)}
                                    _ = fi23
                                    _ = v24
                                }
                            }
                            v25 := []int{fi12, (fi12 * (fi12 << ((455) & 63))), -9223372036854775808, (fi12 | 32767)}
                            switch fi12 {
                            case 32767:
                                for fi26 := 0; fi26 < 19; fi26++ {
                                    var v27 int = fi26
                                    var v28 float64 = 3.14159
                                    _ = fi26
                                    _ = v27
                                    _ = v28
                                }
                            case -1:
                                p1 = p2
                            }
                            _ = v25
                        } else {
                            for fi29 := 0; fi29 < 16; fi29++ {
                                fmt.Println(p1, fi12)
                                var v30 uint32 = 669
                                _ = fi29
                                _ = v30
                            }
                            if false {
                                v31 := GenS2{F0: 444}
                                _ = v31
                            }
                            fmt.Println(p2, p0, p2)
                        }
                    }
                    _ = fi12
                }
                p0 = (("a" + p0) + (("a" + "\n\t") + (("foo bar" + p0) + p0)))
            } else {
                if true {
                    switch -719 {
                    case 121:
                        v32 := GenS3{F0: -938, F1: 2147483647, F2: 0}
                        _ = v32
                    default:
                        if false {
                            p0 = p0
                            v33 := []int{}
                            if false {
                                v34 := GenS1{F0: false}
                                _ = v34
                            }
                            _ = v33
                        } else {
                            v35 := GenS4{F0: 883, F1: 42, F2: 127}
                            if false {
                                fmt.Println(p0, p1, p1)
                                if (true || (-1 <= 1)) {
                                    v36 := GenS5{F0: 65535}
                                    p1 = (p2 - p1)
                                    v37 := GenS6{F0: 32767, F1: 255, F2: 42}
                                    _ = v36
                                    _ = v37
                                }
                            }
                            v38 := GenS3{F0: 1000, F1: 1, F2: 1}
                            _ = v35
                            _ = v38
                        }
                    }
                    var v39 uint = 18446744073709551615
                    for fi40 := 0; fi40 < 10; fi40++ {
                        if true {
                            p2 -= (p2 >> ((v39) & 63))
                        }
                        if false {
                            if ((42 != 291) && true) {
                                p1 -= (p1 & (p2 << ((v39) & 63)))
                                for fi41 := 0; fi41 < 17; fi41++ {
                                    if false {
                                        var v42 uint64 = 613
                                        _ = v42
                                    } else {
                                        switch (fi41 << ((v39) & 63)) {
                                        case 2147483647:
                                            switch (p0 + "foo bar") {
                                            case "hello":
                                                v43 := GenS7{F0: -18, F1: -427}
                                                _ = v43
                                            }
                                            fmt.Println(p0, v39, p2)
                                        case -2147483648:
                                            v44 := []int{fi41}
                                            _ = v44
                                        case 32767:
                                            v39 = v39
                                        default:
                                            var v45 uint16 = 93
                                            _ = v45
                                        }
                                        p2 -= p2
                                        v46 := GenS8{F0: p1}
                                        fmt.Println(p1, p2, v39)
                                        _ = v46
                                    }
                                    var v47 int = (fi40 - (fi41 + fi41))
                                    fmt.Println(fi40, fi40)
                                    _ = fi41
                                    _ = v47
                                }
                                for fi48 := 0; fi48 < 7; fi48++ {
                                    p1 += 1
                                    switch (^(fi48)) {
                                    case 32768:
                                        var v49 uint32 = 42
                                        _ = v49
                                    case 907:
                                        v50 := []int{fi40}
                                        _ = v50
                                    case 113:
                                        var v51 uint64 = 0
                                        _ = v51
                                    default:
                                        v52 := GenS8{F0: (^(p2))}
                                        _ = v52
                                    }
                                    _ = fi48
                                }
                            } else {
                                v39 -= v39
                                switch fi40 {
                                case 128:
                                    var v53 uint16 = 32767
                                    _ = v53
                                case 740:
                                    var v54 float32 = -2.25
                                    _ = v54
                                case 127:
                                    switch (("\n\t" + p0) + p0) {
                                    case "a":
                                        p1 += (p2 % (((^(p2))) | 1))
                                    default:
                                        v55 := []int{fi40, (fi40 ^ (fi40 + fi40))}
                                        _ = v55
                                    }
                                }
                            }
                        } else {
                            fmt.Println(v39, p0)
                            for fi56 := 0; fi56 < 11; fi56++ {
                                v57 := GenS2{F0: 662}
                                v58 := []int{fi56, (fi56 % ((0) | 1))}
                                v59 := GenS8{F0: p1}
                                _ = fi56
                                _ = v57
                                _ = v58
                                _ = v59
                            }
                            v60 := []int{}
                            _ = v60
                        }
                        if true {
                            var v61 float32 = float32(3.14159)
                            for fi62 := 0; fi62 < 17; fi62++ {
                                fmt.Println(p2)
                                p2 += p1
                                p0 = p0
                                _ = fi62
                            }
                            _ = v61
                        }
                        _ = fi40
                    }
                    _ = v39
                } else {
                    if true {
                        var v63 int8 = 100
                        p2 = p2
                        _ = v63
                    }
                    v64 := GenS4{F0: 42, F1: 2, F2: 42}
                    var v65 float64 = float64((3.14159 + (float32(((-2.25 / 1.0) - 0.1)) - (1e10 - 3.14159))))
                    _ = v64
                    _ = v65
                }
                v66 := GenS9{F0: 1}
                _ = v66
            }
        }
        v67 := GenS8{F0: p1}
        v68 := GenS9{F0: 127}
        _ = v67
        _ = v68
    }
    v69 := GenS10{F0: p0, F1: 0, F2: 1000}
    p1 = 128
    _ = p0
    _ = p1
    _ = p2
    _ = v69
    return -122
}

func main() {
    v1 := []int{-1, 9223372036854775807, 127, -1}
    var v2 bool = (((float64((1e10 + float32(1.5))) / -2.25) / 0.1) > -2.25)
    v3 := []string{}
    v2 = (!((((!(!v2)) && v2) || (genFn1("\n\t", 108, 127) >= 1)) && v2))
    _ = len(v3)
    _ = len(v1)
    if len(v1) > 0 {
        v1[0] = -816
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v2 = v2
    var v70 uint64 = 4294967295
    if false {
        v71 := GenS6{F0: 42, F1: 2, F2: (v70 * (v70 % ((v70) | 1)))}
        var v72 uint32 = 2147483647
        _ = v71
        _ = v72
    } else {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        switch "hello" {
        case "\n\t":
            if false {
                v73 := GenS11{F0: 127}
                _ = v73
            }
        case "a":
            var v74 bool = (255 == 563)
            _ = v74
        case "foo bar":
            var v75 string = "a"
            _ = v75
        default:
            var v76 int = -187
            _ = v76
        }
    }
    for fi77 := 0; fi77 < 19; fi77++ {
        if (0 != 672) {
            if v2 {
                v1 = append(v1, 10)
                v78 := GenS8{F0: 127}
                if v2 {
                    var v79 int32 = 1
                    if true {
                        for fi80 := 0; fi80 < 4; fi80++ {
                            if v2 {
                                if len(v1) > 0 {
                                    v1[0] = fi77
                                }
                                switch fi80 {
                                case -46:
                                    var v81 int8 = 42
                                    _ = v81
                                case 127:
                                    var v82 uint8 = 255
                                    _ = v82
                                case 0:
                                    if (128 <= 127) {
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        var v83 int8 = -128
                                        var v84 bool = v2
                                        _ = v83
                                        _ = v84
                                    }
                                }
                                if true {
                                    v70 += 1
                                    fmt.Println(v79)
                                }
                            } else {
                                var v85 int8 = -128
                                _ = v85
                            }
                            fmt.Println(v2, v79)
                            v79 = (^(v79))
                            if v2 {
                                fmt.Println(v2)
                                v70 -= (v70 * v70)
                                fmt.Println(fi80, v70, v79)
                            }
                            _ = fi80
                        }
                        if ((728 == 117) && v2) {
                            v2 = v2
                            switch (fi77 % ((32768) | 1)) {
                            case -1:
                                v70 -= 468
                            case -9223372036854775808:
                                var v86 uint32 = 2
                                _ = v86
                            }
                        } else {
                            switch ("" + (("a" + (("foo bar" + "foo bar") + "hello")) + "")) {
                            case "":
                                for fi87 := 0; fi87 < 2; fi87++ {
                                    v70 += (^(v70))
                                    _ = fi87
                                }
                            default:
                                if v2 {
                                    v88 := GenS12{F0: v2, F1: 4294967295}
                                    if len(v3) > 0 {
                                        v3[0] = ("\n\t" + "\n\t")
                                    }
                                    v89 := GenS13{F0: 65535, F1: 569, F2: 719}
                                    _ = v88
                                    _ = v89
                                } else {
                                    if len(v1) > 0 {
                                        v1[0] = (fi77 & (fi77 * (fi77 / (((fi77 * fi77)) | 1))))
                                    }
                                    var v90 uint32 = 294
                                    _ = v90
                                }
                            }
                            v79 = (v79 | genFn0())
                            v91 := GenS14{F0: 32767, F1: 1}
                            _ = v91
                        }
                    } else {
                        if false {
                            if len(v3) > 0 {
                                v3[0] = "a"
                            }
                            var v92 uint64 = v70
                            v93 := GenS9{F0: 127}
                            _ = v92
                            _ = v93
                        }
                    }
                    if (42 <= -32768) {
                        switch (("\n\t" + "") + "") {
                        case "":
                            for fi94 := 0; fi94 < 19; fi94++ {
                                v70 = v70
                                var v95 int = -184
                                _ = fi94
                                _ = v95
                            }
                        case "a":
                            v70 -= v70
                        }
                        v96 := GenS10{F0: ("\n\t" + "hello"), F1: -15, F2: 0}
                        _ = v96
                    }
                    _ = v79
                }
                _ = v78
            }
            if true {
                v97 := GenS15{F0: int8(v70), F1: 850, F2: ("" + "hello")}
                var v98 uint32 = 42
                _ = v97
                _ = v98
            } else {
                v99 := GenS16{F0: 42, F1: (v70 >> ((65535) & 63)), F2: 127}
                if (v2 || (int32(v70) > int32(v70))) {
                    v100 := GenS17{F0: genFn1("hello", 0, 0), F1: 255}
                    v101 := GenS4{F0: uint(v70), F1: 34, F2: -1}
                    _ = v100
                    _ = v101
                } else {
                    v102 := GenS18{F0: fi77, F1: false, F2: -32768}
                    switch "\n\t" {
                    case "foo bar":
                        var v103 uint64 = (v70 ^ 0)
                        _ = v103
                    default:
                        var v104 uint32 = 503
                        _ = v104
                    }
                    switch ("" + "") {
                    case "foo bar":
                        if (((v70 ^ (v70 >> ((945) & 63))) != v70) && true) {
                            var v105 int32 = 0
                            _ = v105
                        }
                    case "":
                        fmt.Println(v2, v2, v70)
                    default:
                        if len(v3) > 0 {
                            v3[0] = "hello"
                        }
                    }
                    _ = v102
                }
                v70 += v70
                _ = v99
            }
        } else {
            fmt.Println(v2, v70)
            switch ("foo bar" + (("a" + "") + "")) {
            case "\n\t":
                var v106 int16 = 32767
                _ = v106
            case "a":
                fmt.Println(fi77, v2, fi77)
            }
        }
        v107 := GenS17{F0: 782, F1: 450}
        var v108 uint8 = 1
        _ = fi77
        _ = v107
        _ = v108
    }
    var v109 bool = ((("hello" + "hello") == "foo bar") || ((32768 > 32768) || ((((3.14159 + ((-2.25 / -2.25) - ((1.5 - 1.5) - (float32(3.14159) + (float32(1e10) - 0.0))))) - ((1.5 + 1e10) + ((3.14159 / (1.5 / (3.14159 / 1e10))) - 1e10))) / -2.25) <= 3.14159)))
    v110 := GenS19{F0: 0, F1: 32768}
    var v111 int32 = -1
    var v112 uint32 = 261
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v109:", v109)
    fmt.Println("v111:", v111)
    fmt.Println("v112:", v112)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v70:", v70)
    _ = v1
    _ = v109
    _ = v110
    _ = v111
    _ = v112
    _ = v2
    _ = v3
    _ = v70
}
