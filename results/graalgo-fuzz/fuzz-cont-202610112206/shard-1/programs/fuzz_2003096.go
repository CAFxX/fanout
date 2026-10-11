package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int64; F1 int; F2 int16 }
type GenS3 struct { F0 int32; F1 uint32; F2 int8 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 uint64; F1 int32 }
type GenS8 struct { F0 bool; F1 int8 }
type GenS9 struct { F0 int64; F1 string; F2 int16 }
type GenS10 struct { F0 int64; F1 uint64; F2 int32 }
type GenS11 struct { F0 int }
type GenS12 struct { F0 string; F1 uint16; F2 uint }
type GenS13 struct { F0 uint; F1 uint64 }
func genFn0(p0 int8) uint32 {
    for fi7 := 0; fi7 < 4; fi7++ {
        p0 += 68
        for fi8 := 0; fi8 < 17; fi8++ {
            fmt.Println(p0)
            if false {
                p0 = (p0 * p0)
                p0 = p0
                fmt.Println(fi7)
            } else {
                var v9 int64 = 691
                if true {
                    v10 := []int{}
                    if len(v10) >= 2 {
                        v10 = v10[1:len(v10)]
                    }
                    _ = v10
                } else {
                    var v11 int = fi8
                    var v12 string = "foo bar"
                    _ = v11
                    _ = v12
                }
                v13 := []int{fi8, fi8, 0}
                _ = v13
                _ = v9
            }
            _ = fi8
        }
        _ = fi7
    }
    var v14 float32 = 0.1
    v15 := []int{}
    p0 = p0
    _ = p0
    _ = v14
    _ = v15
    return 65535
}
func genFn1(p0 uint, p1 int) int {
    fmt.Println(p1, p0)
    for fi16 := 0; fi16 < 20; fi16++ {
        fmt.Println(p0, p1)
        fmt.Println(p1)
        v17 := GenS1{F0: false}
        _ = fi16
        _ = v17
    }
    for fi18 := 0; fi18 < 10; fi18++ {
        fmt.Println(p1)
        var v19 int8 = 127
        _ = fi18
        _ = v19
    }
    _ = p0
    _ = p1
    return p1
}

func main() {
    v1 := make(map[int]string)
    var v2 int64 = -330
    v3 := []int{}
    var v4 uint = 0
    var v5 uint64 = 744
    var v6 uint16 = 0
    if (v5 > 18446744073709551615) {
        fmt.Println(v6, v5)
        switch 71 {
        case 0:
            var v20 int64 = v2
            _ = v20
        case 32767:
            var v21 uint8 = 0
            _ = v21
        }
        if ((1 > -128) && (239 > uint8(v4))) {
            if true {
                fmt.Println(v4, v6, v4)
            } else {
                for fi22 := 0; fi22 < 5; fi22++ {
                    var v23 uint = 4294967295
                    fmt.Println(fi22)
                    for fi24 := 0; fi24 < 15; fi24++ {
                        if (!(1e10 < (1e10 * 3.14159))) {
                            var v25 float32 = float32((-2.25 * (3.14159 / -2.25)))
                            _ = v25
                        }
                        _ = fi24
                    }
                    _ = fi22
                    _ = v23
                }
            }
            for fi26 := 0; fi26 < 13; fi26++ {
                fmt.Println(v6)
                switch ("hello" + ("\n\t" + ("\n\t" + "foo bar"))) {
                case "":
                    var v27 uint = v4
                    _ = v27
                }
                _ = fi26
            }
        } else {
            var v28 uint8 = 226
            v29 := GenS1{F0: ((^(v5)) != (^(v5)))}
            v5 -= (v5 & (v5 - 18446744073709551615))
            _ = v28
            _ = v29
        }
    }
    v2 = 1
    var v30 string = "foo bar"
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v31 int64 = 9223372036854775807
    if (116 == 42) {
        for fi32 := 0; fi32 < 4; fi32++ {
            v33 := GenS2{F0: (v2 | (v2 + (v31 + v31))), F1: fi32, F2: 0}
            switch fi32 {
            case 128:
                fmt.Println(v5, v31)
            case 2147483647:
                for fi34 := 0; fi34 < 7; fi34++ {
                    var v35 string = v30
                    v5 = (v5 / (((v5 + (v5 << (((v4 ^ v4)) & 63)))) | 1))
                    _ = fi34
                    _ = v35
                }
            }
            v6 = v6
            _ = fi32
            _ = v33
        }
    }
    v5 = v5
    if true {
        for fi36 := 0; fi36 < 20; fi36++ {
            v6 += (^(v6))
            var v37 uint = uint(v6)
            _ = fi36
            _ = v37
        }
        v38 := GenS1{F0: ((!((918 > -32768) || ((false || (-1 < 470)) && true))) || (-468 >= 0))}
        _ = v38
    } else {
        v39 := GenS2{F0: -9223372036854775808, F1: -790, F2: -1}
        _ = v39
    }
    var v40 float32 = 0.0
    var v41 uint64 = 410
    fmt.Println(v5)
    var v42 float32 = v40
    for fi43 := 0; fi43 < 14; fi43++ {
        if true {
            _ = len(v3)
            v44 := GenS3{F0: -514, F1: 65535, F2: -57}
            switch ((v30 + (((v30 + v30) + "foo bar") + (v30 + v30))) + v30) {
            case "hello":
                var v45 uint = v4
                _ = v45
            default:
                if (127 == 100) {
                    if (false || true) {
                        fmt.Println(v4)
                        v46 := GenS3{F0: 1, F1: 2147483648, F2: 127}
                        var v47 uint8 = 127
                        _ = v46
                        _ = v47
                    }
                } else {
                    var v48 int8 = 42
                    if (((!false) || false) || (4294967295 != genFn0(v48))) {
                        fmt.Println(v6, v5, v30)
                        var v49 bool = true
                        var v50 float32 = ((0.0 - v40) - v42)
                        _ = v49
                        _ = v50
                    } else {
                        if true {
                            v51 := GenS3{F0: -2147483648, F1: genFn0((-(v48))), F2: v48}
                            switch "a" {
                            case "":
                                v30 = v30
                            case "a":
                                var v52 uint64 = (v5 / ((148) | 1))
                                _ = v52
                            case "foo bar":
                                v6 -= v6
                            default:
                                v42 += (((v42 * 0.0) + float32(float64(-2.25))) * v42)
                            }
                            fmt.Println(fi43, v40, v2)
                            _ = v51
                        }
                        v41 = (v41 & v5)
                        v53 := GenS1{F0: ((v42 * 3.14159) != (v42 / (v42 * v42)))}
                        _ = v53
                    }
                    _ = v48
                }
            }
            _ = v44
        }
        if (-128 >= -39) {
            var v54 int16 = -427
            v55 := GenS3{F0: -2147483648, F1: genFn0(1), F2: 100}
            for fi56 := 0; fi56 < 10; fi56++ {
                var v57 float64 = -2.25
                v42 = 1.5
                _ = fi56
                _ = v57
            }
            _ = v54
            _ = v55
        } else {
            for fi58 := 0; fi58 < 3; fi58++ {
                var v59 int16 = -32768
                var v60 bool = false
                _ = fi58
                _ = v59
                _ = v60
            }
            switch "" {
            case "\n\t":
                var v61 uint16 = v6
                _ = v61
            }
            v6 = (^(v6))
        }
        for fi62 := 0; fi62 < 9; fi62++ {
            var v63 string = ("a" + (((v30 + v30) + "\n\t") + ""))
            _ = fi62
            _ = v63
        }
        _ = fi43
    }
    v64 := GenS3{F0: -686, F1: 4294967295, F2: int8(v4)}
    if false {
        v3 = append(v3, 32767)
        v30 = "foo bar"
        for fi65 := 0; fi65 < 7; fi65++ {
            if false {
                v30 = v30
                switch ("hello" + "") {
                case "a":
                    if (v5 == 4294967295) {
                        v66 := GenS4{F0: (^(fi65))}
                        _ = v66
                    }
                case "hello":
                    v67 := GenS5{F0: (v6 / ((2) | 1))}
                    _ = v67
                case "":
                    var v68 uint16 = (v6 >> ((v4) & 63))
                    _ = v68
                default:
                    v31 += 9223372036854775807
                }
            } else {
                var v69 int = (fi65 << ((v4) & 63))
                if false {
                    v70 := GenS6{F0: (-(v2))}
                    switch fi65 {
                    case -1:
                        v3 = append(v3, v69)
                    default:
                        v71 := GenS1{F0: true}
                        _ = v71
                    }
                    _ = v70
                } else {
                    v31 = (v31 - 42)
                    var v72 int32 = -412
                    if ((!true) || false) {
                        v73 := GenS5{F0: v6}
                        switch v30 {
                        case "":
                            switch v69 {
                            case 128:
                                fmt.Println(v31, v2, v42)
                            case -124:
                                var v74 uint16 = (v6 ^ v6)
                                _ = v74
                            default:
                                if false {
                                    switch int(v4) {
                                    case -324:
                                        v30 = ""
                                    }
                                }
                            }
                        case "a":
                            v42 += 0.0
                        case "foo bar":
                            var v75 uint8 = 126
                            _ = v75
                        }
                        var v76 uint64 = (v5 ^ 627)
                        _ = v73
                        _ = v76
                    } else {
                        v72 = int32(v4)
                        for fi77 := 0; fi77 < 16; fi77++ {
                            var v78 float64 = (1.5 * -2.25)
                            v3 = append(v3, v69)
                            fmt.Println(v69)
                            _ = fi77
                            _ = v78
                        }
                        fmt.Println(fi65, v5, v42)
                    }
                    _ = v72
                }
                v79 := GenS7{F0: 2, F1: 42}
                _ = v69
                _ = v79
            }
            for fi80 := 0; fi80 < 13; fi80++ {
                switch genFn1(2, int(v31)) {
                case -398:
                    var v81 int8 = 127
                    _ = v81
                case 1:
                    if ((fi65 == fi80) && (!(!true))) {
                        v2 = v31
                        if true {
                            switch v30 {
                            case "foo bar":
                                v40 = 3.14159
                            case "\n\t":
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            case "a":
                                v82 := GenS8{F0: (true || (!false)), F1: 127}
                                _ = v82
                            }
                            if (true && (true && true)) {
                                v5 += v41
                                if true {
                                    var v83 string = "\n\t"
                                    delete(v1, genFn1(v4, fi65))
                                    _ = v83
                                } else {
                                    switch genFn1(v4, genFn1((v4 & 2), (fi65 + (^(fi65))))) {
                                    case 128:
                                        if false {
                                            var v84 int16 = 1000
                                            var v85 uint16 = (v6 % ((uint16(v41)) | 1))
                                            _ = v84
                                            _ = v85
                                        } else {
                                            v4 = v4
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                        }
                                    case 32768:
                                        if true {
                                            switch fi80 {
                                            case -271:
                                                switch v30 {
                                                case "foo bar":
                                                    var v86 int16 = 461
                                                    if (v86 < int16(v5)) {
                                                        switch genFn1(v4, genFn1((v4 >> (((^(v4))) & 63)), (^(fi65)))) {
                                                        case 263:
                                                            v87 := GenS9{F0: v2, F1: v30, F2: (v86 + v86)}
                                                            _ = v87
                                                        case 1:
                                                            v5 = 556
                                                        default:
                                                            fmt.Println(v86)
                                                        }
                                                    }
                                                    if (0 < -2147483648) {
                                                        var v88 uint = (^(v4))
                                                        v3 = append(v3, (fi65 << ((v88) & 63)))
                                                        _ = v88
                                                    } else {
                                                        var v89 uint32 = 4294967295
                                                        _ = v89
                                                    }
                                                    var v90 int64 = (v2 - -464)
                                                    v5 = (v5 * (v41 ^ v41))
                                                    var v91 bool = ((v41 >> (((v4 - 255)) & 63)) < v41)
                                                    _ = v86
                                                    _ = v90
                                                    _ = v91
                                                case "a":
                                                    _ = len(v1)
                                                }
                                                var v92 uint = v4
                                                var v93 uint64 = v41
                                                var v94 uint = 295
                                                _ = v92
                                                _ = v93
                                                _ = v94
                                            case 925:
                                                v30 = v30
                                            case 0:
                                                switch v30 {
                                                case "a":
                                                    _ = len(v1)
                                                case "hello":
                                                    var v95 bool = false
                                                    _ = v95
                                                default:
                                                    switch genFn1(v4, fi65) {
                                                    case 496:
                                                        switch (fi80 & (^(fi80))) {
                                                        case 127:
                                                            switch "\n\t" {
                                                            case "\n\t":
                                                                v2 = v2
                                                            }
                                                        case -9223372036854775808:
                                                            v5 -= (v41 % (((v41 << ((v4) & 63))) | 1))
                                                        default:
                                                            v96 := GenS10{F0: (v2 % ((v31) | 1)), F1: (v5 / ((846) | 1)), F2: 2147483647}
                                                            _ = v96
                                                        }
                                                    case 516:
                                                        v3 = append(v3, fi80)
                                                    case 42:
                                                        var v97 float32 = v42
                                                        _ = v97
                                                    default:
                                                        v98 := GenS11{F0: 110}
                                                        _ = v98
                                                    }
                                                }
                                            }
                                        } else {
                                            fmt.Println(v31, v31)
                                            v99 := GenS12{F0: "hello", F1: (v6 / ((uint16(v4)) | 1)), F2: (v4 / ((v4) | 1))}
                                            _ = v99
                                        }
                                    }
                                    var v100 uint = (v4 * 0)
                                    fmt.Println(v30, v100)
                                    _ = v100
                                }
                            } else {
                                if (!true) {
                                    v31 += (v31 & v31)
                                    var v101 uint32 = uint32(v5)
                                    v31 += v2
                                    _ = v101
                                }
                                v41 += (^(v5))
                                v102 := GenS13{F0: v4, F1: uint64(v6)}
                                _ = v102
                            }
                            var v103 int16 = -630
                            _ = v103
                        }
                    } else {
                        v30 = (v30 + v30)
                        switch fi65 {
                        case -2147483648:
                            fmt.Println(v41, v40)
                        case -9223372036854775808:
                            if (0 > -1) {
                                switch "foo bar" {
                                case "foo bar":
                                    delete(v1, -2147483648)
                                case "a":
                                    switch "foo bar" {
                                    case "\n\t":
                                        fmt.Println(v6, v30)
                                    case "":
                                        delete(v1, genFn1((v4 / ((v4) | 1)), 0))
                                    case "foo bar":
                                        _ = len(v3)
                                    }
                                }
                            } else {
                                v104 := GenS9{F0: 42, F1: (v30 + v30), F2: 32767}
                                _ = v104
                            }
                        case -1:
                            if (true || true) {
                                v4 += (v4 * (v4 - (v4 % ((18446744073709551615) | 1))))
                                if true {
                                    v40 -= 1e10
                                } else {
                                    v105 := v1[fi80]
                                    _ = v105
                                }
                            } else {
                                var v106 int = -2147483648
                                _ = v106
                            }
                        default:
                            v6 = v6
                        }
                    }
                default:
                    var v107 uint16 = v6
                    _ = v107
                }
                _ = fi80
            }
            _ = fi65
        }
    }
    var v108 int8 = 1
    v42 = (float32(((-2.25 / (-2.25 / 1e10)) + 1.5)) * float32((3.14159 * (((-2.25 * float64(1e10)) * (-2.25 + 3.14159)) + 1e10))))
    delete(v1, -2147483648)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v108:", v108)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v108
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v5
    _ = v6
    _ = v64
}
