package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8; F2 uint8 }
type GenS2 struct { F0 int64; F1 int }
type GenS3 struct { F0 bool; F1 uint16 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 bool; F1 string; F2 int64 }
type GenS9 struct { F0 uint32; F1 int32; F2 uint16 }
type GenS10 struct { F0 int }
type GenS11 struct { F0 uint32; F1 int16 }
type GenS12 struct { F0 bool; F1 uint64 }
func genFn0(p0 uint64, p1 uint8) uint32 {
    for fi7 := 0; fi7 < 13; fi7++ {
        fmt.Println(p0)
        _ = fi7
    }
    var v8 uint32 = 2147483647
    fmt.Println(v8)
    _ = p0
    _ = p1
    _ = v8
    return v8
}

func main() {
    var v1 int16 = -639
    v2 := make(map[uint32]int)
    var v3 uint = 4294967295
    v4 := make(map[string]string)
    v5 := []int{}
    v6 := make(map[uint32]string)
    var v9 int64 = 0
    switch 563 {
    case 1:
        var v10 int64 = 1
        _ = v10
    case -2147483648:
        var v11 float32 = (3.14159 / 1.5)
        _ = v11
    default:
        _ = len(v5)
    }
    fmt.Println(v3)
    if len(v5) > 0 {
        v5[0] = int(v1)
    }
    var v12 uint = 822
    if (3.14159 > 3.14159) {
        v13 := GenS1{F0: 4294967295, F1: 100, F2: 1}
        if true {
            switch 32767 {
            case -2147483648:
                var v14 bool = (100 >= 100)
                _ = v14
            case -760:
                v15 := GenS2{F0: v9, F1: 42}
                _ = v15
            }
            for fi16 := 0; fi16 < 15; fi16++ {
                var v17 int64 = v9
                _ = fi16
                _ = v17
            }
        }
        _ = v13
    }
    var v18 string = "hello"
    var v19 int = 355
    if false {
        switch (v19 % ((-992) | 1)) {
        case 128:
            v18 = v18
        case 32768:
            v20 := GenS3{F0: true, F1: 255}
            _ = v20
        }
    }
    v4[v18] = v18
    var v21 int16 = (v1 & (-(v1)))
    if (2 > 127) {
        for fi22 := 0; fi22 < 15; fi22++ {
            delete(v4, v18)
            switch "foo bar" {
            case "foo bar":
                var v23 bool = false
                _ = v23
            case "hello":
                for fi24 := 0; fi24 < 6; fi24++ {
                    var v25 string = v18
                    var v26 int = (v19 ^ (fi24 / ((-1) | 1)))
                    _ = fi24
                    _ = v25
                    _ = v26
                }
            case "":
                v21 = v1
            default:
                for fi27 := 0; fi27 < 14; fi27++ {
                    var v28 int = (fi22 % ((v19) | 1))
                    if len(v5) > 0 {
                        v5[0] = fi27
                    }
                    _ = fi27
                    _ = v28
                }
            }
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            _ = fi22
        }
        for fi29 := 0; fi29 < 9; fi29++ {
            if true {
                if false {
                    for fi30 := 0; fi30 < 3; fi30++ {
                        var v31 int = (-(v19))
                        var v32 int16 = (v1 / ((v1) | 1))
                        _ = fi30
                        _ = v31
                        _ = v32
                    }
                    if true {
                        if false {
                            switch fi29 {
                            case 637:
                                for fi33 := 0; fi33 < 8; fi33++ {
                                    v34 := GenS3{F0: (false && (!false)), F1: 464}
                                    delete(v2, 0)
                                    _ = fi33
                                    _ = v34
                                }
                            case -318:
                                v35 := GenS4{F0: 42}
                                _ = v35
                            default:
                                var v36 int = v19
                                _ = v36
                            }
                            for fi37 := 0; fi37 < 12; fi37++ {
                                v21 = v1
                                v12 += v3
                                fmt.Println(v12, v9, v3)
                                var v38 int = v19
                                switch -399 {
                                case 691:
                                    v4["foo bar"] = (("foo bar" + v18) + (("foo bar" + v18) + "a"))
                                case 0:
                                    if true {
                                        if ((-2.25 != (-2.25 / 1e10)) && ((true || true) || false)) {
                                            var v39 float64 = float64((0.0 / -2.25))
                                            _ = v39
                                        }
                                    }
                                    var v40 bool = false
                                    _ = v40
                                case 522:
                                    v19 = v19
                                default:
                                    if true {
                                        fmt.Println(v18)
                                        if (32767 == v1) {
                                            var v41 uint8 = 128
                                            _ = v41
                                        } else {
                                            v42 := GenS2{F0: -136, F1: 9223372036854775807}
                                            if true {
                                                var v43 uint64 = 4294967295
                                                switch -958 {
                                                case 0:
                                                    switch 127 {
                                                    case 0:
                                                        var v44 int8 = 127
                                                        _ = v44
                                                    case 128:
                                                        v12 -= 255
                                                    case 9223372036854775807:
                                                        var v45 float32 = 0.1
                                                        _ = v45
                                                    default:
                                                        v46 := GenS1{F0: (v43 >> (((v12 & v3)) & 63)), F1: 1, F2: 239}
                                                        _ = v46
                                                    }
                                                    if (fi37 == 32767) {
                                                        v3 = v12
                                                        if false {
                                                            v9 += 885
                                                        }
                                                    }
                                                    v5 = append(v5, (fi29 | -9223372036854775808))
                                                }
                                                _ = v43
                                            } else {
                                                fmt.Println(v3, fi29, v38)
                                            }
                                            v3 -= (v3 * v12)
                                            _ = len(v5)
                                            delete(v4, (v18 + ("\n\t" + "foo bar")))
                                            fmt.Println(v21, v9, fi37)
                                            v47 := GenS5{F0: (0 > 18446744073709551615)}
                                            _ = v42
                                            _ = v47
                                        }
                                    }
                                }
                                var v48 int64 = 42
                                var v49 uint8 = 128
                                _ = fi37
                                _ = v38
                                _ = v48
                                _ = v49
                            }
                            var v50 string = ""
                            _ = v50
                        }
                        var v51 int8 = 127
                        _ = v51
                    } else {
                        v18 = v18
                        var v52 string = v18
                        fmt.Println(fi29, v3)
                        _ = v52
                    }
                    v12 += (v3 << (((v12 << ((v3) & 63))) & 63))
                }
                for fi53 := 0; fi53 < 9; fi53++ {
                    if ((v9 >= v9) && (0 < 42)) {
                        if true {
                            v54 := GenS4{F0: 976}
                            var v55 int8 = -16
                            switch (fi53 & 9223372036854775807) {
                            case 0:
                                v56 := GenS6{F0: 2}
                                _ = v56
                            case 32767:
                                v57 := GenS1{F0: 18446744073709551615, F1: (^(v55)), F2: 2}
                                _ = v57
                            case 470:
                                switch fi29 {
                                case 4294967295:
                                    v21 += 1
                                default:
                                    fmt.Println(fi29, v3)
                                    switch 32768 {
                                    case 803:
                                        v19 = fi53
                                    case -2147483648:
                                        switch 885 {
                                        case 436:
                                            v58 := GenS2{F0: v9, F1: fi53}
                                            _ = v58
                                        }
                                    case 530:
                                        if true {
                                            switch "a" {
                                            case "a":
                                                v1 -= 0
                                                v59 := GenS6{F0: 1}
                                                v12 = v3
                                                _ = v59
                                            case "hello":
                                                v5 = append(v5, (fi53 ^ (fi29 * 2147483647)))
                                            }
                                            if false {
                                                v18 = (v18 + v18)
                                                fmt.Println(v18, fi29)
                                            }
                                            v12 -= 0
                                        } else {
                                            var v60 uint8 = 0
                                            v61 := GenS2{F0: v9, F1: fi53}
                                            _ = v60
                                            _ = v61
                                        }
                                    default:
                                        fmt.Println(fi53, fi53)
                                    }
                                    v9 -= 42
                                }
                            }
                            _ = v54
                            _ = v55
                        }
                        var v62 uint64 = 420
                        _ = v62
                    } else {
                        switch "\n\t" {
                        case "hello":
                            v63 := v4["a"]
                            _ = v63
                        }
                        fmt.Println(v18, v12)
                    }
                    _ = fi53
                }
                fmt.Println(fi29, v3)
            } else {
                v64 := GenS1{F0: 18446744073709551615, F1: -54, F2: 127}
                var v65 int = v19
                _ = v64
                _ = v65
            }
            if true {
                if true {
                    v66 := GenS5{F0: true}
                    _ = v66
                } else {
                    v9 = v9
                    v6[genFn0(4294967295, 128)] = v18
                }
                var v67 uint16 = 65535
                _ = v67
            } else {
                for fi68 := 0; fi68 < 5; fi68++ {
                    var v69 int8 = 127
                    if (v69 > v69) {
                        v18 = "\n\t"
                        var v70 int16 = v21
                        switch -153 {
                        case -790:
                            var v71 uint = (^(v12))
                            _ = v71
                        case 2147483647:
                            v72 := GenS7{F0: 2}
                            _ = v72
                        default:
                            fmt.Println(v69, v1, fi29)
                        }
                        _ = v70
                    } else {
                        var v73 int32 = -2147483648
                        _ = v73
                    }
                    _ = fi68
                    _ = v69
                }
                if true {
                    for fi74 := 0; fi74 < 5; fi74++ {
                        if (!((v18 + v18) > ("hello" + v18))) {
                            var v75 int = (-(v19))
                            _ = v75
                        } else {
                            v9 = -9223372036854775808
                            if len(v5) > 0 {
                                v5[0] = 42
                            }
                            var v76 uint64 = 73
                            _ = v76
                        }
                        fmt.Println(v1, v19, v3)
                        _ = fi74
                    }
                    if ((genFn0(481, 0) < 300) || ((((v9 >= v9) || (0.0 >= 0.0)) || false) || true)) {
                        if (!false) {
                            v18 = (("a" + "foo bar") + "hello")
                        } else {
                            _ = len(v5)
                            var v77 int64 = (^(v9))
                            _ = v77
                        }
                        for fi78 := 0; fi78 < 17; fi78++ {
                            v3 -= v3
                            if true {
                                fmt.Println(fi78)
                                var v79 float64 = float64(0.0)
                                v5 = append(v5, v19)
                                _ = v79
                            }
                            var v80 uint32 = genFn0(2, 127)
                            _ = fi78
                            _ = v80
                        }
                        for fi81 := 0; fi81 < 9; fi81++ {
                            switch "" {
                            case "a":
                                var v82 int32 = -1
                                _ = v82
                            case "\n\t":
                                var v83 uint8 = 0
                                _ = v83
                            }
                            _ = fi81
                        }
                    } else {
                        v84 := GenS8{F0: false, F1: v18, F2: 1}
                        var v85 int16 = 103
                        _ = v84
                        _ = v85
                    }
                    switch v19 {
                    case 32768:
                        if (((false || (153 < 0)) || (1 >= 39)) && true) {
                            if (1.5 >= 0.0) {
                                fmt.Println(v9, v12, v19)
                                v12 -= (v3 & 567)
                                v2[1] = (-(fi29))
                            }
                        } else {
                            for fi86 := 0; fi86 < 9; fi86++ {
                                delete(v2, genFn0(18446744073709551615, 128))
                                if false {
                                    v19 = (fi86 & fi86)
                                    switch (v18 + "foo bar") {
                                    case "foo bar":
                                        var v87 int = -9223372036854775808
                                        _ = v87
                                    case "\n\t":
                                        var v88 int32 = 1
                                        _ = v88
                                    default:
                                        v21 += -32768
                                    }
                                }
                                _ = fi86
                            }
                            if (230 == 2) {
                                if false {
                                    switch (v19 << (((^(v3))) & 63)) {
                                    case 128:
                                        switch (v18 + v18) {
                                        case "":
                                            for fi89 := 0; fi89 < 2; fi89++ {
                                                fmt.Println(fi29, fi29, fi29)
                                                switch "a" {
                                                case "hello":
                                                    switch -663 {
                                                    case 32767:
                                                        var v90 float64 = ((float64(1.5) * 0.0) - ((0.1 + 3.14159) - 1.5))
                                                        _ = v90
                                                    case 128:
                                                        switch fi29 {
                                                        case 221:
                                                            v9 = (v9 + int64(v12))
                                                        }
                                                    }
                                                case "\n\t":
                                                    v19 -= (v19 << ((v12) & 63))
                                                case "a":
                                                    var v91 uint16 = 42
                                                    _ = v91
                                                }
                                                _ = fi89
                                            }
                                        }
                                    case 32768:
                                        var v92 bool = (!true)
                                        _ = v92
                                    default:
                                        fmt.Println(v9, v21, fi29)
                                    }
                                }
                                var v93 uint64 = 4294967295
                                if (v12 <= 65535) {
                                    var v94 bool = false
                                    _ = v94
                                }
                                _ = v93
                            }
                        }
                    case 1:
                        fmt.Println(v9, v21)
                    default:
                        for fi95 := 0; fi95 < 1; fi95++ {
                            var v96 uint32 = 1
                            fmt.Println(v96, v12)
                            switch 32768 {
                            case -233:
                                v97 := GenS9{F0: v96, F1: 8, F2: 42}
                                _ = v97
                            }
                            v98 := GenS3{F0: true, F1: 32767}
                            v99 := v4[(v18 + v18)]
                            fmt.Println(v9, v18, v12)
                            _ = fi95
                            _ = v96
                            _ = v98
                            _ = v99
                        }
                    }
                } else {
                    v100 := GenS3{F0: false, F1: 1}
                    var v101 int = fi29
                    _ = v100
                    _ = v101
                }
            }
            if (!true) {
                var v102 int16 = int16(fi29)
                if (-1 < 42) {
                    v1 = v21
                    var v103 int64 = (v9 / ((-1) | 1))
                    _ = v103
                }
                _ = v102
            } else {
                switch "\n\t" {
                case "hello":
                    v104 := GenS10{F0: (fi29 << (((v12 << (((v12 / (((v12 - 255)) | 1))) & 63))) & 63))}
                    _ = v104
                case "foo bar":
                    _ = len(v5)
                }
            }
            _ = fi29
        }
        v105 := GenS11{F0: 4294967295, F1: 1000}
        _ = v105
    } else {
        v106 := GenS12{F0: false, F1: 2}
        var v107 uint16 = 528
        _ = v106
        _ = v107
    }
    var v108 int8 = int8(v1)
    fmt.Println("v1:", v1)
    fmt.Println("v108:", v108)
    fmt.Println("v12:", v12)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v108
    _ = v12
    _ = v18
    _ = v19
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
