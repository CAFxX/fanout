package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32 }
type GenS2 struct { F0 string; F1 uint32 }
type GenS3 struct { F0 uint64; F1 bool }
type GenS4 struct { F0 int16; F1 string }
type GenS5 struct { F0 int32; F1 int8; F2 uint16 }
type GenS6 struct { F0 int8; F1 int }
type GenS7 struct { F0 int16; F1 string; F2 bool }
type GenS8 struct { F0 int; F1 string }
type GenS9 struct { F0 int32; F1 uint64; F2 int32 }
type GenS10 struct { F0 string }
type GenS11 struct { F0 int }
func genFn0(p0 int64, p1 uint64) uint8 {
    if (!((p0 + p0) > -111)) {
        v5 := GenS1{F0: 877, F1: 2147483648}
        p1 += p1
        _ = v5
    } else {
        p1 += p1
    }
    for fi6 := 0; fi6 < 10; fi6++ {
        switch "foo bar" {
        case "hello":
            v7 := []int{1, (^(fi6)), (fi6 | fi6)}
            _ = v7
        }
        var v8 int32 = -132
        _ = fi6
        _ = v8
    }
    fmt.Println(p0)
    var v9 bool = (-1 != 0)
    var v10 uint8 = 1
    _ = p0
    _ = p1
    _ = v10
    _ = v9
    return v10
}
func genFn1() uint {
    if (76 <= 1) {
        for fi11 := 0; fi11 < 12; fi11++ {
            var v12 uint = 1
            _ = fi11
            _ = v12
        }
        var v13 int64 = 753
        for fi14 := 0; fi14 < 18; fi14++ {
            for fi15 := 0; fi15 < 3; fi15++ {
                switch fi15 {
                case 630:
                    fmt.Println(v13, v13)
                }
                _ = fi15
            }
            v13 = v13
            var v16 int8 = -33
            _ = fi14
            _ = v16
        }
        _ = v13
    }
    var v17 uint16 = 1
    for fi18 := 0; fi18 < 15; fi18++ {
        switch "\n\t" {
        case "hello":
            fmt.Println(v17)
        case "\n\t":
            var v19 int16 = 32767
            _ = v19
        }
        var v20 bool = true
        _ = fi18
        _ = v20
    }
    switch 42 {
    case -369:
        var v21 int64 = -1
        _ = v21
    default:
        v22 := GenS1{F0: 305, F1: 744}
        _ = v22
    }
    var v23 uint64 = 467
    _ = v17
    _ = v23
    return 42
}

func main() {
    var v1 int32 = 373
    var v2 uint = 2
    var v3 int = 54
    var v4 float64 = -2.25
    var v24 int8 = 16
    if true {
        if true {
            for fi25 := 0; fi25 < 15; fi25++ {
                var v26 int16 = -980
                v26 = v26
                var v27 int32 = v1
                _ = fi25
                _ = v26
                _ = v27
            }
            v28 := []int{v3, (^(v3)), (v3 >> (((v2 - (v2 % (((v2 ^ genFn1())) | 1)))) & 63))}
            for fi29 := 0; fi29 < 3; fi29++ {
                var v30 int8 = 127
                _ = fi29
                _ = v30
            }
            _ = v28
        } else {
            v4 = v4
        }
        v31 := []int{}
        _ = v31
    } else {
        for fi32 := 0; fi32 < 11; fi32++ {
            switch "hello" {
            case "hello":
                for fi33 := 0; fi33 < 5; fi33++ {
                    v24 = v24
                    v2 -= v2
                    _ = fi33
                }
            case "":
                switch int(v24) {
                case -215:
                    for fi34 := 0; fi34 < 13; fi34++ {
                        v1 = v1
                        v35 := GenS2{F0: "foo bar", F1: 0}
                        _ = fi34
                        _ = v35
                    }
                default:
                    if (1 != 1) {
                        v36 := GenS1{F0: 65535, F1: 65535}
                        var v37 int64 = 1
                        v24 = v24
                        _ = v36
                        _ = v37
                    } else {
                        var v38 float64 = -2.25
                        _ = v38
                    }
                }
            case "foo bar":
                v39 := []int{(v3 ^ (v3 % ((fi32) | 1))), (v3 << ((v2) & 63))}
                _ = v39
            default:
                fmt.Println(v1)
            }
            for fi40 := 0; fi40 < 9; fi40++ {
                v4 -= 1.5
                if (-1 > 9223372036854775807) {
                    v24 -= (^(v24))
                    fmt.Println(v4, v4)
                    fmt.Println(fi40, v2)
                } else {
                    if true {
                        var v41 int = fi40
                        var v42 int = fi40
                        v43 := GenS1{F0: 1, F1: 894}
                        _ = v41
                        _ = v42
                        _ = v43
                    } else {
                        fmt.Println(v2, v4, v24)
                        switch "foo bar" {
                        case "\n\t":
                            v4 += (float64((0.0 / 1.0)) * -2.25)
                        case "foo bar":
                            var v44 float64 = (v4 / v4)
                            _ = v44
                        case "a":
                            if false {
                                var v45 uint16 = 42
                                _ = v45
                            }
                        }
                    }
                    var v46 float64 = (v4 - float64(0.1))
                    if (((true || true) || (!true)) || true) {
                        if (true && true) {
                            v47 := GenS1{F0: 42, F1: 65535}
                            v3 += 128
                            v48 := []int{127}
                            _ = v47
                            _ = v48
                        }
                        if true {
                            v49 := GenS1{F0: 4294967295, F1: uint32(v2)}
                            _ = v49
                        } else {
                            fmt.Println(v1, v3, v3)
                            var v50 bool = (-9223372036854775808 < 9223372036854775807)
                            fmt.Println(v50, v24, fi32)
                            _ = v50
                        }
                    } else {
                        v4 -= v4
                        if false {
                            v51 := GenS3{F0: 42, F1: false}
                            _ = v51
                        }
                    }
                    _ = v46
                }
                v4 = (((((v4 / v4) / v4) * v4) / v4) - v4)
                fmt.Println(fi40, v4)
                switch ("a" + "a") {
                case "hello":
                    if true {
                        if false {
                            v24 += v24
                        }
                        if ((v24 - (v24 & v24)) != v24) {
                            var v52 string = "a"
                            v53 := []int{fi40, (fi40 ^ int(v24)), fi32, -2147483648}
                            switch fi40 {
                            case 842:
                                switch (fi40 >> ((v2) & 63)) {
                                case 2147483647:
                                    v1 = (-(v1))
                                case 42:
                                    var v54 int32 = (v1 ^ (v1 & -685))
                                    _ = v54
                                case 9223372036854775807:
                                    v24 = -1
                                default:
                                    switch (^(v3)) {
                                    case -188:
                                        switch (^(v3)) {
                                        case 32767:
                                            switch fi40 {
                                            case 2147483647:
                                                v52 = v52
                                            }
                                        case -663:
                                            v1 = 2147483647
                                        case 391:
                                            var v55 uint8 = genFn0(9223372036854775807, 2)
                                            _ = v55
                                        default:
                                            if true {
                                                v2 += genFn1()
                                                if (fi32 < fi32) {
                                                    if false {
                                                        fmt.Println(v3)
                                                        if len(v53) >= 2 {
                                                            v53 = v53[1:len(v53)]
                                                        }
                                                    }
                                                    v56 := GenS3{F0: 0, F1: (-30 == -450)}
                                                    v24 = (v24 + (v24 & (^(v24))))
                                                    _ = v56
                                                } else {
                                                    if ((4294967295 < 151) && true) {
                                                        switch (v3 >> ((283) & 63)) {
                                                        case 32768:
                                                            if (1e10 == v4) {
                                                                v52 = v52
                                                            } else {
                                                                switch fi40 {
                                                                case -2147483648:
                                                                    if len(v53) > 0 {
                                                                        v53[0] = -222
                                                                    }
                                                                }
                                                                if len(v53) >= 2 {
                                                                    v53 = v53[1:len(v53)]
                                                                }
                                                            }
                                                        case 0:
                                                            v24 = 1
                                                        case 32767:
                                                            var v57 int8 = v24
                                                            _ = v57
                                                        default:
                                                            v3 = (fi40 + (fi32 * (fi40 ^ 1)))
                                                        }
                                                    } else {
                                                        var v58 uint = genFn1()
                                                        var v59 int8 = v24
                                                        _ = v58
                                                        _ = v59
                                                    }
                                                }
                                                fmt.Println(v3)
                                            }
                                        }
                                    }
                                }
                            case 2147483647:
                                _ = len(v53)
                                if len(v53) >= 2 {
                                    v53 = v53[1:len(v53)]
                                }
                            }
                            switch v52 {
                            case "a":
                                v53 = append(v53, fi32)
                            case "hello":
                                fmt.Println(fi32)
                            case "foo bar":
                                if (609 <= 2147483647) {
                                    v53 = append(v53, (fi32 | fi40))
                                    v60 := GenS4{F0: -403, F1: v52}
                                    v61 := GenS5{F0: v1, F1: (-(v24)), F2: 1}
                                    _ = v60
                                    _ = v61
                                } else {
                                    fmt.Println(v24, v2, fi32)
                                    v62 := GenS4{F0: 0, F1: v52}
                                    if true {
                                        fmt.Println(v4, fi32)
                                        switch fi40 {
                                        case 0:
                                            v63 := GenS4{F0: 562, F1: (v52 + "a")}
                                            _ = v63
                                        case 1:
                                            v1 = (^(v1))
                                        case -593:
                                            if ((v24 ^ (v24 / (((v24 ^ 89)) | 1))) != (v24 & v24)) {
                                                if (!(!(!(42 < -9223372036854775808)))) {
                                                    fmt.Println(fi40, fi32)
                                                    if len(v53) >= 2 {
                                                        v53 = v53[1:len(v53)]
                                                    }
                                                    var v64 int16 = 1
                                                    _ = v64
                                                }
                                                _ = len(v53)
                                            }
                                        }
                                        fmt.Println(v4, v4)
                                        v65 := GenS4{F0: 1, F1: v52}
                                        _ = v65
                                    } else {
                                        fmt.Println(v1, fi40)
                                        if len(v53) >= 2 {
                                            v53 = v53[1:len(v53)]
                                        }
                                    }
                                    _ = v62
                                }
                            default:
                                var v66 uint = 2
                                _ = v66
                            }
                            fmt.Println(v3)
                            v67 := GenS6{F0: (v24 + (v24 ^ v24)), F1: fi40}
                            v1 = v1
                            _ = v52
                            _ = v53
                            _ = v67
                        } else {
                            if true {
                                v68 := []int{32767}
                                v4 = v4
                                v69 := GenS6{F0: v24, F1: (^(fi32))}
                                _ = v68
                                _ = v69
                            }
                            fmt.Println(v2, v1, fi40)
                            var v70 int32 = (-(v1))
                            _ = v70
                        }
                        var v71 int32 = v1
                        _ = v71
                    }
                case "a":
                    v72 := GenS2{F0: "hello", F1: 0}
                    _ = v72
                default:
                    var v73 uint32 = 545
                    _ = v73
                }
                _ = fi40
            }
            _ = fi32
        }
        if false {
            v74 := GenS7{F0: 821, F1: "foo bar", F2: true}
            if false {
                if true {
                    for fi75 := 0; fi75 < 19; fi75++ {
                        v76 := GenS8{F0: (v3 + -609), F1: ("hello" + (("\n\t" + ("hello" + "")) + ""))}
                        v77 := []int{}
                        switch (fi75 + (fi75 - (fi75 - (v3 * (fi75 >> ((v2) & 63)))))) {
                        case 381:
                            _ = len(v77)
                        case 0:
                            fmt.Println(v24, v3)
                        case 42:
                            var v78 bool = false
                            _ = v78
                        }
                        _ = fi75
                        _ = v76
                        _ = v77
                    }
                }
                v79 := GenS1{F0: 0, F1: 2147483647}
                var v80 float32 = 3.14159
                _ = v79
                _ = v80
            }
            var v81 float64 = v4
            _ = v74
            _ = v81
        } else {
            var v82 uint64 = 42
            _ = v82
        }
        switch -985 {
        case -2147483648:
            v83 := []int{(v3 % ((-9223372036854775808) | 1)), 1}
            _ = v83
        case -509:
            v84 := GenS6{F0: 1, F1: v3}
            _ = v84
        case 0:
            var v85 int64 = 360
            _ = v85
        default:
            if false {
                v86 := []int{v3, (v3 & (v3 | v3)), (v3 ^ (v3 / (((v3 + (v3 * v3))) | 1))), v3}
                for fi87 := 0; fi87 < 20; fi87++ {
                    var v88 float32 = -2.25
                    for fi89 := 0; fi89 < 3; fi89++ {
                        var v90 int8 = (-(v24))
                        v3 = 1
                        _ = len(v86)
                        _ = fi89
                        _ = v90
                    }
                    _ = fi87
                    _ = v88
                }
                _ = v86
            }
        }
    }
    v91 := GenS7{F0: 1000, F1: "a", F2: true}
    var v92 int64 = -1
    v93 := GenS8{F0: (v3 % ((0) | 1)), F1: ""}
    v94 := GenS9{F0: -1, F1: 1, F2: (v1 | -2147483648)}
    v95 := GenS10{F0: ("hello" + ((("\n\t" + ("hello" + "")) + "hello") + "hello"))}
    v4 -= (1e10 + (v4 * v4))
    var v96 int64 = int64(v2)
    var v97 int64 = -9223372036854775808
    v98 := []int{v3, -2147483648, 580}
    v4 -= 1.5
    if true {
        var v99 uint64 = 1
        _ = v99
    } else {
        switch "foo bar" {
        case "":
            if len(v98) > 0 {
                v98[0] = (v3 << ((v2) & 63))
            }
        }
        v100 := GenS6{F0: 0, F1: v3}
        _ = v100
    }
    if len(v98) >= 2 {
        v98 = v98[1:len(v98)]
    }
    v101 := GenS11{F0: v3}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v92:", v92)
    fmt.Println("v96:", v96)
    fmt.Println("v97:", v97)
    fmt.Println("slicelen v98:", len(v98), v98)
    _ = v1
    _ = v101
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v91
    _ = v92
    _ = v93
    _ = v94
    _ = v95
    _ = v96
    _ = v97
    _ = v98
}
