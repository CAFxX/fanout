package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int16; F1 int; F2 int32 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int32; F1 string; F2 bool }
type GenS6 struct { F0 bool; F1 uint16; F2 int64 }
type GenS7 struct { F0 uint8; F1 int }
type GenS8 struct { F0 bool; F1 uint16; F2 int8 }
type GenS9 struct { F0 int8; F1 int }
type GenS10 struct { F0 uint; F1 uint8 }
type GenS11 struct { F0 int8; F1 uint16; F2 int8 }
type GenS12 struct { F0 uint32; F1 int8; F2 uint64 }
func genFn0(p0 int64, p1 int64, p2 int16) uint16 {
    v6 := GenS1{F0: 424}
    if (32767 <= 65535) {
        var v7 uint = 4294967295
        _ = v7
    }
    p1 = p1
    switch "" {
    case "\n\t":
        switch "\n\t" {
        case "hello":
            p1 = p1
        case "foo bar":
            fmt.Println(p2)
        case "a":
            p2 -= p2
        default:
            v8 := GenS1{F0: 42}
            _ = v8
        }
    case "":
        v9 := []int{-500, -676, -9223372036854775808}
        _ = v9
    case "a":
        p0 += p0
    }
    if true {
        v10 := []int{2147483647, 2147483647, 1, 127}
        for fi11 := 0; fi11 < 20; fi11++ {
            var v12 int8 = -1
            v13 := GenS1{F0: 65535}
            v14 := GenS1{F0: 4294967295}
            _ = fi11
            _ = v12
            _ = v13
            _ = v14
        }
        _ = v10
    } else {
        for fi15 := 0; fi15 < 19; fi15++ {
            var v16 float64 = 0.0
            _ = fi15
            _ = v16
        }
        var v17 int8 = -128
        _ = v17
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 255
}

func main() {
    var v1 int = 128
    var v2 string = "a"
    var v3 int = (v1 & (v1 >> ((65535) & 63)))
    var v4 float64 = float64(3.14159)
    var v5 float64 = (float64(0.1) / (v4 * v4))
    v3 -= v1
    switch (v1 - v1) {
    case 42:
        switch v2 {
        case "a":
            v18 := []int{(-(v3))}
            _ = v18
        default:
            for fi19 := 0; fi19 < 5; fi19++ {
                v4 = v4
                v5 += ((3.14159 - 3.14159) / v5)
                v3 = (v3 >> ((349) & 63))
                _ = fi19
            }
        }
    default:
        var v20 int32 = 0
        _ = v20
    }
    for fi21 := 0; fi21 < 19; fi21++ {
        for fi22 := 0; fi22 < 3; fi22++ {
            var v23 uint32 = 2147483647
            if (false && false) {
                v1 += (fi21 << ((uint(v23)) & 63))
                var v24 float64 = 1.5
                _ = v24
            }
            _ = fi22
            _ = v23
        }
        var v25 int = 42
        _ = fi21
        _ = v25
    }
    for fi26 := 0; fi26 < 18; fi26++ {
        switch ((("hello" + v2) + v2) + ((v2 + v2) + "foo bar")) {
        case "hello":
            if ((v2 + v2) <= "\n\t") {
                switch v2 {
                case "\n\t":
                    v4 = (3.14159 + 0.1)
                case "a":
                    v27 := GenS2{F0: 32767, F1: (-(v3)), F2: -1}
                    _ = v27
                case "hello":
                    v28 := []int{(^(fi26)), (v1 + (v3 ^ -1))}
                    _ = v28
                default:
                    for fi29 := 0; fi29 < 14; fi29++ {
                        fmt.Println(fi26, v2)
                        _ = fi29
                    }
                }
                for fi30 := 0; fi30 < 14; fi30++ {
                    if true {
                        v31 := GenS3{F0: 2}
                        _ = v31
                    }
                    _ = fi30
                }
                v5 -= v5
            }
        case "":
            var v32 int = (^(v3))
            _ = v32
        case "\n\t":
            if true {
                v1 -= 36
                if true {
                    for fi33 := 0; fi33 < 6; fi33++ {
                        if true {
                            if (-1 <= 42) {
                                var v34 int = fi33
                                _ = v34
                            } else {
                                v1 -= (fi26 ^ fi33)
                                v35 := GenS2{F0: 1000, F1: (fi33 & v3), F2: -2147483648}
                                _ = v35
                            }
                            if ((v3 / ((fi26) | 1)) >= fi33) {
                                v36 := GenS3{F0: 0}
                                var v37 uint64 = 2
                                _ = v36
                                _ = v37
                            } else {
                                switch ((v2 + "foo bar") + (v2 + (v2 + v2))) {
                                case "":
                                    fmt.Println(v1)
                                case "hello":
                                    v38 := GenS2{F0: 227, F1: fi33, F2: 0}
                                    _ = v38
                                default:
                                    var v39 uint64 = 1
                                    _ = v39
                                }
                                v40 := GenS3{F0: 811}
                                _ = v40
                            }
                            fmt.Println(v3, fi33)
                            v41 := []int{(fi26 + fi33), (fi26 + 9223372036854775807), (v3 << ((449) & 63)), 127}
                            if (v4 > (v4 - (0.1 - float64(3.14159)))) {
                                switch ("foo bar" + v2) {
                                case "a":
                                    var v42 int16 = 1
                                    _ = v42
                                case "":
                                    var v43 int64 = 0
                                    _ = v43
                                case "foo bar":
                                    _ = len(v41)
                                default:
                                    if len(v41) > 0 {
                                        v41[0] = v3
                                    }
                                }
                                v5 += (((float64(float32((-2.25 / v5))) * 0.1) / v4) + v5)
                            } else {
                                v3 = (v1 & v1)
                                if (!(0 >= 32767)) {
                                    var v44 float64 = ((v5 * (1e10 + v5)) * (3.14159 / (1e10 / (v5 / 0.1))))
                                    var v45 float32 = 0.1
                                    var v46 uint32 = 2147483647
                                    _ = v44
                                    _ = v45
                                    _ = v46
                                } else {
                                    v2 = v2
                                    v1 = v1
                                    switch (v3 ^ 32768) {
                                    case 1:
                                        switch 32767 {
                                        case 32767:
                                            v5 = v4
                                        case 0:
                                            fmt.Println(v3)
                                        default:
                                            v1 = (v1 >> ((2) & 63))
                                            v4 = float64(1e10)
                                            var v47 int8 = 34
                                            v4 = float64(0.0)
                                            v48 := GenS4{F0: 2}
                                            switch v3 {
                                            case -2147483648:
                                                v49 := GenS2{F0: -320, F1: (fi26 + 800), F2: 0}
                                                _ = v49
                                            case 2147483647:
                                                switch (v3 % ((-616) | 1)) {
                                                case 415:
                                                    if false {
                                                        var v50 uint = 917
                                                        _ = v50
                                                    }
                                                case 0:
                                                    fmt.Println(fi33, v47, v1)
                                                case 32767:
                                                    v2 = (v2 + ("a" + "a"))
                                                    v3 = int(v47)
                                                default:
                                                    var v51 int64 = 0
                                                    _ = v51
                                                }
                                            }
                                            fmt.Println(v1, v4, fi26)
                                            if ((!(v47 >= (^(v47)))) || (false || false)) {
                                                v5 -= 1.5
                                                switch v2 {
                                                case "":
                                                    switch "\n\t" {
                                                    case "hello":
                                                        switch (v2 + v2) {
                                                        case "":
                                                            v52 := GenS5{F0: -538, F1: v2, F2: false}
                                                            _ = v52
                                                        default:
                                                            switch (v3 << ((255) & 63)) {
                                                            case -871:
                                                                var v53 uint8 = 80
                                                                _ = v53
                                                            case -1:
                                                                v54 := GenS4{F0: 162}
                                                                _ = v54
                                                            default:
                                                                fmt.Println(v47, v3, v4)
                                                            }
                                                        }
                                                    }
                                                case "\n\t":
                                                    v55 := GenS6{F0: false, F1: 255, F2: 42}
                                                    _ = v55
                                                default:
                                                    if true {
                                                        v5 = (v5 - v5)
                                                        switch "" {
                                                        case "foo bar":
                                                            if len(v41) >= 2 {
                                                                v41 = v41[1:len(v41)]
                                                            }
                                                        case "hello":
                                                            fmt.Println(v3, v1, fi33)
                                                        case "\n\t":
                                                            switch -881 {
                                                            case 212:
                                                                fmt.Println(v5, fi26)
                                                                var v56 float32 = (-2.25 / 1e10)
                                                                _ = v56
                                                            case -9223372036854775808:
                                                                v41 = append(v41, (v3 >> ((0) & 63)))
                                                            case 1:
                                                                v5 = ((v5 * v5) - (v5 * (0.1 + v4)))
                                                            }
                                                        }
                                                        switch -9223372036854775808 {
                                                        case 285:
                                                            v1 += fi26
                                                        case 128:
                                                            fmt.Println(v4, v3)
                                                        case 42:
                                                            v2 = v2
                                                        }
                                                    } else {
                                                        var v57 int8 = (v47 ^ -128)
                                                        switch v1 {
                                                        case -2147483648:
                                                            fmt.Println(v47, v57, v3)
                                                        case 2147483647:
                                                            if true {
                                                                var v58 int32 = -2147483648
                                                                switch fi26 {
                                                                case -470:
                                                                    switch v3 {
                                                                    case 0:
                                                                        v59 := GenS3{F0: 909}
                                                                        _ = v59
                                                                    }
                                                                default:
                                                                    var v60 string = "a"
                                                                    _ = v60
                                                                }
                                                                v4 -= (3.14159 * v4)
                                                                var v61 int32 = (v58 >> ((4294967295) & 63))
                                                                _ = v58
                                                                _ = v61
                                                            } else {
                                                                fmt.Println(v1, v47, v3)
                                                                if true {
                                                                    var v62 bool = true
                                                                    var v63 int8 = 100
                                                                    v64 := GenS7{F0: 1, F1: (fi33 & v1)}
                                                                    v65 := GenS1{F0: 237}
                                                                    _ = v62
                                                                    _ = v63
                                                                    _ = v64
                                                                    _ = v65
                                                                } else {
                                                                    v2 = "a"
                                                                    var v66 uint64 = 4294967295
                                                                    _ = v66
                                                                }
                                                                _ = len(v41)
                                                            }
                                                        }
                                                        _ = v57
                                                    }
                                                }
                                                var v67 int = fi33
                                                v4 += (1e10 * 0.0)
                                                v67 += (v67 * v1)
                                                _ = v67
                                            }
                                            _ = v47
                                            _ = v48
                                        }
                                    case -719:
                                        if false {
                                            fmt.Println(v3, fi33)
                                            v3 = -993
                                        }
                                    }
                                    fmt.Println(v1)
                                }
                            }
                            _ = v41
                        }
                        if true {
                            switch ("hello" + "\n\t") {
                            case "foo bar":
                                var v68 int16 = -789
                                _ = v68
                            case "\n\t":
                                v69 := []int{fi33, -2147483648, (-(fi26))}
                                _ = v69
                            }
                            var v70 int16 = -915
                            var v71 float32 = 3.14159
                            _ = v70
                            _ = v71
                        } else {
                            v72 := []int{}
                            if true {
                                var v73 int = (v1 << ((1) & 63))
                                if ((32767 >= genFn0(-32, 217, 42)) && (255 != 431)) {
                                    v74 := GenS8{F0: true, F1: 32767, F2: 1}
                                    switch (v1 % (((fi33 + -2147483648)) | 1)) {
                                    case 32768:
                                        var v75 uint = 0
                                        _ = v75
                                    case 767:
                                        var v76 uint64 = 42
                                        _ = v76
                                    default:
                                        var v77 int = (v1 ^ (v3 + -2147483648))
                                        _ = v77
                                    }
                                    if ((0 >= 0) && (95 < -9223372036854775808)) {
                                        v4 = 3.14159
                                    }
                                    _ = v74
                                } else {
                                    var v78 int32 = 321
                                    v79 := GenS9{F0: 54, F1: (v73 / ((v1) | 1))}
                                    if (false && (42 != 32767)) {
                                        var v80 uint32 = 1
                                        _ = v80
                                    }
                                    _ = v78
                                    _ = v79
                                }
                                _ = v73
                            }
                            _ = v72
                        }
                        _ = fi33
                    }
                    var v81 string = ""
                    var v82 int32 = -796
                    _ = v81
                    _ = v82
                }
                for fi83 := 0; fi83 < 6; fi83++ {
                    switch "hello" {
                    case "\n\t":
                        v2 = (("" + (v2 + ("foo bar" + (v2 + v2)))) + v2)
                        v1 = (-(fi26))
                        v3 -= -889
                    default:
                        v84 := GenS10{F0: 65535, F1: 0}
                        _ = v84
                    }
                    var v85 int8 = 4
                    var v86 uint32 = 2147483648
                    v87 := GenS11{F0: v85, F1: 65535, F2: v85}
                    if false {
                        v3 -= fi26
                        switch 2147483647 {
                        case 32767:
                            v4 = v4
                        }
                        v88 := GenS4{F0: 183}
                        _ = v88
                    }
                    _ = fi83
                    _ = v85
                    _ = v86
                    _ = v87
                }
            } else {
                switch "hello" {
                case "a":
                    v89 := []int{-794, fi26, (v3 - v3), v3}
                    _ = v89
                case "\n\t":
                    v90 := GenS12{F0: 1, F1: 1, F2: 841}
                    _ = v90
                default:
                    var v91 bool = (3.14159 != v5)
                    _ = v91
                }
                var v92 uint16 = 2
                fmt.Println(fi26, v1, v2)
                _ = v92
            }
        }
        _ = fi26
    }
    fmt.Println(v1, v5, v2)
    v4 -= (v4 - 3.14159)
    for fi93 := 0; fi93 < 8; fi93++ {
        var v94 int64 = -9223372036854775808
        v4 = ((v5 / v4) + v5)
        v95 := []int{fi93, 42, 127}
        _ = fi93
        _ = v94
        _ = v95
    }
    if true {
        v96 := GenS11{F0: -128, F1: 32767, F2: 127}
        _ = v96
    } else {
        var v97 int16 = -399
        _ = v97
    }
    fmt.Println(v1)
    switch -9223372036854775808 {
    case -896:
        v98 := []int{v1, (^(v3)), -356}
        _ = v98
    default:
        if false {
            var v99 uint16 = 0
            _ = v99
        } else {
            if true {
                v100 := GenS11{F0: 1, F1: 437, F2: 127}
                _ = v100
            } else {
                var v101 uint8 = 255
                for fi102 := 0; fi102 < 15; fi102++ {
                    var v103 uint8 = (v101 << ((65535) & 63))
                    switch v2 {
                    case "":
                        for fi104 := 0; fi104 < 19; fi104++ {
                            v105 := GenS1{F0: 18446744073709551615}
                            v106 := []int{(v3 ^ v3), (v3 >> ((255) & 63)), (v3 >> ((42) & 63)), 698}
                            _ = fi104
                            _ = v105
                            _ = v106
                        }
                    }
                    var v107 int32 = 1
                    _ = fi102
                    _ = v103
                    _ = v107
                }
                fmt.Println(v3)
                _ = v101
            }
            fmt.Println(v4)
        }
    }
    var v108 uint8 = 2
    v109 := []int{}
    fmt.Println(v1, v5)
    fmt.Println("v1:", v1)
    fmt.Println("v108:", v108)
    fmt.Println("slicelen v109:", len(v109), v109)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v108
    _ = v109
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
