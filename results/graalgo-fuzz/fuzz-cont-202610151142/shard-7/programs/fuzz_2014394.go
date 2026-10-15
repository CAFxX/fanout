package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64 }
type GenS2 struct { F0 int; F1 int64 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 bool; F1 uint64 }
type GenS5 struct { F0 int; F1 uint8 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 int; F1 uint; F2 int }
type GenS8 struct { F0 uint32; F1 uint16 }
type GenS9 struct { F0 uint16 }

func main() {
    var v1 float64 = (1.5 * 0.1)
    var v2 int32 = 42
    var v3 float64 = 3.14159
    var v4 uint = 549
    v5 := make(map[uint32]uint32)
    var v6 int16 = -710
    fmt.Println(v2, v2, v2)
    fmt.Println(v2, v1)
    var v7 int32 = v2
    v7 -= v2
    var v8 uint64 = 1
    v9 := GenS1{F0: 0, F1: 42}
    if true {
        v8 -= v8
        switch "hello" {
        case "a":
            switch "foo bar" {
            case "a":
                var v10 uint8 = 1
                _ = v10
            case "":
                var v11 float64 = (v3 / (v3 * -2.25))
                _ = v11
            case "foo bar":
                var v12 int8 = -34
                _ = v12
            }
        case "hello":
            v8 -= 2
        case "\n\t":
            switch 0 {
            case -9223372036854775808:
                switch "hello" {
                case "a":
                    v2 -= (v7 + (v7 & (v2 + v7)))
                case "foo bar":
                    if false {
                        for fi13 := 0; fi13 < 15; fi13++ {
                            switch "\n\t" {
                            case "a":
                                for fi14 := 0; fi14 < 2; fi14++ {
                                    var v15 string = "a"
                                    var v16 float64 = 1e10
                                    v17 := []int{fi14, 32768}
                                    _ = fi14
                                    _ = v15
                                    _ = v16
                                    _ = v17
                                }
                            case "hello":
                                var v18 float64 = v3
                                _ = v18
                            }
                            v8 = 4294967295
                            v19 := v5[1]
                            _ = fi13
                            _ = v19
                        }
                    } else {
                        fmt.Println(v7, v4, v8)
                    }
                case "":
                    var v20 int8 = 100
                    _ = v20
                }
            case 128:
                fmt.Println(v7, v1, v1)
            case -637:
                var v21 uint = (v4 % ((v4) | 1))
                _ = v21
            default:
                if true {
                    switch "a" {
                    case "":
                        var v22 float32 = 0.0
                        _ = v22
                    case "foo bar":
                        v23 := GenS1{F0: 2, F1: (v8 & v8)}
                        _ = v23
                    case "hello":
                        if false {
                            switch -9223372036854775808 {
                            case 32768:
                                var v24 uint64 = (v8 - (v8 / (((v8 * (v8 * 2))) | 1)))
                                _ = v24
                            case 4294967295:
                                var v25 int32 = 1
                                _ = v25
                            default:
                                var v26 string = "foo bar"
                                _ = v26
                            }
                        }
                    }
                } else {
                    var v27 int64 = 60
                    fmt.Println(v7, v27, v2)
                    _ = v27
                }
            }
        }
        if (!(1.5 == 1e10)) {
            for fi28 := 0; fi28 < 20; fi28++ {
                v29 := GenS1{F0: uint32(v8), F1: v8}
                for fi30 := 0; fi30 < 11; fi30++ {
                    v1 += v3
                    _ = fi30
                }
                v31 := []int{-2147483648, fi28, fi28, (fi28 ^ (-(fi28)))}
                _ = fi28
                _ = v29
                _ = v31
            }
            var v32 int32 = (v7 >> ((v4) & 63))
            switch "hello" {
            case "foo bar":
                switch 127 {
                case 128:
                    var v33 uint = 297
                    _ = v33
                }
            case "hello":
                fmt.Println(v7, v4, v8)
            case "a":
                if true {
                    if (648 >= 42) {
                        var v34 int32 = 1
                        v3 += v1
                        _ = v34
                    }
                    var v35 uint = (v4 & (v4 * (v4 ^ v4)))
                    var v36 uint32 = 42
                    _ = v35
                    _ = v36
                } else {
                    v5[65535] = 2
                    switch ("" + (("foo bar" + (("" + ("foo bar" + "\n\t")) + "foo bar")) + "")) {
                    case "hello":
                        if false {
                            v37 := []int{-2147483648, 127, -1}
                            var v38 uint64 = v8
                            _ = v37
                            _ = v38
                        }
                    case "":
                        switch "" {
                        case "hello":
                            switch 948 {
                            case 1:
                                fmt.Println(v8)
                            case 42:
                                v39 := GenS2{F0: 1, F1: 1}
                                _ = v39
                            }
                        default:
                            if (!(!(128 < 42))) {
                                if false {
                                    if (-9223372036854775808 != -1) {
                                        if true {
                                            switch (("hello" + "") + "") {
                                            case "\n\t":
                                                v7 = 1
                                            case "":
                                                v1 = v1
                                            default:
                                                var v40 float32 = 0.0
                                                _ = v40
                                            }
                                            if false {
                                                v41 := GenS1{F0: 0, F1: v8}
                                                var v42 float32 = (0.0 / (0.0 + (-2.25 + 3.14159)))
                                                _ = v41
                                                _ = v42
                                            }
                                            for fi43 := 0; fi43 < 2; fi43++ {
                                                v44 := GenS3{F0: 1}
                                                v45 := GenS1{F0: 0, F1: 1}
                                                _ = fi43
                                                _ = v44
                                                _ = v45
                                            }
                                        } else {
                                            v3 -= 1e10
                                        }
                                        if true {
                                            if ((v4 * 954) <= 0) {
                                                if true {
                                                    v8 = v8
                                                    if false {
                                                        var v46 uint64 = v8
                                                        fmt.Println(v6, v8, v6)
                                                        _ = v46
                                                    }
                                                    var v47 float32 = 0.0
                                                    _ = v47
                                                } else {
                                                    v7 -= 0
                                                }
                                                switch 42 {
                                                case 1:
                                                    if true {
                                                        for fi48 := 0; fi48 < 9; fi48++ {
                                                            for fi49 := 0; fi49 < 18; fi49++ {
                                                                if (!(1 < (v4 - v4))) {
                                                                    switch (fi49 >> ((v4) & 63)) {
                                                                    case 282:
                                                                        v50 := []int{(fi48 / ((fi49) | 1)), -20, (fi49 ^ (^(fi49)))}
                                                                        var v51 int16 = v6
                                                                        _ = v50
                                                                        _ = v51
                                                                    default:
                                                                        var v52 float32 = 1e10
                                                                        _ = v52
                                                                    }
                                                                }
                                                                _ = fi49
                                                            }
                                                            _ = fi48
                                                        }
                                                        if true {
                                                            v4 = 42
                                                            var v53 int8 = 100
                                                            _ = v53
                                                        } else {
                                                            v32 = -1
                                                            if (!(false && true)) {
                                                                v7 = 2147483647
                                                                if true {
                                                                    var v54 int64 = -1
                                                                    if true {
                                                                        v55 := GenS3{F0: v2}
                                                                        v56 := GenS3{F0: (v2 | -1)}
                                                                        v5[690] = 2
                                                                        _ = v55
                                                                        _ = v56
                                                                    } else {
                                                                        var v57 int = 32767
                                                                        _ = v57
                                                                    }
                                                                    v2 = v32
                                                                    _ = v54
                                                                }
                                                                v4 = (v4 / (((v4 / (((^(v4))) | 1))) | 1))
                                                            } else {
                                                                var v58 int32 = (v2 | (^(v7)))
                                                                if (v32 > v7) {
                                                                    if true {
                                                                        for fi59 := 0; fi59 < 5; fi59++ {
                                                                            switch "\n\t" {
                                                                            case "hello":
                                                                                v60 := GenS4{F0: true, F1: (v8 ^ v8)}
                                                                                _ = v60
                                                                            case "a":
                                                                                fmt.Println(v3)
                                                                            default:
                                                                                v7 = v58
                                                                            }
                                                                            _ = fi59
                                                                        }
                                                                        if (!(128 < 0)) {
                                                                            for fi61 := 0; fi61 < 13; fi61++ {
                                                                                var v62 uint64 = (v8 >> ((v4) & 63))
                                                                                v63 := GenS2{F0: fi61, F1: -836}
                                                                                _ = fi61
                                                                                _ = v62
                                                                                _ = v63
                                                                            }
                                                                        } else {
                                                                            for fi64 := 0; fi64 < 15; fi64++ {
                                                                                if true {
                                                                                    fmt.Println(v58)
                                                                                }
                                                                                v65 := GenS2{F0: (fi64 + (-(fi64))), F1: 1}
                                                                                v66 := GenS1{F0: 4294967295, F1: (^(v8))}
                                                                                _ = fi64
                                                                                _ = v65
                                                                                _ = v66
                                                                            }
                                                                            v1 = (v3 + (v3 * 0.0))
                                                                            if ((false || (false || false)) && ((v8 ^ 1) == (v8 % ((1) | 1)))) {
                                                                                v67 := GenS3{F0: v7}
                                                                                v68 := GenS2{F0: 9223372036854775807, F1: 9223372036854775807}
                                                                                var v69 string = "hello"
                                                                                _ = v67
                                                                                _ = v68
                                                                                _ = v69
                                                                            } else {
                                                                                var v70 uint = v4
                                                                                v71 := v5[4294967295]
                                                                                _ = v70
                                                                                _ = v71
                                                                            }
                                                                        }
                                                                        fmt.Println(v1, v4, v2)
                                                                    } else {
                                                                        v72 := GenS3{F0: v58}
                                                                        if true {
                                                                            if true {
                                                                                var v73 uint = 18446744073709551615
                                                                                var v74 int = 192
                                                                                switch ("\n\t" + "foo bar") {
                                                                                case "foo bar":
                                                                                    v8 = v8
                                                                                }
                                                                                _ = v73
                                                                                _ = v74
                                                                            } else {
                                                                                var v75 uint64 = v8
                                                                                v8 -= 4294967295
                                                                                _ = v75
                                                                            }
                                                                        }
                                                                        for fi76 := 0; fi76 < 16; fi76++ {
                                                                            var v77 int32 = int32(v4)
                                                                            _ = fi76
                                                                            _ = v77
                                                                        }
                                                                        _ = v72
                                                                    }
                                                                    fmt.Println(v8, v32, v8)
                                                                } else {
                                                                    for fi78 := 0; fi78 < 4; fi78++ {
                                                                        v79 := GenS1{F0: 42, F1: 4294967295}
                                                                        switch (^(fi78)) {
                                                                        case 9223372036854775807:
                                                                            if (147 == 128) {
                                                                                var v80 uint32 = 625
                                                                                _ = v80
                                                                            }
                                                                        case 53:
                                                                            var v81 int16 = (v6 + v6)
                                                                            _ = v81
                                                                        case 2147483647:
                                                                            if false {
                                                                                if true {
                                                                                    var v82 uint16 = 2
                                                                                    v83 := GenS5{F0: fi78, F1: 128}
                                                                                    _ = v82
                                                                                    _ = v83
                                                                                }
                                                                                v7 -= (v2 ^ (v32 << ((255) & 63)))
                                                                            } else {
                                                                                v84 := []int{}
                                                                                _ = len(v84)
                                                                                _ = v84
                                                                            }
                                                                        default:
                                                                            v85 := v5[0]
                                                                            _ = v85
                                                                        }
                                                                        delete(v5, 2)
                                                                        _ = fi78
                                                                        _ = v79
                                                                    }
                                                                    if (false || false) {
                                                                        if (4294967295 >= 4294967295) {
                                                                            switch 2147483647 {
                                                                            case -494:
                                                                                fmt.Println(v1, v4, v8)
                                                                            }
                                                                            for fi86 := 0; fi86 < 18; fi86++ {
                                                                                fmt.Println(v1, v4)
                                                                                v87 := GenS6{F0: 112}
                                                                                var v88 int64 = 1
                                                                                _ = fi86
                                                                                _ = v87
                                                                                _ = v88
                                                                            }
                                                                        }
                                                                        var v89 int16 = -750
                                                                        _ = v89
                                                                    } else {
                                                                        var v90 int16 = (^(v6))
                                                                        _ = v90
                                                                    }
                                                                }
                                                                if false {
                                                                    v91 := GenS7{F0: -1, F1: uint(v8), F2: 831}
                                                                    _ = v91
                                                                } else {
                                                                    if (-2147483648 < 9223372036854775807) {
                                                                        v8 = v8
                                                                        v2 += -1
                                                                    }
                                                                }
                                                                _ = v58
                                                            }
                                                        }
                                                    } else {
                                                        if false {
                                                            v3 = 1e10
                                                            v4 = 220
                                                            fmt.Println(v8)
                                                        } else {
                                                            delete(v5, 4294967295)
                                                            v92 := GenS8{F0: 4294967295, F1: 42}
                                                            _ = v92
                                                        }
                                                        if true {
                                                            if ((1e10 + (0.1 - float32((v3 * v3)))) == (1e10 * 0.0)) {
                                                                if ((v8 <= 4294967295) && (float64((((((float32(v1) - 0.1) * 0.0) - -2.25) - float32(float64((float32((-2.25 / v3)) / (0.0 - (3.14159 - (-2.25 / 3.14159))))))) / (1e10 / 1.0))) >= float64(3.14159))) {
                                                                    v4 += uint(v8)
                                                                    switch -2147483648 {
                                                                    case 128:
                                                                        var v93 float64 = (((float64((0.0 / 1.5)) + (v1 / -2.25)) + v1) + 1e10)
                                                                        _ = v93
                                                                    default:
                                                                        v94 := []int{4294967295}
                                                                        _ = v94
                                                                    }
                                                                } else {
                                                                    fmt.Println(v8, v6, v3)
                                                                }
                                                                switch ("hello" + "") {
                                                                case "foo bar":
                                                                    var v95 int = 105
                                                                    _ = v95
                                                                default:
                                                                    if (42 < v2) {
                                                                        delete(v5, 2147483647)
                                                                    } else {
                                                                        v5[1] = 992
                                                                        var v96 uint8 = 127
                                                                        _ = v96
                                                                    }
                                                                }
                                                            }
                                                        } else {
                                                            var v97 int32 = (-(v32))
                                                            var v98 uint32 = 4294967295
                                                            _ = v97
                                                            _ = v98
                                                        }
                                                        v2 = (v32 << (((^(v4))) & 63))
                                                    }
                                                }
                                                var v99 uint64 = (v8 | (v8 * (v8 + 915)))
                                                _ = v99
                                            } else {
                                                v100 := []int{-308}
                                                _ = v100
                                            }
                                            for fi101 := 0; fi101 < 19; fi101++ {
                                                var v102 int64 = 0
                                                if true {
                                                    var v103 uint64 = v8
                                                    if false {
                                                        for fi104 := 0; fi104 < 1; fi104++ {
                                                            fmt.Println(fi101, v1)
                                                            if (!(v2 < v7)) {
                                                                var v105 bool = true
                                                                var v106 int8 = 42
                                                                _ = v105
                                                                _ = v106
                                                            }
                                                            fmt.Println(v2, v6, fi101)
                                                            _ = fi104
                                                        }
                                                        var v107 uint8 = 2
                                                        fmt.Println(v2, v8, v102)
                                                        _ = v107
                                                    }
                                                    v8 = v8
                                                    _ = v103
                                                } else {
                                                    v108 := v5[875]
                                                    v109 := GenS9{F0: 2}
                                                    _ = v108
                                                    _ = v109
                                                }
                                                _ = fi101
                                                _ = v102
                                            }
                                            fmt.Println(v4, v6, v3)
                                        }
                                        var v110 int64 = -1
                                        _ = v110
                                    }
                                }
                            } else {
                                if (v4 > (v4 * (v4 - (v4 | v4)))) {
                                    v111 := GenS2{F0: 127, F1: -1}
                                    _ = v111
                                } else {
                                    var v112 bool = (true && true)
                                    _ = v112
                                }
                                v4 -= (v4 ^ (^(v4)))
                            }
                        }
                    default:
                        var v113 uint16 = 1
                        _ = v113
                    }
                    fmt.Println(v3, v2, v8)
                }
            default:
                var v114 int16 = 0
                _ = v114
            }
            _ = v32
        }
    }
    if (!false) {
        var v115 uint = 730
        _ = v115
    }
    v2 = (^(v2))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
