package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint }
type GenS2 struct { F0 int64; F1 int32; F2 uint16 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 bool; F1 int }
type GenS5 struct { F0 uint16; F1 uint8 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 string; F1 int }
type GenS8 struct { F0 uint64 }
type GenS9 struct { F0 uint64; F1 uint16; F2 bool }
type GenS10 struct { F0 uint8; F1 uint16 }
func genFn0(p0 bool, p1 bool) int32 {
    if p1 {
        if p0 {
            switch "" {
            case "\n\t":
                v7 := []int{0, 4294967295, 0}
                _ = v7
            default:
                p0 = false
            }
            v8 := GenS1{F0: 2147483647, F1: 0}
            v9 := GenS1{F0: 67, F1: 255}
            _ = v8
            _ = v9
        } else {
            p1 = p0
        }
    } else {
        p0 = (-93 <= -32768)
    }
    var v10 string = (("\n\t" + "foo bar") + "")
    v11 := GenS2{F0: 0, F1: 42, F2: 65535}
    if p0 {
        var v12 int8 = 42
        if (17 < 42) {
            v13 := []int{32768}
            if len(v13) >= 2 {
                v13 = v13[1:len(v13)]
            }
            v10 = v10
            _ = v13
        } else {
            if (0 == 42) {
                var v14 int = 42
                _ = v14
            } else {
                v15 := []int{128}
                _ = v15
            }
            v16 := []int{-1}
            v17 := GenS3{F0: -2147483648}
            _ = v16
            _ = v17
        }
        _ = v12
    }
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    return 0
}
func genFn1() int32 {
    var v18 uint64 = 297
    var v19 uint = 1
    v20 := []int{}
    v19 = v19
    _ = v18
    _ = v19
    _ = v20
    return genFn0(true, (false || (false && ((((3.14159 * 0.0) - ((1e10 / 3.14159) + 1e10)) * 1e10) > 1e10))))
}

func main() {
    var v1 uint = 0
    var v2 string = "a"
    var v3 uint8 = 2
    v4 := []uint32{}
    var v5 int8 = 101
    v6 := make(map[uint32]bool)
    for fi21 := 0; fi21 < 9; fi21++ {
        v22 := GenS3{F0: fi21}
        switch "foo bar" {
        case "a":
            for fi23 := 0; fi23 < 10; fi23++ {
                var v24 int64 = 298
                v25 := GenS1{F0: 2, F1: (v1 % ((v1) | 1))}
                v3 -= v3
                _ = fi23
                _ = v24
                _ = v25
            }
        case "foo bar":
            v26 := GenS4{F0: false, F1: fi21}
            _ = v26
        default:
            if (!((570 != (fi21 ^ fi21)) && true)) {
                var v27 uint8 = (v3 + uint8(v1))
                if true {
                    switch (fi21 % (((fi21 << ((uint(v27)) & 63))) | 1)) {
                    case 2147483647:
                        v28 := GenS5{F0: 1, F1: (v3 & v3)}
                        _ = v28
                    default:
                        v2 = v2
                    }
                    for fi29 := 0; fi29 < 13; fi29++ {
                        v2 = ""
                        if (!true) {
                            var v30 uint = 255
                            v3 -= (v27 << (((v1 + (v30 >> ((v1) & 63)))) & 63))
                            _ = len(v6)
                            _ = v30
                        } else {
                            var v31 float64 = 0.1
                            switch ("foo bar" + v2) {
                            case "\n\t":
                                if true {
                                    var v32 uint8 = (v3 | (v27 ^ 1))
                                    _ = v32
                                } else {
                                    v27 = v3
                                    v2 = v2
                                }
                            case "hello":
                                _ = len(v4)
                            case "a":
                                var v33 uint64 = 0
                                _ = v33
                            default:
                                if len(v4) > 0 {
                                    v4[0] = 1
                                }
                            }
                            var v34 uint64 = 2
                            _ = v31
                            _ = v34
                        }
                        _ = fi29
                    }
                } else {
                    for fi35 := 0; fi35 < 3; fi35++ {
                        v3 = (^(v3))
                        fmt.Println(fi35)
                        _ = fi35
                    }
                }
                var v36 uint64 = 2
                _ = v27
                _ = v36
            } else {
                v37 := GenS6{F0: 112}
                if true {
                    v2 = v2
                    v38 := GenS4{F0: (!false), F1: int(v1)}
                    _ = len(v6)
                    _ = v38
                }
                switch v2 {
                case "\n\t":
                    for fi39 := 0; fi39 < 17; fi39++ {
                        var v40 uint16 = 0
                        switch (((v2 + ("\n\t" + ("" + ("hello" + v2)))) + "a") + v2) {
                        case "foo bar":
                            var v41 string = "\n\t"
                            _ = v41
                        case "":
                            _ = len(v6)
                        case "a":
                            v3 -= (v3 << ((v1) & 63))
                        }
                        _ = fi39
                        _ = v40
                    }
                }
                _ = v37
            }
        }
        _ = fi21
        _ = v22
    }
    for fi42 := 0; fi42 < 19; fi42++ {
        for fi43 := 0; fi43 < 17; fi43++ {
            v2 = ("\n\t" + v2)
            v44 := GenS7{F0: ("" + v2), F1: fi42}
            _ = fi43
            _ = v44
        }
        var v45 int8 = v5
        _ = fi42
        _ = v45
    }
    var v46 uint32 = 65535
    for fi47 := 0; fi47 < 6; fi47++ {
        for fi48 := 0; fi48 < 3; fi48++ {
            var v49 float32 = 3.14159
            switch (^(fi47)) {
            case 1:
                v1 += 42
            case 0:
                v49 = v49
            case -1:
                if false {
                    var v50 int8 = v5
                    var v51 uint = v1
                    v5 -= int8(fi48)
                    _ = v50
                    _ = v51
                }
            }
            _ = fi48
            _ = v49
        }
        _ = len(v6)
        _ = fi47
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v2 = ("\n\t" + v2)
    v52 := v6[v46]
    switch v2 {
    case "\n\t":
        var v53 bool = v52
        _ = v53
    case "":
        for fi54 := 0; fi54 < 13; fi54++ {
            if (-1 != -999) {
                switch v2 {
                case "":
                    switch fi54 {
                    case 127:
                        var v55 uint32 = v46
                        _ = v55
                    case 676:
                        for fi56 := 0; fi56 < 10; fi56++ {
                            var v57 int32 = 1
                            _ = fi56
                            _ = v57
                        }
                    }
                case "\n\t":
                    var v58 float32 = -2.25
                    _ = v58
                }
                v5 += (v5 >> ((v1) & 63))
                v5 += v5
            }
            _ = fi54
        }
    default:
        v59 := GenS6{F0: v46}
        _ = v59
    }
    switch -1 {
    case -59:
        for fi60 := 0; fi60 < 10; fi60++ {
            for fi61 := 0; fi61 < 10; fi61++ {
                if len(v4) > 0 {
                    v4[0] = 2147483648
                }
                switch 939 {
                case 32768:
                    var v62 float64 = ((0.1 + 0.0) - (1.5 / -2.25))
                    _ = v62
                default:
                    var v63 bool = false
                    _ = v63
                }
                var v64 float64 = 0.1
                switch fi60 {
                case -828:
                    switch (fi61 & fi60) {
                    case 1:
                        v65 := GenS8{F0: 598}
                        _ = v65
                    }
                    if len(v4) > 0 {
                        v4[0] = 2
                    }
                }
                _ = fi61
                _ = v64
            }
            _ = fi60
        }
    case 856:
        var v66 uint64 = 2
        _ = v66
    }
    if v52 {
        for fi67 := 0; fi67 < 1; fi67++ {
            for fi68 := 0; fi68 < 6; fi68++ {
                fmt.Println(v2, v46, v46)
                _ = fi68
            }
            fmt.Println(v52, v1, v5)
            var v69 int8 = int8(v1)
            _ = fi67
            _ = v69
        }
    } else {
        v5 = v5
        var v70 int8 = v5
        _ = v70
    }
    v2 = (v2 + v2)
    v5 -= v5
    switch v2 {
    case "\n\t":
        var v71 int16 = -23
        _ = v71
    case "":
        v46 -= (v46 % ((v46) | 1))
    case "hello":
        v1 = 549
    default:
        v72 := GenS9{F0: 2, F1: 309, F2: (!((18446744073709551615 > 0) || (v46 != (^(v46)))))}
        _ = v72
    }
    if false {
        v5 = v5
        _ = len(v4)
        var v73 int = 168
        _ = v73
    } else {
        for fi74 := 0; fi74 < 9; fi74++ {
            for fi75 := 0; fi75 < 7; fi75++ {
                var v76 uint32 = (^(v46))
                switch (-(fi74)) {
                case 4294967295:
                    switch fi75 {
                    case 438:
                        var v77 uint8 = 127
                        _ = v77
                    case -932:
                        var v78 uint8 = v3
                        _ = v78
                    default:
                        var v79 uint8 = 2
                        _ = v79
                    }
                case 0:
                    var v80 int16 = -104
                    _ = v80
                case 1:
                    v81 := GenS3{F0: fi75}
                    _ = v81
                default:
                    if v52 {
                        var v82 string = v2
                        var v83 float64 = 1.5
                        _ = v82
                        _ = v83
                    } else {
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        fmt.Println(v46, v76)
                        v84 := GenS2{F0: 0, F1: 1, F2: 2}
                        if v52 {
                            v5 -= v5
                        }
                        _ = v84
                    }
                }
                _ = fi75
                _ = v76
            }
            if len(v4) > 0 {
                v4[0] = v46
            }
            switch v2 {
            case "foo bar":
                if len(v4) > 0 {
                    v4[0] = 510
                }
            default:
                var v85 uint = (v1 + (^(v1)))
                _ = v85
            }
            _ = fi74
        }
        v2 = (("hello" + ("" + "")) + v2)
    }
    v52 = (1 == 0)
    for fi86 := 0; fi86 < 5; fi86++ {
        if true {
            var v87 int = fi86
            for fi88 := 0; fi88 < 19; fi88++ {
                v5 = v5
                var v89 int8 = v5
                var v90 uint = v1
                var v91 uint = v1
                _ = fi88
                _ = v89
                _ = v90
                _ = v91
            }
            _ = v87
        }
        v92 := GenS10{F0: (v3 / (((v3 - 2)) | 1)), F1: 2}
        var v93 string = (v2 + v2)
        _ = fi86
        _ = v92
        _ = v93
    }
    switch "hello" {
    case "foo bar":
        var v94 uint8 = 136
        _ = v94
    case "":
        var v95 uint16 = 2
        _ = v95
    case "hello":
        v4 = append(v4, v46)
    default:
        v1 += (v1 & (v1 * (v1 * (v1 >> ((v1) & 63)))))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    fmt.Println("v52:", v52)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v5
    _ = v52
    _ = v6
}
