package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool; F2 int }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int16; F1 int16; F2 int }
type GenS4 struct { F0 int32; F1 uint; F2 uint32 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 bool; F1 int8; F2 uint32 }
type GenS9 struct { F0 uint32; F1 uint16 }
type GenS10 struct { F0 int64 }
func genFn0(p0 int8, p1 int64) uint8 {
    if (!(0 == 826)) {
        p0 += 0
    } else {
        for fi7 := 0; fi7 < 19; fi7++ {
            v8 := []int{}
            var v9 uint16 = 151
            _ = fi7
            _ = v8
            _ = v9
        }
        fmt.Println(p1, p1)
        p1 -= (-(p1))
    }
    v10 := []int{}
    for fi11 := 0; fi11 < 2; fi11++ {
        switch "foo bar" {
        case "":
            switch "" {
            case "a":
                fmt.Println(p1, p0)
            }
        }
        p1 = (^(p1))
        _ = fi11
    }
    var v12 float64 = (-2.25 / float64((3.14159 + 0.0)))
    var v13 uint32 = 42
    _ = p0
    _ = p1
    _ = v10
    _ = v12
    _ = v13
    return 42
}
func genFn1() uint {
    var v14 int32 = -2147483648
    for fi15 := 0; fi15 < 17; fi15++ {
        v16 := []int{fi15, (fi15 % ((4294967295) | 1)), fi15}
        if (0 >= 2) {
            for fi17 := 0; fi17 < 1; fi17++ {
                v14 = -1
                _ = fi17
            }
        } else {
            var v18 int32 = v14
            if (1000 < -258) {
                var v19 uint64 = 2
                _ = v19
            } else {
                fmt.Println(fi15, fi15, v14)
            }
            for fi20 := 0; fi20 < 16; fi20++ {
                v21 := GenS1{F0: 127, F1: false, F2: fi20}
                if false {
                    v18 -= -547
                    if (-9223372036854775808 > -102) {
                        if len(v16) > 0 {
                            v16[0] = (fi20 & 128)
                        }
                        var v22 string = ("hello" + "a")
                        v18 += (v18 % ((v14) | 1))
                        _ = v22
                    } else {
                        var v23 int32 = v18
                        v16 = append(v16, fi20)
                        _ = v23
                    }
                }
                _ = fi20
                _ = v21
            }
            _ = v18
        }
        if len(v16) > 0 {
            v16[0] = fi15
        }
        _ = fi15
        _ = v16
    }
    v24 := GenS1{F0: 1, F1: true, F2: 389}
    var v25 float64 = 0.0
    _ = v14
    _ = v24
    _ = v25
    return 1
}

func main() {
    var v1 uint32 = 2147483647
    v2 := make(map[uint32]uint32)
    var v3 int64 = 0
    var v4 uint = 766
    var v5 uint32 = (v1 >> ((364) & 63))
    var v6 uint8 = 255
    var v26 uint16 = 65535
    fmt.Println(v1, v1)
    v3 = v3
    if (!true) {
        if (42 != 669) {
            v27 := GenS1{F0: (v6 / ((v6) | 1)), F1: true, F2: 32768}
            if false {
                var v28 int = 128
                for fi29 := 0; fi29 < 12; fi29++ {
                    var v30 int16 = 32767
                    v6 = genFn0(0, (v3 % ((v3) | 1)))
                    fmt.Println(v26, v3, v3)
                    _ = fi29
                    _ = v30
                }
                _ = v28
            } else {
                for fi31 := 0; fi31 < 16; fi31++ {
                    fmt.Println(v5)
                    _ = fi31
                }
            }
            if true {
                v3 -= (v3 & v3)
            } else {
                fmt.Println(v3, v6, v4)
                if true {
                    v32 := GenS2{F0: 127}
                    var v33 string = (("\n\t" + ("foo bar" + "a")) + "\n\t")
                    var v34 uint16 = (v26 ^ v26)
                    _ = v32
                    _ = v33
                    _ = v34
                } else {
                    switch 32768 {
                    case 32767:
                        var v35 uint32 = 1
                        _ = v35
                    case 682:
                        var v36 int = 9223372036854775807
                        _ = v36
                    default:
                        v37 := GenS3{F0: 1, F1: 467, F2: -2147483648}
                        _ = v37
                    }
                    switch 42 {
                    case -9223372036854775808:
                        v1 = (v1 % ((v1) | 1))
                    case 32768:
                        switch "foo bar" {
                        case "a":
                            v38 := GenS2{F0: v6}
                            _ = v38
                        case "":
                            fmt.Println(v26, v3)
                        default:
                            v2[(v1 - (^(v1)))] = (^(v1))
                        }
                    case -2147483648:
                        var v39 float64 = (0.0 - 0.0)
                        _ = v39
                    }
                    fmt.Println(v6, v4)
                }
            }
            _ = v27
        }
    } else {
        if (42 == 0) {
            var v40 int8 = 49
            v41 := GenS3{F0: -1, F1: -1, F2: 9223372036854775807}
            _ = v40
            _ = v41
        }
        v42 := []int{530, 32767, -295, -559}
        _ = v42
    }
    var v43 bool = true
    v44 := GenS3{F0: -1, F1: 32767, F2: 128}
    for fi45 := 0; fi45 < 15; fi45++ {
        v1 = v5
        v2[v5] = 0
        _ = fi45
    }
    switch "\n\t" {
    case "":
        if (v43 || v43) {
            v46 := GenS1{F0: v6, F1: v43, F2: -1}
            v47 := GenS1{F0: (^(v6)), F1: v43, F2: 446}
            _ = v46
            _ = v47
        } else {
            if true {
                if v43 {
                    v26 += v26
                    v43 = (v3 > 1)
                    var v48 int64 = v3
                    _ = v48
                } else {
                    v49 := GenS4{F0: -978, F1: 4294967295, F2: 464}
                    if (!false) {
                        v4 = (^(v4))
                        v43 = (!true)
                    }
                    if (!true) {
                        v50 := GenS4{F0: 1, F1: 552, F2: v1}
                        v4 += v4
                        _ = v50
                    }
                    _ = v49
                }
                v51 := []int{882}
                v5 += (v1 << ((v4) & 63))
                _ = v51
            }
            v6 -= genFn0(0, v3)
        }
    case "hello":
        if (v43 && true) {
            for fi52 := 0; fi52 < 12; fi52++ {
                v53 := v2[(v5 >> ((99) & 63))]
                var v54 uint8 = (v6 - genFn0(0, (v3 << (((^(v4))) & 63))))
                if false {
                    v55 := GenS4{F0: int32(v5), F1: genFn1(), F2: 2147483647}
                    _ = v55
                }
                _ = fi52
                _ = v53
                _ = v54
            }
        }
    case "foo bar":
        for fi56 := 0; fi56 < 9; fi56++ {
            if (v43 || (0.1 > -2.25)) {
                var v57 float32 = float32(((3.14159 * 3.14159) / -2.25))
                _ = v57
            }
            var v58 float64 = (float64(-2.25) + (1e10 + 1e10))
            _ = fi56
            _ = v58
        }
    }
    for fi59 := 0; fi59 < 13; fi59++ {
        delete(v2, 42)
        var v60 int32 = -2147483648
        v61 := GenS5{F0: (-(v60))}
        _ = fi59
        _ = v60
        _ = v61
    }
    var v62 float64 = float64((-2.25 * 3.14159))
    switch "foo bar" {
    case "":
        for fi63 := 0; fi63 < 3; fi63++ {
            var v64 uint16 = 738
            switch 9223372036854775807 {
            case -1:
                v65 := GenS5{F0: -859}
                _ = v65
            }
            _ = fi63
            _ = v64
        }
    case "hello":
        switch "\n\t" {
        case "hello":
            v62 = ((v62 + float64(1.5)) + 1e10)
        case "foo bar":
            v43 = (true && v43)
        case "":
            var v66 int8 = 42
            _ = v66
        }
    case "\n\t":
        switch -770 {
        case 127:
            v67 := []int{460, 2147483647, 127, 63}
            _ = v67
        case 32768:
            var v68 uint = (v4 + 0)
            _ = v68
        }
    }
    if ((1 == 0) || v43) {
        v69 := GenS6{F0: false}
        fmt.Println(v43)
        _ = v69
    } else {
        for fi70 := 0; fi70 < 8; fi70++ {
            v5 = (^(v5))
            v1 += (v5 & v5)
            for fi71 := 0; fi71 < 19; fi71++ {
                v3 = (v3 | (v3 & v3))
                if true {
                    v4 += (v4 << (((v4 & (v4 - 65535))) & 63))
                    v72 := GenS7{F0: 32767}
                    delete(v2, v1)
                    _ = v72
                }
                _ = fi71
            }
            _ = fi70
        }
    }
    v73 := GenS8{F0: v43, F1: 0, F2: (v5 % ((v5) | 1))}
    v4 -= 42
    var v74 float32 = 1.5
    for fi75 := 0; fi75 < 1; fi75++ {
        fmt.Println(v5, fi75)
        _ = fi75
    }
    if v43 {
        if false {
            if (v6 != 255) {
                if (!true) {
                    fmt.Println(v1)
                    switch "a" {
                    case "\n\t":
                        switch "\n\t" {
                        case "foo bar":
                            for fi76 := 0; fi76 < 7; fi76++ {
                                for fi77 := 0; fi77 < 20; fi77++ {
                                    switch "" {
                                    case "\n\t":
                                        fmt.Println(v5, v4)
                                    default:
                                        switch fi77 {
                                        case 32768:
                                            v78 := GenS6{F0: true}
                                            _ = v78
                                        case 1:
                                            v79 := []int{32768, fi77, 42, (-(fi76))}
                                            _ = v79
                                        default:
                                            v4 = genFn1()
                                        }
                                    }
                                    _ = fi77
                                }
                                _ = fi76
                            }
                        case "hello":
                            v80 := []int{-1}
                            _ = v80
                        }
                    default:
                        fmt.Println(v1, v1, v62)
                    }
                } else {
                    if v43 {
                        switch "a" {
                        case "a":
                            fmt.Println(v1, v26, v4)
                        default:
                            for fi81 := 0; fi81 < 12; fi81++ {
                                for fi82 := 0; fi82 < 18; fi82++ {
                                    v3 = (v3 / ((v3) | 1))
                                    if (!true) {
                                        var v83 float64 = ((v62 * 1.5) * 0.1)
                                        _ = v83
                                    } else {
                                        var v84 int = (fi82 - 32767)
                                        if ((v62 / float64(v74)) == v62) {
                                            if false {
                                                v85 := []int{fi82}
                                                _ = v85
                                            }
                                        }
                                        var v86 uint32 = v5
                                        v26 = uint16(v5)
                                        _ = v84
                                        _ = v86
                                    }
                                    v87 := []int{-1}
                                    _ = fi82
                                    _ = v87
                                }
                                for fi88 := 0; fi88 < 20; fi88++ {
                                    var v89 float32 = v74
                                    _ = fi88
                                    _ = v89
                                }
                                _ = fi81
                            }
                        }
                        delete(v2, v5)
                    } else {
                        var v90 int8 = -1
                        v91 := GenS9{F0: v1, F1: (v26 | v26)}
                        _ = v90
                        _ = v91
                    }
                    if v43 {
                        v4 = 886
                        var v92 int64 = (v3 / (((v3 * (v3 * (v3 << (((v4 >> (((v4 * (^(v4)))) & 63))) & 63))))) | 1))
                        _ = v92
                    } else {
                        v26 = uint16(v5)
                        v4 = (^(v4))
                    }
                }
                v43 = ((v62 * v62) == (v62 / (((v62 - -2.25) - v62) + 1.5)))
                fmt.Println(v26, v5, v62)
            } else {
                v26 -= v26
            }
            switch "a" {
            case "hello":
                var v93 bool = v43
                _ = v93
            case "a":
                var v94 uint64 = 18446744073709551615
                _ = v94
            default:
                switch 887 {
                case 4294967295:
                    v62 -= v62
                case 234:
                    for fi95 := 0; fi95 < 17; fi95++ {
                        var v96 uint16 = uint16(v5)
                        for fi97 := 0; fi97 < 1; fi97++ {
                            var v98 uint8 = (v6 << ((v4) & 63))
                            v74 = v74
                            var v99 int16 = 42
                            var v100 int16 = (v99 >> (((v4 & genFn1())) & 63))
                            var v101 float32 = v74
                            fmt.Println(v96)
                            v74 = float32(((3.14159 - v62) + (((-2.25 + (((0.0 * float64(float32(1e10))) - v62) + 3.14159)) + 1e10) / v62)))
                            _ = fi97
                            _ = v100
                            _ = v101
                            _ = v98
                            _ = v99
                        }
                        v2[(v5 | 2147483648)] = 556
                        _ = fi95
                        _ = v96
                    }
                }
            }
        }
        v26 += v26
        v102 := v2[uint32(v26)]
        _ = v102
    } else {
        var v103 int64 = -871
        var v104 bool = false
        fmt.Println(v103)
        _ = v103
        _ = v104
    }
    v105 := GenS10{F0: (v3 - v3)}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v62:", v62)
    fmt.Println("v74:", v74)
    _ = v1
    _ = v105
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v5
    _ = v6
    _ = v62
    _ = v73
    _ = v74
}
