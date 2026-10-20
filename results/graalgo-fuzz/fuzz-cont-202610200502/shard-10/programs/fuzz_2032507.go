package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool; F2 bool }
type GenS2 struct { F0 uint8; F1 int8 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint; F1 int8 }
type GenS5 struct { F0 int16; F1 uint16; F2 bool }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 string; F1 uint8; F2 uint8 }
type GenS8 struct { F0 int64; F1 int64; F2 int64 }
type GenS9 struct { F0 uint32 }
type GenS10 struct { F0 int }
type GenS11 struct { F0 uint64; F1 uint16; F2 int8 }
type GenS12 struct { F0 int8 }
type GenS13 struct { F0 uint64 }
type GenS14 struct { F0 uint8; F1 int8 }
func genFn0(p0 uint64, p1 int) uint16 {
    switch (p1 >> ((0) & 63)) {
    case 32767:
        v4 := GenS1{F0: 255, F1: (false || (2147483647 >= 411)), F2: true}
        _ = v4
    }
    var v5 uint32 = 0
    _ = p0
    _ = p1
    _ = v5
    return 32767
}
func genFn1(p0 int32, p1 uint16, p2 uint) uint {
    switch -830 {
    case 54:
        v6 := GenS1{F0: 42, F1: true, F2: true}
        _ = v6
    default:
        p1 += p1
    }
    var v7 uint = 4294967295
    for fi8 := 0; fi8 < 4; fi8++ {
        fmt.Println(p2, p2, p0)
        var v9 bool = true
        _ = fi8
        _ = v9
    }
    v10 := GenS1{F0: 0, F1: false, F2: ((1e10 * float32((0.1 * 0.1))) <= (3.14159 + 1.5))}
    p0 = 2147483647
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v7
    return uint(p1)
}

func main() {
    var v1 uint8 = 10
    var v2 int16 = -951
    v3 := make(map[uint32]int)
    v2 = 0
    switch 42 {
    case 62:
        if false {
            switch 32767 {
            case 128:
                fmt.Println(v1, v1)
            case 42:
                var v11 float32 = 1e10
                _ = v11
            case -865:
                if (!((42 != 0) && false)) {
                    for fi12 := 0; fi12 < 5; fi12++ {
                        switch fi12 {
                        case 127:
                            for fi13 := 0; fi13 < 19; fi13++ {
                                v2 = (v2 * (v2 + 1))
                                var v14 int32 = -2147483648
                                switch ("\n\t" + "hello") {
                                case "\n\t":
                                    if true {
                                        var v15 uint32 = 968
                                        if (!true) {
                                            var v16 int32 = v14
                                            fmt.Println(fi13)
                                            _ = v16
                                        } else {
                                            var v17 int16 = v2
                                            v14 -= 42
                                            _ = v17
                                        }
                                        var v18 int32 = (v14 + (v14 - v14))
                                        _ = v15
                                        _ = v18
                                    } else {
                                        switch -791 {
                                        case 121:
                                            if true {
                                                var v19 int8 = 40
                                                delete(v3, 942)
                                                var v20 float32 = ((-2.25 + float32((((float64(0.0) + -2.25) + 0.0) * 1.5))) - 1e10)
                                                _ = v19
                                                _ = v20
                                            } else {
                                                var v21 float32 = (1.5 - (1e10 + (3.14159 + 0.0)))
                                                _ = v21
                                            }
                                            v1 = v1
                                            v22 := []int{(-(fi13)), -382}
                                            v23 := GenS2{F0: v1, F1: -128}
                                            var v24 int32 = (v14 ^ -1)
                                            _ = v22
                                            _ = v23
                                            _ = v24
                                        }
                                    }
                                }
                                _ = fi13
                                _ = v14
                            }
                        }
                        for fi25 := 0; fi25 < 3; fi25++ {
                            fmt.Println(v1, fi25, fi25)
                            switch (fi12 & (fi25 - 32768)) {
                            case 9223372036854775807:
                                if false {
                                    var v26 bool = true
                                    var v27 uint = genFn1(2147483647, genFn0(4294967295, (fi12 * (fi12 % ((fi25) | 1)))), 4294967295)
                                    _ = v26
                                    _ = v27
                                }
                            case -9223372036854775808:
                                var v28 uint64 = 0
                                _ = v28
                            case 0:
                                v2 -= 32767
                            default:
                                v1 += v1
                                fmt.Println(v1)
                            }
                            var v29 int64 = -9223372036854775808
                            _ = fi25
                            _ = v29
                        }
                        _ = fi12
                    }
                    if false {
                        var v30 float32 = 1.5
                        v31 := GenS1{F0: v1, F1: true, F2: (127 != 207)}
                        fmt.Println(v1, v2, v2)
                        _ = v30
                        _ = v31
                    }
                }
            default:
                if (!false) {
                    var v32 int64 = int64(v1)
                    _ = v32
                }
            }
            var v33 int64 = 936
            v1 = (v1 - v1)
            _ = v33
        }
    case 691:
        if false {
            for fi34 := 0; fi34 < 17; fi34++ {
                fmt.Println(fi34)
                var v35 string = "a"
                fmt.Println(v35)
                _ = fi34
                _ = v35
            }
        }
    }
    v36 := []int{32767, -9223372036854775808, 32768}
    if len(v36) > 0 {
        v36[0] = 32768
    }
    for fi37 := 0; fi37 < 14; fi37++ {
        for fi38 := 0; fi38 < 2; fi38++ {
            var v39 int32 = 238
            var v40 int16 = v2
            _ = fi38
            _ = v39
            _ = v40
        }
        switch ("hello" + "") {
        case "a":
            if (!true) {
                v41 := GenS1{F0: 128, F1: (!false), F2: (fi37 >= (fi37 - (fi37 * (fi37 * fi37))))}
                if true {
                    for fi42 := 0; fi42 < 1; fi42++ {
                        if false {
                            fmt.Println(v2, v1)
                        }
                        switch (fi37 / ((fi37) | 1)) {
                        case 688:
                            var v43 int16 = (v2 ^ 1)
                            _ = v43
                        }
                        _ = fi42
                    }
                    if (42 >= 2147483647) {
                        v44 := GenS3{F0: v2}
                        if ((923 <= 65535) && true) {
                            var v45 int32 = -14
                            _ = v45
                        }
                        var v46 float32 = float32((0.1 + 3.14159))
                        _ = v44
                        _ = v46
                    } else {
                        switch "" {
                        case "hello":
                            if true {
                                v2 -= v2
                            } else {
                                for fi47 := 0; fi47 < 3; fi47++ {
                                    fmt.Println(fi37)
                                    fmt.Println(v2, v2, fi37)
                                    if (0.1 != float64((1.5 - 1e10))) {
                                        if ((3.14159 < 1e10) && false) {
                                            v2 = -621
                                            if (3.14159 == 1e10) {
                                                v48 := v3[42]
                                                v2 -= -1
                                                if false {
                                                    var v49 int8 = 4
                                                    v48 += (fi47 & (fi47 & fi37))
                                                    _ = v49
                                                }
                                                _ = v48
                                            }
                                            if false {
                                                v50 := GenS4{F0: uint(v1), F1: 1}
                                                fmt.Println(v1, v2, fi47)
                                                var v51 uint16 = 1
                                                _ = v50
                                                _ = v51
                                            } else {
                                                var v52 int8 = 127
                                                v52 -= (v52 / ((v52) | 1))
                                                var v53 float64 = (((((0.0 + 0.0) * -2.25) / ((0.0 + (float64(3.14159) - 0.0)) + float64(((1.5 - 1e10) - (0.0 - 1.5))))) * 1e10) - 0.1)
                                                _ = v52
                                                _ = v53
                                            }
                                        }
                                        v1 += 125
                                        if (32767 != v2) {
                                            v36 = append(v36, (fi47 + (-(fi37))))
                                            v3[2] = (^(fi47))
                                        }
                                    }
                                    _ = fi47
                                }
                            }
                        }
                    }
                    v2 -= -1
                }
                _ = v41
            }
        case "hello":
            v54 := GenS5{F0: -32768, F1: 623, F2: (1 <= 2147483647)}
            _ = v54
        case "\n\t":
            if len(v36) > 0 {
                v36[0] = 32768
            }
        }
        _ = fi37
    }
    var v55 string = "\n\t"
    var v56 float32 = 3.14159
    for fi57 := 0; fi57 < 10; fi57++ {
        if (-2147483648 == -235) {
            if (100 >= 42) {
                var v58 uint8 = 222
                _ = v58
            } else {
                if len(v36) >= 2 {
                    v36 = v36[1:len(v36)]
                }
                var v59 uint64 = 680
                if (!(true || (1e10 != 0.1))) {
                    switch v55 {
                    case "\n\t":
                        if true {
                            switch fi57 {
                            case 32767:
                                v60 := GenS3{F0: v2}
                                _ = v60
                            case -57:
                                for fi61 := 0; fi61 < 20; fi61++ {
                                    var v62 uint = 65535
                                    _ = fi61
                                    _ = v62
                                }
                            case -140:
                                var v63 uint = 42
                                _ = v63
                            }
                        }
                    case "":
                        var v64 uint32 = 65535
                        _ = v64
                    case "foo bar":
                        v65 := GenS2{F0: 0, F1: 127}
                        _ = v65
                    default:
                        if (v55 != (v55 + "")) {
                            for fi66 := 0; fi66 < 4; fi66++ {
                                var v67 uint64 = v59
                                delete(v3, 0)
                                v59 -= v59
                                switch (^(fi57)) {
                                case -1:
                                    switch (("\n\t" + ((v55 + (("foo bar" + v55) + v55)) + "a")) + v55) {
                                    case "a":
                                        _ = len(v36)
                                    case "":
                                        switch v55 {
                                        case "hello":
                                            switch "hello" {
                                            case "\n\t":
                                                var v68 uint8 = (v1 >> ((genFn1(65, 934, 946)) & 63))
                                                _ = v68
                                            case "hello":
                                                v56 -= v56
                                            case "":
                                                switch fi57 {
                                                case 42:
                                                    v69 := GenS3{F0: (v2 << ((588) & 63))}
                                                    v70 := GenS6{F0: 65535}
                                                    _ = v69
                                                    _ = v70
                                                case 0:
                                                    v36 = append(v36, (-(fi66)))
                                                case 128:
                                                    switch 290 {
                                                    case -347:
                                                        fmt.Println(v59, v1, fi66)
                                                        switch v55 {
                                                        case "a":
                                                            v71 := GenS5{F0: (v2 - 915), F1: 259, F2: true}
                                                            _ = v71
                                                        case "\n\t":
                                                            switch fi66 {
                                                            case -717:
                                                                v72 := GenS2{F0: v1, F1: 127}
                                                                _ = v72
                                                            default:
                                                                var v73 uint32 = 829
                                                                _ = v73
                                                            }
                                                        }
                                                        var v74 string = v55
                                                        switch fi66 {
                                                        case 42:
                                                            var v75 int16 = -32768
                                                            _ = v75
                                                        case 19:
                                                            var v76 int64 = -526
                                                            _ = v76
                                                        }
                                                        _ = v74
                                                    default:
                                                        fmt.Println(v2, v2)
                                                    }
                                                }
                                            default:
                                                v2 -= (v2 * v2)
                                            }
                                        case "\n\t":
                                            if false {
                                                fmt.Println(fi57, fi57)
                                            }
                                        case "a":
                                            fmt.Println(v67)
                                        default:
                                            v59 = v59
                                        }
                                    default:
                                        if ((v55 + v55) < v55) {
                                            if false {
                                                var v77 uint16 = 1
                                                switch fi57 {
                                                case 127:
                                                    _ = len(v36)
                                                case 0:
                                                    var v78 int16 = (v2 - (v2 * v2))
                                                    v2 = v2
                                                    _ = v78
                                                }
                                                _ = v77
                                            } else {
                                                v56 -= (1.5 + float32((0.0 - 1.5)))
                                                v36 = append(v36, (fi57 + -2147483648))
                                                delete(v3, 65535)
                                            }
                                            v2 = 32767
                                        }
                                    }
                                case 127:
                                    var v79 int64 = -1
                                    _ = v79
                                case 128:
                                    v80 := GenS7{F0: "\n\t", F1: v1, F2: v1}
                                    _ = v80
                                default:
                                    v2 = 1000
                                }
                                v81 := GenS8{F0: 1, F1: 1, F2: -9223372036854775808}
                                _ = fi66
                                _ = v67
                                _ = v81
                            }
                            var v82 uint = genFn1(42, uint16(v1), 2)
                            _ = v82
                        } else {
                            fmt.Println(v1, fi57)
                            switch "foo bar" {
                            case "hello":
                                v59 = v59
                            case "foo bar":
                                var v83 uint = 2
                                _ = v83
                            case "a":
                                v1 += v1
                            }
                            var v84 uint16 = 42
                            _ = v84
                        }
                    }
                    var v85 int = 32767
                    fmt.Println(v85, fi57, fi57)
                    _ = v85
                } else {
                    if (0.1 != 1e10) {
                        v59 = v59
                    }
                    v86 := GenS9{F0: 2147483647}
                    if ((fi57 * (fi57 % ((fi57) | 1))) <= (^(fi57))) {
                        if (285 > fi57) {
                            var v87 int16 = (v2 & v2)
                            v1 = 74
                            _ = v87
                        } else {
                            if true {
                                if (!true) {
                                    if len(v36) > 0 {
                                        v36[0] = 2147483647
                                    }
                                } else {
                                    var v88 uint32 = 65535
                                    _ = v88
                                }
                            } else {
                                var v89 uint16 = 255
                                fmt.Println(v1, v2)
                                _ = v89
                            }
                            v90 := GenS8{F0: 1, F1: 278, F2: -9223372036854775808}
                            fmt.Println(fi57, v56, v1)
                            _ = v90
                        }
                    }
                    _ = v86
                }
                _ = v59
            }
            v91 := GenS10{F0: fi57}
            fmt.Println(v56)
            _ = v91
        }
        _ = fi57
    }
    for fi92 := 0; fi92 < 20; fi92++ {
        v56 = (1e10 - (v56 - v56))
        v55 = "hello"
        if (18446744073709551615 <= 65535) {
            if (0.0 != ((v56 + (((v56 / v56) - (v56 - 0.1)) / v56)) * v56)) {
                for fi93 := 0; fi93 < 18; fi93++ {
                    fmt.Println(fi92, fi92)
                    v1 = 1
                    fmt.Println(v1)
                    if false {
                        v94 := GenS11{F0: 2, F1: 255, F2: 100}
                        _ = v94
                    }
                    _ = fi93
                }
                v2 += (v2 ^ (v2 & -194))
                fmt.Println(v56)
            }
        } else {
            v95 := GenS6{F0: 32767}
            var v96 string = v55
            _ = v95
            _ = v96
        }
        _ = fi92
    }
    if (!(42 >= 0)) {
        for fi97 := 0; fi97 < 6; fi97++ {
            fmt.Println(v56)
            if false {
                var v98 uint8 = 42
                var v99 uint32 = 42
                _ = v98
                _ = v99
            } else {
                v100 := GenS1{F0: (^(v1)), F1: false, F2: true}
                _ = v100
            }
            if true {
                switch -598 {
                case -1:
                    if false {
                        v55 = (v55 + (("\n\t" + "hello") + ""))
                    } else {
                        switch fi97 {
                        case 1:
                            fmt.Println(fi97, fi97)
                        case -81:
                            v2 += v2
                        default:
                            v2 += 1
                        }
                        var v101 int16 = -661
                        switch "hello" {
                        case "foo bar":
                            v102 := GenS9{F0: 734}
                            _ = v102
                        case "\n\t":
                            for fi103 := 0; fi103 < 15; fi103++ {
                                if (-2147483648 < -1) {
                                    if false {
                                        v101 += v101
                                        delete(v3, 521)
                                        var v104 uint16 = 2
                                        _ = v104
                                    } else {
                                        v56 = (float32((0.0 * -2.25)) * float32(((1.5 - (-2.25 - 0.0)) * 0.1)))
                                        if len(v36) > 0 {
                                            v36[0] = 2147483647
                                        }
                                    }
                                }
                                fmt.Println(v101, v55)
                                switch v55 {
                                case "a":
                                    v105 := GenS12{F0: 0}
                                    _ = v105
                                }
                                _ = fi103
                            }
                        case "a":
                            if (true || true) {
                                v106 := v3[613]
                                delete(v3, 800)
                                v36 = append(v36, v106)
                                _ = v106
                            }
                        default:
                            v107 := GenS12{F0: -128}
                            _ = v107
                        }
                        _ = v101
                    }
                }
                v56 -= v56
            } else {
                v2 = (v2 + (v2 ^ 1000))
                v55 = "a"
            }
            _ = fi97
        }
        v108 := GenS13{F0: 201}
        _ = v108
    } else {
        fmt.Println(v56, v56)
        v55 = "foo bar"
        fmt.Println(v55, v56)
    }
    fmt.Println(v55, v1, v1)
    v109 := GenS14{F0: v1, F1: 1}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v36:", len(v36), v36)
    fmt.Println("v55:", v55)
    fmt.Println("v56:", v56)
    _ = v1
    _ = v109
    _ = v2
    _ = v3
    _ = v36
    _ = v55
    _ = v56
}
