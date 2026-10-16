package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint16; F1 int16; F2 int }
type GenS3 struct { F0 int32; F1 int64; F2 int16 }
type GenS4 struct { F0 bool; F1 uint32; F2 int16 }
type GenS5 struct { F0 int32; F1 int16; F2 int16 }
type GenS6 struct { F0 bool; F1 uint8; F2 int8 }
type GenS7 struct { F0 int }
type GenS8 struct { F0 bool; F1 int }
type GenS9 struct { F0 uint32; F1 uint8 }
type GenS10 struct { F0 int64; F1 bool }
type GenS11 struct { F0 uint8 }
type GenS12 struct { F0 uint32 }
type GenS13 struct { F0 uint32 }
type GenS14 struct { F0 int16 }

func main() {
    v1 := make(map[int]int)
    v2 := make(map[int]bool)
    var v3 int16 = 32767
    v4 := make(map[uint32]int)
    v5 := make(map[int]uint32)
    v3 = -458
    if true {
        v3 += 1
        v3 -= v3
    }
    if (!false) {
        if true {
            for fi6 := 0; fi6 < 14; fi6++ {
                if false {
                    if true {
                        if false {
                            v7 := GenS1{F0: 1}
                            fmt.Println(v3, fi6)
                            v8 := GenS2{F0: 674, F1: (v3 - v3), F2: (fi6 / (((fi6 ^ fi6)) | 1))}
                            _ = v7
                            _ = v8
                        }
                        switch fi6 {
                        case 42:
                            var v9 int32 = 42
                            _ = v9
                        case -952:
                            v3 = v3
                        default:
                            if (((!true) || true) || true) {
                                for fi10 := 0; fi10 < 9; fi10++ {
                                    v3 = (^(v3))
                                    _ = fi10
                                }
                                var v11 int8 = -1
                                _ = v11
                            }
                        }
                        var v12 int8 = 82
                        _ = v12
                    }
                } else {
                    for fi13 := 0; fi13 < 7; fi13++ {
                        v14 := []int{fi13}
                        v3 -= v3
                        v15 := GenS3{F0: 42, F1: -113, F2: (v3 << ((1) & 63))}
                        v3 += (v3 & (v3 + (v3 >> ((18446744073709551615) & 63))))
                        v3 = v3
                        _ = fi13
                        _ = v14
                        _ = v15
                    }
                    v16 := GenS4{F0: true, F1: 42, F2: v3}
                    _ = len(v2)
                    _ = v16
                }
                fmt.Println(v3)
                fmt.Println(v3)
                _ = fi6
            }
            v17 := GenS4{F0: (true || (0 >= 2)), F1: 2147483648, F2: (v3 & 1)}
            _ = v17
        } else {
            v3 -= (v3 & v3)
            fmt.Println(v3)
        }
    } else {
        v3 = v3
        if (!true) {
            if (2147483647 < -434) {
                var v18 float64 = 0.1
                _ = v18
            }
        } else {
            if true {
                switch -666 {
                case 192:
                    for fi19 := 0; fi19 < 5; fi19++ {
                        for fi20 := 0; fi20 < 17; fi20++ {
                            v21 := GenS5{F0: -1, F1: (v3 * v3), F2: 0}
                            _ = fi20
                            _ = v21
                        }
                        _ = fi19
                    }
                default:
                    switch -201 {
                    case -9223372036854775808:
                        switch "a" {
                        case "":
                            var v22 uint32 = 386
                            _ = v22
                        case "foo bar":
                            fmt.Println(v3)
                        }
                    case 0:
                        switch "a" {
                        case "hello":
                            for fi23 := 0; fi23 < 15; fi23++ {
                                fmt.Println(fi23)
                                fmt.Println(fi23, fi23)
                                if ("" > "") {
                                    v3 += 731
                                } else {
                                    for fi24 := 0; fi24 < 5; fi24++ {
                                        v25 := GenS1{F0: 575}
                                        if true {
                                            if (3.14159 != -2.25) {
                                                switch "\n\t" {
                                                case "foo bar":
                                                    v26 := []int{fi23}
                                                    _ = v26
                                                case "a":
                                                    v3 += 1
                                                default:
                                                    fmt.Println(v3, fi23)
                                                }
                                            }
                                            v3 += v3
                                            fmt.Println(fi23, fi24, v3)
                                        }
                                        v3 = v3
                                        _ = fi24
                                        _ = v25
                                    }
                                    v27 := GenS4{F0: (!true), F1: 838, F2: v3}
                                    v3 += (v3 - 0)
                                    _ = v27
                                }
                                _ = fi23
                            }
                        case "a":
                            fmt.Println(v3)
                        case "foo bar":
                            v28 := GenS6{F0: (-897 == -700), F1: 1, F2: -98}
                            _ = v28
                        }
                    default:
                        if false {
                            v29 := GenS7{F0: 42}
                            _ = v29
                        } else {
                            if true {
                                v30 := GenS8{F0: true, F1: 1}
                                _ = v30
                            } else {
                                var v31 int16 = v3
                                _ = v31
                            }
                            v3 += (v3 / ((v3) | 1))
                        }
                    }
                }
                v32 := []int{}
                for fi33 := 0; fi33 < 8; fi33++ {
                    if (2147483647 >= 740) {
                        for fi34 := 0; fi34 < 9; fi34++ {
                            var v35 uint8 = 42
                            v36 := GenS3{F0: -1, F1: 9223372036854775807, F2: 1}
                            _ = fi34
                            _ = v35
                            _ = v36
                        }
                        for fi37 := 0; fi37 < 4; fi37++ {
                            switch ("a" + "") {
                            case "a":
                                if false {
                                    var v38 uint32 = 42
                                    _ = v38
                                } else {
                                    var v39 uint = 6
                                    v40 := GenS7{F0: fi37}
                                    fmt.Println(v39)
                                    _ = v39
                                    _ = v40
                                }
                            case "hello":
                                switch ("hello" + "foo bar") {
                                case "hello":
                                    _ = len(v32)
                                default:
                                    v3 = v3
                                }
                            }
                            _ = fi37
                        }
                    } else {
                        if false {
                            v41 := GenS9{F0: 377, F1: 42}
                            for fi42 := 0; fi42 < 4; fi42++ {
                                if len(v32) > 0 {
                                    v32[0] = fi33
                                }
                                _ = fi42
                            }
                            _ = v41
                        } else {
                            v43 := GenS1{F0: 4294967295}
                            for fi44 := 0; fi44 < 17; fi44++ {
                                if (!false) {
                                    switch "\n\t" {
                                    case "foo bar":
                                        var v45 int = (fi44 ^ fi33)
                                        _ = v45
                                    case "":
                                        if ((!(941 < 2147483648)) && ((609 >= 4294967295) || (42 < 349))) {
                                            fmt.Println(fi44, v3)
                                            switch fi33 {
                                            case 32767:
                                                v3 += (v3 * v3)
                                            case -655:
                                                fmt.Println(fi44, fi33, fi44)
                                            default:
                                                if false {
                                                    v3 = -937
                                                } else {
                                                    fmt.Println(v3)
                                                }
                                            }
                                            var v46 uint8 = 222
                                            _ = v46
                                        } else {
                                            var v47 int8 = 127
                                            if true {
                                                v48 := GenS7{F0: (^(fi33))}
                                                v3 -= 1
                                                v3 = (v3 + v3)
                                                _ = v48
                                            }
                                            _ = v47
                                        }
                                    default:
                                        if (1 < 4294967295) {
                                            v3 = (v3 / ((v3) | 1))
                                        } else {
                                            if false {
                                                fmt.Println(fi44, fi44, fi33)
                                            } else {
                                                v3 += 42
                                            }
                                            var v49 float32 = 3.14159
                                            v3 -= (-(v3))
                                            var v50 bool = false
                                            v3 -= (v3 & v3)
                                            _ = v49
                                            _ = v50
                                        }
                                    }
                                }
                                var v51 int16 = v3
                                _ = fi44
                                _ = v51
                            }
                            v52 := v5[(^(fi33))]
                            _ = v43
                            _ = v52
                        }
                        var v53 string = ("hello" + ("foo bar" + "\n\t"))
                        _ = v53
                    }
                    switch "\n\t" {
                    case "hello":
                        v54 := GenS3{F0: 1, F1: 741, F2: (v3 * (v3 - v3))}
                        _ = v54
                    default:
                        var v55 uint32 = 921
                        _ = v55
                    }
                    v56 := GenS6{F0: (!true), F1: 1, F2: -57}
                    _ = fi33
                    _ = v56
                }
                _ = v32
            }
            if (!true) {
                if false {
                    var v57 uint = 15
                    v58 := GenS2{F0: 65535, F1: -46, F2: 744}
                    if (0.0 >= 0.0) {
                        var v59 float32 = (float32((0.1 * 0.0)) * 0.1)
                        _ = v59
                    } else {
                        switch 9223372036854775807 {
                        case 32767:
                            v3 -= v3
                        }
                        fmt.Println(v3, v57)
                    }
                    _ = v57
                    _ = v58
                } else {
                    v60 := GenS10{F0: 0, F1: true}
                    var v61 int16 = (v3 ^ (v3 ^ (v3 ^ (-(v3)))))
                    v3 += -464
                    _ = v60
                    _ = v61
                }
            } else {
                switch "\n\t" {
                case "\n\t":
                    v3 -= (v3 % ((v3) | 1))
                case "a":
                    switch 128 {
                    case 127:
                        v3 = 42
                    default:
                        for fi62 := 0; fi62 < 2; fi62++ {
                            v63 := GenS2{F0: 189, F1: v3, F2: 4294967295}
                            for fi64 := 0; fi64 < 10; fi64++ {
                                v3 = (v3 / ((0) | 1))
                                _ = fi64
                            }
                            _ = fi62
                            _ = v63
                        }
                    }
                default:
                    if (-792 < -266) {
                        v65 := []int{32768}
                        for fi66 := 0; fi66 < 7; fi66++ {
                            _ = len(v2)
                            switch ("\n\t" + "") {
                            case "hello":
                                switch int(v3) {
                                case -2147483648:
                                    v67 := v1[fi66]
                                    _ = v67
                                case 32767:
                                    if ("foo bar" == "\n\t") {
                                        if (v3 < v3) {
                                            var v68 int32 = -2147483648
                                            var v69 float64 = 1e10
                                            _ = v68
                                            _ = v69
                                        } else {
                                            if (42 == -2147483648) {
                                                if (fi66 == (fi66 * -848)) {
                                                    v65 = append(v65, fi66)
                                                } else {
                                                    fmt.Println(v3, fi66)
                                                    for fi70 := 0; fi70 < 2; fi70++ {
                                                        v3 += (^(v3))
                                                        fmt.Println(v3)
                                                        _ = fi70
                                                    }
                                                    if len(v65) >= 2 {
                                                        v65 = v65[1:len(v65)]
                                                    }
                                                }
                                                var v71 int32 = 1
                                                _ = v71
                                            } else {
                                                fmt.Println(v3)
                                                switch fi66 {
                                                case 42:
                                                    fmt.Println(fi66)
                                                case 1:
                                                    if true {
                                                        v72 := GenS11{F0: 2}
                                                        var v73 int64 = 0
                                                        for fi74 := 0; fi74 < 9; fi74++ {
                                                            _ = len(v65)
                                                            if true {
                                                                v73 += v73
                                                            }
                                                            _ = fi74
                                                        }
                                                        _ = v72
                                                        _ = v73
                                                    } else {
                                                        v75 := GenS9{F0: 1, F1: 0}
                                                        var v76 uint16 = 714
                                                        _ = v75
                                                        _ = v76
                                                    }
                                                case 32768:
                                                    for fi77 := 0; fi77 < 10; fi77++ {
                                                        var v78 int = fi66
                                                        _ = fi77
                                                        _ = v78
                                                    }
                                                default:
                                                    _ = len(v65)
                                                }
                                            }
                                            if false {
                                                if true {
                                                    for fi79 := 0; fi79 < 6; fi79++ {
                                                        fmt.Println(fi66, fi79)
                                                        _ = fi79
                                                    }
                                                } else {
                                                    switch (fi66 + fi66) {
                                                    case -1:
                                                        var v80 int64 = 500
                                                        _ = v80
                                                    case -2147483648:
                                                        var v81 int16 = (-(v3))
                                                        _ = v81
                                                    case 128:
                                                        v3 = (^(v3))
                                                    default:
                                                        var v82 uint64 = 1
                                                        _ = v82
                                                    }
                                                    v3 -= v3
                                                }
                                            }
                                        }
                                    } else {
                                        var v83 float64 = 0.1
                                        for fi84 := 0; fi84 < 8; fi84++ {
                                            v83 -= float64(3.14159)
                                            _ = fi84
                                        }
                                        var v85 float64 = float64((3.14159 + 1e10))
                                        _ = v83
                                        _ = v85
                                    }
                                case 431:
                                    if ((!false) || true) {
                                        v86 := GenS3{F0: 0, F1: 9223372036854775807, F2: v3}
                                        for fi87 := 0; fi87 < 20; fi87++ {
                                            var v88 uint32 = 938
                                            v88 += 65535
                                            _ = fi87
                                            _ = v88
                                        }
                                        for fi89 := 0; fi89 < 9; fi89++ {
                                            var v90 int8 = 42
                                            _ = fi89
                                            _ = v90
                                        }
                                        _ = v86
                                    } else {
                                        v2[-9223372036854775808] = true
                                    }
                                }
                            case "\n\t":
                                if true {
                                    v3 = 441
                                    if (0 >= 42) {
                                        for fi91 := 0; fi91 < 11; fi91++ {
                                            var v92 float64 = 1.5
                                            _ = fi91
                                            _ = v92
                                        }
                                        if false {
                                            switch (fi66 & fi66) {
                                            case 548:
                                                v3 -= v3
                                            case 9223372036854775807:
                                                if false {
                                                    v3 -= (v3 << ((65535) & 63))
                                                } else {
                                                    var v93 int16 = -843
                                                    for fi94 := 0; fi94 < 1; fi94++ {
                                                        var v95 uint = 18446744073709551615
                                                        v3 = (v93 % ((32767) | 1))
                                                        _ = fi94
                                                        _ = v95
                                                    }
                                                    _ = v93
                                                }
                                            default:
                                                v65 = append(v65, (fi66 * (-(fi66))))
                                            }
                                            var v96 int16 = v3
                                            _ = v96
                                        }
                                    }
                                } else {
                                    var v97 int = (fi66 * fi66)
                                    _ = v97
                                }
                            case "a":
                                v98 := GenS12{F0: 1}
                                _ = v98
                            default:
                                fmt.Println(v3, fi66)
                            }
                            if true {
                                v3 += (v3 % ((v3) | 1))
                            }
                            _ = fi66
                        }
                        _ = v65
                    } else {
                        v99 := GenS13{F0: 4294967295}
                        _ = v99
                    }
                }
            }
            v100 := GenS14{F0: v3}
            _ = v100
        }
        var v101 int8 = 127
        _ = v101
    }
    if (2 >= 2147483647) {
        fmt.Println(v3)
    }
    switch "" {
    case "hello":
        var v102 bool = ((v3 % ((-1) | 1)) >= (v3 + (v3 ^ v3)))
        _ = v102
    case "":
        fmt.Println(v3)
    default:
        var v103 int = 128
        _ = v103
    }
    v104 := GenS5{F0: 1, F1: (v3 / (((-(v3))) | 1)), F2: v3}
    var v105 float64 = 1.5
    v3 += (v3 * v3)
    v3 = (^(v3))
    v105 -= v105
    v3 -= v3
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v105:", v105)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v104
    _ = v105
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
