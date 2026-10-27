package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8; F2 int64 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint32; F1 uint8 }
type GenS4 struct { F0 int16; F1 uint }
type GenS5 struct { F0 int8; F1 int8; F2 uint64 }
type GenS6 struct { F0 int16; F1 int16 }
type GenS7 struct { F0 string }
type GenS8 struct { F0 bool; F1 uint }
type GenS9 struct { F0 int16 }
type GenS10 struct { F0 int; F1 uint8; F2 int8 }
func genFn0(p0 bool, p1 int16, p2 uint) int64 {
    fmt.Println(p2, p2)
    if (true && p0) {
        if (32768 == -674) {
            fmt.Println(p2, p0, p1)
            p2 = 4294967295
            var v7 int64 = -1
            _ = v7
        }
        if (-104 <= 42) {
            if p0 {
                var v8 bool = (p0 || true)
                var v9 string = ("a" + "")
                v8 = (v8 || ((!("foo bar" >= "foo bar")) || v8))
                _ = v8
                _ = v9
            } else {
                var v10 uint = (p2 ^ (^(p2)))
                _ = v10
            }
        } else {
            if false {
                p1 -= p1
                if p0 {
                    v11 := GenS1{F0: 605, F1: 42, F2: 42}
                    switch -925 {
                    case -2147483648:
                        fmt.Println(p1, p2)
                    case 42:
                        var v12 float64 = float64(-2.25)
                        _ = v12
                    case -902:
                        v13 := GenS1{F0: 2, F1: 42, F2: -1}
                        _ = v13
                    default:
                        v14 := GenS2{F0: 42}
                        _ = v14
                    }
                    _ = v11
                } else {
                    v15 := []int{-781, -1, -2147483648}
                    _ = v15
                }
            } else {
                fmt.Println(p0, p2)
                if false {
                    var v16 int16 = (p1 >> ((p2) & 63))
                    fmt.Println(p1, p2, p1)
                    _ = v16
                } else {
                    var v17 int = 867
                    _ = v17
                }
                var v18 int16 = (p1 - 42)
                _ = v18
            }
            fmt.Println(p0, p1, p2)
        }
        var v19 string = (("foo bar" + "\n\t") + "hello")
        _ = v19
    }
    fmt.Println(p2, p0, p0)
    var v20 float64 = 0.0
    _ = p0
    _ = p1
    _ = p2
    _ = v20
    return -644
}
func genFn1(p0 uint, p1 uint, p2 uint32) int {
    if (!false) {
        switch 417 {
        case 858:
            switch "hello" {
            case "":
                if (42 < -1) {
                    p0 += (^(p0))
                    switch "foo bar" {
                    case "hello":
                        var v21 float64 = 1e10
                        _ = v21
                    case "\n\t":
                        fmt.Println(p0)
                    case "a":
                        if (127 <= -77) {
                            var v22 string = ((("a" + "hello") + "hello") + "hello")
                            _ = v22
                        } else {
                            var v23 int16 = 425
                            p2 = (p2 ^ p2)
                            for fi24 := 0; fi24 < 6; fi24++ {
                                p0 -= (p1 * p0)
                                for fi25 := 0; fi25 < 3; fi25++ {
                                    v26 := []int{2147483647, fi24, 9223372036854775807, 0}
                                    var v27 int = fi24
                                    if len(v26) > 0 {
                                        v26[0] = fi24
                                    }
                                    fmt.Println(v27, p2)
                                    _ = fi25
                                    _ = v26
                                    _ = v27
                                }
                                _ = fi24
                            }
                            _ = v23
                        }
                    }
                    v28 := []int{1}
                    _ = v28
                } else {
                    p1 = (p0 % ((p1) | 1))
                }
            case "a":
                var v29 bool = false
                _ = v29
            default:
                fmt.Println(p0, p0)
            }
        case 32768:
            v30 := GenS2{F0: 42}
            _ = v30
        case 577:
            p2 = 2
        default:
            for fi31 := 0; fi31 < 10; fi31++ {
                p2 -= (p2 & p2)
                v32 := []int{fi31, fi31, fi31}
                var v33 int16 = 1000
                _ = fi31
                _ = v32
                _ = v33
            }
        }
        for fi34 := 0; fi34 < 6; fi34++ {
            var v35 uint16 = 0
            _ = fi34
            _ = v35
        }
        switch "foo bar" {
        case "\n\t":
            var v36 uint16 = 31
            _ = v36
        case "foo bar":
            fmt.Println(p0, p1)
        default:
            v37 := []int{0}
            _ = v37
        }
    } else {
        v38 := GenS2{F0: 2}
        v39 := []int{9223372036854775807, 2147483647, 276}
        _ = len(v39)
        _ = v38
        _ = v39
    }
    if ((30 == 4294967295) && false) {
        v40 := []int{-760}
        _ = v40
    } else {
        switch 32768 {
        case 1:
            v41 := GenS2{F0: 970}
            _ = v41
        }
    }
    switch 9223372036854775807 {
    case 230:
        var v42 int8 = int8(p0)
        _ = v42
    case -9223372036854775808:
        v43 := []int{9223372036854775807, 679}
        _ = v43
    default:
        v44 := []int{}
        _ = v44
    }
    _ = p0
    _ = p1
    _ = p2
    return -1
}

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[uint32]uint32)
    v3 := make(map[int]bool)
    var v4 float64 = (3.14159 / 0.1)
    v5 := make(map[int]int)
    var v6 int32 = -1
    v4 += (1.5 + ((v4 + 0.0) / v4))
    if true {
        if (true && (301 < 18446744073709551615)) {
            if (false || true) {
                v3[32767] = (-1 == -128)
                if true {
                    var v45 bool = (2 <= 434)
                    var v46 int8 = -128
                    _ = v45
                    _ = v46
                } else {
                    if true {
                        fmt.Println(v6, v6)
                        if (false || (42 != 42)) {
                            v6 = (v6 % (((^(v6))) | 1))
                            var v47 bool = false
                            _ = v47
                        } else {
                            if (((107 >= 0) && false) || true) {
                                var v48 uint8 = 128
                                _ = v48
                            } else {
                                v4 = v4
                            }
                            var v49 uint8 = 1
                            _ = v49
                        }
                        v50 := GenS1{F0: 162, F1: -128, F2: -1}
                        _ = v50
                    } else {
                        v6 += v6
                    }
                    fmt.Println(v4, v4)
                    for fi51 := 0; fi51 < 8; fi51++ {
                        switch "hello" {
                        case "\n\t":
                            if (1 < 42) {
                                var v52 int16 = 32767
                                if false {
                                    var v53 int32 = v6
                                    var v54 string = (("foo bar" + ("hello" + "foo bar")) + ("hello" + ("\n\t" + (("\n\t" + "hello") + "a"))))
                                    _ = v53
                                    _ = v54
                                }
                                for fi55 := 0; fi55 < 18; fi55++ {
                                    var v56 int32 = 2147483647
                                    var v57 bool = true
                                    _ = fi55
                                    _ = v56
                                    _ = v57
                                }
                                _ = v52
                            } else {
                                fmt.Println(v6)
                                var v58 uint64 = 2
                                _ = v58
                            }
                        case "a":
                            for fi59 := 0; fi59 < 13; fi59++ {
                                var v60 int16 = 0
                                v61 := GenS1{F0: 42, F1: 125, F2: 9223372036854775807}
                                v60 = (v60 & v60)
                                _ = fi59
                                _ = v60
                                _ = v61
                            }
                        }
                        _ = fi51
                    }
                }
                fmt.Println(v6, v4)
            } else {
                fmt.Println(v6, v6)
                for fi62 := 0; fi62 < 1; fi62++ {
                    v63 := GenS3{F0: 165, F1: 127}
                    _ = fi62
                    _ = v63
                }
                fmt.Println(v4)
            }
        }
        v64 := []int{genFn1(831, 100, 2147483648)}
        var v65 bool = true
        _ = v64
        _ = v65
    } else {
        v66 := GenS4{F0: 826, F1: 288}
        _ = v66
    }
    v6 += v6
    v4 = 3.14159
    if (77 <= 1) {
        delete(v2, 0)
        for fi67 := 0; fi67 < 3; fi67++ {
            switch (fi67 & -2147483648) {
            case 4294967295:
                v68 := []int{}
                _ = v68
            case 2147483647:
                v6 = v6
            case 1:
                v69 := GenS5{F0: 42, F1: 1, F2: 890}
                _ = v69
            default:
                for fi70 := 0; fi70 < 15; fi70++ {
                    var v71 float32 = (((1e10 - 0.0) / 1.0) / 0.1)
                    if ((float32(3.14159) + v71) == (v71 + v71)) {
                        var v72 int8 = 127
                        _ = v72
                    } else {
                        fmt.Println(v6, fi67, v6)
                        var v73 float32 = float32(v4)
                        _ = v73
                    }
                    v2[2147483648] = 2
                    _ = fi70
                    _ = v71
                }
            }
            var v74 int64 = -106
            for fi75 := 0; fi75 < 14; fi75++ {
                if false {
                    var v76 string = ""
                    v77 := []int{genFn1(2, 42, 836), fi67, genFn1(2, 255, 343), fi75}
                    if (!true) {
                        switch (v76 + (v76 + v76)) {
                        case "hello":
                            v78 := GenS6{F0: 1000, F1: -540}
                            _ = v78
                        }
                        v6 = (v6 % ((v6) | 1))
                    } else {
                        switch (fi67 >> ((42) & 63)) {
                        case 32767:
                            if len(v77) > 0 {
                                v77[0] = genFn1(65535, 255, 1)
                            }
                        case 149:
                            var v79 int64 = genFn0((v76 >= ""), 1000, 271)
                            _ = v79
                        case 9223372036854775807:
                            v80 := GenS6{F0: 0, F1: 928}
                            _ = v80
                        default:
                            var v81 float64 = (v4 * v4)
                            _ = v81
                        }
                        v82 := GenS7{F0: v76}
                        if (448 < 0) {
                            v83 := GenS8{F0: true, F1: 676}
                            switch fi75 {
                            case 4294967295:
                                v6 = v6
                            case -2147483648:
                                fmt.Println(fi67)
                            }
                            v6 -= v6
                            _ = v83
                        }
                        v4 = v4
                        _ = v82
                    }
                    _ = v76
                    _ = v77
                }
                switch "a" {
                case "\n\t":
                    fmt.Println(fi75, v6)
                case "hello":
                    if (v6 != (v6 / ((-1) | 1))) {
                        v84 := v1[42]
                        if ((0 == 880) && false) {
                            fmt.Println(v4, v4, fi75)
                            fmt.Println(v6)
                        }
                        _ = v84
                    }
                case "":
                    v74 = (v74 % (((v74 << ((42) & 63))) | 1))
                }
                if (true || (((v4 / -2.25) - v4) != v4)) {
                    if (815 < 236) {
                        v85 := GenS9{F0: -32768}
                        _ = v85
                    }
                }
                if true {
                    var v86 int64 = -9223372036854775808
                    if ((0.0 - 1.5) < ((0.0 - 1e10) * (0.1 + (1.5 / 1e10)))) {
                        v6 = int32(v86)
                        fmt.Println(v6, v6, v74)
                        switch fi67 {
                        case 128:
                            switch "" {
                            case "foo bar":
                                if ((1e10 - -2.25) >= float32((1e10 * (float64(-2.25) * 1e10)))) {
                                    fmt.Println(v74, v74, v6)
                                    if true {
                                        var v87 uint8 = 255
                                        var v88 uint16 = 852
                                        _ = v87
                                        _ = v88
                                    }
                                }
                            case "":
                                fmt.Println(fi67, fi67)
                            }
                        case -1:
                            var v89 uint64 = 607
                            _ = v89
                        }
                    } else {
                        fmt.Println(fi67, v74, v86)
                        if (65535 > 2147483648) {
                            v90 := v5[fi67]
                            fmt.Println(v86, v4, v74)
                            var v91 float32 = 3.14159
                            _ = v90
                            _ = v91
                        } else {
                            var v92 int16 = 42
                            v86 = (^(v74))
                            var v93 int = (fi67 & genFn1(42, 0, 783))
                            _ = v92
                            _ = v93
                        }
                    }
                    _ = v86
                } else {
                    var v94 uint8 = 1
                    v4 -= float64(1.5)
                    _ = v94
                }
                var v95 uint16 = 255
                _ = fi75
                _ = v95
            }
            _ = fi67
            _ = v74
        }
    } else {
        for fi96 := 0; fi96 < 13; fi96++ {
            v97 := v5[fi96]
            fmt.Println(v97, v6)
            v97 = (fi96 ^ fi96)
            _ = fi96
            _ = v97
        }
        switch "a" {
        case "hello":
            var v98 int8 = 127
            _ = v98
        case "a":
            v4 += v4
        default:
            v99 := GenS3{F0: 964, F1: 0}
            _ = v99
        }
        if true {
            v100 := []int{776, -2147483648}
            _ = v100
        }
    }
    switch "foo bar" {
    case "a":
        fmt.Println(v6)
    case "foo bar":
        if (127 != 60) {
            var v101 int32 = -1
            v4 -= v4
            v6 -= v101
            _ = v101
        } else {
            v6 = (v6 - (v6 | (v6 % (((v6 / ((-20) | 1))) | 1))))
            v102 := GenS10{F0: 128, F1: 42, F2: 29}
            _ = v102
        }
    }
    v103 := GenS6{F0: -1, F1: 1000}
    v4 -= v4
    for fi104 := 0; fi104 < 13; fi104++ {
        v6 -= v6
        v105 := v3[int(v6)]
        v105 = (v105 || (true && (2 <= 18446744073709551615)))
        _ = fi104
        _ = v105
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v103
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
