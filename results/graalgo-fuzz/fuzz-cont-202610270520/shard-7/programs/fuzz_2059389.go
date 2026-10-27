package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8 }
type GenS2 struct { F0 int32; F1 int }
type GenS3 struct { F0 uint64; F1 int64 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 uint32; F1 uint32; F2 uint }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 uint64; F1 int; F2 uint16 }
type GenS9 struct { F0 string }
func genFn0() int8 {
    if false {
        if true {
            for fi4 := 0; fi4 < 13; fi4++ {
                var v5 uint = 947
                _ = fi4
                _ = v5
            }
            if (1 >= 163) {
                v6 := GenS1{F0: 2147483647, F1: 128}
                _ = v6
            }
            v7 := []int{752}
            _ = v7
        } else {
            v8 := GenS2{F0: 1, F1: 9223372036854775807}
            v9 := []int{32768, 32767, -1}
            _ = v8
            _ = v9
        }
        var v10 float32 = 1.5
        _ = v10
    } else {
        var v11 float64 = 1e10
        v11 = 1.5
        _ = v11
    }
    var v12 uint = 18446744073709551615
    _ = v12
    return 0
}
func genFn1(p0 uint8, p1 uint8) uint64 {
    var v13 bool = true
    var v14 int64 = 0
    if v13 {
        v14 += v14
        for fi15 := 0; fi15 < 16; fi15++ {
            if (2147483647 > 639) {
                switch "hello" {
                case "foo bar":
                    v13 = (!(0.1 < (3.14159 - 0.0)))
                case "a":
                    var v16 bool = v13
                    _ = v16
                }
                v17 := GenS3{F0: 4294967295, F1: 42}
                if true {
                    fmt.Println(p1, p1)
                } else {
                    v13 = (p1 > (p0 << ((uint(p0)) & 63)))
                    v18 := GenS3{F0: 4294967295, F1: v14}
                    _ = v18
                }
                _ = v17
            } else {
                v19 := []int{(fi15 * 0)}
                for fi20 := 0; fi20 < 1; fi20++ {
                    if len(v19) > 0 {
                        v19[0] = fi15
                    }
                    fmt.Println(v13)
                    fmt.Println(v14, fi20)
                    _ = fi20
                }
                v21 := GenS2{F0: 536, F1: (fi15 << ((42) & 63))}
                _ = v19
                _ = v21
            }
            _ = fi15
        }
    } else {
        for fi22 := 0; fi22 < 9; fi22++ {
            v23 := []int{0, 200, (fi22 % (((-(fi22))) | 1)), (-(fi22))}
            fmt.Println(v14, p0, p0)
            _ = fi22
            _ = v23
        }
        var v24 string = ""
        _ = v24
    }
    fmt.Println(p1, v14)
    _ = p0
    _ = p1
    _ = v13
    _ = v14
    return 285
}

func main() {
    var v1 string = "foo bar"
    var v2 int = 703
    var v3 uint = 255
    for fi25 := 0; fi25 < 2; fi25++ {
        fmt.Println(fi25)
        v26 := GenS4{F0: 0}
        var v27 int = -9223372036854775808
        _ = fi25
        _ = v26
        _ = v27
    }
    if false {
        for fi28 := 0; fi28 < 13; fi28++ {
            v1 = v1
            switch v2 {
            case 479:
                v29 := GenS2{F0: 162, F1: v2}
                _ = v29
            case 32768:
                var v30 uint64 = 42
                _ = v30
            }
            if true {
                v3 -= v3
                if true {
                    if (v2 > (v2 << (((v3 ^ (v3 & 255))) & 63))) {
                        v1 = v1
                    }
                    switch fi28 {
                    case -159:
                        v31 := GenS3{F0: 42, F1: 1}
                        _ = v31
                    }
                } else {
                    if false {
                        switch (^(v2)) {
                        case 1:
                            v1 = ((((v1 + ("foo bar" + ((v1 + v1) + "foo bar"))) + v1) + v1) + v1)
                        case 704:
                            var v32 int8 = 1
                            _ = v32
                        default:
                            v33 := GenS5{F0: 724}
                            _ = v33
                        }
                        v2 = (^(fi28))
                    } else {
                        for fi34 := 0; fi34 < 1; fi34++ {
                            v35 := []int{fi28}
                            var v36 uint16 = 266
                            var v37 int8 = genFn0()
                            _ = fi34
                            _ = v35
                            _ = v36
                            _ = v37
                        }
                        v38 := GenS1{F0: 2147483647, F1: 0}
                        v2 = 2147483647
                        _ = v38
                    }
                    if (-1 > 961) {
                        var v39 float32 = (0.1 / 3.14159)
                        switch (v2 >> (((v3 ^ v3)) & 63)) {
                        case -1:
                            if false {
                                var v40 int16 = -369
                                v2 = fi28
                                for fi41 := 0; fi41 < 11; fi41++ {
                                    if (127 >= 127) {
                                        v1 = v1
                                        v39 -= (1e10 * v39)
                                    } else {
                                        var v42 uint64 = 18446744073709551615
                                        v2 -= (fi28 - v2)
                                        _ = v42
                                    }
                                    if true {
                                        switch "hello" {
                                        case "\n\t":
                                            switch v1 {
                                            case "foo bar":
                                                if false {
                                                    switch v1 {
                                                    case "a":
                                                        var v43 bool = (-56 < genFn0())
                                                        _ = v43
                                                    }
                                                } else {
                                                    v40 += 42
                                                }
                                                v2 -= 1
                                                v2 = fi41
                                                var v44 uint8 = 128
                                                v45 := []int{-9223372036854775808, int(v3), fi41, 2147483647}
                                                _ = v44
                                                _ = v45
                                            }
                                        case "a":
                                            switch v1 {
                                            case "foo bar":
                                                var v46 int16 = v40
                                                _ = v46
                                            default:
                                                var v47 uint8 = 1
                                                _ = v47
                                            }
                                        default:
                                            if true {
                                                v3 = (v3 % ((v3) | 1))
                                                var v48 float32 = -2.25
                                                _ = v48
                                            } else {
                                                fmt.Println(v39)
                                                var v49 uint = v3
                                                if false {
                                                    v50 := []int{(fi41 | v2), (v2 >> (((v49 % (((v3 / ((v3) | 1))) | 1))) & 63)), v2, (fi41 % (((^(fi41))) | 1))}
                                                    v3 = v3
                                                    fmt.Println(fi41, fi28)
                                                    var v51 string = v1
                                                    if (42 > 128) {
                                                        v2 = (v2 + 9223372036854775807)
                                                    } else {
                                                        if (255 != 0) {
                                                            v49 -= 255
                                                            v2 = fi41
                                                            v52 := GenS6{F0: 1, F1: 8, F2: (^(v49))}
                                                            _ = v52
                                                        } else {
                                                            var v53 int8 = 100
                                                            _ = v53
                                                        }
                                                        var v54 uint64 = 18446744073709551615
                                                        fmt.Println(fi28, v3, v1)
                                                        switch v2 {
                                                        case 32767:
                                                            v3 = v49
                                                        default:
                                                            var v55 bool = true
                                                            _ = v55
                                                        }
                                                        _ = v54
                                                    }
                                                    v50 = append(v50, (fi28 >> ((v49) & 63)))
                                                    _ = v50
                                                    _ = v51
                                                }
                                                _ = v49
                                            }
                                            switch (-(fi28)) {
                                            case 2147483647:
                                                if (1 <= (v40 / ((42) | 1))) {
                                                    fmt.Println(fi28, fi28)
                                                } else {
                                                    var v56 uint32 = 1
                                                    switch fi28 {
                                                    case 0:
                                                        var v57 float64 = (1.5 / ((3.14159 / float64(float32((float64(3.14159) * ((((0.0 + 1e10) + ((3.14159 + (0.1 + ((float64(float32(0.1)) + -2.25) - -2.25))) * 0.1)) / float64((3.14159 * v39))) + (0.1 + float64(((((1e10 + (v39 + (v39 * (float32(1e10) - 3.14159)))) * v39) / v39) + 0.1)))))))) - 3.14159))
                                                        _ = v57
                                                    case -9223372036854775808:
                                                        var v58 bool = (880 != 0)
                                                        _ = v58
                                                    case 738:
                                                        var v59 string = "foo bar"
                                                        _ = v59
                                                    default:
                                                        v40 = (v40 & 984)
                                                        var v60 uint16 = 255
                                                        _ = v60
                                                    }
                                                    v61 := []int{(fi41 % ((fi28) | 1)), (fi41 + 1)}
                                                    _ = v56
                                                    _ = v61
                                                }
                                            }
                                        }
                                        if (1000 == 124) {
                                            var v62 float64 = 0.0
                                            var v63 float64 = (v62 / ((float64(v39) - v62) / 1.5))
                                            switch 32767 {
                                            case 2147483647:
                                                var v64 int8 = 42
                                                _ = v64
                                            case 557:
                                                v63 = 1e10
                                            case -1:
                                                v65 := GenS6{F0: 65535, F1: 1, F2: v3}
                                                _ = v65
                                            default:
                                                switch (v1 + v1) {
                                                case "hello":
                                                    v63 = v63
                                                case "":
                                                    if (0 == 663) {
                                                        v66 := []int{(fi41 % (((fi41 - (fi28 ^ v2))) | 1))}
                                                        if len(v66) >= 2 {
                                                            v66 = v66[1:len(v66)]
                                                        }
                                                        if len(v66) > 0 {
                                                            v66[0] = fi41
                                                        }
                                                        _ = v66
                                                    }
                                                    v67 := []int{(v2 >> (((v3 % ((903) | 1))) & 63)), (fi41 | fi41), (fi41 * -855), (fi28 >> (((v3 << ((255) & 63))) & 63))}
                                                    v68 := GenS7{F0: 232}
                                                    _ = v67
                                                    _ = v68
                                                default:
                                                    v69 := []int{(^(fi28)), (fi41 << (((^(v3))) & 63)), fi28}
                                                    _ = v69
                                                }
                                                v2 = v2
                                                fmt.Println(v1, v3, fi28)
                                            }
                                            _ = v62
                                            _ = v63
                                        } else {
                                            v39 += -2.25
                                            fmt.Println(v1, v40)
                                            var v70 int64 = 1
                                            v71 := GenS8{F0: 4294967295, F1: (v2 + 127), F2: 32767}
                                            _ = v70
                                            _ = v71
                                        }
                                        v39 -= v39
                                    } else {
                                        v72 := []int{}
                                        _ = v72
                                    }
                                    _ = fi41
                                }
                                _ = v40
                            } else {
                                var v73 float64 = ((1.5 * 0.1) / -2.25)
                                _ = v73
                            }
                        case 128:
                            fmt.Println(v39, v39, v3)
                        }
                        _ = v39
                    }
                }
                if (true || (0 <= 521)) {
                    v74 := GenS4{F0: 127}
                    for fi75 := 0; fi75 < 4; fi75++ {
                        v76 := GenS3{F0: 18446744073709551615, F1: 1}
                        fmt.Println(fi75, v1)
                        v2 += fi28
                        _ = fi75
                        _ = v76
                    }
                    _ = v74
                } else {
                    for fi77 := 0; fi77 < 13; fi77++ {
                        v3 += v3
                        if (-780 != int64(fi28)) {
                            v78 := GenS2{F0: 1, F1: fi28}
                            var v79 int8 = -78
                            _ = v78
                            _ = v79
                        }
                        _ = fi77
                    }
                    v80 := []int{238}
                    v1 = "a"
                    _ = v80
                }
            } else {
                if true {
                    v81 := GenS4{F0: 2}
                    var v82 int8 = genFn0()
                    v83 := []int{(v2 % ((v2) | 1))}
                    _ = v81
                    _ = v82
                    _ = v83
                } else {
                    for fi84 := 0; fi84 < 6; fi84++ {
                        var v85 uint32 = 2
                        _ = fi84
                        _ = v85
                    }
                    var v86 int8 = 107
                    var v87 bool = false
                    _ = v86
                    _ = v87
                }
                v88 := []int{fi28, 578, 251, (^(v2))}
                _ = v88
            }
            _ = fi28
        }
    }
    v2 -= (^(v2))
    if ("\n\t" > "hello") {
        var v89 uint64 = 676
        _ = v89
    }
    var v90 float64 = -2.25
    for fi91 := 0; fi91 < 18; fi91++ {
        var v92 int32 = int32(v3)
        _ = fi91
        _ = v92
    }
    if (2 >= 102) {
        for fi93 := 0; fi93 < 2; fi93++ {
            var v94 int = 128
            switch fi93 {
            case -255:
                v3 += 2
            case 9223372036854775807:
                for fi95 := 0; fi95 < 18; fi95++ {
                    if (!(true || true)) {
                        v96 := GenS1{F0: 0, F1: 128}
                        _ = v96
                    }
                    var v97 float32 = 0.0
                    v90 = (v90 - float64(((v97 + (-2.25 + ((v97 * (v97 * (v97 + float32(float64((v97 + v97)))))) * 0.1))) - 1.5)))
                    _ = fi95
                    _ = v97
                }
            case -521:
                switch fi93 {
                case 0:
                    for fi98 := 0; fi98 < 8; fi98++ {
                        switch "" {
                        case "hello":
                            switch "" {
                            case "foo bar":
                                fmt.Println(v90, v90, v90)
                            case "hello":
                                var v99 uint32 = 2147483647
                                _ = v99
                            case "a":
                                var v100 uint16 = 384
                                _ = v100
                            }
                        }
                        fmt.Println(fi98, fi98, v1)
                        v101 := []int{fi98, 1}
                        _ = fi98
                        _ = v101
                    }
                case 4294967295:
                    var v102 int64 = -1
                    _ = v102
                default:
                    fmt.Println(v3, v3)
                }
            default:
                var v103 uint = v3
                _ = v103
            }
            _ = fi93
            _ = v94
        }
        switch v1 {
        case "":
            if (1 != -9223372036854775808) {
                v104 := []int{v2, int(v3)}
                _ = v104
            }
        case "\n\t":
            for fi105 := 0; fi105 < 14; fi105++ {
                var v106 float64 = 1.5
                switch v1 {
                case "foo bar":
                    var v107 uint8 = 128
                    _ = v107
                case "\n\t":
                    v90 -= v106
                case "":
                    if false {
                        v90 = (v90 + float64((0.1 / 1.5)))
                    } else {
                        var v108 int = (v2 << (((v3 % ((v3) | 1))) & 63))
                        _ = v108
                    }
                default:
                    v109 := GenS3{F0: 18446744073709551615, F1: 1}
                    _ = v109
                }
                fmt.Println(v1, fi105)
                _ = fi105
                _ = v106
            }
        default:
            v110 := []int{(v2 & -2147483648), v2}
            _ = v110
        }
    }
    var v111 uint64 = 656
    fmt.Println(v111, v3, v3)
    switch (v1 + ((v1 + "") + v1)) {
    case "\n\t":
        switch (v1 + (((v1 + v1) + "\n\t") + v1)) {
        case "\n\t":
            fmt.Println(v1)
        case "":
            var v112 float64 = ((float64(3.14159) - v90) / v90)
            _ = v112
        default:
            var v113 int32 = -2147483648
            _ = v113
        }
    case "":
        if (((!(626 >= (v111 + (v111 / (((v111 | 2)) | 1))))) && false) || true) {
            v1 = v1
        } else {
            v111 = v111
            v114 := GenS2{F0: -1, F1: (v2 - (v2 ^ (v2 + v2)))}
            var v115 uint = (v3 ^ v3)
            _ = v114
            _ = v115
        }
    case "foo bar":
        v116 := []int{(v2 << (((^(v3))) & 63)), v2, v2, 4294967295}
        _ = v116
    }
    if ((v111 <= v111) && ((!(false && true)) && (98 == 506))) {
        v2 = v2
    }
    v117 := GenS5{F0: genFn1(143, 234)}
    v118 := GenS9{F0: v1}
    fmt.Println("v1:", v1)
    fmt.Println("v111:", v111)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v90:", v90)
    _ = v1
    _ = v111
    _ = v117
    _ = v118
    _ = v2
    _ = v3
    _ = v90
}
