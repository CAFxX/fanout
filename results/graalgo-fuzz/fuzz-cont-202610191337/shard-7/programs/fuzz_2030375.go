package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64; F2 int }
type GenS2 struct { F0 uint; F1 uint; F2 uint16 }
type GenS3 struct { F0 uint16; F1 int16; F2 uint32 }
type GenS4 struct { F0 uint32; F1 int }
type GenS5 struct { F0 bool; F1 int32; F2 uint32 }
type GenS6 struct { F0 bool; F1 int8 }
type GenS7 struct { F0 uint16; F1 uint16; F2 int32 }
func genFn0(p0 uint8, p1 int, p2 uint8) uint64 {
    for fi5 := 0; fi5 < 19; fi5++ {
        fmt.Println(p0)
        for fi6 := 0; fi6 < 18; fi6++ {
            p0 += 128
            v7 := []int{}
            _ = fi6
            _ = v7
        }
        _ = fi5
    }
    p2 = (p2 - (^(p2)))
    for fi8 := 0; fi8 < 16; fi8++ {
        switch "a" {
        case "foo bar":
            p2 += (p2 * 128)
        case "hello":
            for fi9 := 0; fi9 < 13; fi9++ {
                p0 += p0
                _ = fi9
            }
        case "\n\t":
            if false {
                switch "\n\t" {
                case "a":
                    var v10 int16 = 1
                    _ = v10
                case "":
                    var v11 int8 = 1
                    _ = v11
                case "foo bar":
                    if false {
                        for fi12 := 0; fi12 < 5; fi12++ {
                            if (true && false) {
                                v13 := []int{(fi12 | fi12)}
                                fmt.Println(p1, p0, fi12)
                                fmt.Println(p1, p0, fi8)
                                _ = v13
                            } else {
                                switch fi12 {
                                case 32767:
                                    var v14 int = (fi12 % ((32767) | 1))
                                    _ = v14
                                case 116:
                                    if (!false) {
                                        fmt.Println(fi12)
                                        if (1.5 == 0.1) {
                                            v15 := GenS1{F0: -119, F1: 0, F2: p1}
                                            switch (fi12 - fi8) {
                                            case 9223372036854775807:
                                                switch 9223372036854775807 {
                                                case -9223372036854775808:
                                                    var v16 int16 = int16(fi8)
                                                    _ = v16
                                                case 42:
                                                    var v17 uint8 = p2
                                                    _ = v17
                                                }
                                            default:
                                                p2 -= (p2 >> ((2) & 63))
                                            }
                                            v18 := []int{2147483647, p1}
                                            _ = v15
                                            _ = v18
                                        }
                                        switch "foo bar" {
                                        case "a":
                                            var v19 int8 = -1
                                            _ = v19
                                        case "":
                                            var v20 int16 = -32768
                                            _ = v20
                                        }
                                    } else {
                                        p2 += (p0 % ((71) | 1))
                                        v21 := GenS2{F0: 18446744073709551615, F1: 242, F2: 896}
                                        var v22 float64 = 1.5
                                        v23 := []int{}
                                        _ = v21
                                        _ = v22
                                        _ = v23
                                    }
                                }
                            }
                            switch (fi8 | (fi8 << ((2) & 63))) {
                            case 127:
                                p1 = -2147483648
                            case 32768:
                                var v24 int32 = -62
                                _ = v24
                            case 485:
                                var v25 int16 = int16(p1)
                                _ = v25
                            default:
                                v26 := GenS1{F0: 100, F1: 0, F2: fi12}
                                _ = v26
                            }
                            switch (fi12 / ((fi8) | 1)) {
                            case 2147483647:
                                v27 := GenS1{F0: 42, F1: 4294967295, F2: (fi8 * -489)}
                                _ = v27
                            case -218:
                                if false {
                                    var v28 bool = (-32768 != int16(p2))
                                    var v29 uint16 = 42
                                    _ = v28
                                    _ = v29
                                } else {
                                    var v30 bool = (false && true)
                                    _ = v30
                                }
                            case 127:
                                v31 := []int{449}
                                _ = v31
                            }
                            _ = fi12
                        }
                        v32 := GenS1{F0: -128, F1: 42, F2: p1}
                        for fi33 := 0; fi33 < 8; fi33++ {
                            var v34 float64 = 3.14159
                            p2 = (p0 | 42)
                            _ = fi33
                            _ = v34
                        }
                        _ = v32
                    }
                }
                p1 += p1
            }
        default:
            v35 := GenS1{F0: 127, F1: 18446744073709551615, F2: (p1 >> ((4294967295) & 63))}
            _ = v35
        }
        _ = fi8
    }
    _ = p0
    _ = p1
    _ = p2
    return 128
}
func genFn1(p0 int, p1 uint32, p2 uint) uint32 {
    v36 := []int{p0, (p0 / (((^(p0))) | 1)), p0}
    var v37 uint8 = 1
    v36 = append(v36, p0)
    if len(v36) >= 2 {
        v36 = v36[1:len(v36)]
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v36
    _ = v37
    return uint32(p2)
}

func main() {
    v1 := make(map[uint32]int)
    v2 := []string{"hello", ("a" + "hello"), "hello", "foo bar"}
    var v3 bool = false
    v4 := make(map[int]string)
    v38 := GenS2{F0: 65535, F1: 1, F2: 929}
    var v39 uint = 246
    _ = len(v4)
    v3 = v3
    v4[-2147483648] = ("a" + "")
    if true {
        for fi40 := 0; fi40 < 8; fi40++ {
            if v3 {
                var v41 int8 = -1
                v4[(fi40 - int(v41))] = ("a" + ("hello" + ("foo bar" + "a")))
                v39 += (v39 & v39)
                _ = v41
            }
            var v42 uint8 = 255
            var v43 bool = true
            _ = fi40
            _ = v42
            _ = v43
        }
    }
    v39 = v39
    if (!true) {
        var v44 int8 = 88
        _ = v44
    }
    v39 -= v39
    for fi45 := 0; fi45 < 13; fi45++ {
        if v3 {
            var v46 int = fi45
            for fi47 := 0; fi47 < 10; fi47++ {
                if false {
                    _ = len(v4)
                } else {
                    fmt.Println(fi47, v46, v3)
                }
                if (2147483648 == 2) {
                    if (!true) {
                        if len(v2) > 0 {
                            v2[0] = "hello"
                        }
                        var v48 float32 = ((((float32(0.0) - (float32(0.0) - (((0.1 + (3.14159 * (0.0 + 3.14159))) + 0.1) * 0.0))) - (1.5 * 3.14159)) / 3.14159) + float32(((((3.14159 - float64(1.5)) / ((0.1 * -2.25) / 1.5)) * 1e10) - ((3.14159 + 3.14159) / 3.14159))))
                        _ = v48
                    }
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    v49 := GenS1{F0: -128, F1: 4294967295, F2: (^(fi45))}
                    _ = v49
                } else {
                    v3 = true
                }
                var v50 uint32 = 846
                v50 -= (^(v50))
                _ = len(v2)
                _ = fi47
                _ = v50
            }
            _ = v46
        }
        for fi51 := 0; fi51 < 16; fi51++ {
            v39 += v39
            _ = fi51
        }
        v52 := GenS3{F0: 703, F1: 0, F2: 1}
        _ = fi45
        _ = v52
    }
    var v53 uint8 = 2
    var v54 bool = v3
    var v55 float64 = 1.5
    v54 = v54
    if (2147483648 >= 2147483647) {
        v54 = (uint32(v53) > genFn1(32768, 4294967295, v39))
        v39 += v39
        v53 -= 1
    } else {
        for fi56 := 0; fi56 < 15; fi56++ {
            switch "hello" {
            case "\n\t":
                if (42 != -1) {
                    var v57 float32 = 0.1
                    _ = v57
                } else {
                    fmt.Println(v54, v55, v55)
                    for fi58 := 0; fi58 < 20; fi58++ {
                        switch fi58 {
                        case 127:
                            v53 = v53
                        default:
                            var v59 uint64 = genFn0((v53 ^ (v53 & v53)), (fi58 << ((205) & 63)), 127)
                            if (68 < -128) {
                                if true {
                                    var v60 uint8 = v53
                                    _ = v60
                                }
                            } else {
                                switch 127 {
                                case 32767:
                                    fmt.Println(fi56)
                                case 883:
                                    fmt.Println(fi58, v3, v59)
                                }
                                v59 = v59
                            }
                            v39 -= v39
                            _ = v59
                        }
                        _ = fi58
                    }
                    v61 := GenS1{F0: -58, F1: 42, F2: fi56}
                    _ = v61
                }
            case "hello":
                switch (^(fi56)) {
                case -106:
                    v62 := GenS1{F0: 1, F1: 2, F2: (fi56 - 367)}
                    _ = v62
                }
            case "foo bar":
                if ((v3 && v3) || (!false)) {
                    for fi63 := 0; fi63 < 11; fi63++ {
                        fmt.Println(v55, v53, fi56)
                        _ = fi63
                    }
                    if v54 {
                        if true {
                            switch fi56 {
                            case 0:
                                for fi64 := 0; fi64 < 4; fi64++ {
                                    v65 := GenS4{F0: 2147483647, F1: 0}
                                    fmt.Println(v55, fi56, v53)
                                    var v66 uint16 = 1
                                    _ = fi64
                                    _ = v65
                                    _ = v66
                                }
                            case -146:
                                if (!(42 > 880)) {
                                    var v67 int64 = 1
                                    for fi68 := 0; fi68 < 16; fi68++ {
                                        var v69 string = ""
                                        _ = fi68
                                        _ = v69
                                    }
                                    _ = v67
                                } else {
                                    for fi70 := 0; fi70 < 15; fi70++ {
                                        v71 := GenS1{F0: 100, F1: 595, F2: -764}
                                        _ = fi70
                                        _ = v71
                                    }
                                    var v72 int64 = 42
                                    switch fi56 {
                                    case 82:
                                        v72 = int64(fi56)
                                    case 448:
                                        fmt.Println(v54, v72)
                                    default:
                                        v73 := GenS4{F0: 405, F1: -1}
                                        _ = v73
                                    }
                                    _ = v72
                                }
                            case 802:
                                var v74 int = fi56
                                _ = v74
                            default:
                                if (!v3) {
                                    var v75 uint16 = 32767
                                    _ = v75
                                } else {
                                    if (false || false) {
                                        v53 -= v53
                                    } else {
                                        for fi76 := 0; fi76 < 1; fi76++ {
                                            v2 = append(v2, "a")
                                            v39 -= 4294967295
                                            _ = fi76
                                        }
                                        if len(v2) > 0 {
                                            v2[0] = "foo bar"
                                        }
                                    }
                                }
                            }
                            fmt.Println(v3)
                        } else {
                            fmt.Println(v55)
                            v54 = ("foo bar" < "foo bar")
                        }
                        switch "\n\t" {
                        case "":
                            v39 = 221
                        case "a":
                            fmt.Println(v53, v3)
                        case "foo bar":
                            switch "\n\t" {
                            case "":
                                for fi77 := 0; fi77 < 17; fi77++ {
                                    switch 545 {
                                    case -2147483648:
                                        v78 := GenS2{F0: v39, F1: 666, F2: 65535}
                                        _ = v78
                                    case 4294967295:
                                        fmt.Println(v39, v55, v39)
                                    case 2147483647:
                                        var v79 uint8 = v53
                                        _ = v79
                                    default:
                                        if (!((((v54 && ((v55 == v55) && (((fi77 + (fi77 & fi56)) != 1) && (42 != 2147483647)))) && v3) && (((false || v54) || true) || false)) && v3)) {
                                            v53 = (v53 << (((v39 - 1)) & 63))
                                            v80 := GenS5{F0: true, F1: 2147483647, F2: 4294967295}
                                            fmt.Println(fi77)
                                            v3 = true
                                            _ = v80
                                        }
                                        v2 = append(v2, "hello")
                                    }
                                    _ = fi77
                                }
                            case "foo bar":
                                v54 = (uint64(v53) >= 377)
                            default:
                                _ = len(v1)
                            }
                        default:
                            switch "foo bar" {
                            case "\n\t":
                                var v81 int32 = 1
                                _ = v81
                            default:
                                if true {
                                    switch "" {
                                    case "":
                                        if v54 {
                                            switch fi56 {
                                            case 2147483647:
                                                fmt.Println(v55)
                                            case 583:
                                                var v82 int64 = 1
                                                _ = v82
                                            case 128:
                                                v55 -= v55
                                            }
                                            var v83 uint8 = v53
                                            for fi84 := 0; fi84 < 3; fi84++ {
                                                v83 -= (v53 / (((v83 + (v53 << ((255) & 63)))) | 1))
                                                _ = fi84
                                            }
                                            _ = v83
                                        } else {
                                            var v85 int32 = -1
                                            v86 := GenS5{F0: (244 > -1), F1: v85, F2: genFn1(fi56, 2147483648, (v39 | 2))}
                                            v85 = -2147483648
                                            _ = v85
                                            _ = v86
                                        }
                                    case "hello":
                                        var v87 string = ((("hello" + (("a" + "foo bar") + "foo bar")) + "hello") + "foo bar")
                                        _ = v87
                                    }
                                    for fi88 := 0; fi88 < 1; fi88++ {
                                        var v89 uint64 = 4294967295
                                        _ = fi88
                                        _ = v89
                                    }
                                    var v90 uint64 = 1
                                    _ = v90
                                } else {
                                    for fi91 := 0; fi91 < 20; fi91++ {
                                        switch fi56 {
                                        case 649:
                                            v3 = false
                                        default:
                                            v53 = (^(v53))
                                        }
                                        _ = fi91
                                    }
                                    v92 := v4[787]
                                    if (!true) {
                                        v3 = v54
                                    }
                                    _ = v92
                                }
                            }
                        }
                    } else {
                        var v93 uint8 = v53
                        if v54 {
                            v94 := GenS4{F0: 4294967295, F1: (fi56 / (((-(fi56))) | 1))}
                            var v95 uint64 = 18446744073709551615
                            v39 = (v39 | (^(v39)))
                            _ = v94
                            _ = v95
                        } else {
                            v53 -= (v53 >> ((v39) & 63))
                        }
                        var v96 uint64 = 1
                        _ = v93
                        _ = v96
                    }
                } else {
                    if true {
                        v97 := GenS6{F0: (!v3), F1: 1}
                        v3 = (genFn1(fi56, 1, v39) != 4294967295)
                        for fi98 := 0; fi98 < 7; fi98++ {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            fmt.Println(v54)
                            v39 = v39
                            _ = fi98
                        }
                        _ = v97
                    } else {
                        for fi99 := 0; fi99 < 3; fi99++ {
                            v3 = ((v54 || (0 >= 0)) && v3)
                            var v100 uint = (v39 - (^(v39)))
                            v3 = (-52 < 125)
                            _ = fi99
                            _ = v100
                        }
                        v101 := GenS5{F0: v54, F1: -2147483648, F2: 405}
                        _ = v101
                    }
                }
            default:
                fmt.Println(v55, v54)
            }
            fmt.Println(v3)
            _ = fi56
        }
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v102 := GenS7{F0: 527, F1: 255, F2: 0}
    v54 = ((-2.25 + ((1e10 + 3.14159) / 1e10)) >= 1e10)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v53:", v53)
    fmt.Println("v54:", v54)
    fmt.Println("v55:", v55)
    _ = v1
    _ = v102
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v53
    _ = v54
    _ = v55
}
