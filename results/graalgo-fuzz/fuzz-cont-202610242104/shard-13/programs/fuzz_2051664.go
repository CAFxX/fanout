package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint; F2 int32 }
type GenS2 struct { F0 uint; F1 string }
type GenS3 struct { F0 int8; F1 int8; F2 uint64 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint64; F1 uint8 }
type GenS6 struct { F0 uint32; F1 int; F2 int8 }
type GenS7 struct { F0 uint32; F1 int16; F2 int8 }
type GenS8 struct { F0 string; F1 uint16 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 string = "foo bar"
    var v3 int16 = 32767
    var v4 uint8 = 127
    var v5 uint32 = 728
    switch 1 {
    case -2147483648:
        var v6 uint8 = 63
        _ = v6
    case -402:
        var v7 uint = 4294967295
        _ = v7
    }
    if (-110 > 127) {
        fmt.Println(v3, v2, v5)
        v8 := []int{529, -64}
        if (false || (!false)) {
            if true {
                var v9 bool = true
                if false {
                    var v10 float32 = -2.25
                    _ = v10
                }
                if v9 {
                    if len(v8) > 0 {
                        v8[0] = -50
                    }
                    for fi11 := 0; fi11 < 17; fi11++ {
                        for fi12 := 0; fi12 < 9; fi12++ {
                            if (false || (795 <= 2147483647)) {
                                v2 = (("\n\t" + v2) + ("a" + "foo bar"))
                                _ = len(v8)
                                switch 1 {
                                case 1:
                                    var v13 bool = (v9 || v9)
                                    _ = v13
                                case 0:
                                    v3 += 0
                                }
                            } else {
                                var v14 int = 873
                                var v15 string = ("foo bar" + (v2 + v2))
                                _ = v14
                                _ = v15
                            }
                            var v16 int64 = -857
                            if len(v8) > 0 {
                                v8[0] = (fi12 - fi12)
                            }
                            _ = fi12
                            _ = v16
                        }
                        var v17 int32 = 363
                        _ = fi11
                        _ = v17
                    }
                    if v9 {
                        var v18 int64 = -1
                        fmt.Println(v4, v9, v4)
                        v18 -= v18
                        _ = v18
                    }
                }
                _ = v9
            }
            v3 += v3
            v19 := GenS1{F0: (("foo bar" + v2) + (v2 + (((v2 + (v2 + v2)) + (v2 + "a")) + v2))), F1: 18446744073709551615, F2: -2147483648}
            _ = v19
        }
        _ = v8
    }
    var v20 int8 = 42
    v21 := GenS1{F0: v2, F1: 87, F2: 820}
    for fi22 := 0; fi22 < 17; fi22++ {
        if (718 >= 255) {
            if false {
                delete(v1, (v5 << ((844) & 63)))
                for fi23 := 0; fi23 < 13; fi23++ {
                    v4 = (^(v4))
                    var v24 uint8 = (v4 | (^(v4)))
                    switch (fi23 | fi23) {
                    case 1:
                        if false {
                            if (("hello" + (v2 + ((v2 + "foo bar") + v2))) == v2) {
                                fmt.Println(fi23)
                                v3 -= 1000
                                fmt.Println(v20)
                            } else {
                                delete(v1, 47)
                                _ = len(v1)
                                fmt.Println(v3, v3)
                                delete(v1, (^(v5)))
                                if true {
                                    v25 := GenS2{F0: 117, F1: v2}
                                    v2 = (("" + v2) + "foo bar")
                                    v26 := GenS2{F0: 42, F1: v2}
                                    _ = v25
                                    _ = v26
                                }
                                var v27 uint32 = 1
                                v5 = v5
                                v28 := v1[v27]
                                _ = v27
                                _ = v28
                            }
                            var v29 uint8 = v4
                            if true {
                                switch v2 {
                                case "foo bar":
                                    var v30 int64 = 864
                                    _ = v30
                                case "\n\t":
                                    var v31 int16 = v3
                                    _ = v31
                                case "a":
                                    if false {
                                        v5 = (v5 >> ((1) & 63))
                                        v20 += (^(v20))
                                    }
                                }
                                v2 = "hello"
                            } else {
                                if (v2 != v2) {
                                    var v32 bool = (0.0 < 0.1)
                                    v29 = 128
                                    v33 := v1[v5]
                                    v34 := []int{(fi23 | int(v20)), fi22, fi22}
                                    var v35 int = 2147483647
                                    _ = v32
                                    _ = v33
                                    _ = v34
                                    _ = v35
                                } else {
                                    switch fi23 {
                                    case 867:
                                        v4 = 255
                                    case -2147483648:
                                        v24 -= v29
                                    case 32768:
                                        v3 = -222
                                    default:
                                        var v36 float64 = 3.14159
                                        _ = v36
                                    }
                                    if false {
                                        if false {
                                            v37 := []int{(-(fi23)), fi23}
                                            _ = v37
                                        } else {
                                            _ = len(v1)
                                        }
                                        fmt.Println(v29, v4, fi22)
                                    } else {
                                        switch 2147483647 {
                                        case 79:
                                            var v38 uint32 = 282
                                            _ = v38
                                        default:
                                            v39 := GenS3{F0: (-(v20)), F1: (v20 - (v20 & v20)), F2: uint64(v24)}
                                            _ = v39
                                        }
                                        var v40 float32 = -2.25
                                        v5 = (^(v5))
                                        _ = v40
                                    }
                                }
                            }
                            _ = v29
                        } else {
                            v2 = v2
                            v5 -= v5
                        }
                    default:
                        switch ("\n\t" + (((("\n\t" + v2) + "a") + v2) + v2)) {
                        case "\n\t":
                            var v41 int32 = 0
                            _ = v41
                        case "":
                            var v42 uint = 493
                            _ = v42
                        }
                    }
                    _ = fi23
                    _ = v24
                }
                var v43 bool = false
                _ = v43
            } else {
                var v44 int64 = -1
                _ = v44
            }
            if false {
                v45 := []int{fi22, fi22}
                _ = v45
            }
            for fi46 := 0; fi46 < 13; fi46++ {
                switch fi46 {
                case -348:
                    var v47 int = 840
                    if true {
                        fmt.Println(v47)
                        if (false && true) {
                            if (!true) {
                                var v48 int64 = -650
                                v49 := GenS1{F0: v2, F1: 65535, F2: 167}
                                _ = v48
                                _ = v49
                            }
                            if false {
                                if false {
                                    v3 = (^(v3))
                                } else {
                                    var v50 float32 = (0.1 / 1e10)
                                    v20 = (v20 % (((v20 ^ (v20 ^ v20))) | 1))
                                    _ = v50
                                }
                                v20 = (v20 + 127)
                                v51 := GenS3{F0: v20, F1: 27, F2: 4294967295}
                                _ = v51
                            } else {
                                v52 := []int{2147483647, v47, fi22}
                                _ = v52
                            }
                        }
                        switch fi22 {
                        case 128:
                            if (true && true) {
                                v53 := GenS4{F0: 42}
                                v54 := GenS5{F0: 0, F1: 42}
                                _ = v53
                                _ = v54
                            } else {
                                var v55 uint64 = 625
                                _ = v55
                            }
                        case -1:
                            fmt.Println(v2, v5, v5)
                        }
                        v5 += (^(v5))
                    }
                    _ = v47
                }
                if true {
                    v56 := []int{(^(fi22)), (fi46 & (fi46 << ((4294967295) & 63))), fi22}
                    _ = v56
                } else {
                    switch "" {
                    case "hello":
                        v57 := v1[v5]
                        _ = v57
                    }
                    switch (fi46 - (fi22 - (fi46 >> ((38) & 63)))) {
                    case 32767:
                        var v58 float64 = 0.0
                        _ = v58
                    case -2147483648:
                        var v59 string = "foo bar"
                        _ = v59
                    default:
                        v20 = (v20 * v20)
                    }
                }
                if false {
                    v1[2147483648] = false
                    fmt.Println(v3)
                    var v60 uint64 = 438
                    _ = v60
                }
                _ = fi46
            }
        } else {
            fmt.Println(v4, v4, v20)
        }
        _ = fi22
    }
    if false {
        if true {
            fmt.Println(v20, v3)
            for fi61 := 0; fi61 < 10; fi61++ {
                v62 := []int{fi61, 127}
                v2 = ((v2 + (v2 + (v2 + (("hello" + ("hello" + "foo bar")) + v2)))) + v2)
                _ = fi61
                _ = v62
            }
            var v63 int = -2147483648
            _ = v63
        }
        v64 := GenS6{F0: v5, F1: 32767, F2: (v20 - v20)}
        _ = v64
    }
    v5 -= v5
    if (1 == 766) {
        if false {
            v2 = "hello"
            if (495 != 900) {
                v4 = (v4 % ((v4) | 1))
            } else {
                fmt.Println(v5, v20, v4)
                for fi65 := 0; fi65 < 14; fi65++ {
                    for fi66 := 0; fi66 < 9; fi66++ {
                        switch (fi65 + (fi66 % ((fi66) | 1))) {
                        case 451:
                            v67 := v1[v5]
                            _ = v67
                        case 8:
                            var v68 int64 = int64(fi66)
                            _ = v68
                        }
                        var v69 int16 = v3
                        switch (fi65 % ((-244) | 1)) {
                        case 9223372036854775807:
                            v69 -= 779
                        case -1:
                            var v70 uint16 = 32767
                            _ = v70
                        case 128:
                            var v71 int32 = -1
                            _ = v71
                        }
                        if (false || (534 <= 337)) {
                            var v72 int16 = int16(v20)
                            fmt.Println(v4, v72)
                            var v73 uint32 = v5
                            _ = v72
                            _ = v73
                        }
                        v20 = (v20 * (v20 + (v20 / ((127) | 1))))
                        v74 := GenS7{F0: 0, F1: v3, F2: (v20 ^ (-(v20)))}
                        _ = fi66
                        _ = v69
                        _ = v74
                    }
                    _ = fi65
                }
                fmt.Println(v2)
            }
            if false {
                fmt.Println(v3, v4, v3)
                if false {
                    v75 := GenS5{F0: 4294967295, F1: v4}
                    _ = v75
                } else {
                    if (-516 == 0) {
                        if true {
                            delete(v1, v5)
                        }
                        var v76 uint32 = v5
                        _ = v76
                    }
                }
            } else {
                var v77 int = 32767
                if false {
                    var v78 int16 = (v3 - -29)
                    var v79 int = v77
                    _ = v78
                    _ = v79
                }
                _ = v77
            }
        }
        v20 += v20
        v80 := GenS4{F0: -1}
        _ = v80
    } else {
        if (65535 != 65535) {
            v4 += v4
        } else {
            for fi81 := 0; fi81 < 11; fi81++ {
                for fi82 := 0; fi82 < 14; fi82++ {
                    v83 := []int{fi81, (-(fi81)), (fi82 / ((int(v20)) | 1))}
                    v3 -= v3
                    _ = fi82
                    _ = v83
                }
                if (302 < fi81) {
                    switch fi81 {
                    case -9223372036854775808:
                        v4 -= (^(v4))
                    default:
                        for fi84 := 0; fi84 < 18; fi84++ {
                            fmt.Println(v20)
                            switch 709 {
                            case 419:
                                v85 := v1[v5]
                                _ = v85
                            case 338:
                                v5 += v5
                            default:
                                v5 = (v5 & (^(v5)))
                            }
                            var v86 int8 = v20
                            fmt.Println(v5, fi81)
                            _ = fi84
                            _ = v86
                        }
                    }
                    if false {
                        if ((v2 <= v2) || true) {
                            v87 := GenS7{F0: v5, F1: (v3 & (v3 >> ((42) & 63))), F2: -107}
                            _ = v87
                        } else {
                            v2 = "hello"
                            v4 = (^(v4))
                        }
                        var v88 int8 = v20
                        _ = v88
                    }
                } else {
                    switch v2 {
                    case "foo bar":
                        for fi89 := 0; fi89 < 10; fi89++ {
                            v3 += (v3 & v3)
                            if (!(!true)) {
                                v1[(v5 | (v5 >> ((40) & 63)))] = false
                            } else {
                                switch 0 {
                                case 32767:
                                    v2 = v2
                                case -919:
                                    v90 := []int{}
                                    _ = v90
                                case 699:
                                    v5 = 65535
                                }
                            }
                            _ = fi89
                        }
                    case "a":
                        for fi91 := 0; fi91 < 10; fi91++ {
                            v92 := GenS8{F0: v2, F1: 42}
                            if false {
                                if true {
                                    if true {
                                        if (v3 == v3) {
                                            var v93 uint64 = 315
                                            _ = v93
                                        } else {
                                            fmt.Println(v5, v2)
                                            switch 32768 {
                                            case -960:
                                                var v94 int16 = (v3 ^ v3)
                                                _ = v94
                                            case 0:
                                                switch "hello" {
                                                case "a":
                                                    v95 := v1[523]
                                                    _ = v95
                                                default:
                                                    v96 := GenS1{F0: v2, F1: 4294967295, F2: -550}
                                                    _ = v96
                                                }
                                            case -799:
                                                v3 = (v3 | -1)
                                            }
                                            v2 = "\n\t"
                                            v5 = (v5 * 62)
                                        }
                                        v2 = "hello"
                                        switch v2 {
                                        case "":
                                            v3 = (v3 % ((v3) | 1))
                                        case "a":
                                            v5 = (v5 * v5)
                                        }
                                    }
                                    v3 = v3
                                }
                                v20 = v20
                                v97 := GenS5{F0: 1, F1: v4}
                                delete(v1, (v5 + (v5 ^ v5)))
                                _ = v97
                            } else {
                                v5 = (v5 & v5)
                            }
                            delete(v1, v5)
                            _ = fi91
                            _ = v92
                        }
                    case "\n\t":
                        v2 = v2
                    default:
                        v3 = v3
                    }
                }
                _ = fi81
            }
            for fi98 := 0; fi98 < 1; fi98++ {
                v99 := GenS3{F0: 0, F1: (v20 % (((v20 >> ((1) & 63))) | 1)), F2: 1}
                if (!((v2 + "hello") == ("" + "a"))) {
                    var v100 uint = 583
                    if (((uint32(v100) >= (v5 * (^(v5)))) || false) || true) {
                        var v101 uint8 = uint8(v100)
                        v102 := v1[(v5 % ((2147483647) | 1))]
                        for fi103 := 0; fi103 < 4; fi103++ {
                            v104 := GenS5{F0: 18446744073709551615, F1: (v4 ^ (^(v101)))}
                            v105 := []int{(fi98 - (fi98 << ((4294967295) & 63))), (fi103 | (fi103 * fi103)), (fi103 * fi98), -199}
                            _ = fi103
                            _ = v104
                            _ = v105
                        }
                        _ = v101
                        _ = v102
                    } else {
                        for fi106 := 0; fi106 < 15; fi106++ {
                            if (42 < 0) {
                                var v107 uint = v100
                                _ = v107
                            }
                            _ = fi106
                        }
                    }
                    _ = v100
                } else {
                    var v108 string = v2
                    _ = v108
                }
                _ = fi98
                _ = v99
            }
            v2 = "a"
        }
        v109 := GenS6{F0: 783, F1: 9223372036854775807, F2: -128}
        _ = v109
    }
    switch 32767 {
    case 128:
        for fi110 := 0; fi110 < 4; fi110++ {
            if true {
                var v111 uint32 = (v5 / ((v5) | 1))
                var v112 uint64 = 4294967295
                switch v2 {
                case "foo bar":
                    var v113 uint64 = (v112 ^ (v112 - (v112 >> ((65535) & 63))))
                    _ = v113
                case "":
                    var v114 int32 = 42
                    _ = v114
                case "\n\t":
                    v115 := []int{(-(fi110)), fi110}
                    _ = v115
                }
                _ = v111
                _ = v112
            } else {
                v2 = (("" + ((v2 + v2) + v2)) + v2)
                v116 := []int{fi110}
                _ = v116
            }
            for fi117 := 0; fi117 < 3; fi117++ {
                v118 := GenS2{F0: 2, F1: v2}
                _ = fi117
                _ = v118
            }
            _ = fi110
        }
    case 127:
        for fi119 := 0; fi119 < 19; fi119++ {
            v120 := []int{int(v5), fi119, fi119, fi119}
            switch ((v2 + v2) + v2) {
            case "\n\t":
                v121 := GenS5{F0: 0, F1: v4}
                _ = v121
            default:
                var v122 int8 = v20
                _ = v122
            }
            fmt.Println(v2, v5)
            _ = fi119
            _ = v120
        }
    }
    if false {
        var v123 uint = 255
        var v124 uint = (v123 & (v123 << ((160) & 63)))
        _ = v123
        _ = v124
    }
    for fi125 := 0; fi125 < 5; fi125++ {
        var v126 int64 = int64(v3)
        var v127 uint8 = (^(v4))
        _ = fi125
        _ = v126
        _ = v127
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
