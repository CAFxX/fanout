package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int; F1 string; F2 uint8 }
type GenS3 struct { F0 uint64; F1 int64; F2 uint8 }
type GenS4 struct { F0 int; F1 bool; F2 int }
type GenS5 struct { F0 uint16; F1 uint; F2 uint }
type GenS6 struct { F0 int8; F1 int8 }
type GenS7 struct { F0 uint64; F1 int8; F2 int16 }
type GenS8 struct { F0 bool }
func genFn0(p0 bool, p1 bool, p2 uint16) uint {
    for fi7 := 0; fi7 < 7; fi7++ {
        var v8 uint64 = 0
        _ = fi7
        _ = v8
    }
    var v9 uint16 = p2
    v10 := []int{-2147483648, 695}
    if p0 {
        switch ((("hello" + "foo bar") + "foo bar") + ("foo bar" + "")) {
        case "":
            switch 4294967295 {
            case -139:
                fmt.Println(p1, v9)
            case -18:
                switch 26 {
                case 0:
                    for fi11 := 0; fi11 < 1; fi11++ {
                        v12 := GenS1{F0: -128}
                        fmt.Println(fi11, p0, p0)
                        _ = fi11
                        _ = v12
                    }
                case 9223372036854775807:
                    if (302 == 4294967295) {
                        v13 := GenS1{F0: -128}
                        _ = v13
                    } else {
                        for fi14 := 0; fi14 < 17; fi14++ {
                            p1 = (!(42 != 275))
                            if (!false) {
                                var v15 uint = 18446744073709551615
                                for fi16 := 0; fi16 < 6; fi16++ {
                                    var v17 int = 42
                                    _ = fi16
                                    _ = v17
                                }
                                for fi18 := 0; fi18 < 4; fi18++ {
                                    fmt.Println(p2, fi18)
                                    v19 := GenS2{F0: (^(fi18)), F1: "a", F2: uint8(v9)}
                                    _ = fi18
                                    _ = v19
                                }
                                _ = v15
                            }
                            p1 = (42 <= -9223372036854775808)
                            _ = fi14
                        }
                        for fi20 := 0; fi20 < 18; fi20++ {
                            _ = len(v10)
                            if len(v10) > 0 {
                                v10[0] = 127
                            }
                            _ = fi20
                        }
                    }
                case -297:
                    var v21 bool = p1
                    _ = v21
                }
            default:
                _ = len(v10)
            }
        case "a":
            if (-1 != 42) {
                for fi22 := 0; fi22 < 4; fi22++ {
                    fmt.Println(p1, fi22, p0)
                    for fi23 := 0; fi23 < 16; fi23++ {
                        p2 = 881
                        v9 -= v9
                        v24 := GenS2{F0: (fi22 | fi22), F1: "", F2: 2}
                        p1 = (!p0)
                        _ = fi23
                        _ = v24
                    }
                    switch fi22 {
                    case 4294967295:
                        var v25 int64 = 42
                        _ = v25
                    }
                    _ = fi22
                }
            }
        default:
            for fi26 := 0; fi26 < 11; fi26++ {
                if (!(p0 || ((p1 || (!true)) && p1))) {
                    var v27 float32 = -2.25
                    _ = v27
                } else {
                    var v28 uint8 = 42
                    if ((-938 < -384) || p1) {
                        var v29 int64 = 453
                        _ = v29
                    } else {
                        switch 0 {
                        case 42:
                            if p0 {
                                v30 := GenS2{F0: fi26, F1: (("a" + "foo bar") + "hello"), F2: v28}
                                _ = v30
                            } else {
                                p2 -= (v9 << ((0) & 63))
                                for fi31 := 0; fi31 < 4; fi31++ {
                                    fmt.Println(fi26)
                                    _ = fi31
                                }
                            }
                        }
                        _ = len(v10)
                    }
                    for fi32 := 0; fi32 < 17; fi32++ {
                        switch -606 {
                        case -703:
                            p0 = ("\n\t" != "\n\t")
                        default:
                            v33 := GenS3{F0: 2, F1: 9223372036854775807, F2: (v28 & (v28 * v28))}
                            _ = v33
                        }
                        if len(v10) >= 2 {
                            v10 = v10[1:len(v10)]
                        }
                        v28 -= v28
                        var v34 int32 = 42
                        var v35 bool = (true || p0)
                        v36 := GenS2{F0: fi26, F1: ("a" + ""), F2: 2}
                        var v37 int = (fi26 & fi26)
                        _ = fi32
                        _ = v34
                        _ = v35
                        _ = v36
                        _ = v37
                    }
                    _ = v28
                }
                switch (fi26 ^ fi26) {
                case 32767:
                    p1 = true
                default:
                    for fi38 := 0; fi38 < 14; fi38++ {
                        var v39 int = (-(fi38))
                        var v40 string = "a"
                        _ = fi38
                        _ = v39
                        _ = v40
                    }
                }
                _ = fi26
            }
        }
    } else {
        p1 = (2 >= 893)
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v9
    return 42
}

func main() {
    var v1 int32 = 1
    var v2 uint = 2
    v3 := make(map[string]string)
    v4 := make(map[uint32]int)
    var v5 int8 = -1
    var v6 int = -2147483648
    fmt.Println(v6, v5)
    var v41 uint8 = 2
    v6 += v6
    var v42 string = ""
    switch v6 {
    case 4294967295:
        v42 = (v42 + (("\n\t" + ((v42 + (v42 + v42)) + v42)) + v42))
    case 32768:
        v6 += v6
    default:
        var v43 string = (v42 + "foo bar")
        _ = v43
    }
    v42 = ("a" + ((("\n\t" + v42) + v42) + v42))
    if true {
        v41 = v41
        switch (^(v6)) {
        case 9223372036854775807:
            if (v5 <= -1) {
                v44 := GenS2{F0: (v6 | (v6 >> ((v2) & 63))), F1: ("" + v42), F2: 1}
                _ = v44
            }
        case 128:
            if false {
                fmt.Println(v5, v41, v41)
                v45 := []int{(v6 - v6)}
                for fi46 := 0; fi46 < 8; fi46++ {
                    if len(v45) > 0 {
                        v45[0] = 32768
                    }
                    _ = fi46
                }
                _ = v45
            } else {
                fmt.Println(v5)
            }
        }
    }
    v6 += (v6 % (((v6 ^ int(v5))) | 1))
    delete(v3, (v42 + "a"))
    if (v5 <= v5) {
        if (!(false && (!(3.14159 <= -2.25)))) {
            var v47 int64 = 806
            _ = v47
        }
        switch "\n\t" {
        case "":
            var v48 int16 = -139
            _ = v48
        case "hello":
            for fi49 := 0; fi49 < 5; fi49++ {
                for fi50 := 0; fi50 < 18; fi50++ {
                    v51 := GenS2{F0: -544, F1: v42, F2: (v41 & (v41 | (v41 >> ((genFn0(false, false, 241)) & 63))))}
                    var v52 int8 = (v5 | 1)
                    if false {
                        switch "foo bar" {
                        case "\n\t":
                            if true {
                                v6 += fi50
                            }
                        default:
                            var v53 int32 = (-(v1))
                            delete(v3, ("" + v42))
                            v41 += (^(v41))
                            var v54 string = v42
                            _ = v53
                            _ = v54
                        }
                        v55 := GenS1{F0: v52}
                        _ = v55
                    } else {
                        switch v42 {
                        case "":
                            v56 := []int{(fi49 + fi50), fi50, fi49}
                            _ = v56
                        case "\n\t":
                            v57 := GenS4{F0: (-(fi49)), F1: (!(42 <= -9223372036854775808)), F2: (fi50 ^ 42)}
                            _ = v57
                        case "hello":
                            v1 = 2147483647
                        }
                        v5 = v5
                    }
                    _ = fi50
                    _ = v51
                    _ = v52
                }
                var v58 float32 = 1.5
                var v59 int8 = (v5 ^ v5)
                _ = fi49
                _ = v58
                _ = v59
            }
        case "foo bar":
            v60 := GenS2{F0: (v6 | 127), F1: "", F2: v41}
            _ = v60
        }
        fmt.Println(v6, v5, v1)
    } else {
        var v61 uint8 = (v41 ^ uint8(v2))
        v41 = uint8(v2)
        var v62 uint16 = 141
        _ = v61
        _ = v62
    }
    for fi63 := 0; fi63 < 15; fi63++ {
        v64 := GenS3{F0: 2, F1: -1, F2: v41}
        _ = fi63
        _ = v64
    }
    v65 := GenS5{F0: 65535, F1: (v2 * (v2 >> ((42) & 63))), F2: (v2 ^ v2)}
    switch (v6 >> ((genFn0(false, true, 32767)) & 63)) {
    case 9223372036854775807:
        v66 := GenS6{F0: (v5 + v5), F1: v5}
        _ = v66
    }
    v6 += (v6 & v6)
    var v67 int32 = v1
    v68 := []int{v6, (^(v6)), v6, v6}
    if true {
        v1 -= v67
        v68 = append(v68, (v6 >> ((v2) & 63)))
        if (!true) {
            for fi69 := 0; fi69 < 10; fi69++ {
                v2 = 165
                _ = fi69
            }
        } else {
            v5 += int8(v41)
            if (487 == 2) {
                if (209 < 144) {
                    var v70 int32 = (v1 / (((v1 % (((v67 + (v1 & v67))) | 1))) | 1))
                    var v71 uint64 = 662
                    v41 += v41
                    _ = v70
                    _ = v71
                } else {
                    var v72 int32 = v67
                    if len(v68) > 0 {
                        v68[0] = v6
                    }
                    _ = v72
                }
                var v73 int64 = -9223372036854775808
                _ = v73
            } else {
                v68 = append(v68, (v6 * 4294967295))
                var v74 uint64 = 1
                v75 := GenS2{F0: (v6 * v6), F1: (v42 + "hello"), F2: v41}
                _ = v74
                _ = v75
            }
        }
    } else {
        var v76 int8 = 1
        for fi77 := 0; fi77 < 19; fi77++ {
            if false {
                if (228 != 47) {
                    for fi78 := 0; fi78 < 1; fi78++ {
                        if false {
                            switch fi77 {
                            case 32768:
                                v5 = (v5 + (v76 - v5))
                            case 2147483647:
                                if len(v68) > 0 {
                                    v68[0] = (^(fi77))
                                }
                            default:
                                if true {
                                    var v79 int = v6
                                    v76 = (v5 % (((v76 % ((v5) | 1))) | 1))
                                    v80 := GenS1{F0: (v5 | (v76 - (v76 << ((255) & 63))))}
                                    _ = v79
                                    _ = v80
                                } else {
                                    var v81 string = v42
                                    var v82 uint32 = 2147483647
                                    if (!((-(v6)) == int(v76))) {
                                        v1 = int32(v41)
                                        v83 := GenS3{F0: 4294967295, F1: 0, F2: 127}
                                        if len(v68) > 0 {
                                            v68[0] = 1
                                        }
                                        _ = v83
                                    }
                                    _ = v81
                                    _ = v82
                                }
                            }
                            if true {
                                if (v1 > (v1 / ((2147483647) | 1))) {
                                    var v84 float32 = -2.25
                                    _ = len(v3)
                                    _ = v84
                                }
                                v2 -= 18446744073709551615
                                v85 := GenS4{F0: (fi78 / (((-(fi77))) | 1)), F1: false, F2: fi78}
                                _ = v85
                            } else {
                                v86 := GenS7{F0: 981, F1: (v5 + 0), F2: -1}
                                v41 = v41
                                _ = v86
                            }
                            v42 = v42
                            var v87 uint8 = 128
                            v88 := GenS6{F0: (v5 & 100), F1: -128}
                            _ = v87
                            _ = v88
                        }
                        _ = fi78
                    }
                    if (0 == 4294967295) {
                        if false {
                            var v89 uint8 = v41
                            var v90 uint = v2
                            _ = v89
                            _ = v90
                        } else {
                            v5 = (v5 >> ((0) & 63))
                            fmt.Println(v5, v5, v42)
                            var v91 int64 = -1
                            _ = v91
                        }
                        if len(v68) > 0 {
                            v68[0] = 0
                        }
                    } else {
                        for fi92 := 0; fi92 < 3; fi92++ {
                            v6 -= fi92
                            _ = fi92
                        }
                    }
                    var v93 float64 = 1e10
                    _ = v93
                } else {
                    if (false && (v42 > "hello")) {
                        v41 += v41
                    }
                }
                v68 = append(v68, (v6 / (((fi77 | v6)) | 1)))
                v2 = v2
            } else {
                v94 := GenS2{F0: (fi77 - (fi77 * v6)), F1: "a", F2: (v41 % (((^(v41))) | 1))}
                var v95 int16 = -32768
                v2 += genFn0((fi77 <= (fi77 ^ 2147483647)), false, 464)
                _ = v94
                _ = v95
            }
            _ = fi77
        }
        for fi96 := 0; fi96 < 10; fi96++ {
            var v97 uint64 = 42
            if (750 > (v97 & (v97 - (v97 | v97)))) {
                fmt.Println(v1, v67, v5)
            } else {
                if true {
                    fmt.Println(v5)
                    fmt.Println(v42, v2, v42)
                    var v98 uint16 = 2
                    _ = v98
                } else {
                    v76 = 127
                    v99 := GenS8{F0: (v67 == (v67 >> (((v2 - (^(v2)))) & 63)))}
                    _ = v99
                }
                v4[1] = (v6 << ((v2) & 63))
                switch "hello" {
                case "hello":
                    v2 = v2
                case "a":
                    var v100 int64 = int64(v2)
                    _ = v100
                default:
                    var v101 uint32 = 1
                    _ = v101
                }
            }
            fmt.Println(v76, v2)
            _ = fi96
            _ = v97
        }
        _ = v76
    }
    var v102 uint8 = (^(v41))
    fmt.Println("v1:", v1)
    fmt.Println("v102:", v102)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v67:", v67)
    fmt.Println("slicelen v68:", len(v68), v68)
    _ = v1
    _ = v102
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v5
    _ = v6
    _ = v65
    _ = v67
    _ = v68
}
