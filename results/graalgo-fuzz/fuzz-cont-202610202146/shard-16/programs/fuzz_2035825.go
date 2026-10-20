package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32; F2 uint64 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int64; F1 uint }
type GenS4 struct { F0 uint32; F1 uint16 }
type GenS5 struct { F0 uint64; F1 uint32 }
type GenS6 struct { F0 uint8; F1 uint16 }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 uint; F1 int }
type GenS10 struct { F0 uint16; F1 int16; F2 int16 }
type GenS11 struct { F0 uint16 }
type GenS12 struct { F0 int64; F1 uint; F2 uint64 }

func main() {
    v1 := make(map[uint32]int)
    var v2 uint16 = 277
    v3 := make(map[string]int)
    for fi4 := 0; fi4 < 13; fi4++ {
        for fi5 := 0; fi5 < 5; fi5++ {
            var v6 float32 = ((1.5 - ((0.0 / float32(3.14159)) - ((-2.25 + 3.14159) + float32((0.1 - 0.0))))) + 3.14159)
            _ = fi5
            _ = v6
        }
        v2 += 65535
        _ = fi4
    }
    for fi7 := 0; fi7 < 10; fi7++ {
        for fi8 := 0; fi8 < 20; fi8++ {
            v2 = v2
            if (255 <= 2) {
                var v9 int = (fi7 << ((290) & 63))
                _ = v9
            } else {
                v2 = (v2 - 42)
                v10 := v1[2]
                if false {
                    v11 := []int{(v10 * fi7)}
                    v12 := GenS1{F0: 1, F1: -850, F2: 4294967295}
                    v13 := GenS1{F0: 42, F1: -2147483648, F2: 864}
                    _ = v11
                    _ = v12
                    _ = v13
                } else {
                    v10 = fi7
                }
                v2 -= (v2 % ((366) | 1))
                _ = v10
            }
            if false {
                var v14 float32 = 0.1
                var v15 uint = 18446744073709551615
                var v16 uint16 = (^(v2))
                _ = v14
                _ = v15
                _ = v16
            } else {
                var v17 float64 = 1e10
                _ = v17
            }
            _ = fi8
        }
        fmt.Println(v2)
        _ = fi7
    }
    v2 -= v2
    v1[0] = 0
    for fi18 := 0; fi18 < 6; fi18++ {
        for fi19 := 0; fi19 < 1; fi19++ {
            var v20 uint32 = 2
            v21 := GenS2{F0: -100}
            _ = fi19
            _ = v20
            _ = v21
        }
        v22 := GenS1{F0: 303, F1: 0, F2: 1}
        v2 -= v2
        _ = fi18
        _ = v22
    }
    v2 -= 32767
    v23 := v1[2147483647]
    var v24 uint32 = 677
    if ((3.14159 != (((3.14159 / 1.0) * ((1e10 * 1.5) * ((float32(1.5) * 1.5) * 1.5))) / (0.1 / 1e10))) && (42 >= 4294967295)) {
        if false {
            var v25 uint16 = 220
            switch "foo bar" {
            case "\n\t":
                v2 += v2
            default:
                v23 -= (v23 % ((9223372036854775807) | 1))
            }
            v25 = (v25 | v2)
            _ = v25
        }
        var v26 int8 = 0
        v24 -= 4294967295
        _ = v26
    } else {
        switch "a" {
        case "":
            v23 = (^(v23))
        case "\n\t":
            for fi27 := 0; fi27 < 12; fi27++ {
                if true {
                    v28 := v1[(v24 | (v24 - 42))]
                    _ = v28
                }
                _ = fi27
            }
        case "foo bar":
            if false {
                v23 = -2147483648
                v3["hello"] = v23
                for fi29 := 0; fi29 < 11; fi29++ {
                    var v30 int16 = 42
                    switch "\n\t" {
                    case "a":
                        var v31 uint16 = (v2 << ((1) & 63))
                        _ = v31
                    case "foo bar":
                        switch "\n\t" {
                        case "":
                            v30 = v30
                        case "hello":
                            for fi32 := 0; fi32 < 15; fi32++ {
                                var v33 uint16 = 202
                                _ = fi32
                                _ = v33
                            }
                        default:
                            v34 := []int{fi29, fi29, -746, int(v24)}
                            _ = v34
                        }
                    case "hello":
                        v1[(v24 - 2147483648)] = 42
                    default:
                        v35 := GenS3{F0: -1, F1: 1}
                        _ = v35
                    }
                    if false {
                        v36 := GenS3{F0: -391, F1: 4294967295}
                        var v37 int8 = 42
                        _ = v36
                        _ = v37
                    }
                    _ = fi29
                    _ = v30
                }
            } else {
                v38 := GenS4{F0: (v24 % (((v24 % (((v24 * (v24 + v24))) | 1))) | 1)), F1: 464}
                _ = v38
            }
        }
        for fi39 := 0; fi39 < 1; fi39++ {
            for fi40 := 0; fi40 < 3; fi40++ {
                switch fi39 {
                case 9223372036854775807:
                    if (false && (4294967295 == 2)) {
                        v24 = (v24 ^ 2)
                        v24 -= v24
                    }
                case 441:
                    switch (fi40 | fi40) {
                    case -988:
                        if false {
                            var v41 uint = 18446744073709551615
                            v42 := GenS2{F0: -128}
                            var v43 uint = (^(v41))
                            _ = v41
                            _ = v42
                            _ = v43
                        }
                    case 210:
                        v44 := []int{fi39}
                        _ = v44
                    case -195:
                        var v45 int = fi40
                        _ = v45
                    default:
                        v46 := GenS5{F0: 1, F1: (v24 & v24)}
                        _ = v46
                    }
                default:
                    v23 = (fi39 * 32768)
                }
                _ = fi40
            }
            var v47 bool = ((true || false) || ((!true) && true))
            _ = fi39
            _ = v47
        }
        var v48 string = "hello"
        _ = v48
    }
    v49 := GenS3{F0: -9223372036854775808, F1: uint(v24)}
    for fi50 := 0; fi50 < 6; fi50++ {
        for fi51 := 0; fi51 < 16; fi51++ {
            switch "hello" {
            case "foo bar":
                v52 := GenS3{F0: 382, F1: 512}
                _ = v52
            case "\n\t":
                var v53 int32 = 42
                _ = v53
            case "a":
                var v54 uint32 = v24
                if ((v2 + 1) <= (v2 + v2)) {
                    var v55 uint32 = v24
                    var v56 string = "foo bar"
                    _ = v55
                    _ = v56
                } else {
                    v57 := GenS6{F0: 86, F1: uint16(v54)}
                    _ = v57
                }
                if true {
                    if ("\n\t" >= "hello") {
                        switch (fi51 << ((0) & 63)) {
                        case 2147483647:
                            v58 := GenS4{F0: 706, F1: (v2 % (((v2 / (((^(v2))) | 1))) | 1))}
                            _ = v58
                        }
                    }
                    var v59 uint32 = (v54 - 620)
                    switch "" {
                    case "hello":
                        v23 = -9223372036854775808
                    case "a":
                        if false {
                            var v60 int8 = 100
                            if true {
                                if (255 > uint(v2)) {
                                    v2 += v2
                                    v61 := GenS1{F0: -753, F1: -27, F2: 2}
                                    _ = v61
                                } else {
                                    switch 9223372036854775807 {
                                    case 85:
                                        v62 := v1[(v24 << ((0) & 63))]
                                        _ = v62
                                    case -9223372036854775808:
                                        if ((0 != v60) || (v24 == v24)) {
                                            var v63 int = v23
                                            fmt.Println(v23, v54)
                                            _ = v63
                                        } else {
                                            v64 := v1[(^(v24))]
                                            _ = v64
                                        }
                                    case 282:
                                        fmt.Println(v23)
                                    default:
                                        if false {
                                            v65 := GenS1{F0: 42, F1: 42, F2: 369}
                                            if false {
                                                v66 := []int{fi51, fi51, 32768, -1}
                                                _ = v66
                                            }
                                            v67 := GenS6{F0: 1, F1: v2}
                                            _ = v65
                                            _ = v67
                                        } else {
                                            v3["a"] = v23
                                        }
                                        v68 := v1[v24]
                                        v69 := []int{}
                                        _ = v68
                                        _ = v69
                                    }
                                    var v70 uint64 = 4294967295
                                    _ = v70
                                }
                                fmt.Println(v23)
                            }
                            switch ("foo bar" + (((("a" + "foo bar") + "hello") + "a") + "\n\t")) {
                            case "":
                                delete(v1, (^(v59)))
                                if false {
                                    _ = len(v3)
                                    if (8 != 32767) {
                                        if (255 != 45) {
                                            switch fi50 {
                                            case 127:
                                                if false {
                                                    v71 := []int{(fi50 % ((0) | 1)), (-(fi51))}
                                                    if len(v71) > 0 {
                                                        v71[0] = 616
                                                    }
                                                    _ = v71
                                                }
                                            case -645:
                                                v2 += 2
                                            }
                                            v72 := GenS7{F0: (v60 >> ((953) & 63))}
                                            _ = v72
                                        } else {
                                            var v73 int64 = -1
                                            v60 -= 127
                                            fmt.Println(v54, v2, v73)
                                            _ = v73
                                        }
                                    } else {
                                        if false {
                                            if false {
                                                var v74 float64 = (3.14159 / 3.14159)
                                                _ = v74
                                            }
                                            v75 := GenS7{F0: 42}
                                            v54 += (v54 >> ((18446744073709551615) & 63))
                                            _ = v75
                                        }
                                        _ = len(v1)
                                        var v76 float32 = -2.25
                                        _ = v76
                                    }
                                    v77 := []int{(fi51 - fi51)}
                                    var v78 int32 = -41
                                    switch ("foo bar" + "hello") {
                                    case "\n\t":
                                        var v79 uint16 = (v2 * (v2 % (((^(v2))) | 1)))
                                        _ = v79
                                    case "foo bar":
                                        if false {
                                            if len(v77) > 0 {
                                                v77[0] = v23
                                            }
                                        } else {
                                            v60 += 1
                                            if len(v77) >= 2 {
                                                v77 = v77[1:len(v77)]
                                            }
                                            v24 -= (v54 >> ((789) & 63))
                                        }
                                    case "":
                                        v59 = (v59 ^ v24)
                                    }
                                    _ = v77
                                    _ = v78
                                } else {
                                    switch "foo bar" {
                                    case "":
                                        v80 := v3["foo bar"]
                                        _ = v80
                                    case "foo bar":
                                        var v81 int8 = (v60 >> ((982) & 63))
                                        _ = v81
                                    }
                                }
                            }
                            v82 := []int{(fi50 | (fi51 * fi50))}
                            v83 := GenS1{F0: -225, F1: -89, F2: 18446744073709551615}
                            _ = v60
                            _ = v82
                            _ = v83
                        } else {
                            v84 := GenS1{F0: 42, F1: 42, F2: 1}
                            _ = v84
                        }
                    case "\n\t":
                        v24 += (v24 % (((v59 >> ((889) & 63))) | 1))
                        if false {
                            v85 := GenS4{F0: v24, F1: (^(v2))}
                            _ = v85
                        }
                        v86 := []int{fi50, 2147483647, 0, fi51}
                        _ = v86
                    default:
                        if true {
                            v87 := GenS8{F0: 42}
                            v88 := GenS7{F0: -128}
                            _ = v87
                            _ = v88
                        }
                    }
                    _ = v59
                } else {
                    if true {
                        v23 += (fi50 >> ((4294967295) & 63))
                        var v89 int64 = 125
                        v90 := []int{fi50, (v23 + v23), (v23 * (-(fi50)))}
                        var v91 int32 = 42
                        fmt.Println(v24, fi50, v24)
                        switch ("" + "\n\t") {
                        case "hello":
                            switch 0 {
                            case 32767:
                                v92 := GenS7{F0: -9}
                                _ = v92
                            default:
                                if len(v90) >= 2 {
                                    v90 = v90[1:len(v90)]
                                }
                            }
                        case "":
                            v23 = 855
                        case "foo bar":
                            v2 -= (^(v2))
                        }
                        switch fi51 {
                        case -2147483648:
                            if len(v90) >= 2 {
                                v90 = v90[1:len(v90)]
                            }
                        }
                        switch (fi51 ^ (v23 - (v23 / (((fi51 & (fi50 % ((v23) | 1)))) | 1)))) {
                        case 544:
                            if (2147483647 >= (v91 & (v91 - (v91 ^ v91)))) {
                                if len(v90) > 0 {
                                    v90[0] = 2147483647
                                }
                                if len(v90) >= 2 {
                                    v90 = v90[1:len(v90)]
                                }
                                if true {
                                    switch "" {
                                    case "foo bar":
                                        v93 := GenS9{F0: 4294967295, F1: (fi51 + fi51)}
                                        _ = v93
                                    case "hello":
                                        var v94 int32 = v91
                                        _ = v94
                                    }
                                }
                            }
                        }
                        fmt.Println(v91, fi51, v23)
                        _ = v89
                        _ = v90
                        _ = v91
                    } else {
                        v2 = v2
                    }
                    v3["\n\t"] = fi51
                    v24 -= 2147483648
                }
                _ = v54
            }
            _ = fi51
        }
        fmt.Println(fi50)
        _ = fi50
    }
    for fi95 := 0; fi95 < 18; fi95++ {
        if true {
            v24 += 594
            v96 := GenS10{F0: (v2 % (((v2 / ((v2) | 1))) | 1)), F1: 0, F2: 173}
            v23 -= v23
            _ = v96
        } else {
            fmt.Println(v2, fi95, fi95)
            v2 -= v2
            fmt.Println(v2, v2)
        }
        v97 := GenS4{F0: v24, F1: (v2 << ((255) & 63))}
        _ = fi95
        _ = v97
    }
    for fi98 := 0; fi98 < 4; fi98++ {
        if true {
            v2 = v2
            if (!false) {
                switch (v23 * (^(v23))) {
                case -1:
                    v24 += 65535
                default:
                    switch 381 {
                    case 128:
                        v24 += 4
                    case 452:
                        switch "\n\t" {
                        case "foo bar":
                            v99 := GenS11{F0: (v2 - 65535)}
                            _ = v99
                        case "hello":
                            if true {
                                v23 -= 127
                                v2 = v2
                                fmt.Println(fi98, v2)
                            } else {
                                var v100 int16 = 1
                                _ = v100
                            }
                        }
                    default:
                        v24 = 2147483647
                    }
                }
                v101 := []int{v23}
                _ = v101
            }
        } else {
            v102 := GenS12{F0: -1, F1: 65535, F2: uint64(v2)}
            _ = v102
        }
        _ = fi98
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v49
}
