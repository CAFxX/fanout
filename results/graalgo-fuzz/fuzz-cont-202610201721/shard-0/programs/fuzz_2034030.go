package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 bool; F1 uint8 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int; F1 int; F2 int32 }
type GenS5 struct { F0 int64; F1 uint32 }
type GenS6 struct { F0 bool; F1 int; F2 int64 }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 int }
type GenS9 struct { F0 int16; F1 uint }
type GenS10 struct { F0 uint64 }
type GenS11 struct { F0 uint16 }
type GenS12 struct { F0 uint64 }
type GenS13 struct { F0 string; F1 uint16 }
type GenS14 struct { F0 int16 }
type GenS15 struct { F0 uint32; F1 int8; F2 uint8 }
func genFn0(p0 int16, p1 uint, p2 uint64) uint16 {
    fmt.Println(p1, p1)
    p2 = (^(p2))
    _ = p0
    _ = p1
    _ = p2
    return 255
}

func main() {
    v1 := []uint32{460, 485, 65535, 4294967295}
    var v2 int = 127
    v3 := []uint32{4294967295, 975, 4294967295}
    v4 := make(map[uint32]uint32)
    var v5 float32 = (3.14159 * 0.1)
    switch "a" {
    case "\n\t":
        var v6 int16 = 1000
        _ = v6
    }
    v7 := GenS1{F0: -482}
    v8 := GenS1{F0: (v2 | v2)}
    v9 := GenS2{F0: true, F1: 84}
    var v10 string = "hello"
    v11 := GenS1{F0: -2147483648}
    v2 = (v2 + (v2 + (v2 * 4294967295)))
    if ((1e10 + 1e10) == 3.14159) {
        v5 = (v5 + v5)
        v12 := GenS2{F0: true, F1: 1}
        switch v2 {
        case 4294967295:
            v4[2] = 2147483648
        case -643:
            fmt.Println(v10)
        default:
            var v13 uint = 196
            _ = v13
        }
        _ = v12
    } else {
        for fi14 := 0; fi14 < 19; fi14++ {
            v10 = v10
            if false {
                if (-26 > 63) {
                    var v15 int16 = -32768
                    if (false && ((((-2.25 * -2.25) + 1e10) / 0.1) <= 1e10)) {
                        if false {
                            v16 := GenS2{F0: false, F1: 0}
                            _ = v16
                        }
                        v17 := GenS1{F0: 165}
                        _ = v17
                    } else {
                        fmt.Println(v5, v5, v10)
                    }
                    v18 := GenS1{F0: (v2 - v2)}
                    _ = v15
                    _ = v18
                }
                var v19 int64 = 0
                switch ("foo bar" + (v10 + ("hello" + (v10 + v10)))) {
                case "\n\t":
                    var v20 int32 = 21
                    _ = v20
                case "a":
                    var v21 bool = (!(!(false && (((62 == 321) && false) && true))))
                    _ = v21
                }
                _ = v19
            }
            _ = fi14
        }
        fmt.Println(v2, v2, v5)
    }
    if false {
        if (203 == 688) {
            for fi22 := 0; fi22 < 18; fi22++ {
                switch (-(v2)) {
                case -493:
                    v10 = "foo bar"
                case -626:
                    if (true || (true && (!false))) {
                        var v23 int16 = 32767
                        _ = v23
                    }
                case -924:
                    var v24 uint32 = 0
                    _ = v24
                }
                _ = fi22
            }
            v25 := GenS1{F0: -845}
            _ = v25
        } else {
            for fi26 := 0; fi26 < 1; fi26++ {
                _ = len(v1)
                for fi27 := 0; fi27 < 16; fi27++ {
                    v5 = v5
                    _ = fi27
                }
                switch v10 {
                case "":
                    v10 = v10
                }
                _ = fi26
            }
            v5 = 1e10
        }
    }
    if (!true) {
        if false {
            for fi28 := 0; fi28 < 20; fi28++ {
                var v29 int16 = -32768
                delete(v4, 65535)
                _ = fi28
                _ = v29
            }
        } else {
            fmt.Println(v10)
            var v30 string = ("\n\t" + "\n\t")
            _ = v30
        }
        var v31 float64 = -2.25
        _ = v31
    } else {
        fmt.Println(v10, v5)
        switch "" {
        case "":
            for fi32 := 0; fi32 < 3; fi32++ {
                v5 += 3.14159
                v5 = v5
                if (!true) {
                    switch v10 {
                    case "a":
                        fmt.Println(v10)
                    case "\n\t":
                        v2 = (v2 ^ (v2 + (fi32 * (v2 + (v2 | fi32)))))
                    default:
                        for fi33 := 0; fi33 < 18; fi33++ {
                            fmt.Println(v10)
                            v34 := v4[716]
                            var v35 bool = false
                            switch 2147483647 {
                            case 4294967295:
                                switch (v2 ^ (v2 | -9223372036854775808)) {
                                case 32767:
                                    if true {
                                        fmt.Println(v5)
                                        var v36 bool = (v35 || v35)
                                        v5 = v5
                                        _ = v36
                                    } else {
                                        var v37 int16 = 1000
                                        _ = v37
                                    }
                                case 4294967295:
                                    _ = len(v3)
                                }
                            default:
                                if (898 >= 878) {
                                    fmt.Println(fi32, v34, v2)
                                    if (v34 >= v34) {
                                        var v38 int32 = 1
                                        v39 := GenS2{F0: (!v35), F1: 210}
                                        switch 42 {
                                        case 1:
                                            v34 = (v34 ^ v34)
                                        case 9223372036854775807:
                                            v34 = v34
                                        case -452:
                                            if v35 {
                                                fmt.Println(v2)
                                                v40 := GenS3{F0: true}
                                                fmt.Println(v34, v2)
                                                _ = v40
                                            } else {
                                                var v41 uint8 = 2
                                                _ = v41
                                            }
                                        }
                                        _ = v38
                                        _ = v39
                                    }
                                } else {
                                    v2 = (fi32 - v2)
                                    var v42 float32 = v5
                                    _ = v42
                                }
                            }
                            delete(v4, (v34 + (v34 ^ v34)))
                            switch (v10 + "a") {
                            case "":
                                if (-9223372036854775808 > 0) {
                                    var v43 int64 = -1
                                    _ = v43
                                }
                            case "a":
                                v34 -= 493
                            case "\n\t":
                                v35 = true
                                var v44 int8 = 76
                                _ = v44
                            default:
                                v45 := GenS2{F0: v35, F1: 2}
                                _ = v45
                            }
                            _ = fi33
                            _ = v34
                            _ = v35
                        }
                    }
                    v46 := GenS4{F0: v2, F1: -889, F2: 2147483647}
                    var v47 bool = (!false)
                    _ = v46
                    _ = v47
                }
                _ = fi32
            }
        case "foo bar":
            for fi48 := 0; fi48 < 8; fi48++ {
                v49 := GenS4{F0: fi48, F1: (v2 ^ v2), F2: -53}
                var v50 uint8 = 42
                for fi51 := 0; fi51 < 15; fi51++ {
                    if (!(!(16 > 2147483647))) {
                        v52 := GenS5{F0: 9223372036854775807, F1: 2}
                        v1 = append(v1, 4294967295)
                        v4[4294967295] = 939
                        _ = v52
                    } else {
                        if (false && true) {
                            v53 := GenS4{F0: 255, F1: 42, F2: -1}
                            switch "" {
                            case "hello":
                                v54 := GenS4{F0: fi48, F1: 0, F2: -1}
                                _ = v54
                            }
                            if (156 <= (v50 << ((1) & 63))) {
                                fmt.Println(fi48, v10, v2)
                            }
                            _ = v53
                        } else {
                            if false {
                                var v55 float32 = ((-2.25 / (0.1 - 1e10)) / v5)
                                switch fi48 {
                                case 42:
                                    fmt.Println(fi51, fi48)
                                }
                                var v56 uint = 0
                                _ = v55
                                _ = v56
                            } else {
                                v57 := GenS6{F0: (!(!false)), F1: (fi51 + 0), F2: 231}
                                if len(v1) > 0 {
                                    v1[0] = 2147483647
                                }
                                if false {
                                    if false {
                                        if true {
                                            switch ((v10 + (v10 + ((v10 + ("hello" + (v10 + (v10 + v10)))) + ("a" + "a")))) + "a") {
                                            case "":
                                                if (1000 >= 1) {
                                                    var v58 bool = (genFn0(-1, 18446744073709551615, 42) <= 718)
                                                    if v58 {
                                                        switch (fi48 | (-(fi51))) {
                                                        case 2147483647:
                                                            if len(v3) > 0 {
                                                                v3[0] = 2
                                                            }
                                                        case -942:
                                                            fmt.Println(v5)
                                                        default:
                                                            switch fi51 {
                                                            case 32767:
                                                                v59 := GenS7{F0: 4294967295}
                                                                _ = v59
                                                            case 32768:
                                                                var v60 uint8 = v50
                                                                _ = v60
                                                            case 42:
                                                                v5 = v5
                                                            default:
                                                                var v61 uint8 = v50
                                                                _ = v61
                                                            }
                                                        }
                                                        v62 := GenS8{F0: 1}
                                                        var v63 uint = 539
                                                        _ = v62
                                                        _ = v63
                                                    }
                                                    switch ((("hello" + v10) + v10) + v10) {
                                                    case "":
                                                        if (2147483647 < uint32(v50)) {
                                                            if len(v3) > 0 {
                                                                v3[0] = 331
                                                            }
                                                            var v64 uint8 = (v50 >> ((1) & 63))
                                                            _ = v64
                                                        } else {
                                                            v65 := GenS4{F0: (v2 + (fi48 | 2147483647)), F1: 9223372036854775807, F2: -2147483648}
                                                            var v66 uint16 = 2
                                                            v58 = (float32(1e10) > float32(-2.25))
                                                            _ = v65
                                                            _ = v66
                                                        }
                                                    }
                                                    _ = v58
                                                }
                                            case "hello":
                                                var v67 int32 = 0
                                                _ = v67
                                            default:
                                                if (127 < int8(v2)) {
                                                    fmt.Println(v50, fi51)
                                                }
                                            }
                                            if true {
                                                switch "foo bar" {
                                                case "foo bar":
                                                    var v68 float32 = float32(1.5)
                                                    _ = v68
                                                case "a":
                                                    v2 -= v2
                                                case "hello":
                                                    var v69 int64 = 42
                                                    _ = v69
                                                }
                                                if true {
                                                    v50 -= 244
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                    if (42 != 32767) {
                                                        var v70 string = v10
                                                        fmt.Println(v50, fi51, v70)
                                                        _ = v70
                                                    } else {
                                                        var v71 float64 = 1.5
                                                        if true {
                                                            _ = len(v1)
                                                            if ((1 > -28) || false) {
                                                                var v72 float32 = (1e10 * 0.1)
                                                                switch "foo bar" {
                                                                case "a":
                                                                    var v73 int32 = 0
                                                                    _ = v73
                                                                case "\n\t":
                                                                    if (18446744073709551615 > 533) {
                                                                        v74 := GenS9{F0: 1000, F1: 2}
                                                                        _ = v74
                                                                    } else {
                                                                        if len(v3) > 0 {
                                                                            v3[0] = 4294967295
                                                                        }
                                                                    }
                                                                case "foo bar":
                                                                    v75 := GenS10{F0: 0}
                                                                    _ = v75
                                                                default:
                                                                    v71 += v71
                                                                }
                                                                _ = v72
                                                            }
                                                        }
                                                        v76 := GenS11{F0: 42}
                                                        _ = len(v4)
                                                        v71 = v71
                                                        _ = v71
                                                        _ = v76
                                                    }
                                                } else {
                                                    fmt.Println(v50)
                                                }
                                                fmt.Println(v2, v10)
                                            }
                                            if len(v1) > 0 {
                                                v1[0] = 42
                                            }
                                            v4[2147483647] = 1
                                        } else {
                                            fmt.Println(v5, fi48, v10)
                                            var v77 int16 = 0
                                            v78 := GenS4{F0: fi51, F1: v2, F2: -633}
                                            _ = v77
                                            _ = v78
                                        }
                                    } else {
                                        var v79 int16 = -1
                                        var v80 int8 = 3
                                        v81 := GenS12{F0: 413}
                                        _ = v79
                                        _ = v80
                                        _ = v81
                                    }
                                    if false {
                                        fmt.Println(fi48, fi48, fi51)
                                        v2 = (fi48 / ((fi48) | 1))
                                    } else {
                                        var v82 int = (^(fi48))
                                        fmt.Println(v50, v10)
                                        if true {
                                            var v83 bool = ((v50 | v50) > 2)
                                            v4[0] = 42
                                            if true {
                                                fmt.Println(v82)
                                                v50 += v50
                                            } else {
                                                v4[4294967295] = 4294967295
                                                switch "\n\t" {
                                                case "hello":
                                                    var v84 uint = 255
                                                    if (!(9223372036854775807 >= 1)) {
                                                        v4[592] = 2
                                                        v85 := GenS13{F0: v10, F1: 255}
                                                        _ = v85
                                                    }
                                                    v10 = ((v10 + "hello") + ((v10 + "") + v10))
                                                    _ = v84
                                                }
                                                switch (-(fi48)) {
                                                case -1:
                                                    v50 += (^(v50))
                                                }
                                                v86 := GenS14{F0: 725}
                                                _ = v86
                                            }
                                            _ = v83
                                        }
                                        _ = v82
                                    }
                                    var v87 int8 = 0
                                    _ = v87
                                }
                                fmt.Println(fi51)
                                v88 := GenS12{F0: 1}
                                _ = v57
                                _ = v88
                            }
                            var v89 int64 = -9223372036854775808
                            v90 := GenS7{F0: 42}
                            _ = v89
                            _ = v90
                        }
                    }
                    var v91 bool = (4294967295 <= 0)
                    _ = fi51
                    _ = v91
                }
                _ = fi48
                _ = v49
                _ = v50
            }
        }
        v2 -= 2147483647
    }
    v10 = "hello"
    var v92 int32 = 1
    for fi93 := 0; fi93 < 11; fi93++ {
        v92 += v92
        v94 := GenS14{F0: -32768}
        v95 := GenS15{F0: 434, F1: 0, F2: 128}
        _ = fi93
        _ = v94
        _ = v95
    }
    fmt.Println(v2, v92, v92)
    switch "\n\t" {
    case "a":
        if false {
            var v96 float32 = v5
            v2 = 1
            _ = v96
        } else {
            v92 = v92
        }
    case "\n\t":
        var v97 float32 = v5
        _ = v97
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v92:", v92)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
    _ = v92
}
