package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint; F1 uint8 }
type GenS3 struct { F0 int64; F1 uint; F2 uint32 }
type GenS4 struct { F0 uint32; F1 int8 }
type GenS5 struct { F0 bool; F1 bool; F2 int8 }
type GenS6 struct { F0 int; F1 int32 }
type GenS7 struct { F0 uint16; F1 int }

func main() {
    v1 := make(map[int]string)
    var v2 uint32 = 1
    var v3 float64 = 1e10
    v3 = float64(0.1)
    switch 9223372036854775807 {
    case -2147483648:
        var v4 float32 = -2.25
        _ = v4
    case 9223372036854775807:
        var v5 uint = 1
        _ = v5
    }
    v2 = 2147483648
    v3 = v3
    var v6 float32 = -2.25
    switch 924 {
    case -9223372036854775808:
        if (false || true) {
            v7 := GenS1{F0: 42}
            _ = v7
        } else {
            v6 -= v6
            if (1 > 127) {
                fmt.Println(v6)
                v3 += v3
            } else {
                v6 = v6
                var v8 int8 = 115
                var v9 uint32 = 2147483647
                _ = v8
                _ = v9
            }
            v3 = (v3 / float64(v6))
        }
    case 0:
        if true {
            v6 = (float32(v3) * v6)
        } else {
            if true {
                for fi10 := 0; fi10 < 15; fi10++ {
                    v11 := []int{}
                    for fi12 := 0; fi12 < 7; fi12++ {
                        if (false && (32767 <= -32768)) {
                            v6 += (((v6 + v6) * -2.25) + (-2.25 / v6))
                            var v13 float32 = 3.14159
                            v13 -= v6
                            fmt.Println(v2, v6, fi10)
                            if true {
                                fmt.Println(fi12)
                            }
                            var v14 bool = (42 >= 159)
                            if (426 != 42) {
                                v15 := GenS2{F0: 0, F1: 127}
                                switch "hello" {
                                case "\n\t":
                                    v13 -= (v6 * (v13 * v13))
                                default:
                                    if len(v11) > 0 {
                                        v11[0] = -9223372036854775808
                                    }
                                }
                                _ = v15
                            } else {
                                if len(v11) >= 2 {
                                    v11 = v11[1:len(v11)]
                                }
                                var v16 uint = 0
                                var v17 uint = 986
                                _ = v16
                                _ = v17
                            }
                            var v18 uint = 985
                            switch ("" + "foo bar") {
                            case "foo bar":
                                if (v14 && v14) {
                                    v2 = (^(v2))
                                    var v19 int16 = -490
                                    _ = v19
                                } else {
                                    switch 32767 {
                                    case -364:
                                        var v20 bool = v14
                                        _ = v20
                                    case 42:
                                        switch "" {
                                        case "foo bar":
                                            if ((!(v14 && (0 > 2))) && (!(v13 == v6))) {
                                                _ = len(v11)
                                                v21 := v1[(^(fi12))]
                                                if ((fi12 >> (((v18 >> ((1) & 63))) & 63)) == (-(fi10))) {
                                                    var v22 uint8 = 68
                                                    _ = v22
                                                }
                                                _ = v21
                                            } else {
                                                fmt.Println(v6)
                                                if (1000 == 1000) {
                                                    fmt.Println(v14)
                                                    v23 := GenS1{F0: 18446744073709551615}
                                                    var v24 uint8 = 0
                                                    _ = v23
                                                    _ = v24
                                                }
                                            }
                                        default:
                                            switch ("\n\t" + "hello") {
                                            case "a":
                                                switch (fi10 | fi12) {
                                                case -1:
                                                    v18 -= 42
                                                default:
                                                    if (v3 <= (((v3 * 0.0) / ((v3 * 3.14159) / (1.5 + 0.0))) * v3)) {
                                                        var v25 uint64 = 668
                                                        _ = v25
                                                    }
                                                }
                                            case "":
                                                switch "a" {
                                                case "":
                                                    fmt.Println(v14)
                                                case "foo bar":
                                                    v18 -= v18
                                                case "hello":
                                                    if (v2 != v2) {
                                                        fmt.Println(v13, v2)
                                                    } else {
                                                        v26 := GenS3{F0: 1, F1: (v18 ^ 770), F2: (v2 ^ v2)}
                                                        v11 = append(v11, (fi10 & -9223372036854775808))
                                                        v3 = 0.0
                                                        _ = v26
                                                    }
                                                default:
                                                    v13 -= (v13 + 1.5)
                                                }
                                            case "\n\t":
                                                switch fi10 {
                                                case -784:
                                                    if (913 <= 385) {
                                                        fmt.Println(fi12)
                                                        v27 := GenS3{F0: -394, F1: (v18 * (v18 ^ (v18 * v18))), F2: 65535}
                                                        switch "foo bar" {
                                                        case "hello":
                                                            v28 := GenS3{F0: 1, F1: v18, F2: (v2 & (v2 ^ v2))}
                                                            _ = v28
                                                        default:
                                                            if len(v11) > 0 {
                                                                v11[0] = fi10
                                                            }
                                                        }
                                                        _ = v27
                                                    } else {
                                                        v18 -= 0
                                                        switch -824 {
                                                        case -2147483648:
                                                            fmt.Println(v18, fi12)
                                                            v2 -= (v2 >> ((609) & 63))
                                                        case -9223372036854775808:
                                                            if (803 == 1) {
                                                                v18 -= (v18 << ((255) & 63))
                                                                switch ("foo bar" + "") {
                                                                case "a":
                                                                    var v29 bool = false
                                                                    _ = v29
                                                                case "":
                                                                    if (true && ((!(0 == 127)) && v14)) {
                                                                        if (v14 && v14) {
                                                                            v11 = append(v11, fi12)
                                                                        }
                                                                        fmt.Println(v3)
                                                                    } else {
                                                                        if (!(v13 < v13)) {
                                                                            v30 := GenS4{F0: (v2 << (((v18 + v18)) & 63)), F1: 0}
                                                                            v13 = 0.0
                                                                            v6 -= float32(v3)
                                                                            _ = v30
                                                                        } else {
                                                                            fmt.Println(v2, v6, v18)
                                                                            v31 := GenS1{F0: 715}
                                                                            if (2 >= 18446744073709551615) {
                                                                                fmt.Println(v14)
                                                                                var v32 string = "foo bar"
                                                                                if ((-2.25 != ((0.0 - (-2.25 * v13)) * 0.1)) && ((440 != -368) && v14)) {
                                                                                    var v33 int16 = 1
                                                                                    if v14 {
                                                                                        fmt.Println(v2, fi10, v3)
                                                                                        switch (fi12 & (fi12 + 9223372036854775807)) {
                                                                                        case -1:
                                                                                            var v34 uint64 = 1
                                                                                            _ = v34
                                                                                        default:
                                                                                            fmt.Println(v2, v6, v33)
                                                                                        }
                                                                                        v35 := GenS3{F0: -1, F1: 2, F2: 2}
                                                                                        _ = v35
                                                                                    }
                                                                                    _ = v33
                                                                                }
                                                                                _ = v32
                                                                            } else {
                                                                                if ((v18 % ((65535) | 1)) <= (v18 * (^(v18)))) {
                                                                                    var v36 int64 = -193
                                                                                    _ = v36
                                                                                } else {
                                                                                    v37 := GenS2{F0: (v18 - (v18 ^ (v18 >> ((646) & 63)))), F1: 162}
                                                                                    v11 = append(v11, -9223372036854775808)
                                                                                    if true {
                                                                                        v13 += ((v6 + 0.1) - v13)
                                                                                        if len(v11) >= 2 {
                                                                                            v11 = v11[1:len(v11)]
                                                                                        }
                                                                                        if len(v11) >= 2 {
                                                                                            v11 = v11[1:len(v11)]
                                                                                        }
                                                                                        var v38 float64 = v3
                                                                                        var v39 uint64 = 252
                                                                                        v3 = v38
                                                                                        _ = v38
                                                                                        _ = v39
                                                                                    } else {
                                                                                        v40 := GenS3{F0: 0, F1: 65535, F2: v2}
                                                                                        v41 := GenS4{F0: v2, F1: 1}
                                                                                        var v42 uint32 = (^(v2))
                                                                                        _ = v40
                                                                                        _ = v41
                                                                                        _ = v42
                                                                                    }
                                                                                    fmt.Println(v2, v14)
                                                                                    _ = v37
                                                                                }
                                                                                v13 = float32(v3)
                                                                            }
                                                                            _ = v31
                                                                        }
                                                                    }
                                                                case "foo bar":
                                                                    _ = len(v11)
                                                                }
                                                            } else {
                                                                if (-76 <= 127) {
                                                                    fmt.Println(v6)
                                                                    var v43 float32 = ((v6 * ((v6 - 1.5) - 1e10)) / (v13 - v13))
                                                                    switch (fi12 - fi10) {
                                                                    case 457:
                                                                        var v44 string = ("\n\t" + ("\n\t" + "hello"))
                                                                        _ = v44
                                                                    case 32768:
                                                                        var v45 int8 = -128
                                                                        _ = v45
                                                                    case -821:
                                                                        v14 = ((v2 << ((v18) & 63)) < (v2 << ((v18) & 63)))
                                                                    default:
                                                                        var v46 int = (fi12 >> ((78) & 63))
                                                                        _ = v46
                                                                    }
                                                                    _ = v43
                                                                }
                                                                var v47 uint16 = 65535
                                                                _ = v47
                                                            }
                                                        case 1:
                                                            switch fi12 {
                                                            case 32767:
                                                                v1[0] = "a"
                                                            }
                                                        default:
                                                            switch ("a" + "") {
                                                            case "a":
                                                                switch fi12 {
                                                                case -9223372036854775808:
                                                                    v6 = float32(v3)
                                                                case 127:
                                                                    fmt.Println(v18, v3)
                                                                    v48 := GenS3{F0: 0, F1: 255, F2: (v2 * v2)}
                                                                    if (498 <= 18446744073709551615) {
                                                                        v49 := GenS1{F0: 268}
                                                                        _ = v49
                                                                    } else {
                                                                        if v14 {
                                                                            v13 = v6
                                                                            fmt.Println(fi10)
                                                                        } else {
                                                                            v50 := GenS4{F0: (^(v2)), F1: -59}
                                                                            v3 -= v3
                                                                            fmt.Println(v18, v14, fi10)
                                                                            _ = v50
                                                                        }
                                                                        if v14 {
                                                                            v6 = ((v13 / v13) * 1e10)
                                                                            var v51 float64 = 0.0
                                                                            v52 := GenS5{F0: (v18 != (v18 + v18)), F1: true, F2: 21}
                                                                            _ = v51
                                                                            _ = v52
                                                                        }
                                                                    }
                                                                    switch "foo bar" {
                                                                    case "\n\t":
                                                                        v53 := GenS4{F0: v2, F1: -101}
                                                                        _ = v53
                                                                    }
                                                                    v18 += v18
                                                                    _ = v48
                                                                case 9223372036854775807:
                                                                    switch (fi10 & fi10) {
                                                                    case 4294967295:
                                                                        fmt.Println(v14)
                                                                    case 32768:
                                                                        switch ("hello" + "foo bar") {
                                                                        case "hello":
                                                                            v13 = -2.25
                                                                        default:
                                                                            v6 += ((v13 / ((v6 * ((1.5 * 3.14159) / ((1.5 - ((-2.25 - v6) / v6)) + v6))) + v13)) / 1.5)
                                                                        }
                                                                    case -9223372036854775808:
                                                                        fmt.Println(v14, v2, fi10)
                                                                    }
                                                                }
                                                            case "\n\t":
                                                                var v54 float64 = 1e10
                                                                _ = v54
                                                            default:
                                                                v3 = v3
                                                            }
                                                        }
                                                        if len(v11) >= 2 {
                                                            v11 = v11[1:len(v11)]
                                                        }
                                                        var v55 uint64 = 0
                                                        _ = v55
                                                    }
                                                default:
                                                    var v56 uint32 = (v2 + v2)
                                                    v13 = ((1.5 / 1e10) + (v6 / v13))
                                                    if false {
                                                        _ = len(v11)
                                                        v57 := GenS6{F0: fi12, F1: -1}
                                                        v13 -= float32((((v3 * 1.5) + v3) / 3.14159))
                                                        _ = v57
                                                    }
                                                    switch fi10 {
                                                    case 725:
                                                        var v58 int16 = 32767
                                                        _ = v58
                                                    case 42:
                                                        var v59 int64 = 9223372036854775807
                                                        _ = v59
                                                    }
                                                    var v60 int8 = int8(v18)
                                                    _ = v56
                                                    _ = v60
                                                }
                                            }
                                        }
                                    default:
                                        if (!v14) {
                                            var v61 uint32 = 873
                                            var v62 uint64 = 758
                                            _ = v61
                                            _ = v62
                                        }
                                    }
                                    v3 -= float64(-2.25)
                                }
                            case "a":
                                if len(v11) >= 2 {
                                    v11 = v11[1:len(v11)]
                                }
                            default:
                                switch fi12 {
                                case -575:
                                    var v63 uint16 = 0
                                    _ = v63
                                case -1:
                                    var v64 uint = (v18 % ((v18) | 1))
                                    _ = v64
                                case -800:
                                    v65 := GenS7{F0: 692, F1: fi10}
                                    _ = v65
                                default:
                                    var v66 int16 = 0
                                    _ = v66
                                }
                            }
                            fmt.Println(v18, v2)
                            var v67 int8 = -128
                            v68 := GenS2{F0: (^(v18)), F1: 193}
                            v18 -= (v18 << (((v18 >> ((v18) & 63))) & 63))
                            _ = v13
                            _ = v14
                            _ = v18
                            _ = v67
                            _ = v68
                        }
                        _ = fi12
                    }
                    _ = fi10
                    _ = v11
                }
                if false {
                    var v69 float32 = v6
                    delete(v1, -959)
                    _ = v69
                }
            }
            v70 := []int{42, -930}
            var v71 int8 = -99
            _ = v70
            _ = v71
        }
    default:
        var v72 float64 = 3.14159
        _ = v72
    }
    v6 = ((v6 + v6) - (float32(float64((v6 - (float32(-2.25) + float32(v3))))) * v6))
    for fi73 := 0; fi73 < 19; fi73++ {
        if ((42 <= 1) && false) {
            for fi74 := 0; fi74 < 4; fi74++ {
                fmt.Println(v2, v6)
                _ = fi74
            }
            v75 := GenS5{F0: true, F1: ((!(int8(fi73) > 127)) || false), F2: -94}
            var v76 int32 = 42
            _ = v75
            _ = v76
        }
        v6 = v6
        if true {
            v3 = 0.1
        }
        _ = fi73
    }
    var v77 int64 = 518
    switch 9223372036854775807 {
    case -2147483648:
        v78 := []int{766, -2147483648, 609, 537}
        _ = v78
    }
    switch "foo bar" {
    case "foo bar":
        var v79 int = 2147483647
        _ = v79
    case "a":
        if true {
            if true {
                v80 := []int{-444}
                for fi81 := 0; fi81 < 8; fi81++ {
                    switch "hello" {
                    case "foo bar":
                        v82 := v1[4294967295]
                        _ = v82
                    case "\n\t":
                        v6 = float32((3.14159 + (((-2.25 / v3) + float64((v6 - (v6 - float32(((3.14159 + (0.0 / 1.5)) / ((v3 / ((v3 + v3) / v3)) / 1.0))))))) + 3.14159)))
                    case "hello":
                        var v83 float32 = (1.5 / 3.14159)
                        _ = v83
                    }
                    if (4294967295 <= 4294967295) {
                        v84 := GenS4{F0: 65535, F1: 127}
                        _ = v84
                    } else {
                        v77 = v77
                    }
                    _ = len(v1)
                    _ = fi81
                }
                _ = v80
            }
        } else {
            v3 += float64(float32((v3 / v3)))
            fmt.Println(v77, v3, v2)
        }
    }
    for fi85 := 0; fi85 < 19; fi85++ {
        for fi86 := 0; fi86 < 3; fi86++ {
            delete(v1, (fi85 % ((fi86) | 1)))
            _ = fi86
        }
        for fi87 := 0; fi87 < 16; fi87++ {
            var v88 uint64 = 0
            v88 = 18446744073709551615
            switch fi85 {
            case 742:
                var v89 int16 = 1
                _ = v89
            default:
                v88 = (v88 & (v88 + (v88 << ((4294967295) & 63))))
            }
            _ = fi87
            _ = v88
        }
        var v90 uint = 65535
        _ = fi85
        _ = v90
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    fmt.Println("v77:", v77)
    _ = v1
    _ = v2
    _ = v3
    _ = v6
    _ = v77
}
