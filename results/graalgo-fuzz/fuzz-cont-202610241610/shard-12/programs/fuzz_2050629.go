package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64; F2 int }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int; F1 bool }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint8; F1 uint }
type GenS6 struct { F0 uint16; F1 uint8; F2 int64 }
type GenS7 struct { F0 bool; F1 int16; F2 string }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 string; F1 int64 }
type GenS10 struct { F0 uint8; F1 int16 }
type GenS11 struct { F0 uint16; F1 int }
type GenS12 struct { F0 int64; F1 uint; F2 int8 }
type GenS13 struct { F0 int32 }
type GenS14 struct { F0 int }

func main() {
    v1 := []uint32{1, 4294967295}
    v2 := []uint32{65535, 2147483647, 42, 42}
    v3 := make(map[string]int)
    var v4 bool = (-1 > 1)
    v5 := []int{125, 4294967295, -9223372036854775808}
    v6 := GenS1{F0: 128, F1: -275, F2: 128}
    var v7 string = "a"
    var v8 int = 32768
    v4 = true
    if ((false && (v4 || false)) || (42 < 640)) {
        for fi9 := 0; fi9 < 5; fi9++ {
            v7 = "a"
            var v10 float64 = (-2.25 * 0.1)
            _ = fi9
            _ = v10
        }
        switch "hello" {
        case "":
            switch v8 {
            case 42:
                if v4 {
                    switch "foo bar" {
                    case "\n\t":
                        v8 -= (v8 & (v8 << ((65535) & 63)))
                    case "foo bar":
                        for fi11 := 0; fi11 < 14; fi11++ {
                            v7 = "hello"
                            _ = fi11
                        }
                    default:
                        var v12 int64 = -1
                        _ = v12
                    }
                    if (!v4) {
                        switch v7 {
                        case "":
                            var v13 string = v7
                            _ = v13
                        default:
                            if (0.1 > (0.1 * 0.1)) {
                                var v14 float64 = (((1e10 * 0.1) - (0.0 * 0.1)) + 0.1)
                                _ = v14
                            }
                        }
                        fmt.Println(v8, v8)
                        for fi15 := 0; fi15 < 14; fi15++ {
                            var v16 int8 = 127
                            v17 := GenS1{F0: (fi15 ^ -1000), F1: -1, F2: fi15}
                            var v18 int = fi15
                            _ = fi15
                            _ = v16
                            _ = v17
                            _ = v18
                        }
                    }
                    for fi19 := 0; fi19 < 20; fi19++ {
                        var v20 uint32 = 4294967295
                        _ = fi19
                        _ = v20
                    }
                }
            }
        case "\n\t":
            if v4 {
                if false {
                    if v4 {
                        if false {
                            v7 = v7
                        } else {
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                            var v21 uint64 = 575
                            _ = v21
                        }
                    }
                }
                var v22 int8 = 127
                if (!(2 >= 1)) {
                    for fi23 := 0; fi23 < 9; fi23++ {
                        fmt.Println(v22, v4)
                        switch 4294967295 {
                        case -67:
                            v4 = v4
                        case 9223372036854775807:
                            if v4 {
                                if false {
                                    var v24 uint32 = 0
                                    if true {
                                        v25 := v3[(v7 + "")]
                                        var v26 int32 = 2147483647
                                        _ = v25
                                        _ = v26
                                    }
                                    _ = v24
                                }
                                v8 += v8
                            }
                        default:
                            switch "foo bar" {
                            case "a":
                                var v27 uint64 = 0
                                _ = v27
                            case "\n\t":
                                v7 = "hello"
                            default:
                                v4 = (327 < 0)
                            }
                        }
                        _ = fi23
                    }
                    if len(v1) > 0 {
                        v1[0] = 2147483647
                    }
                } else {
                    switch v7 {
                    case "hello":
                        v28 := GenS2{F0: 42}
                        _ = v28
                    case "a":
                        if v4 {
                            _ = len(v3)
                            var v29 string = (v7 + "")
                            switch ("foo bar" + (v29 + (v7 + v29))) {
                            case "a":
                                if len(v2) > 0 {
                                    v2[0] = 42
                                }
                            case "\n\t":
                                v30 := GenS3{F0: (v8 * v8), F1: (1e10 <= (3.14159 / 0.1))}
                                _ = v30
                            case "foo bar":
                                switch (v29 + v29) {
                                case "foo bar":
                                    v31 := GenS4{F0: 18446744073709551615}
                                    _ = v31
                                }
                            }
                            _ = v29
                        } else {
                            v4 = v4
                            v8 = (v8 % ((v8) | 1))
                            var v32 int32 = 42
                            _ = v32
                        }
                    }
                    v33 := GenS5{F0: 128, F1: 972}
                    _ = v33
                }
                _ = v22
            }
        }
    } else {
        switch v7 {
        case "":
            if v4 {
                switch (v8 % (((v8 % ((v8) | 1))) | 1)) {
                case -962:
                    v8 = (v8 << ((255) & 63))
                default:
                    for fi34 := 0; fi34 < 1; fi34++ {
                        fmt.Println(v7, v8)
                        _ = fi34
                    }
                }
                switch v8 {
                case -1:
                    if v4 {
                        for fi35 := 0; fi35 < 5; fi35++ {
                            if v4 {
                                v36 := GenS4{F0: 518}
                                switch (v8 & v8) {
                                case 1:
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                }
                                _ = v36
                            } else {
                                for fi37 := 0; fi37 < 3; fi37++ {
                                    var v38 bool = v4
                                    v39 := GenS6{F0: 32767, F1: 53, F2: 1}
                                    _ = fi37
                                    _ = v38
                                    _ = v39
                                }
                            }
                            var v40 string = "a"
                            v1 = append(v1, 42)
                            _ = fi35
                            _ = v40
                        }
                    }
                default:
                    if (32767 > 65535) {
                        if v4 {
                            var v41 int64 = -9223372036854775808
                            if v4 {
                                var v42 int64 = 42
                                v42 -= 563
                                _ = v42
                            }
                            _ = v41
                        } else {
                            for fi43 := 0; fi43 < 8; fi43++ {
                                var v44 uint8 = 1
                                _ = fi43
                                _ = v44
                            }
                            v7 = ((("\n\t" + "\n\t") + "") + (v7 + ""))
                            v4 = true
                        }
                        var v45 uint = 65535
                        var v46 float32 = 0.1
                        _ = v45
                        _ = v46
                    }
                }
            }
        case "foo bar":
            if (18446744073709551615 <= 0) {
                v47 := GenS7{F0: (1 > -32768), F1: -32768, F2: v7}
                _ = v47
            } else {
                var v48 string = "foo bar"
                _ = v48
            }
        }
    }
    var v49 int64 = 9223372036854775807
    var v50 int64 = -1
    for fi51 := 0; fi51 < 12; fi51++ {
        switch ((v7 + v7) + v7) {
        case "foo bar":
            if (!(1 < 100)) {
                switch (v8 << ((42) & 63)) {
                case 127:
                    switch "a" {
                    case "":
                        v50 = (v49 * (v50 + v50))
                    default:
                        for fi52 := 0; fi52 < 13; fi52++ {
                            fmt.Println(fi52, v7)
                            _ = fi52
                        }
                    }
                case -121:
                    for fi53 := 0; fi53 < 10; fi53++ {
                        v4 = (false && v4)
                        v4 = (false && v4)
                        var v54 int64 = v50
                        var v55 int8 = 42
                        v4 = (42 >= 42)
                        v56 := GenS8{F0: 1}
                        var v57 uint32 = 2147483647
                        var v58 float64 = 0.0
                        _ = fi53
                        _ = v54
                        _ = v55
                        _ = v56
                        _ = v57
                        _ = v58
                    }
                }
                v59 := GenS2{F0: 2}
                _ = v59
            }
        }
        switch (v7 + v7) {
        case "a":
            fmt.Println(v49, fi51, v4)
        case "hello":
            if true {
                fmt.Println(v49, v4)
                v60 := GenS9{F0: ("hello" + (v7 + (((v7 + ("a" + "\n\t")) + (v7 + "\n\t")) + (v7 + (("\n\t" + v7) + v7))))), F1: (v49 - (-(v49)))}
                _ = v60
            } else {
                if len(v5) > 0 {
                    v5[0] = fi51
                }
                if false {
                    switch (("a" + ((("foo bar" + "\n\t") + ("" + "a")) + ("\n\t" + v7))) + v7) {
                    case "\n\t":
                        switch v7 {
                        case "foo bar":
                            fmt.Println(v7, v49)
                        case "\n\t":
                            switch v7 {
                            case "\n\t":
                                var v61 uint64 = 547
                                _ = v61
                            }
                        case "a":
                            v62 := GenS10{F0: 4, F1: 253}
                            _ = v62
                        default:
                            fmt.Println(v7, v7, v4)
                        }
                    default:
                        switch fi51 {
                        case 32768:
                            var v63 int32 = -676
                            _ = v63
                        }
                    }
                } else {
                    var v64 float32 = float32(1.5)
                    v8 += (fi51 - v8)
                    _ = v64
                }
                if ("hello" < (v7 + "hello")) {
                    if (v4 && (!true)) {
                        var v65 uint16 = 916
                        var v66 float32 = float32((float64(0.1) + 1e10))
                        var v67 uint8 = 249
                        _ = v65
                        _ = v66
                        _ = v67
                    } else {
                        if v4 {
                            for fi68 := 0; fi68 < 14; fi68++ {
                                v7 = ((("foo bar" + v7) + v7) + v7)
                                _ = fi68
                            }
                            fmt.Println(v4, v8)
                            v4 = (255 >= 42)
                        } else {
                            v50 = (v49 << ((39) & 63))
                            switch 9223372036854775807 {
                            case -9223372036854775808:
                                for fi69 := 0; fi69 < 4; fi69++ {
                                    delete(v3, ((v7 + "\n\t") + v7))
                                    var v70 uint8 = 255
                                    _ = fi69
                                    _ = v70
                                }
                            case 42:
                                var v71 int32 = 1
                                _ = v71
                            }
                        }
                        for fi72 := 0; fi72 < 2; fi72++ {
                            v73 := GenS5{F0: 136, F1: 255}
                            delete(v3, v7)
                            _ = fi72
                            _ = v73
                        }
                    }
                    if (!(!((42 >= 2147483647) && (((!(167 >= 1)) && v4) && true)))) {
                        var v74 uint64 = 4294967295
                        switch "\n\t" {
                        case "a":
                            v7 = "hello"
                        }
                        _ = v74
                    } else {
                        if (!false) {
                            var v75 int16 = 32767
                            fmt.Println(v49, fi51, v4)
                            v76 := GenS10{F0: 255, F1: (v75 / ((v75) | 1))}
                            _ = v75
                            _ = v76
                        } else {
                            _ = len(v3)
                        }
                        v8 -= 128
                    }
                    v3[v7] = (v8 / ((fi51) | 1))
                } else {
                    if (!v4) {
                        var v77 string = v7
                        _ = v77
                    } else {
                        v4 = (!((v8 << ((0) & 63)) == (fi51 % (((fi51 * fi51)) | 1))))
                        _ = len(v3)
                        if (!(v50 > (v49 | v50))) {
                            delete(v3, v7)
                            switch 254 {
                            case 0:
                                delete(v3, (v7 + v7))
                            case 127:
                                switch v7 {
                                case "\n\t":
                                    for fi78 := 0; fi78 < 2; fi78++ {
                                        if (169 != 255) {
                                            v1 = append(v1, 2147483648)
                                            v49 -= (v50 & v50)
                                        } else {
                                            v4 = false
                                        }
                                        if v4 {
                                            if (false && (!v4)) {
                                                fmt.Println(v50, v4)
                                            }
                                            var v79 uint64 = 196
                                            var v80 uint64 = v79
                                            v80 -= (v79 * v79)
                                            var v81 uint64 = v79
                                            _ = v79
                                            _ = v80
                                            _ = v81
                                        }
                                        _ = fi78
                                    }
                                default:
                                    if (412 > 1000) {
                                        if (true && ((419 >= 0) && v4)) {
                                            v4 = (65535 <= 209)
                                        } else {
                                            switch v7 {
                                            case "foo bar":
                                                if v4 {
                                                    fmt.Println(v4, v4)
                                                    fmt.Println(v50, v4, fi51)
                                                    var v82 uint = 2
                                                    _ = v82
                                                }
                                            default:
                                                var v83 uint32 = 753
                                                _ = v83
                                            }
                                            v84 := GenS11{F0: 0, F1: (v8 * fi51)}
                                            fmt.Println(v7)
                                            _ = v84
                                        }
                                        if (0.0 < 1.5) {
                                            var v85 int64 = -1
                                            _ = v85
                                        } else {
                                            var v86 string = "hello"
                                            for fi87 := 0; fi87 < 18; fi87++ {
                                                v88 := GenS12{F0: v49, F1: 4294967295, F2: -1}
                                                switch (^(fi87)) {
                                                case -633:
                                                    delete(v3, "")
                                                }
                                                v8 -= fi51
                                                _ = fi87
                                                _ = v88
                                            }
                                            fmt.Println(v4, v4)
                                            _ = v86
                                        }
                                        for fi89 := 0; fi89 < 17; fi89++ {
                                            _ = len(v1)
                                            if true {
                                                if true {
                                                    v8 = (fi89 + (fi51 / ((fi51) | 1)))
                                                    var v90 uint64 = 0
                                                    var v91 uint = 2
                                                    v7 = v7
                                                    _ = v90
                                                    _ = v91
                                                }
                                                var v92 uint16 = 472
                                                v92 += v92
                                                _ = v92
                                            } else {
                                                var v93 string = v7
                                                if ((v49 - 42) < v49) {
                                                    v49 += 427
                                                    if v4 {
                                                        if ((!(-2.25 != ((0.1 / float32(1.5)) * (float32((0.1 + (0.1 - 3.14159))) / (1.5 * float32(1.5)))))) && (42 > 0)) {
                                                            v94 := GenS7{F0: (881 <= 1), F1: -1, F2: "foo bar"}
                                                            _ = len(v3)
                                                            if (2 <= 4294967295) {
                                                                var v95 uint64 = 18446744073709551615
                                                                _ = v95
                                                            }
                                                            _ = v94
                                                        } else {
                                                            if len(v2) >= 2 {
                                                                v2 = v2[1:len(v2)]
                                                            }
                                                            v3["a"] = 9223372036854775807
                                                            v2 = append(v2, 42)
                                                        }
                                                        v96 := GenS4{F0: 63}
                                                        _ = v96
                                                    } else {
                                                        var v97 uint8 = 42
                                                        var v98 string = v93
                                                        v99 := GenS13{F0: -1}
                                                        _ = v97
                                                        _ = v98
                                                        _ = v99
                                                    }
                                                }
                                                fmt.Println(v93, v50, v50)
                                                _ = v93
                                            }
                                            v49 += (v49 + (v49 - (v50 ^ 190)))
                                            _ = fi89
                                        }
                                    }
                                }
                            case 32767:
                                var v100 float64 = (0.0 * (1.5 + float64(0.1)))
                                _ = v100
                            default:
                                fmt.Println(v7)
                            }
                            for fi101 := 0; fi101 < 16; fi101++ {
                                var v102 int64 = (v50 >> ((0) & 63))
                                v103 := GenS8{F0: 0}
                                if v4 {
                                    var v104 uint64 = 4294967295
                                    v3["hello"] = (fi51 >> ((42) & 63))
                                    _ = v104
                                }
                                _ = fi101
                                _ = v102
                                _ = v103
                            }
                        } else {
                            for fi105 := 0; fi105 < 9; fi105++ {
                                var v106 bool = ((v49 >> ((1) & 63)) > (v50 * v50))
                                var v107 int = (fi51 >> ((2) & 63))
                                v50 = 602
                                _ = fi105
                                _ = v106
                                _ = v107
                            }
                            v108 := v3[v7]
                            switch -85 {
                            case 94:
                                fmt.Println(v4)
                            }
                            _ = v108
                        }
                    }
                    v3[v7] = (fi51 | v8)
                }
            }
        }
        var v109 int8 = -122
        _ = fi51
        _ = v109
    }
    fmt.Println(v50, v4)
    for fi110 := 0; fi110 < 13; fi110++ {
        v49 -= v49
        fmt.Println(v4)
        v111 := GenS1{F0: v8, F1: (v49 * (v50 & 0)), F2: v8}
        _ = fi110
        _ = v111
    }
    for fi112 := 0; fi112 < 16; fi112++ {
        var v113 int16 = -1
        delete(v3, "hello")
        var v114 int8 = 127
        _ = fi112
        _ = v113
        _ = v114
    }
    v115 := GenS14{F0: (v8 + v8)}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v49:", v49)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v50:", v50)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v115
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v50
    _ = v6
    _ = v7
    _ = v8
}
