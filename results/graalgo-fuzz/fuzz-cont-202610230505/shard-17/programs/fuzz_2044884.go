package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string; F2 uint64 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint64; F1 int32; F2 uint8 }
type GenS4 struct { F0 int64; F1 int32; F2 uint16 }
type GenS5 struct { F0 int; F1 uint8; F2 uint }
type GenS6 struct { F0 int16; F1 uint; F2 uint32 }
type GenS7 struct { F0 uint64; F1 int32 }
type GenS8 struct { F0 int64; F1 uint32; F2 int32 }
type GenS9 struct { F0 string; F1 int16; F2 uint32 }
type GenS10 struct { F0 int32; F1 uint8; F2 int }
type GenS11 struct { F0 string; F1 int }
type GenS12 struct { F0 uint; F1 int }
type GenS13 struct { F0 int32; F1 int8; F2 uint8 }
func genFn0() int64 {
    v6 := []int{9223372036854775807, 2147483647, 0, -247}
    var v7 float32 = 0.1
    var v8 bool = (!true)
    v8 = ((!true) || v8)
    var v9 int8 = 42
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 9223372036854775807
}

func main() {
    var v1 float32 = -2.25
    var v2 uint8 = 0
    var v3 int = 4294967295
    var v4 int = v3
    var v5 int = (v3 + v3)
    var v10 uint = 18446744073709551615
    if (!(!((v1 >= v1) || (((413 != 65535) && true) && true)))) {
        switch "foo bar" {
        case "\n\t":
            fmt.Println(v10, v2)
        case "foo bar":
            if false {
                v11 := GenS1{F0: 0, F1: ("foo bar" + (((("foo bar" + "hello") + "\n\t") + ("\n\t" + "foo bar")) + "foo bar")), F2: 42}
                v3 = v3
                v10 = v10
                _ = v11
            } else {
                switch "\n\t" {
                case "foo bar":
                    v12 := GenS1{F0: -32768, F1: "foo bar", F2: 465}
                    _ = v12
                case "hello":
                    for fi13 := 0; fi13 < 19; fi13++ {
                        v2 = 0
                        v14 := GenS2{F0: -472}
                        _ = fi13
                        _ = v14
                    }
                default:
                    var v15 int = v3
                    _ = v15
                }
                for fi16 := 0; fi16 < 19; fi16++ {
                    v17 := GenS2{F0: fi16}
                    for fi18 := 0; fi18 < 10; fi18++ {
                        v5 += (-(fi16))
                        _ = fi18
                    }
                    _ = fi16
                    _ = v17
                }
            }
        default:
            v19 := GenS3{F0: 367, F1: -866, F2: v2}
            _ = v19
        }
        for fi20 := 0; fi20 < 5; fi20++ {
            switch "" {
            case "\n\t":
                fmt.Println(v10)
            }
            _ = fi20
        }
    }
    switch "hello" {
    case "a":
        switch ((("\n\t" + ("\n\t" + "a")) + "a") + "hello") {
        case "\n\t":
            if ((false || false) && (!false)) {
                var v21 bool = (500 >= 42)
                v1 -= v1
                _ = v21
            }
        case "hello":
            v22 := []int{(-(v3)), 127, -9223372036854775808, -1}
            _ = v22
        }
    case "hello":
        switch ("" + "hello") {
        case "hello":
            fmt.Println(v5, v2, v2)
        default:
            if true {
                fmt.Println(v10)
            } else {
                var v23 int8 = 0
                fmt.Println(v4)
                _ = v23
            }
        }
    default:
        if true {
            v3 = (v3 << (((v10 + (v10 & v10))) & 63))
        } else {
            for fi24 := 0; fi24 < 12; fi24++ {
                var v25 uint64 = 934
                _ = fi24
                _ = v25
            }
        }
    }
    var v26 int64 = 371
    for fi27 := 0; fi27 < 10; fi27++ {
        v28 := []int{}
        _ = fi27
        _ = v28
    }
    if false {
        v26 = v26
        for fi29 := 0; fi29 < 18; fi29++ {
            if true {
                var v30 string = "hello"
                v31 := GenS3{F0: 952, F1: -1, F2: (v2 >> (((v10 >> ((uint(v2)) & 63))) & 63))}
                _ = v30
                _ = v31
            }
            var v32 float32 = (float32(0.1) - 1.5)
            if false {
                v3 = 2147483647
                v1 -= v1
                if (18446744073709551615 != 0) {
                    switch "foo bar" {
                    case "a":
                        var v33 uint32 = 925
                        _ = v33
                    default:
                        v34 := []int{v4, (v4 - (v3 | (v3 << ((0) & 63)))), (fi29 | v5)}
                        _ = v34
                    }
                } else {
                    if ((v26 >> (((v10 >> (((^(v10))) & 63))) & 63)) != v26) {
                        v2 = v2
                        fmt.Println(v4)
                    }
                    for fi35 := 0; fi35 < 13; fi35++ {
                        v36 := GenS3{F0: 347, F1: 42, F2: 231}
                        if (!false) {
                            v37 := []int{-1, fi35, fi35, (fi29 >> (((^(v10))) & 63))}
                            _ = v37
                        } else {
                            fmt.Println(fi29, v3)
                            switch (("foo bar" + "a") + (("hello" + "hello") + "a")) {
                            case "\n\t":
                                v38 := GenS1{F0: -174, F1: ("\n\t" + "a"), F2: 42}
                                _ = v38
                            case "hello":
                                v26 = v26
                            case "foo bar":
                                var v39 float64 = ((0.0 / ((0.1 + float64(v32)) - float64(v1))) / ((1e10 / float64((3.14159 / (0.0 - v32)))) / (3.14159 * -2.25)))
                                _ = v39
                            default:
                                if ((false && (!(true && (true || (fi29 <= (v4 - fi35)))))) || false) {
                                    var v40 bool = false
                                    var v41 uint64 = 18446744073709551615
                                    _ = v40
                                    _ = v41
                                } else {
                                    fmt.Println(v1, v1, v32)
                                }
                            }
                            v42 := GenS4{F0: v26, F1: 1, F2: 1}
                            v10 = (^(v10))
                            _ = v42
                        }
                        _ = fi35
                        _ = v36
                    }
                    switch ("" + "a") {
                    case "a":
                        for fi43 := 0; fi43 < 13; fi43++ {
                            var v44 float64 = 0.1
                            v45 := GenS2{F0: (v5 << (((v10 * 2)) & 63))}
                            if true {
                                var v46 uint32 = 689
                                var v47 int32 = -194
                                fmt.Println(v10, v2)
                                v48 := GenS1{F0: -1, F1: "", F2: 131}
                                switch ("" + "hello") {
                                case "":
                                    if true {
                                        switch 42 {
                                        case 42:
                                            var v49 int = v5
                                            switch 9223372036854775807 {
                                            case -2147483648:
                                                var v50 int16 = -32768
                                                _ = v50
                                            }
                                            var v51 float64 = (float64(1e10) - v44)
                                            _ = v49
                                            _ = v51
                                        default:
                                            var v52 uint8 = v2
                                            _ = v52
                                        }
                                        v53 := []int{}
                                        var v54 float32 = v32
                                        _ = v53
                                        _ = v54
                                    }
                                case "foo bar":
                                    v55 := []int{}
                                    _ = v55
                                }
                                v32 += (((v1 * v1) + (0.1 * (((0.1 - v32) + 0.1) * (v1 / v32)))) - v1)
                                v56 := []int{fi29, int(v26), -217}
                                v57 := GenS5{F0: (fi43 + v4), F1: (v2 - v2), F2: 0}
                                if (0 != 0) {
                                    if len(v56) >= 2 {
                                        v56 = v56[1:len(v56)]
                                    }
                                    fmt.Println(v46)
                                    if true {
                                        switch "foo bar" {
                                        case "a":
                                            if len(v56) > 0 {
                                                v56[0] = -258
                                            }
                                        case "foo bar":
                                            var v58 float32 = (v1 / 0.1)
                                            _ = v58
                                        case "":
                                            switch ("hello" + "\n\t") {
                                            case "a":
                                                v26 = v26
                                            case "hello":
                                                fmt.Println(v46, v44)
                                            default:
                                                switch "foo bar" {
                                                case "foo bar":
                                                    if true {
                                                        if false {
                                                            var v59 uint64 = 2
                                                            _ = v59
                                                        } else {
                                                            fmt.Println(v1, fi29, fi43)
                                                        }
                                                        switch (fi29 << ((v10) & 63)) {
                                                        case 1:
                                                            if (0 != 2) {
                                                                v56 = append(v56, (fi29 + 32767))
                                                                v44 += (1.5 - (1.5 - 1e10))
                                                                fmt.Println(fi29, v32)
                                                            } else {
                                                                var v60 int16 = 769
                                                                if len(v56) >= 2 {
                                                                    v56 = v56[1:len(v56)]
                                                                }
                                                                v5 -= v5
                                                                _ = v60
                                                            }
                                                        case 4294967295:
                                                            switch "foo bar" {
                                                            case "a":
                                                                v32 = v32
                                                                v32 -= 1e10
                                                            case "\n\t":
                                                                if len(v56) > 0 {
                                                                    v56[0] = (v3 - (v5 >> ((18446744073709551615) & 63)))
                                                                }
                                                            case "":
                                                                var v61 uint16 = 973
                                                                _ = v61
                                                            }
                                                            var v62 int64 = 1
                                                            if (false && ((v10 / ((v10) | 1)) != (v10 >> (((v10 >> ((2) & 63))) & 63)))) {
                                                                v44 += float64(v32)
                                                            }
                                                            _ = v62
                                                        case -583:
                                                            var v63 int16 = 42
                                                            _ = v63
                                                        default:
                                                            var v64 float64 = float64((v32 - v1))
                                                            _ = v64
                                                        }
                                                        v5 = v3
                                                    }
                                                case "":
                                                    v65 := GenS6{F0: 1000, F1: 965, F2: v46}
                                                    _ = v65
                                                }
                                                v66 := GenS2{F0: fi29}
                                                if len(v56) > 0 {
                                                    v56[0] = (v5 / ((v4) | 1))
                                                }
                                                _ = v66
                                            }
                                        default:
                                            switch "" {
                                            case "a":
                                                if true {
                                                    fmt.Println(v32)
                                                    var v67 uint = (^(v10))
                                                    var v68 float64 = (1e10 / v44)
                                                    v69 := GenS7{F0: 42, F1: -221}
                                                    if ((v2 + v2) == v2) {
                                                        switch "" {
                                                        case "":
                                                            var v70 uint32 = v46
                                                            _ = v70
                                                        default:
                                                            v10 = v10
                                                            var v71 uint64 = 0
                                                            _ = v71
                                                        }
                                                        if true {
                                                            switch -108 {
                                                            case -78:
                                                                v5 = (v5 | v4)
                                                            case -2147483648:
                                                                if (v1 > 0.0) {
                                                                    v32 -= 1e10
                                                                } else {
                                                                    v72 := GenS1{F0: 0, F1: "hello", F2: 1}
                                                                    v26 = (v26 % ((genFn0()) | 1))
                                                                    _ = v72
                                                                }
                                                            case 127:
                                                                switch ("\n\t" + "\n\t") {
                                                                case "hello":
                                                                    v73 := GenS8{F0: (v26 % ((int64(v47)) | 1)), F1: uint32(v67), F2: (v47 << ((v10) & 63))}
                                                                    _ = v73
                                                                default:
                                                                    v74 := GenS9{F0: "a", F1: -964, F2: 2}
                                                                    var v75 bool = (!(!(1 > 0)))
                                                                    _ = v74
                                                                    _ = v75
                                                                }
                                                            default:
                                                                if (false || (v44 < v44)) {
                                                                    if false {
                                                                        if len(v56) >= 2 {
                                                                            v56 = v56[1:len(v56)]
                                                                        }
                                                                    } else {
                                                                        v67 = v10
                                                                        v3 += 812
                                                                        var v76 bool = ((v26 > 9223372036854775807) && true)
                                                                        _ = v76
                                                                    }
                                                                    if (((uint16(v67) >= 1) || true) || (!(!false))) {
                                                                        v44 = v68
                                                                        v1 += v32
                                                                        v77 := GenS10{F0: v47, F1: uint8(v10), F2: -637}
                                                                        v78 := GenS11{F0: ("hello" + (("hello" + (("a" + (("a" + "hello") + "\n\t")) + ((("foo bar" + "hello") + "\n\t") + ""))) + "a")), F1: fi29}
                                                                        fmt.Println(fi43)
                                                                        _ = v77
                                                                        _ = v78
                                                                    }
                                                                }
                                                            }
                                                            if false {
                                                                fmt.Println(v44, v1)
                                                            }
                                                        }
                                                    } else {
                                                        var v79 int32 = v47
                                                        _ = v79
                                                    }
                                                    _ = v67
                                                    _ = v68
                                                    _ = v69
                                                }
                                                var v80 int32 = 208
                                                v81 := GenS9{F0: "\n\t", F1: 1000, F2: (v46 % (((v46 | (^(v46)))) | 1))}
                                                _ = v80
                                                _ = v81
                                            case "":
                                                v56 = append(v56, fi29)
                                            case "foo bar":
                                                var v82 bool = ("a" == "\n\t")
                                                _ = v82
                                            }
                                        }
                                    }
                                }
                                _ = v46
                                _ = v47
                                _ = v48
                                _ = v56
                                _ = v57
                            }
                            _ = fi43
                            _ = v44
                            _ = v45
                        }
                    case "hello":
                        v83 := []int{}
                        _ = v83
                    case "":
                        if false {
                            var v84 int8 = 100
                            fmt.Println(v84)
                            fmt.Println(v3)
                            _ = v84
                        }
                    }
                }
            } else {
                switch fi29 {
                case -1:
                    v1 -= 3.14159
                default:
                    var v85 string = "hello"
                    _ = v85
                }
                v86 := []int{-197}
                _ = v86
            }
            _ = fi29
            _ = v32
        }
    } else {
        if (!(true && true)) {
            v1 += (0.0 + (((v1 / (v1 * v1)) - ((0.1 / (1.5 - -2.25)) / 0.1)) / 3.14159))
        } else {
            switch ("foo bar" + "a") {
            case "a":
                v87 := []int{(v4 >> (((^(v10))) & 63))}
                _ = v87
            }
        }
        fmt.Println(v10, v3)
    }
    var v88 uint16 = uint16(v2)
    if false {
        v88 -= (^(v88))
        for fi89 := 0; fi89 < 6; fi89++ {
            for fi90 := 0; fi90 < 16; fi90++ {
                v5 = (fi90 * (fi90 / (((v4 / ((v3) | 1))) | 1)))
                var v91 int8 = -34
                v5 -= (v3 << (((v10 | (^(v10)))) & 63))
                _ = fi90
                _ = v91
            }
            if false {
                var v92 float64 = 3.14159
                v2 += v2
                _ = v92
            } else {
                v93 := GenS12{F0: 1, F1: fi89}
                if false {
                    v94 := GenS10{F0: -1, F1: v2, F2: 128}
                    _ = v94
                } else {
                    var v95 uint32 = 2147483647
                    _ = v95
                }
                switch "foo bar" {
                case "\n\t":
                    var v96 int32 = -1
                    _ = v96
                default:
                    var v97 float32 = v1
                    _ = v97
                }
                _ = v93
            }
            var v98 int64 = -748
            _ = fi89
            _ = v98
        }
        switch "" {
        case "foo bar":
            for fi99 := 0; fi99 < 18; fi99++ {
                for fi100 := 0; fi100 < 14; fi100++ {
                    if false {
                        v101 := []int{}
                        _ = v101
                    } else {
                        v4 = fi99
                        var v102 uint32 = 2
                        switch ("hello" + "\n\t") {
                        case "a":
                            switch v3 {
                            case 128:
                                v102 += v102
                            }
                        case "hello":
                            v103 := GenS9{F0: (("\n\t" + "\n\t") + (("hello" + "hello") + "")), F1: 1000, F2: v102}
                            _ = v103
                        }
                        _ = v102
                    }
                    if true {
                        switch (-(fi100)) {
                        case -220:
                            switch "foo bar" {
                            case "":
                                switch v5 {
                                case 963:
                                    var v104 int16 = 0
                                    _ = v104
                                default:
                                    var v105 uint64 = 984
                                    _ = v105
                                }
                            }
                        case 42:
                            v26 = v26
                        }
                        v5 += v3
                    } else {
                        v2 = v2
                        v106 := []int{}
                        v5 = fi99
                        _ = v106
                    }
                    var v107 float64 = (0.0 + 0.1)
                    _ = fi100
                    _ = v107
                }
                for fi108 := 0; fi108 < 15; fi108++ {
                    if (false || false) {
                        var v109 float64 = 1.5
                        v110 := GenS13{F0: -1, F1: 100, F2: (^(v2))}
                        fmt.Println(v1, v2)
                        v1 = (v1 + v1)
                        v88 -= v88
                        v109 = (v109 + v109)
                        v3 -= (^(fi108))
                        _ = v109
                        _ = v110
                    } else {
                        v4 += (v3 * v5)
                        v1 += (v1 - 0.1)
                        v111 := []int{v4, fi99}
                        _ = v111
                    }
                    v26 -= genFn0()
                    _ = fi108
                }
                for fi112 := 0; fi112 < 12; fi112++ {
                    v88 -= (v88 / ((v88) | 1))
                    _ = fi112
                }
                _ = fi99
            }
        }
    }
    v3 = 4294967295
    fmt.Println(v3, v5)
    v26 += 1
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v88:", v88)
    _ = v1
    _ = v10
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v88
}
