package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32; F2 uint64 }
type GenS2 struct { F0 int16; F1 uint32 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int; F1 uint64; F2 int32 }
type GenS5 struct { F0 int8; F1 uint32 }
type GenS6 struct { F0 uint; F1 uint; F2 int64 }
type GenS7 struct { F0 string; F1 bool }
type GenS8 struct { F0 uint16; F1 int; F2 int8 }
func genFn0(p0 uint16) int8 {
    for fi6 := 0; fi6 < 6; fi6++ {
        p0 += p0
        var v7 int8 = 1
        switch fi6 {
        case -2147483648:
            if true {
                for fi8 := 0; fi8 < 1; fi8++ {
                    v7 = (^(v7))
                    if true {
                        v9 := GenS1{F0: fi6, F1: 69, F2: 1}
                        _ = v9
                    } else {
                        v10 := GenS1{F0: -39, F1: 744, F2: 1}
                        fmt.Println(v7)
                        switch fi6 {
                        case 4294967295:
                            var v11 int = fi6
                            _ = v11
                        case -714:
                            var v12 bool = (99 < 2)
                            _ = v12
                        }
                        _ = v10
                    }
                    var v13 float64 = 3.14159
                    var v14 uint32 = 24
                    switch ("" + ((("a" + "a") + "foo bar") + "")) {
                    case "":
                        if (((-13 < 9223372036854775807) || ((v7 - (v7 | v7)) < (v7 ^ 3))) && true) {
                            v13 += v13
                        } else {
                            p0 = 210
                        }
                    case "foo bar":
                        var v15 uint8 = 86
                        _ = v15
                    case "\n\t":
                        var v16 float32 = (1.5 + 1e10)
                        _ = v16
                    }
                    _ = fi8
                    _ = v13
                    _ = v14
                }
                for fi17 := 0; fi17 < 19; fi17++ {
                    p0 -= (p0 % ((p0) | 1))
                    _ = fi17
                }
            }
        case 42:
            v18 := []int{}
            _ = v18
        }
        _ = fi6
        _ = v7
    }
    v19 := GenS1{F0: 965, F1: 0, F2: 4294967295}
    p0 += (^(p0))
    _ = p0
    _ = v19
    return 100
}
func genFn1(p0 int64) int8 {
    var v20 float64 = -2.25
    switch -9223372036854775808 {
    case -9223372036854775808:
        if true {
            var v21 string = ("foo bar" + "")
            p0 = (p0 ^ -9223372036854775808)
            _ = v21
        }
    default:
        v22 := GenS2{F0: 1000, F1: 1}
        _ = v22
    }
    var v23 int32 = 486
    _ = p0
    _ = v20
    _ = v23
    return -128
}

func main() {
    v1 := make(map[uint32]string)
    var v2 float64 = 0.0
    var v3 int32 = 1
    var v4 bool = true
    var v5 bool = v4
    v24 := GenS1{F0: -882, F1: 692, F2: 436}
    v25 := []int{-9223372036854775808, -2147483648, -9223372036854775808, -1}
    for fi26 := 0; fi26 < 7; fi26++ {
        if false {
            for fi27 := 0; fi27 < 10; fi27++ {
                if (42 <= 251) {
                    if v4 {
                        if (0.1 <= -2.25) {
                            delete(v1, 0)
                            v2 -= ((((float64(0.0) * v2) / (v2 - ((-2.25 * 0.0) - ((float64((1.5 * -2.25)) + v2) + v2)))) - ((v2 / (3.14159 * v2)) / 1.0)) - v2)
                        }
                    }
                    v28 := GenS3{F0: (("foo bar" + "\n\t") + "")}
                    v2 = float64(0.1)
                    delete(v1, 45)
                    v29 := GenS4{F0: fi27, F1: 42, F2: v3}
                    if ((fi27 & (fi27 % ((fi26) | 1))) >= fi27) {
                        v25 = append(v25, (fi26 & (fi26 | (fi26 / (((fi27 << ((158) & 63))) | 1)))))
                    } else {
                        v4 = (1 < 536)
                    }
                    _ = v28
                    _ = v29
                }
                _ = fi27
            }
            v5 = ("\n\t" > "\n\t")
            v5 = v5
        }
        if v4 {
            if len(v25) >= 2 {
                v25 = v25[1:len(v25)]
            }
        }
        v2 = v2
        _ = fi26
    }
    v5 = (112 <= 0)
    fmt.Println(v2, v4, v4)
    for fi30 := 0; fi30 < 5; fi30++ {
        v25 = append(v25, fi30)
        _ = fi30
    }
    switch -1 {
    case 32768:
        for fi31 := 0; fi31 < 5; fi31++ {
            fmt.Println(fi31, v3)
            var v32 uint32 = 2147483647
            for fi33 := 0; fi33 < 15; fi33++ {
                v34 := v1[(^(v32))]
                _ = len(v25)
                switch "a" {
                case "hello":
                    v25 = append(v25, fi33)
                case "":
                    switch 0 {
                    case 9223372036854775807:
                        var v35 string = v34
                        _ = v35
                    case 4294967295:
                        switch 4294967295 {
                        case 628:
                            v1[2] = v34
                        case -1:
                            var v36 int16 = 32767
                            _ = v36
                        case 4294967295:
                            v32 = (v32 % (((v32 >> ((58) & 63))) | 1))
                        default:
                            var v37 int = 32768
                            _ = v37
                        }
                    case 42:
                        var v38 float64 = v2
                        _ = v38
                    }
                case "a":
                    if (-1 > -1) {
                        fmt.Println(v3, fi33)
                    }
                    _ = len(v25)
                    if v4 {
                        switch (^(fi31)) {
                        case 32767:
                            var v39 uint = uint(v32)
                            _ = v39
                        case 1:
                            switch v34 {
                            case "foo bar":
                                v40 := v1[(v32 % (((^(v32))) | 1))]
                                _ = v40
                            case "hello":
                                v41 := GenS2{F0: -1, F1: 921}
                                _ = v41
                            }
                        default:
                            var v42 bool = (!v4)
                            _ = v42
                        }
                        fmt.Println(v3, v34)
                        var v43 bool = v4
                        _ = v43
                    } else {
                        var v44 float32 = (3.14159 + float32(v2))
                        fmt.Println(v5)
                        _ = v44
                    }
                }
                _ = fi33
                _ = v34
            }
            _ = fi31
            _ = v32
        }
    default:
        switch "\n\t" {
        case "":
            fmt.Println(v2)
        case "hello":
            if v5 {
                if v5 {
                    switch -2147483648 {
                    case 679:
                        v2 = -2.25
                    case 128:
                        fmt.Println(v3, v4, v5)
                    case -920:
                        v4 = v4
                    }
                } else {
                    var v45 uint64 = 18446744073709551615
                    _ = v45
                }
                _ = len(v25)
                fmt.Println(v4)
            }
        case "\n\t":
            v2 = v2
        default:
            for fi46 := 0; fi46 < 14; fi46++ {
                var v47 uint = 42
                var v48 int8 = genFn1(-192)
                if false {
                    switch 2147483647 {
                    case -2147483648:
                        for fi49 := 0; fi49 < 16; fi49++ {
                            fmt.Println(v5, fi49, v48)
                            switch "a" {
                            case "hello":
                                var v50 bool = (((v48 >= int8(fi46)) || v5) && (572 <= 75))
                                _ = v50
                            case "\n\t":
                                var v51 int64 = 1
                                var v52 string = "\n\t"
                                if v5 {
                                    if v4 {
                                        var v53 float64 = 0.1
                                        fmt.Println(fi46, v47, v4)
                                        _ = v53
                                    }
                                    v51 = v51
                                    v54 := GenS3{F0: ("\n\t" + (v52 + (v52 + v52)))}
                                    _ = v54
                                } else {
                                    v2 -= v2
                                    v55 := GenS3{F0: v52}
                                    v51 += (v51 << (((v47 + 1)) & 63))
                                    _ = v55
                                }
                                if len(v25) > 0 {
                                    v25[0] = int(v47)
                                }
                                _ = v51
                                _ = v52
                            }
                            _ = fi49
                        }
                    case 42:
                        if len(v25) > 0 {
                            v25[0] = fi46
                        }
                    case 2147483647:
                        v47 = v47
                    }
                }
                _ = fi46
                _ = v47
                _ = v48
            }
        }
    }
    switch -1 {
    case 128:
        v4 = v5
    case -591:
        if len(v25) >= 2 {
            v25 = v25[1:len(v25)]
        }
    default:
        v56 := GenS2{F0: 1000, F1: 2}
        _ = v56
    }
    _ = len(v25)
    switch "a" {
    case "\n\t":
        for fi57 := 0; fi57 < 9; fi57++ {
            if len(v25) > 0 {
                v25[0] = fi57
            }
            var v58 uint8 = 127
            if v4 {
                fmt.Println(v3, v4, fi57)
                for fi59 := 0; fi59 < 18; fi59++ {
                    if (!(!v5)) {
                        v60 := GenS3{F0: ""}
                        _ = len(v25)
                        var v61 float32 = 3.14159
                        _ = v60
                        _ = v61
                    } else {
                        var v62 uint32 = 2147483648
                        _ = v62
                    }
                    v2 += ((((-2.25 / 1.0) - 1e10) - -2.25) + 3.14159)
                    fmt.Println(v5, v3)
                    var v63 uint64 = 0
                    _ = fi59
                    _ = v63
                }
                if v5 {
                    if v5 {
                        v4 = (uint32(v58) > 2147483648)
                    } else {
                        for fi64 := 0; fi64 < 19; fi64++ {
                            if true {
                                if ((!(88 <= 0)) || ("a" != ("\n\t" + "a"))) {
                                    v4 = (v3 < 290)
                                    switch ("\n\t" + (("foo bar" + "foo bar") + "\n\t")) {
                                    case "foo bar":
                                        v65 := GenS5{F0: -86, F1: 2147483648}
                                        if (908 < 575) {
                                            if true {
                                                delete(v1, 1)
                                                var v66 int = (fi64 | fi57)
                                                _ = len(v25)
                                                if v5 {
                                                    fmt.Println(fi57, v3, v4)
                                                    v25 = append(v25, 2147483647)
                                                } else {
                                                    v3 -= v3
                                                }
                                                var v67 string = ("hello" + "hello")
                                                _ = v66
                                                _ = v67
                                            } else {
                                                v68 := GenS6{F0: 2, F1: 18446744073709551615, F2: -1}
                                                switch ((("\n\t" + "a") + "hello") + ("\n\t" + "")) {
                                                case "a":
                                                    v69 := GenS5{F0: -128, F1: 42}
                                                    _ = v69
                                                case "foo bar":
                                                    v70 := v1[4294967295]
                                                    _ = v70
                                                }
                                                _ = v68
                                            }
                                        } else {
                                            var v71 float32 = 1.5
                                            v3 += (v3 * (v3 & -1))
                                            var v72 int = 32767
                                            _ = v71
                                            _ = v72
                                        }
                                        v3 = -1
                                        fmt.Println(fi64)
                                        switch fi57 {
                                        case -636:
                                            v73 := GenS2{F0: 1000, F1: 1}
                                            _ = v73
                                        case -9223372036854775808:
                                            v74 := GenS4{F0: -29, F1: 4294967295, F2: v3}
                                            _ = v74
                                        case -440:
                                            var v75 int = fi64
                                            _ = v75
                                        }
                                        _ = v65
                                    case "hello":
                                        v76 := GenS7{F0: "a", F1: (v4 && v5)}
                                        v3 += (v3 + -2147483648)
                                        v4 = (!v5)
                                        _ = v76
                                    case "a":
                                        var v77 uint8 = v58
                                        _ = v77
                                    default:
                                        _ = len(v25)
                                        v25 = append(v25, int(v3))
                                        v1[562] = ""
                                        if len(v25) > 0 {
                                            v25[0] = 940
                                        }
                                        v78 := GenS2{F0: 183, F1: 42}
                                        _ = v78
                                    }
                                    v79 := GenS3{F0: "hello"}
                                    _ = v79
                                }
                                v2 = 0.0
                            } else {
                                v80 := GenS7{F0: "a", F1: v4}
                                var v81 int64 = -692
                                v5 = ((((0.1 / 3.14159) < 0.1) && (false && true)) || (!true))
                                _ = v80
                                _ = v81
                            }
                            if false {
                                var v82 uint8 = (v58 - 127)
                                _ = v82
                            } else {
                                v83 := v1[1]
                                if ((v83 == "hello") || v5) {
                                    v5 = v5
                                    fmt.Println(v5, v5)
                                } else {
                                    if len(v25) > 0 {
                                        v25[0] = (fi64 & fi57)
                                    }
                                    fmt.Println(fi57, v2)
                                    v4 = true
                                }
                                _ = v83
                            }
                            _ = fi64
                        }
                        v25 = append(v25, (fi57 * 32768))
                    }
                } else {
                    var v84 int = -2147483648
                    _ = v84
                }
            }
            _ = fi57
            _ = v58
        }
    case "hello":
        for fi85 := 0; fi85 < 19; fi85++ {
            switch (fi85 * (fi85 / ((fi85) | 1))) {
            case -9223372036854775808:
                delete(v1, 349)
            case 0:
                v2 += v2
            case 4294967295:
                v25 = append(v25, (-(fi85)))
            default:
                var v86 float32 = ((1e10 - (0.0 + 1.5)) * -2.25)
                _ = v86
            }
            v2 = ((v2 + float64(float32((float64(1e10) * ((1e10 - v2) / (((float64(1.5) - (0.1 * v2)) / float64(1.5)) / (((((v2 - v2) - ((float64(1e10) / v2) / v2)) + (-2.25 + float64(1e10))) / v2) * 0.1))))))) / v2)
            if true {
                for fi87 := 0; fi87 < 3; fi87++ {
                    switch fi85 {
                    case 261:
                        var v88 string = ("" + "a")
                        _ = v88
                    case 4294967295:
                        var v89 int64 = -404
                        _ = v89
                    case -1:
                        fmt.Println(v3)
                    default:
                        if v5 {
                            switch ("hello" + "") {
                            case "\n\t":
                                v90 := GenS8{F0: 255, F1: -9223372036854775808, F2: -128}
                                _ = v90
                            }
                        }
                    }
                    _ = fi87
                }
                v91 := GenS8{F0: 0, F1: -9223372036854775808, F2: -128}
                _ = v91
            } else {
                switch ("hello" + "") {
                case "":
                    v3 = (v3 / (((v3 << ((65535) & 63))) | 1))
                case "hello":
                    v25 = append(v25, -974)
                case "\n\t":
                    for fi92 := 0; fi92 < 7; fi92++ {
                        v3 -= v3
                        var v93 int64 = -459
                        if true {
                            fmt.Println(fi85)
                            fmt.Println(v2)
                        } else {
                            switch fi92 {
                            case -2147483648:
                                var v94 bool = (0.1 == float64(0.0))
                                _ = v94
                            case -882:
                                fmt.Println(fi92, fi92)
                            }
                        }
                        _ = fi92
                        _ = v93
                    }
                }
                var v95 float32 = -2.25
                v95 -= v95
                _ = v95
            }
            _ = fi85
        }
    case "a":
        var v96 int8 = genFn1(-9223372036854775808)
        _ = v96
    }
    v5 = false
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
