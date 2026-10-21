package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint16; F1 uint16 }
type GenS3 struct { F0 uint; F1 int; F2 int16 }
type GenS4 struct { F0 string; F1 uint8; F2 uint64 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 uint64; F1 uint8; F2 uint }
type GenS7 struct { F0 int8; F1 int64; F2 uint32 }
type GenS8 struct { F0 uint; F1 string }
type GenS9 struct { F0 string; F1 int8 }
type GenS10 struct { F0 uint; F1 string }
func genFn0(p0 string, p1 uint16, p2 bool) int8 {
    for fi4 := 0; fi4 < 8; fi4++ {
        var v5 uint64 = 18446744073709551615
        var v6 int8 = 100
        p1 = (p1 + (p1 - 0))
        _ = fi4
        _ = v5
        _ = v6
    }
    for fi7 := 0; fi7 < 14; fi7++ {
        for fi8 := 0; fi8 < 1; fi8++ {
            p2 = (-14 == 90)
            _ = fi8
        }
        for fi9 := 0; fi9 < 2; fi9++ {
            var v10 uint32 = 39
            var v11 float32 = (0.1 * -2.25)
            if (2 == v10) {
                if false {
                    if p2 {
                        switch p0 {
                        case "":
                            var v12 int8 = 127
                            _ = v12
                        case "hello":
                            var v13 int32 = 0
                            _ = v13
                        default:
                            var v14 uint = 1
                            _ = v14
                        }
                        p2 = false
                    }
                    v15 := []int{}
                    _ = v15
                } else {
                    v16 := GenS1{F0: 4294967295}
                    fmt.Println(p2)
                    _ = v16
                }
                if p2 {
                    p2 = (false && p2)
                }
            } else {
                p1 -= p1
                if false {
                    v17 := GenS1{F0: 0}
                    if (128 >= (^(fi9))) {
                        v18 := []int{fi9, 1, (fi7 & fi9)}
                        var v19 string = "a"
                        var v20 uint32 = v10
                        _ = v18
                        _ = v19
                        _ = v20
                    } else {
                        var v21 uint = 140
                        _ = v21
                    }
                    var v22 int = (fi7 * (fi9 - (^(fi7))))
                    _ = v17
                    _ = v22
                }
                v23 := GenS1{F0: 18446744073709551615}
                _ = v23
            }
            v24 := GenS1{F0: 635}
            _ = fi9
            _ = v10
            _ = v11
            _ = v24
        }
        var v25 float64 = 3.14159
        _ = fi7
        _ = v25
    }
    fmt.Println(p2, p2, p2)
    var v26 int64 = -1
    _ = p0
    _ = p1
    _ = p2
    _ = v26
    return -10
}
func genFn1(p0 uint) bool {
    p0 = p0
    var v27 uint = p0
    _ = p0
    _ = v27
    return false
}

func main() {
    var v1 uint64 = 883
    var v2 float64 = -2.25
    var v3 uint16 = 32767
    v3 = v3
    switch 4294967295 {
    case -506:
        v3 -= 2
    case 597:
        v28 := GenS1{F0: 1}
        _ = v28
    case -894:
        if (2 != 42) {
            var v29 uint8 = 128
            v30 := GenS1{F0: 18446744073709551615}
            _ = v29
            _ = v30
        } else {
            v1 = v1
            if (genFn0("hello", v3, true) >= 127) {
                var v31 string = "\n\t"
                v32 := []int{32767, 32768, 2147483647, int(v3)}
                v32 = append(v32, 127)
                _ = v31
                _ = v32
            } else {
                var v33 bool = true
                v33 = v33
                if false {
                    v34 := GenS1{F0: v1}
                    v35 := GenS2{F0: v3, F1: v3}
                    switch "\n\t" {
                    case "a":
                        v36 := GenS2{F0: v3, F1: (^(v3))}
                        _ = v36
                    case "\n\t":
                        v33 = (42 <= 127)
                    default:
                        fmt.Println(v3, v33, v2)
                    }
                    _ = v34
                    _ = v35
                }
                _ = v33
            }
        }
    default:
        var v37 uint16 = v3
        _ = v37
    }
    var v38 int64 = 164
    var v39 int32 = 2147483647
    if (!(true && (2 != (v1 ^ v1)))) {
        v40 := GenS3{F0: 42, F1: 1, F2: -1}
        for fi41 := 0; fi41 < 2; fi41++ {
            for fi42 := 0; fi42 < 3; fi42++ {
                v43 := GenS4{F0: "hello", F1: 42, F2: (v1 & (v1 >> ((18446744073709551615) & 63)))}
                _ = fi42
                _ = v43
            }
            switch fi41 {
            case 127:
                if (1 <= -1) {
                    if false {
                        v1 += v1
                        v1 -= (v1 - 42)
                        var v44 uint8 = 128
                        _ = v44
                    } else {
                        if true {
                            for fi45 := 0; fi45 < 15; fi45++ {
                                var v46 uint32 = 0
                                _ = fi45
                                _ = v46
                            }
                            fmt.Println(v2)
                            var v47 string = "foo bar"
                            _ = v47
                        } else {
                            v38 = v38
                        }
                    }
                    fmt.Println(v3, v3, v38)
                }
            }
            var v48 int16 = 42
            _ = fi41
            _ = v48
        }
        _ = v40
    }
    v49 := []int{131, 9223372036854775807, 32767, 42}
    if ((!((!false) && true)) || ((0 <= 548) || (!true))) {
        if (!true) {
            if false {
                switch -2147483648 {
                case 1:
                    var v50 uint64 = v1
                    _ = v50
                case 32767:
                    fmt.Println(v39)
                }
            }
            for fi51 := 0; fi51 < 10; fi51++ {
                _ = len(v49)
                fmt.Println(v1, v38)
                fmt.Println(v39, v39, v38)
                _ = fi51
            }
        }
        v52 := GenS3{F0: 18, F1: 402, F2: 1}
        if false {
            if true {
                v53 := GenS1{F0: (v1 | v1)}
                if (float32(0.1) > 1.5) {
                    if false {
                        if false {
                            var v54 int = 407
                            if (("a" + "foo bar") > "") {
                                switch (-(v54)) {
                                case 2147483647:
                                    v3 -= (v3 & 255)
                                default:
                                    var v55 float64 = (-2.25 / (1e10 / 1.0))
                                    _ = v55
                                }
                            } else {
                                v54 = -526
                                fmt.Println(v2, v38)
                                fmt.Println(v39)
                            }
                            _ = v54
                        }
                    } else {
                        var v56 int8 = int8(v38)
                        if true {
                            v57 := GenS2{F0: v3, F1: 42}
                            if (v38 != (v38 & (v38 - (v38 >> ((61) & 63))))) {
                                v49 = append(v49, 813)
                                if true {
                                    v56 -= (v56 % (((v56 / ((genFn0(((("\n\t" + "foo bar") + "a") + ""), v3, (true || false))) | 1))) | 1))
                                }
                            } else {
                                var v58 int64 = (v38 & v38)
                                _ = v58
                            }
                            _ = v57
                        } else {
                            var v59 uint8 = 179
                            if (100 >= -1) {
                                if (!((("hello" + ("\n\t" + ((("\n\t" + "a") + "a") + ("foo bar" + ("foo bar" + "\n\t"))))) + "foo bar") >= "")) {
                                    var v60 uint = 463
                                    for fi61 := 0; fi61 < 10; fi61++ {
                                        v1 = uint64(v60)
                                        fmt.Println(v38)
                                        _ = fi61
                                    }
                                    _ = v60
                                } else {
                                    for fi62 := 0; fi62 < 11; fi62++ {
                                        fmt.Println(v59, fi62, v56)
                                        if true {
                                            v49 = append(v49, 128)
                                            for fi63 := 0; fi63 < 13; fi63++ {
                                                fmt.Println(v56, v39, v2)
                                                _ = fi63
                                            }
                                            for fi64 := 0; fi64 < 12; fi64++ {
                                                _ = len(v49)
                                                _ = fi64
                                            }
                                        }
                                        for fi65 := 0; fi65 < 1; fi65++ {
                                            var v66 uint8 = v59
                                            fmt.Println(v1, v38, fi65)
                                            _ = fi65
                                            _ = v66
                                        }
                                        _ = fi62
                                    }
                                    v67 := GenS5{F0: (("hello" + "a") + ((("a" + "\n\t") + "foo bar") + ""))}
                                    v2 = v2
                                    _ = v67
                                }
                                if ((true || true) || false) {
                                    if len(v49) > 0 {
                                        v49[0] = 32767
                                    }
                                }
                            } else {
                                v38 = v38
                            }
                            _ = v59
                        }
                        switch "foo bar" {
                        case "foo bar":
                            switch ("\n\t" + "hello") {
                            case "\n\t":
                                v2 += 1.5
                            case "hello":
                                v1 = (v1 & (v1 ^ 18446744073709551615))
                            }
                        case "a":
                            var v68 int8 = (^(v56))
                            _ = v68
                        }
                        _ = v56
                    }
                    v38 -= (^(v38))
                    if len(v49) >= 2 {
                        v49 = v49[1:len(v49)]
                    }
                } else {
                    v38 += (v38 & (^(v38)))
                    v38 -= v38
                }
                _ = v53
            } else {
                if (32767 != 32767) {
                    var v69 float64 = ((v2 / v2) + v2)
                    _ = v69
                } else {
                    if len(v49) > 0 {
                        v49[0] = -1
                    }
                    switch -1 {
                    case 700:
                        _ = len(v49)
                    default:
                        var v70 bool = (v2 >= (v2 / v2))
                        _ = v70
                    }
                    var v71 bool = (2 < 645)
                    _ = v71
                }
            }
            v1 += v1
            var v72 int32 = 42
            _ = v72
        }
        _ = v52
    }
    v39 = 2147483647
    switch -2147483648 {
    case 0:
        var v73 int8 = -128
        _ = v73
    case 32768:
        v39 = (v39 % (((-(v39))) | 1))
    case 32767:
        fmt.Println(v1, v3, v1)
    default:
        v74 := GenS6{F0: (^(v1)), F1: 0, F2: 927}
        _ = v74
    }
    switch ("\n\t" + "") {
    case "\n\t":
        if ((v1 * v1) <= (v1 ^ v1)) {
            switch "" {
            case "a":
                switch (("\n\t" + "a") + ("" + ("foo bar" + "foo bar"))) {
                case "hello":
                    v39 -= (v39 / (((-(v39))) | 1))
                case "":
                    fmt.Println(v2)
                default:
                    switch 0 {
                    case -165:
                        if len(v49) >= 2 {
                            v49 = v49[1:len(v49)]
                        }
                    case 906:
                        if (false || true) {
                            if false {
                                switch 32768 {
                                case 487:
                                    _ = len(v49)
                                case 0:
                                    for fi75 := 0; fi75 < 10; fi75++ {
                                        var v76 int64 = -13
                                        switch fi75 {
                                        case 835:
                                            switch -706 {
                                            case 127:
                                                v77 := GenS7{F0: genFn0(("\n\t" + ("\n\t" + "hello")), (v3 / (((v3 | (v3 - (v3 - 42)))) | 1)), true), F1: -9223372036854775808, F2: 65535}
                                                _ = v77
                                            case 128:
                                                for fi78 := 0; fi78 < 16; fi78++ {
                                                    var v79 float64 = v2
                                                    _ = fi78
                                                    _ = v79
                                                }
                                            default:
                                                if (int16(v76) != -736) {
                                                    var v80 uint8 = 2
                                                    v39 -= v39
                                                    _ = v80
                                                } else {
                                                    for fi81 := 0; fi81 < 1; fi81++ {
                                                        if true {
                                                            fmt.Println(v2)
                                                        }
                                                        switch (fi81 ^ int(v1)) {
                                                        case -2147483648:
                                                            v39 = (-(v39))
                                                        case 32767:
                                                            var v82 float32 = 1e10
                                                            _ = v82
                                                        case -513:
                                                            if ((32768 >= fi81) || false) {
                                                                var v83 float32 = 1e10
                                                                switch (^(fi75)) {
                                                                case -9223372036854775808:
                                                                    fmt.Println(fi81, v39, fi81)
                                                                default:
                                                                    var v84 int64 = (v38 % ((1) | 1))
                                                                    _ = v84
                                                                }
                                                                v85 := GenS8{F0: 65535, F1: "foo bar"}
                                                                v49 = append(v49, (fi81 / (((fi75 / (((fi75 % ((fi75) | 1))) | 1))) | 1)))
                                                                _ = v83
                                                                _ = v85
                                                            } else {
                                                                switch fi81 {
                                                                case 132:
                                                                    switch 32767 {
                                                                    case 0:
                                                                        fmt.Println(v38)
                                                                    case 4294967295:
                                                                        fmt.Println(v2, v2, v2)
                                                                    default:
                                                                        var v86 uint8 = 171
                                                                        _ = v86
                                                                    }
                                                                case 4294967295:
                                                                    v49 = append(v49, fi75)
                                                                default:
                                                                    v39 = v39
                                                                }
                                                                var v87 int8 = -111
                                                                _ = v87
                                                            }
                                                        default:
                                                            switch (("\n\t" + "hello") + "a") {
                                                            case "":
                                                                v2 += v2
                                                            case "\n\t":
                                                                fmt.Println(v2, fi75, v39)
                                                            case "hello":
                                                                var v88 int64 = (v76 & (v38 | 1))
                                                                _ = v88
                                                            default:
                                                                if ("" == "hello") {
                                                                    var v89 int8 = -128
                                                                    v38 += (v76 >> ((255) & 63))
                                                                    fmt.Println(v76, v2, fi75)
                                                                    if (v38 == (v38 * 42)) {
                                                                        v89 = v89
                                                                        v90 := GenS2{F0: uint16(v1), F1: 2}
                                                                        _ = v90
                                                                    } else {
                                                                        v39 = v39
                                                                        v89 = (v89 % ((v89) | 1))
                                                                        fmt.Println(fi75, v3)
                                                                    }
                                                                    var v91 int32 = (v39 * (v39 << ((861) & 63)))
                                                                    _ = v89
                                                                    _ = v91
                                                                } else {
                                                                    v49 = append(v49, fi75)
                                                                }
                                                            }
                                                        }
                                                        _ = fi81
                                                    }
                                                    if true {
                                                        var v92 int8 = genFn0("a", v3, (v39 > (^(v39))))
                                                        if ((v76 ^ 868) != (v76 ^ int64(v39))) {
                                                            var v93 string = "a"
                                                            _ = v93
                                                        }
                                                        _ = v92
                                                    } else {
                                                        fmt.Println(v38, v76)
                                                        if len(v49) > 0 {
                                                            v49[0] = fi75
                                                        }
                                                        v1 += (v1 - v1)
                                                    }
                                                    var v94 int = 32767
                                                    _ = v94
                                                }
                                            }
                                        case -1:
                                            var v95 uint32 = 0
                                            _ = v95
                                        case 1:
                                            var v96 int = (fi75 * (fi75 + 9223372036854775807))
                                            _ = v96
                                        }
                                        v97 := GenS9{F0: ("a" + "foo bar"), F1: -37}
                                        _ = fi75
                                        _ = v76
                                        _ = v97
                                    }
                                }
                            } else {
                                var v98 float64 = v2
                                _ = v98
                            }
                        } else {
                            var v99 int = int(v39)
                            if false {
                                v49 = append(v49, v99)
                                switch -192 {
                                case -483:
                                    switch "" {
                                    case "foo bar":
                                        var v100 float64 = (((((v2 * 0.1) * (v2 / 1.0)) / v2) * v2) - (3.14159 - 0.0))
                                        _ = v100
                                    case "":
                                        v38 -= 9223372036854775807
                                    default:
                                        var v101 float64 = v2
                                        _ = v101
                                    }
                                case -9223372036854775808:
                                    var v102 int64 = 42
                                    _ = v102
                                }
                                var v103 float64 = 0.1
                                _ = v103
                            }
                            _ = v99
                        }
                    default:
                        for fi104 := 0; fi104 < 2; fi104++ {
                            var v105 int32 = (^(v39))
                            v39 = v39
                            _ = fi104
                            _ = v105
                        }
                    }
                }
            case "\n\t":
                for fi106 := 0; fi106 < 20; fi106++ {
                    v49 = append(v49, fi106)
                    _ = fi106
                }
            }
        } else {
            if true {
                v2 += v2
            } else {
                var v107 uint64 = 454
                v108 := GenS10{F0: 2, F1: "foo bar"}
                v109 := GenS1{F0: (v107 / ((v107) | 1))}
                _ = v107
                _ = v108
                _ = v109
            }
        }
    }
    fmt.Println(v2)
    v110 := GenS5{F0: "a"}
    v49 = append(v49, 128)
    var v111 int64 = v38
    var v112 float32 = (0.1 - 0.1)
    fmt.Println("v1:", v1)
    fmt.Println("v111:", v111)
    fmt.Println("v112:", v112)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v49:", len(v49), v49)
    _ = v1
    _ = v110
    _ = v111
    _ = v112
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v49
}
