package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32; F2 uint64 }
type GenS2 struct { F0 int; F1 int64; F2 uint16 }
type GenS3 struct { F0 int32; F1 uint; F2 string }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint8; F1 uint }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 uint64 }
type GenS9 struct { F0 uint8; F1 int64 }
type GenS10 struct { F0 int8; F1 uint; F2 uint }
type GenS11 struct { F0 int32; F1 int }
type GenS12 struct { F0 uint; F1 uint8; F2 uint16 }
func genFn0(p0 uint16, p1 int64) uint16 {
    fmt.Println(p0, p0)
    for fi5 := 0; fi5 < 17; fi5++ {
        var v6 int64 = p1
        if true {
            if false {
                switch "" {
                case "":
                    var v7 int8 = 127
                    _ = v7
                default:
                    var v8 uint64 = 42
                    _ = v8
                }
            }
        }
        var v9 int64 = p1
        _ = fi5
        _ = v6
        _ = v9
    }
    if false {
        v10 := GenS1{F0: -147, F1: 636, F2: 385}
        if (true || false) {
            switch -9223372036854775808 {
            case 1:
                var v11 uint16 = 2
                _ = v11
            case 32768:
                switch 2147483647 {
                case 4294967295:
                    for fi12 := 0; fi12 < 20; fi12++ {
                        for fi13 := 0; fi13 < 17; fi13++ {
                            var v14 uint = uint(p0)
                            _ = fi13
                            _ = v14
                        }
                        _ = fi12
                    }
                case 32768:
                    var v15 string = "a"
                    _ = v15
                case -727:
                    var v16 float64 = ((0.0 - (float64(-2.25) / 1.0)) / 1.5)
                    _ = v16
                }
            default:
                if false {
                    var v17 uint16 = (^(p0))
                    switch "hello" {
                    case "\n\t":
                        v18 := GenS2{F0: 497, F1: (-(p1)), F2: (p0 | (v17 / ((p0) | 1)))}
                        _ = v18
                    }
                    _ = v17
                } else {
                    var v19 bool = false
                    switch 42 {
                    case 86:
                        var v20 uint = 953
                        _ = v20
                    }
                    _ = v19
                }
            }
        } else {
            p1 -= p1
            if (true && (42 == -858)) {
                v21 := []int{-917, 9223372036854775807}
                _ = v21
            }
        }
        v22 := GenS3{F0: 525, F1: 468, F2: "\n\t"}
        _ = v10
        _ = v22
    }
    var v23 int64 = (p1 - -9223372036854775808)
    _ = p0
    _ = p1
    _ = v23
    return 255
}
func genFn1(p0 uint16, p1 bool, p2 string) int8 {
    p1 = p1
    p2 = p2
    var v24 int64 = -1
    v25 := GenS4{F0: 127}
    var v26 string = p2
    _ = p0
    _ = p1
    _ = p2
    _ = v24
    _ = v25
    _ = v26
    return 100
}

func main() {
    var v1 uint = 65535
    var v2 int8 = -128
    v3 := make(map[string]string)
    var v4 string = "hello"
    v2 += v2
    var v27 int8 = (v2 | v2)
    if ((-797 < -9223372036854775808) || (!false)) {
        var v28 int64 = 9223372036854775807
        switch ((v4 + v4) + "foo bar") {
        case "hello":
            for fi29 := 0; fi29 < 4; fi29++ {
                for fi30 := 0; fi30 < 3; fi30++ {
                    v27 = v27
                    _ = fi30
                }
                if true {
                    var v31 int16 = 165
                    v1 = (^(v1))
                    switch fi29 {
                    case 739:
                        switch fi29 {
                        case -2147483648:
                            var v32 float64 = 0.0
                            _ = v32
                        case 684:
                            v33 := GenS4{F0: 88}
                            _ = v33
                        default:
                            for fi34 := 0; fi34 < 14; fi34++ {
                                var v35 uint = v1
                                _ = fi34
                                _ = v35
                            }
                        }
                    case 550:
                        var v36 int = fi29
                        _ = v36
                    case 469:
                        var v37 int16 = v31
                        _ = v37
                    }
                    _ = v31
                } else {
                    var v38 uint8 = 2
                    v39 := GenS5{F0: (v38 >> ((0) & 63)), F1: 1}
                    _ = v38
                    _ = v39
                }
                _ = fi29
            }
        }
        fmt.Println(v27)
        _ = v28
    }
    switch v4 {
    case "a":
        var v40 int32 = -1
        _ = v40
    case "\n\t":
        var v41 float64 = -2.25
        _ = v41
    default:
        var v42 int8 = v27
        _ = v42
    }
    v27 -= genFn1(genFn0(genFn0(924, -101), -641), (-2.25 >= (0.0 - 1e10)), (v4 + v4))
    for fi43 := 0; fi43 < 8; fi43++ {
        switch (v4 + v4) {
        case "":
            v2 = int8(fi43)
        default:
            v44 := []int{fi43}
            _ = v44
        }
        var v45 float64 = (-2.25 + ((float64(3.14159) + 0.1) + 0.0))
        v46 := GenS2{F0: (^(fi43)), F1: 9223372036854775807, F2: uint16(v1)}
        _ = fi43
        _ = v45
        _ = v46
    }
    v47 := GenS4{F0: 255}
    for fi48 := 0; fi48 < 19; fi48++ {
        switch (v4 + (("" + (v4 + "foo bar")) + "foo bar")) {
        case "hello":
            _ = len(v3)
        case "":
            if (!(true && false)) {
                v27 = genFn1(602, false, ((v4 + (((v4 + ((v4 + v4) + v4)) + v4) + v4)) + "a"))
                if false {
                    switch (("" + "hello") + "a") {
                    case "\n\t":
                        v2 -= genFn1(638, true, "hello")
                    case "":
                        var v49 float64 = (float64(float32(float64((((1e10 + 3.14159) + (1e10 - 3.14159)) + (float32(3.14159) + ((float32(((3.14159 * (1e10 / -2.25)) - -2.25)) / 0.1) * ((float32(3.14159) - 1.5) * 0.0))))))) * (1e10 * 0.0))
                        _ = v49
                    case "a":
                        var v50 int64 = 0
                        _ = v50
                    }
                } else {
                    v51 := GenS6{F0: (0 != 30)}
                    for fi52 := 0; fi52 < 12; fi52++ {
                        if (((genFn1(576, true, v4) >= genFn1(32767, ((42 >= -1) || (!(!false))), v4)) && (2 < 18446744073709551615)) && true) {
                            if (!(((false || (228 <= 255)) || false) || (false || (!false)))) {
                                if ((!true) || true) {
                                    var v53 bool = false
                                    _ = v53
                                }
                                v54 := GenS7{F0: 0}
                                _ = v54
                            } else {
                                v55 := GenS1{F0: 1, F1: 65535, F2: 292}
                                v27 -= v2
                                fmt.Println(v27, v27)
                                _ = v55
                            }
                            var v56 float32 = (0.0 * 0.1)
                            _ = v56
                        }
                        if true {
                            _ = len(v3)
                            _ = len(v3)
                            if true {
                                v57 := []int{0, (fi52 + (fi52 / ((fi52) | 1)))}
                                if ((false || true) && true) {
                                    v2 -= v27
                                    v1 = (v1 << (((^(v1))) & 63))
                                    var v58 uint32 = 315
                                    _ = v58
                                }
                                switch fi52 {
                                case 127:
                                    v59 := GenS8{F0: 18446744073709551615}
                                    _ = v59
                                case 1:
                                    v60 := GenS5{F0: 127, F1: v1}
                                    _ = v60
                                case 531:
                                    var v61 uint32 = 2147483647
                                    _ = v61
                                }
                                _ = v57
                            } else {
                                fmt.Println(fi48)
                                if (1 != 1) {
                                    v4 = v4
                                    v4 = v4
                                }
                            }
                            if false {
                                if true {
                                    var v62 uint16 = 255
                                    _ = v62
                                }
                                v63 := GenS9{F0: 0, F1: int64(v1)}
                                _ = v63
                            }
                        } else {
                            _ = len(v3)
                            var v64 int64 = -9223372036854775808
                            v27 = (v2 % ((v27) | 1))
                            v64 = (v64 ^ (v64 % (((v64 >> ((v1) & 63))) | 1)))
                            _ = v64
                        }
                        var v65 float32 = (0.1 * 1.5)
                        v65 -= (0.0 * -2.25)
                        _ = fi52
                        _ = v65
                    }
                    _ = v51
                }
                var v66 int = (fi48 & (fi48 | 707))
                _ = v66
            } else {
                fmt.Println(v1, v27)
                switch "hello" {
                case "a":
                    var v67 uint8 = 0
                    _ = v67
                default:
                    var v68 uint8 = uint8(v1)
                    _ = v68
                }
                switch (v4 + (v4 + v4)) {
                case "a":
                    var v69 int = 128
                    _ = v69
                default:
                    v70 := GenS1{F0: 1000, F1: 4294967295, F2: 4294967295}
                    _ = v70
                }
            }
        case "a":
            _ = len(v3)
        default:
            var v71 uint64 = 42
            _ = v71
        }
        if true {
            if true {
                var v72 float64 = 3.14159
                v4 = v4
                for fi73 := 0; fi73 < 2; fi73++ {
                    var v74 float32 = 0.0
                    _ = fi73
                    _ = v74
                }
                _ = v72
            }
            if ((0 >= 2147483647) || (!(genFn0(2, 5) >= genFn0(791, 9223372036854775807)))) {
                if true {
                    for fi75 := 0; fi75 < 11; fi75++ {
                        switch 2147483647 {
                        case 42:
                            fmt.Println(fi75, v4, v4)
                        case 2147483647:
                            if false {
                                v76 := GenS10{F0: (v27 ^ -128), F1: (v1 % ((117) | 1)), F2: v1}
                                fmt.Println(fi48, fi75)
                                v4 = v4
                                fmt.Println(v27)
                                _ = v76
                            }
                            v2 = genFn1(genFn0(32767, 666), true, v4)
                        case 151:
                            var v77 uint = (v1 & v1)
                            _ = v77
                        default:
                            v27 += genFn1(255, (665 < 805), (v4 + ("" + "a")))
                        }
                        v27 = (-(v2))
                        var v78 uint16 = 385
                        _ = fi75
                        _ = v78
                    }
                    v27 = (v27 | genFn1(0, false, "\n\t"))
                } else {
                    v1 -= 65535
                    v1 += v1
                }
            }
        }
        _ = fi48
    }
    for fi79 := 0; fi79 < 12; fi79++ {
        v80 := []int{(fi79 - (fi79 ^ fi79)), (^(fi79))}
        _ = fi79
        _ = v80
    }
    switch v4 {
    case "\n\t":
        v2 = -1
    default:
        var v81 int16 = -1
        _ = v81
    }
    for fi82 := 0; fi82 < 12; fi82++ {
        for fi83 := 0; fi83 < 5; fi83++ {
            if false {
                v84 := GenS6{F0: true}
                _ = v84
            } else {
                v4 = ("foo bar" + v4)
                var v85 uint32 = 262
                _ = v85
            }
            _ = fi83
        }
        for fi86 := 0; fi86 < 10; fi86++ {
            v2 = v27
            switch 9223372036854775807 {
            case -49:
                if (true || (755 > 0)) {
                    var v87 uint32 = 2
                    _ = len(v3)
                    var v88 uint = (v1 + v1)
                    _ = v87
                    _ = v88
                } else {
                    v89 := GenS11{F0: 42, F1: -719}
                    var v90 int32 = -2147483648
                    _ = v89
                    _ = v90
                }
                if (fi82 != (^(fi86))) {
                    fmt.Println(fi82, v1)
                    v4 = (v4 + "foo bar")
                } else {
                    fmt.Println(v27, v27, fi82)
                    switch (v4 + "a") {
                    case "\n\t":
                        var v91 int16 = 32767
                        _ = v91
                    case "a":
                        switch ((("hello" + v4) + v4) + (v4 + (v4 + "a"))) {
                        case "":
                            var v92 int64 = 539
                            _ = v92
                        }
                    case "foo bar":
                        v93 := []int{(fi82 % ((-207) | 1)), (fi86 >> ((v1) & 63)), -1}
                        _ = v93
                    default:
                        v94 := GenS4{F0: 47}
                        _ = v94
                    }
                }
                var v95 uint16 = 255
                _ = v95
            }
            _ = fi86
        }
        if ((2 == 0) || (((((-2.25 - ((0.0 * 1.5) / 1.5)) - -2.25) * 3.14159) != (3.14159 - 0.0)) || false)) {
            v96 := GenS8{F0: 814}
            v3[((v4 + (v4 + v4)) + ((v4 + v4) + v4))] = (v4 + "\n\t")
            switch "\n\t" {
            case "hello":
                fmt.Println(v27, v2)
            }
            _ = v96
        }
        _ = fi82
    }
    switch 128 {
    case 652:
        for fi97 := 0; fi97 < 17; fi97++ {
            if true {
                delete(v3, ("hello" + ""))
                var v98 float32 = 0.0
                v99 := GenS3{F0: 0, F1: 42, F2: v4}
                _ = v98
                _ = v99
            } else {
                var v100 string = ((v4 + (v4 + "")) + v4)
                if false {
                    var v101 int = 417
                    _ = v101
                } else {
                    var v102 uint64 = 42
                    switch ("foo bar" + (v4 + v4)) {
                    case "foo bar":
                        v100 = v4
                    case "a":
                        v4 = v100
                    case "\n\t":
                        v103 := GenS6{F0: true}
                        _ = v103
                    default:
                        fmt.Println(v100, v100)
                    }
                    _ = v102
                }
                delete(v3, v100)
                _ = v100
            }
            _ = fi97
        }
    default:
        v27 += genFn1(65535, (float32(0.1) < 3.14159), "\n\t")
    }
    switch 9223372036854775807 {
    case 2147483647:
        fmt.Println(v1, v1)
    case 42:
        if ((0.1 < 0.0) || true) {
            v1 += 42
            v2 += (-(v2))
            var v104 uint64 = 954
            _ = v104
        }
    case 4294967295:
        v105 := GenS12{F0: v1, F1: 189, F2: 826}
        _ = v105
    }
    var v106 int32 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v106:", v106)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v106
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v47
}
