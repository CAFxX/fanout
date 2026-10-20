package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16 }
type GenS2 struct { F0 bool; F1 string }
type GenS3 struct { F0 uint8; F1 uint32; F2 int16 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int64; F1 int64 }
type GenS6 struct { F0 uint32; F1 int }
type GenS7 struct { F0 uint32; F1 int8; F2 uint32 }
type GenS8 struct { F0 uint32 }
func genFn0(p0 int32, p1 uint16) uint {
    if (-32768 < 1) {
        var v7 float32 = 1e10
        v8 := GenS1{F0: 232, F1: (^(p1))}
        for fi9 := 0; fi9 < 4; fi9++ {
            for fi10 := 0; fi10 < 19; fi10++ {
                switch fi9 {
                case 127:
                    v11 := GenS1{F0: 17, F1: p1}
                    _ = v11
                case 419:
                    if ((p0 * -1) != p0) {
                        if (p0 == (p0 % ((p0) | 1))) {
                            switch (("a" + "") + "\n\t") {
                            case "hello":
                                var v12 float64 = (1.5 * 0.1)
                                _ = v12
                            default:
                                var v13 int32 = -2147483648
                                _ = v13
                            }
                        }
                        fmt.Println(p1)
                        fmt.Println(p1)
                    } else {
                        v14 := GenS2{F0: (true && false), F1: (("hello" + "") + "foo bar")}
                        v15 := []int{-253, fi10, fi10, (-(fi10))}
                        _ = v14
                        _ = v15
                    }
                case 32768:
                    var v16 float64 = 3.14159
                    _ = v16
                default:
                    v17 := []int{(fi9 * fi9), fi10, 4294967295}
                    _ = v17
                }
                _ = fi10
            }
            for fi18 := 0; fi18 < 4; fi18++ {
                if ((fi9 + fi18) > fi18) {
                    v19 := []int{fi9, (fi18 - fi18), fi9}
                    switch "\n\t" {
                    case "a":
                        v7 += 0.0
                    case "":
                        var v20 int16 = 1
                        _ = v20
                    default:
                        var v21 uint32 = 767
                        _ = v21
                    }
                    _ = v19
                }
                fmt.Println(p0, p0, fi18)
                if (!(1 < -127)) {
                    var v22 uint32 = 776
                    v23 := []int{(fi9 + fi9), 441, 4294967295, (fi18 - fi9)}
                    _ = v22
                    _ = v23
                } else {
                    v7 += v7
                    v7 -= 0.1
                }
                _ = fi18
            }
            v24 := []int{}
            _ = fi9
            _ = v24
        }
        _ = v7
        _ = v8
    }
    if ((!true) && true) {
        v25 := GenS2{F0: false, F1: "hello"}
        var v26 int64 = 1
        _ = v25
        _ = v26
    } else {
        p0 += (p0 / ((p0) | 1))
        var v27 uint32 = 2147483648
        _ = v27
    }
    var v28 uint8 = 128
    for fi29 := 0; fi29 < 18; fi29++ {
        p0 = -358
        switch "foo bar" {
        case "foo bar":
            if ((p0 < int32(v28)) && (42 >= 784)) {
                v30 := GenS1{F0: (v28 % ((v28) | 1)), F1: (p1 / ((p1) | 1))}
                _ = v30
            } else {
                for fi31 := 0; fi31 < 8; fi31++ {
                    if true {
                        v32 := []int{(fi31 | fi29)}
                        p1 = p1
                        var v33 string = ((("hello" + "") + "a") + "a")
                        _ = v32
                        _ = v33
                    } else {
                        switch 1 {
                        case 32768:
                            v34 := []int{128, 32767, (fi31 - fi29)}
                            _ = v34
                        }
                    }
                    v28 = 1
                    _ = fi31
                }
                v35 := GenS3{F0: 245, F1: uint32(v28), F2: int16(v28)}
                _ = v35
            }
        case "a":
            p0 += (p0 - (p0 - -2147483648))
        default:
            var v36 uint16 = (p1 / ((p1) | 1))
            _ = v36
        }
        switch -1 {
        case 32767:
            for fi37 := 0; fi37 < 17; fi37++ {
                var v38 bool = true
                var v39 uint64 = 2
                _ = fi37
                _ = v38
                _ = v39
            }
        case -99:
            var v40 int32 = p0
            _ = v40
        case 127:
            var v41 int64 = 1
            _ = v41
        default:
            for fi42 := 0; fi42 < 18; fi42++ {
                var v43 uint = 0
                if (!true) {
                    p1 = p1
                    fmt.Println(fi42)
                } else {
                    fmt.Println(p0, fi42, fi29)
                    v44 := []int{}
                    if false {
                        if len(v44) > 0 {
                            v44[0] = (fi42 << ((241) & 63))
                        }
                        v28 -= v28
                        var v45 bool = (-2.25 >= 1e10)
                        _ = v45
                    }
                    _ = v44
                }
                v46 := GenS2{F0: true, F1: "\n\t"}
                _ = fi42
                _ = v43
                _ = v46
            }
        }
        _ = fi29
    }
    v47 := []int{808, 264, -850, 32768}
    _ = p0
    _ = p1
    _ = v28
    _ = v47
    return 686
}

func main() {
    var v1 float64 = 0.1
    var v2 float32 = 3.14159
    v3 := make(map[int]int)
    v4 := make(map[uint32]uint32)
    v5 := make(map[uint32]bool)
    v6 := make(map[int]bool)
    fmt.Println(v1, v1)
    var v48 int = 127
    v2 = 0.1
    switch v48 {
    case 127:
        var v49 int = 4294967295
        _ = v49
    default:
        fmt.Println(v1, v48)
    }
    v2 = (0.1 * v2)
    v48 = -1
    if (128 == 127) {
        if true {
            fmt.Println(v1, v48, v2)
            if true {
                v50 := GenS4{F0: 0}
                var v51 bool = (1 != 145)
                _ = v50
                _ = v51
            }
        }
        var v52 uint16 = 1
        _ = v52
    } else {
        for fi53 := 0; fi53 < 3; fi53++ {
            v54 := []int{fi53}
            if false {
                _ = len(v54)
                if ((61 < -1) || false) {
                    if (255 == 2) {
                        v55 := GenS3{F0: 1, F1: 848, F2: 515}
                        v48 = fi53
                        _ = v55
                    } else {
                        v48 += fi53
                    }
                    v2 += (v2 - ((v2 + -2.25) * v2))
                    if (2 < 0) {
                        for fi56 := 0; fi56 < 19; fi56++ {
                            v57 := GenS5{F0: -9223372036854775808, F1: -9223372036854775808}
                            var v58 int16 = -1
                            _ = fi56
                            _ = v57
                            _ = v58
                        }
                    } else {
                        v1 += float64((3.14159 / v2))
                        v1 = 3.14159
                        fmt.Println(v48, v48, v2)
                    }
                }
            }
            var v59 int32 = 0
            _ = fi53
            _ = v54
            _ = v59
        }
        for fi60 := 0; fi60 < 15; fi60++ {
            v48 = -428
            v61 := GenS6{F0: 0, F1: v48}
            _ = fi60
            _ = v61
        }
    }
    for fi62 := 0; fi62 < 5; fi62++ {
        var v63 uint32 = 0
        switch (v48 / ((v48) | 1)) {
        case -9223372036854775808:
            switch (fi62 * -945) {
            case 785:
                if true {
                    var v64 uint64 = 18446744073709551615
                    fmt.Println(fi62)
                    v63 = (v63 - 0)
                    _ = v64
                } else {
                    for fi65 := 0; fi65 < 4; fi65++ {
                        var v66 uint64 = 0
                        v67 := GenS3{F0: uint8(v63), F1: (v63 | (v63 & 422)), F2: -1}
                        v68 := []int{(v48 & v48)}
                        switch "foo bar" {
                        case "a":
                            var v69 float32 = (3.14159 * 0.1)
                            _ = v69
                        case "foo bar":
                            if (!true) {
                                if true {
                                    v2 -= (v2 - v2)
                                }
                                if false {
                                    fmt.Println(v66, fi65, fi65)
                                    v66 -= 473
                                    var v70 uint8 = 1
                                    _ = v70
                                }
                                _ = len(v6)
                            }
                        case "\n\t":
                            var v71 int8 = 0
                            _ = v71
                        }
                        _ = fi65
                        _ = v66
                        _ = v67
                        _ = v68
                    }
                    if (!(!(100 > -128))) {
                        if true {
                            v63 = (v63 + v63)
                            if (2 < 127) {
                                fmt.Println(v63, v63)
                                for fi72 := 0; fi72 < 19; fi72++ {
                                    v73 := v4[v63]
                                    fmt.Println(fi72, fi62, v48)
                                    fmt.Println(v63, v48, v2)
                                    _ = fi72
                                    _ = v73
                                }
                                v48 += (v48 & v48)
                            }
                            v2 = 0.0
                        } else {
                            v74 := GenS6{F0: (^(v63)), F1: v48}
                            switch (fi62 >> ((255) & 63)) {
                            case 32767:
                                var v75 uint8 = 62
                                _ = v75
                            case 32768:
                                for fi76 := 0; fi76 < 3; fi76++ {
                                    switch (fi62 / ((9223372036854775807) | 1)) {
                                    case 4294967295:
                                        if true {
                                            var v77 float64 = 1e10
                                            _ = v77
                                        } else {
                                            if (!false) {
                                                v2 += v2
                                                if (!(2 < 0)) {
                                                    v78 := []int{}
                                                    v79 := GenS7{F0: (v63 - (v63 % (((v63 | (^(v63)))) | 1))), F1: -1, F2: v63}
                                                    var v80 int64 = 642
                                                    _ = v78
                                                    _ = v79
                                                    _ = v80
                                                } else {
                                                    switch fi62 {
                                                    case 9223372036854775807:
                                                        switch (fi76 << ((4294967295) & 63)) {
                                                        case 128:
                                                            switch ("\n\t" + "hello") {
                                                            case "":
                                                                if true {
                                                                    var v81 uint16 = 746
                                                                    v1 = (v1 * (v1 * (v1 + (3.14159 - 3.14159))))
                                                                    fmt.Println(v1)
                                                                    var v82 uint16 = v81
                                                                    switch fi62 {
                                                                    case 9223372036854775807:
                                                                        var v83 int32 = int32(v81)
                                                                        _ = v83
                                                                    case 127:
                                                                        fmt.Println(v82, fi76, v63)
                                                                    case 2147483647:
                                                                        v84 := v3[127]
                                                                        _ = v84
                                                                    }
                                                                    _ = v81
                                                                    _ = v82
                                                                }
                                                            }
                                                        case 4294967295:
                                                            if false {
                                                                v48 -= (^(fi62))
                                                                var v85 int8 = -1
                                                                _ = v85
                                                            }
                                                        case -2147483648:
                                                            switch -275 {
                                                            case 1:
                                                                if false {
                                                                    if false {
                                                                        v48 = fi62
                                                                    }
                                                                    var v86 int16 = 42
                                                                    _ = v86
                                                                } else {
                                                                    fmt.Println(v2, fi76)
                                                                }
                                                            default:
                                                                var v87 int16 = 1000
                                                                var v88 uint = uint(v63)
                                                                v89 := []int{fi62}
                                                                _ = v87
                                                                _ = v88
                                                                _ = v89
                                                            }
                                                        }
                                                    }
                                                    v90 := []int{(fi62 & 338)}
                                                    if (!(187 != 13)) {
                                                        v63 -= (v63 - (v63 % ((4294967295) | 1)))
                                                    }
                                                    _ = v90
                                                }
                                                v91 := []int{9223372036854775807}
                                                _ = v91
                                            } else {
                                                delete(v5, 1)
                                                var v92 uint = uint(v63)
                                                v92 -= v92
                                                _ = v92
                                            }
                                            v2 += v2
                                        }
                                    case 32768:
                                        v63 -= v63
                                        fmt.Println(v63)
                                        var v93 float32 = (v2 * float32((0.0 * v1)))
                                        v94 := GenS6{F0: (v63 ^ (^(v63))), F1: 32767}
                                        var v95 uint16 = 0
                                        _ = v93
                                        _ = v94
                                        _ = v95
                                    }
                                    _ = fi76
                                }
                            case -443:
                                var v96 uint32 = 236
                                _ = v96
                            default:
                                var v97 bool = (-64 != 0)
                                _ = v97
                            }
                            fmt.Println(v1)
                            _ = v74
                        }
                        for fi98 := 0; fi98 < 9; fi98++ {
                            if (65535 >= 42) {
                                fmt.Println(v2, fi62)
                                if false {
                                    var v99 float32 = (((0.0 * 1e10) / 1.5) / v2)
                                    var v100 uint32 = v63
                                    var v101 float64 = 1.5
                                    _ = v100
                                    _ = v101
                                    _ = v99
                                } else {
                                    v102 := GenS8{F0: 42}
                                    _ = v102
                                }
                                fmt.Println(v1, v63, v1)
                                v103 := v5[v63]
                                v103 = ((fi62 + v48) < -2147483648)
                                _ = v103
                            }
                            v63 = 268
                            _ = fi98
                        }
                    }
                    v104 := GenS3{F0: 128, F1: v63, F2: 0}
                    _ = v104
                }
            default:
                switch -504 {
                case 128:
                    fmt.Println(v2, v63, v48)
                case -139:
                    fmt.Println(v2, v48, v48)
                default:
                    for fi105 := 0; fi105 < 12; fi105++ {
                        var v106 uint8 = 2
                        v107 := v6[v48]
                        _ = fi105
                        _ = v106
                        _ = v107
                    }
                }
            }
        case -77:
            var v108 float32 = v2
            _ = v108
        case -788:
            for fi109 := 0; fi109 < 16; fi109++ {
                switch ("a" + "a") {
                case "a":
                    v1 = 1.5
                default:
                    v1 += 1.5
                }
                _ = fi109
            }
        default:
            for fi110 := 0; fi110 < 7; fi110++ {
                fmt.Println(fi62, v2, v1)
                v111 := GenS5{F0: 1, F1: 1}
                var v112 int64 = -9223372036854775808
                _ = fi110
                _ = v111
                _ = v112
            }
        }
        _ = fi62
        _ = v63
    }
    var v113 uint8 = 30
    for fi114 := 0; fi114 < 10; fi114++ {
        fmt.Println(v48)
        _ = fi114
    }
    fmt.Println("v1:", v1)
    fmt.Println("v113:", v113)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v48:", v48)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v113
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v5
    _ = v6
}
