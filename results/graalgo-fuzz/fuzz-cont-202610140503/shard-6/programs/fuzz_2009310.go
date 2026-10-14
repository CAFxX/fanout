package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8; F2 uint32 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint64; F1 int32; F2 int }
type GenS5 struct { F0 string; F1 int }
type GenS6 struct { F0 uint8; F1 uint8 }
type GenS7 struct { F0 uint8; F1 string; F2 int8 }
type GenS8 struct { F0 string; F1 int64 }
type GenS9 struct { F0 uint32; F1 uint32; F2 uint64 }
type GenS10 struct { F0 int16; F1 uint32 }
type GenS11 struct { F0 uint; F1 uint8 }

func main() {
    var v1 int8 = -37
    v2 := make(map[uint32]string)
    v3 := make(map[uint32]int)
    v4 := make(map[uint32]bool)
    var v5 uint32 = 22
    var v6 uint = 42
    var v7 float64 = 1.5
    v6 = (v6 << (((v6 / ((4294967295) | 1))) & 63))
    if false {
        switch -9223372036854775808 {
        case -1:
            v7 += 0.1
        case 1:
            if (127 < 223) {
                var v8 uint16 = 32767
                _ = v8
            } else {
                var v9 int32 = 0
                v1 -= (v1 / (((^(v1))) | 1))
                _ = v9
            }
        case 42:
            var v10 float32 = float32(((v7 + v7) + v7))
            _ = v10
        default:
            v5 = v5
        }
    }
    var v11 uint64 = 18446744073709551615
    switch 1 {
    case 32768:
        if false {
            var v12 uint16 = 2
            for fi13 := 0; fi13 < 20; fi13++ {
                for fi14 := 0; fi14 < 5; fi14++ {
                    v7 += (v7 + (0.1 / v7))
                    switch 32768 {
                    case 2147483647:
                        _ = len(v3)
                    case 32768:
                        fmt.Println(v5)
                    default:
                        v1 += v1
                    }
                    v15 := GenS1{F0: (fi13 % ((fi13) | 1)), F1: 127, F2: v5}
                    _ = fi14
                    _ = v15
                }
                v16 := GenS1{F0: (fi13 + (fi13 & (fi13 >> (((v6 - v6)) & 63)))), F1: 128, F2: uint32(v12)}
                _ = fi13
                _ = v16
            }
            v7 = v7
            _ = v12
        }
    }
    fmt.Println(v7)
    v17 := GenS2{F0: v11}
    if (!true) {
        switch -925 {
        case 32767:
            var v18 float64 = float64(1.5)
            _ = v18
        case 42:
            for fi19 := 0; fi19 < 18; fi19++ {
                if true {
                    fmt.Println(v7, v5, v1)
                    fmt.Println(fi19, v7)
                }
                v7 = (float64((((-2.25 - float32(v7)) + (float32(float64(((((((float32(v7) - (0.0 + 0.0)) / 1.5) - 0.1) - (0.1 / 3.14159)) - float32(v7)) - 0.1))) - 1e10)) + 1.5)) - ((0.1 / v7) + (3.14159 - (v7 / 3.14159))))
                if ((-(v1)) < v1) {
                    var v20 float32 = -2.25
                    switch "\n\t" {
                    case "foo bar":
                        v21 := GenS1{F0: -312, F1: 1, F2: v5}
                        _ = v21
                    case "a":
                        var v22 int32 = -5
                        _ = v22
                    default:
                        switch (fi19 & 32768) {
                        case 127:
                            v11 = v11
                        case 1:
                            v1 = (v1 ^ 42)
                        case 128:
                            if (4294967295 <= 2) {
                                if (3.14159 == 1.5) {
                                    for fi23 := 0; fi23 < 20; fi23++ {
                                        switch "\n\t" {
                                        case "":
                                            v20 -= 1.5
                                        case "a":
                                            if false {
                                                var v24 float64 = v7
                                                var v25 int32 = -2147483648
                                                _ = v24
                                                _ = v25
                                            }
                                        case "foo bar":
                                            fmt.Println(v11)
                                        default:
                                            var v26 float64 = (v7 + (float64(v20) - float64(float32(((v7 * 3.14159) * (v7 / (v7 * 0.0)))))))
                                            _ = v26
                                        }
                                        v11 = v11
                                        _ = fi23
                                    }
                                    v20 = ((v20 + v20) / -2.25)
                                } else {
                                    v7 = v7
                                    for fi27 := 0; fi27 < 10; fi27++ {
                                        if (42 <= 0) {
                                            var v28 uint32 = (v5 & 581)
                                            v20 = (v20 - v20)
                                            var v29 uint64 = (^(v11))
                                            _ = v28
                                            _ = v29
                                        }
                                        v30 := GenS3{F0: 65535}
                                        _ = fi27
                                        _ = v30
                                    }
                                    for fi31 := 0; fi31 < 16; fi31++ {
                                        if (0 >= 959) {
                                            if true {
                                                v32 := []int{fi31, fi31}
                                                _ = v32
                                            } else {
                                                fmt.Println(v1)
                                                v33 := []int{}
                                                _ = v33
                                            }
                                        }
                                        var v34 bool = (v7 < ((0.1 + v7) / (v7 / 0.1)))
                                        v35 := GenS4{F0: (v11 & v11), F1: 0, F2: (fi31 | (fi19 ^ (fi31 % ((-9223372036854775808) | 1))))}
                                        _ = fi31
                                        _ = v34
                                        _ = v35
                                    }
                                }
                                v20 += (3.14159 + (3.14159 - 0.0))
                                v36 := GenS2{F0: v11}
                                _ = v36
                            }
                        }
                    }
                    _ = v20
                } else {
                    for fi37 := 0; fi37 < 11; fi37++ {
                        v38 := GenS5{F0: "a", F1: (-(fi19))}
                        _ = fi37
                        _ = v38
                    }
                    v6 += 1
                }
                _ = fi19
            }
        case -9223372036854775808:
            v6 += v6
        default:
            fmt.Println(v7)
        }
        for fi39 := 0; fi39 < 1; fi39++ {
            var v40 uint32 = v5
            v41 := []int{fi39}
            v11 = (v11 << ((v6) & 63))
            _ = fi39
            _ = v40
            _ = v41
        }
        for fi42 := 0; fi42 < 2; fi42++ {
            var v43 uint8 = 255
            _ = fi42
            _ = v43
        }
    }
    fmt.Println(v1, v5)
    var v44 int = 0
    v44 = v44
    switch "a" {
    case "hello":
        var v45 string = "\n\t"
        _ = v45
    case "foo bar":
        fmt.Println(v7, v44)
    case "a":
        v6 = uint(v5)
    }
    if true {
        var v46 int8 = (^(v1))
        var v47 uint32 = (^(v5))
        _ = v46
        _ = v47
    } else {
        v7 -= v7
        switch (("" + "") + "foo bar") {
        case "hello":
            for fi48 := 0; fi48 < 4; fi48++ {
                if (!(v11 > (v11 % ((1) | 1)))) {
                    if ((-2.25 > 3.14159) || true) {
                        if (1 != 1) {
                            for fi49 := 0; fi49 < 14; fi49++ {
                                switch 263 {
                                case 2147483647:
                                    v1 -= (v1 | v1)
                                case 408:
                                    var v50 int32 = 658
                                    _ = v50
                                }
                                switch v44 {
                                case 368:
                                    var v51 int64 = 1
                                    _ = v51
                                case -190:
                                    var v52 uint16 = 0
                                    _ = v52
                                case -470:
                                    v53 := v4[(^(v5))]
                                    switch 207 {
                                    case 127:
                                        v54 := GenS6{F0: 233, F1: 0}
                                        _ = v54
                                    case -9223372036854775808:
                                        v55 := []int{(^(fi48)), (fi48 % ((fi49) | 1)), 32768}
                                        _ = v55
                                    }
                                    _ = v53
                                default:
                                    if false {
                                        switch 32768 {
                                        case 2147483647:
                                            v7 -= v7
                                        }
                                        _ = len(v3)
                                        if false {
                                            if true {
                                                var v56 string = "hello"
                                                var v57 uint16 = 65535
                                                var v58 float32 = float32(3.14159)
                                                fmt.Println(v44, v7)
                                                var v59 uint16 = v57
                                                _ = v56
                                                _ = v57
                                                _ = v58
                                                _ = v59
                                            }
                                        } else {
                                            var v60 uint16 = 113
                                            var v61 float32 = ((((0.1 * ((1e10 + 3.14159) / 3.14159)) - 0.0) - 0.1) - -2.25)
                                            _ = v60
                                            _ = v61
                                        }
                                    }
                                }
                                _ = fi49
                            }
                        } else {
                            if false {
                                var v62 int = (^(v44))
                                _ = v62
                            }
                            if false {
                                v44 = (fi48 | fi48)
                            }
                            v63 := GenS7{F0: 76, F1: "foo bar", F2: v1}
                            _ = v63
                        }
                        v64 := GenS3{F0: 255}
                        _ = v64
                    }
                    v65 := v3[(v5 >> ((v6) & 63))]
                    _ = v65
                } else {
                    var v66 uint = (v6 << (((v6 + (v6 + (v6 / ((255) | 1))))) & 63))
                    if ((0.0 - (((((3.14159 + float32(v7)) + 0.0) / 1.5) / 1.0) + 0.0)) < float32((v7 - (v7 / (v7 * 3.14159))))) {
                        v67 := GenS7{F0: 127, F1: ("" + (("a" + "a") + "hello")), F2: (v1 - v1)}
                        _ = v67
                    }
                    _ = v66
                }
                _ = fi48
            }
        case "foo bar":
            for fi68 := 0; fi68 < 20; fi68++ {
                var v69 uint = v6
                for fi70 := 0; fi70 < 18; fi70++ {
                    var v71 float64 = v7
                    _ = fi70
                    _ = v71
                }
                if true {
                    switch "foo bar" {
                    case "a":
                        if (128 > 2) {
                            var v72 uint = 511
                            _ = v72
                        }
                    case "hello":
                        if (0 <= -311) {
                            if true {
                                var v73 uint32 = (v5 ^ 758)
                                _ = v73
                            } else {
                                var v74 int8 = v1
                                for fi75 := 0; fi75 < 7; fi75++ {
                                    fmt.Println(v69, v69, v74)
                                    v7 += v7
                                    switch fi75 {
                                    case 9223372036854775807:
                                        var v76 uint8 = 127
                                        _ = v76
                                    case 128:
                                        v77 := GenS8{F0: ("" + ("hello" + ((("a" + "") + (("hello" + "hello") + (("foo bar" + "hello") + "a"))) + (("\n\t" + "foo bar") + (("" + "") + ("\n\t" + "a")))))), F1: int64(v11)}
                                        _ = v77
                                    case 4294967295:
                                        v78 := GenS7{F0: 127, F1: (("a" + "foo bar") + ("\n\t" + "a")), F2: v1}
                                        _ = v78
                                    default:
                                        if (155 >= 2) {
                                            if true {
                                                fmt.Println(v5)
                                            } else {
                                                v44 += v44
                                                v74 = 100
                                            }
                                            v79 := GenS9{F0: (v5 & (v5 / ((2) | 1))), F1: 4294967295, F2: 115}
                                            fmt.Println(v74, v11)
                                            switch -2147483648 {
                                            case 799:
                                                v80 := GenS8{F0: "foo bar", F1: 0}
                                                _ = v80
                                            case 32767:
                                                switch (("hello" + "a") + "hello") {
                                                case "a":
                                                    v74 = v1
                                                case "":
                                                    var v81 uint = (^(v6))
                                                    _ = v81
                                                }
                                            default:
                                                switch (v44 << ((729) & 63)) {
                                                case 1:
                                                    v5 += 0
                                                }
                                            }
                                            if true {
                                                v82 := GenS10{F0: -1, F1: uint32(v11)}
                                                delete(v4, 2)
                                                _ = v82
                                            } else {
                                                var v83 float32 = 3.14159
                                                v7 -= 1e10
                                                if ((v11 | (v11 % ((v11) | 1))) >= v11) {
                                                    v84 := []int{(fi75 >> (((v69 + v69)) & 63)), (^(fi68)), fi68}
                                                    v83 -= float32((((0.1 - (v7 - v7)) + v7) / (v7 + v7)))
                                                    fmt.Println(fi68, v83, fi75)
                                                    var v85 uint = (^(v69))
                                                    _ = v84
                                                    _ = v85
                                                }
                                                _ = v83
                                            }
                                            v86 := []int{(fi68 * -809), fi75, fi68, (-(fi68))}
                                            v87 := GenS11{F0: (v6 + v69), F1: 168}
                                            var v88 uint = v6
                                            _ = v79
                                            _ = v86
                                            _ = v87
                                            _ = v88
                                        }
                                        v5 -= 65535
                                        var v89 uint = (v69 % (((v6 + v6)) | 1))
                                        _ = v89
                                    }
                                    _ = fi75
                                }
                                for fi90 := 0; fi90 < 11; fi90++ {
                                    fmt.Println(v6)
                                    var v91 uint16 = 0
                                    delete(v2, uint32(v11))
                                    v92 := []int{(fi68 & fi90), (v44 ^ fi68), (fi90 & 2147483647), (v44 & int(v5))}
                                    _ = fi90
                                    _ = v91
                                    _ = v92
                                }
                                _ = v74
                            }
                        } else {
                            for fi93 := 0; fi93 < 20; fi93++ {
                                var v94 float32 = -2.25
                                _ = fi93
                                _ = v94
                            }
                        }
                    }
                    for fi95 := 0; fi95 < 14; fi95++ {
                        v96 := []int{}
                        _ = fi95
                        _ = v96
                    }
                } else {
                    if false {
                        v97 := []int{-9223372036854775808, fi68}
                        _ = len(v97)
                        _ = v97
                    }
                }
                _ = fi68
                _ = v69
            }
        }
    }
    var v98 int = v44
    for fi99 := 0; fi99 < 13; fi99++ {
        var v100 int = 9223372036854775807
        if (!(!false)) {
            for fi101 := 0; fi101 < 14; fi101++ {
                var v102 int32 = -1
                _ = fi101
                _ = v102
            }
            v1 = (-(v1))
            v103 := GenS7{F0: 1, F1: "foo bar", F2: 100}
            _ = v103
        } else {
            v6 = 255
        }
        _ = fi99
        _ = v100
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v98:", v98)
    _ = v1
    _ = v11
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v5
    _ = v6
    _ = v7
    _ = v98
}
