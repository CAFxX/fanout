package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint64; F1 uint64; F2 int8 }
type GenS3 struct { F0 bool; F1 int32 }
type GenS4 struct { F0 uint32; F1 int16; F2 uint16 }
type GenS5 struct { F0 uint16; F1 int16; F2 uint32 }
type GenS6 struct { F0 int32; F1 uint64 }
type GenS7 struct { F0 int8; F1 uint16; F2 uint8 }
type GenS8 struct { F0 int64 }
type GenS9 struct { F0 int16 }
type GenS10 struct { F0 uint; F1 uint8 }
func genFn0() string {
    var v6 uint = 262
    for fi7 := 0; fi7 < 15; fi7++ {
        if true {
            for fi8 := 0; fi8 < 20; fi8++ {
                var v9 float32 = 0.1
                var v10 int = fi8
                v11 := GenS1{F0: int64(v6)}
                _ = fi8
                _ = v10
                _ = v11
                _ = v9
            }
            fmt.Println(fi7, v6)
            fmt.Println(v6, fi7)
        }
        _ = fi7
    }
    var v12 int = 0
    switch "a" {
    case "hello":
        switch (v12 ^ (v12 % (((v12 | (v12 * v12))) | 1))) {
        case -9223372036854775808:
            v6 -= 42
        case 977:
            v6 = (v6 ^ 4294967295)
        default:
            var v13 string = ("\n\t" + ("a" + "\n\t"))
            _ = v13
        }
    case "a":
        v14 := GenS1{F0: 9223372036854775807}
        _ = v14
    case "\n\t":
        v6 = (v6 - 42)
    default:
        switch "hello" {
        case "foo bar":
            v15 := GenS1{F0: 9223372036854775807}
            _ = v15
        case "\n\t":
            v16 := GenS1{F0: -10}
            _ = v16
        case "":
            v12 = 32768
        default:
            for fi17 := 0; fi17 < 6; fi17++ {
                for fi18 := 0; fi18 < 5; fi18++ {
                    v19 := GenS2{F0: 1, F1: 4294967295, F2: 100}
                    var v20 int = (v12 + (^(v12)))
                    _ = fi18
                    _ = v19
                    _ = v20
                }
                for fi21 := 0; fi21 < 17; fi21++ {
                    var v22 int = v12
                    v6 = (^(v6))
                    _ = fi21
                    _ = v22
                }
                v23 := GenS3{F0: (uint32(v6) >= 4294967295), F1: -2147483648}
                _ = fi17
                _ = v23
            }
        }
    }
    fmt.Println(v12)
    _ = v12
    _ = v6
    return "hello"
}
func genFn1(p0 int64, p1 uint) uint8 {
    v24 := []int{0, -9223372036854775808}
    if true {
        for fi25 := 0; fi25 < 3; fi25++ {
            var v26 uint = 1
            v27 := GenS2{F0: 82, F1: 2, F2: 100}
            _ = fi25
            _ = v26
            _ = v27
        }
        fmt.Println(p1, p1)
    } else {
        switch 42 {
        case 32767:
            p0 -= p0
        }
    }
    if false {
        var v28 uint8 = 255
        _ = v28
    }
    _ = p0
    _ = p1
    _ = v24
    return 127
}

func main() {
    var v1 int8 = 1
    v2 := make(map[string]string)
    v3 := make(map[string]string)
    var v4 uint8 = 42
    v5 := make(map[uint32]int)
    v4 = (^(v4))
    v1 = v1
    for fi29 := 0; fi29 < 9; fi29++ {
        if false {
            switch fi29 {
            case 127:
                var v30 int64 = 477
                _ = v30
            case -2147483648:
                var v31 int64 = 9223372036854775807
                _ = v31
            case 32768:
                switch "a" {
                case "":
                    for fi32 := 0; fi32 < 11; fi32++ {
                        var v33 int64 = -913
                        switch (fi32 << ((580) & 63)) {
                        case 714:
                            fmt.Println(v33, fi29, v4)
                        case -2147483648:
                            v1 = 1
                        case 128:
                            var v34 float32 = 3.14159
                            _ = v34
                        }
                        _ = fi32
                        _ = v33
                    }
                }
            default:
                var v35 uint64 = 544
                _ = v35
            }
        }
        for fi36 := 0; fi36 < 14; fi36++ {
            switch ("hello" + "hello") {
            case "hello":
                v4 += 142
            case "":
                v4 -= (^(v4))
            }
            var v37 int16 = 987
            v4 -= (v4 >> ((1) & 63))
            _ = fi36
            _ = v37
        }
        var v38 string = "\n\t"
        _ = fi29
        _ = v38
    }
    for fi39 := 0; fi39 < 11; fi39++ {
        switch ("" + "") {
        case "":
            v1 += v1
        case "foo bar":
            fmt.Println(fi39)
        case "hello":
            v40 := GenS2{F0: 2, F1: 42, F2: v1}
            _ = v40
        }
        _ = fi39
    }
    var v41 bool = ("hello" <= "\n\t")
    v41 = (!(32767 < 2147483647))
    fmt.Println(v1, v4, v41)
    switch "\n\t" {
    case "":
        var v42 string = (("foo bar" + "hello") + "foo bar")
        _ = v42
    case "hello":
        if ((-910 > -763) || v41) {
            v41 = (!true)
        }
    case "a":
        v43 := GenS4{F0: 2, F1: -118, F2: 2}
        _ = v43
    default:
        switch -2147483648 {
        case 747:
            v1 = (-(v1))
        case -9223372036854775808:
            if true {
                v41 = v41
                for fi44 := 0; fi44 < 18; fi44++ {
                    switch "\n\t" {
                    case "\n\t":
                        if v41 {
                            v41 = ((fi44 << ((255) & 63)) > fi44)
                            v4 = v4
                        } else {
                            fmt.Println(v41, fi44, v4)
                            v5[571] = (fi44 / ((fi44) | 1))
                            switch "foo bar" {
                            case "hello":
                                for fi45 := 0; fi45 < 14; fi45++ {
                                    v41 = v41
                                    _ = fi45
                                }
                            }
                        }
                    case "foo bar":
                        v46 := GenS1{F0: 0}
                        _ = v46
                    case "a":
                        for fi47 := 0; fi47 < 17; fi47++ {
                            v48 := GenS5{F0: 32767, F1: 1, F2: 4294967295}
                            v49 := GenS6{F0: -1, F1: 4294967295}
                            _ = fi47
                            _ = v48
                            _ = v49
                        }
                    default:
                        switch ((("a" + ("" + ("foo bar" + ((("hello" + "a") + "\n\t") + "foo bar")))) + "a") + ("foo bar" + "")) {
                        case "":
                            v41 = (!true)
                        case "\n\t":
                            var v50 int = (fi44 % ((602) | 1))
                            _ = v50
                        case "foo bar":
                            v51 := GenS7{F0: int8(v4), F1: 0, F2: (v4 / ((v4) | 1))}
                            _ = v51
                        default:
                            fmt.Println(v41, v41, fi44)
                        }
                    }
                    v5[42] = int(v4)
                    _ = fi44
                }
            } else {
                v41 = (!true)
                for fi52 := 0; fi52 < 12; fi52++ {
                    if (!v41) {
                        var v53 float64 = (0.0 * 1e10)
                        fmt.Println(v53, v4, v4)
                        _ = v53
                    }
                    _ = fi52
                }
                switch 28 {
                case 127:
                    for fi54 := 0; fi54 < 15; fi54++ {
                        switch "" {
                        case "\n\t":
                            var v55 bool = (true || false)
                            _ = v55
                        case "":
                            v56 := GenS4{F0: 4294967295, F1: 674, F2: 42}
                            _ = v56
                        default:
                            for fi57 := 0; fi57 < 18; fi57++ {
                                if (710 >= 288) {
                                    switch fi57 {
                                    case -2147483648:
                                        v58 := GenS8{F0: 42}
                                        _ = v58
                                    case -142:
                                        v2["hello"] = "hello"
                                    case -9223372036854775808:
                                        fmt.Println(fi54, fi57)
                                    }
                                } else {
                                    switch "foo bar" {
                                    case "":
                                        v41 = (v41 || v41)
                                    }
                                }
                                if (!((v41 || true) && true)) {
                                    switch fi54 {
                                    case 128:
                                        v59 := GenS5{F0: 255, F1: 191, F2: 2147483647}
                                        _ = v59
                                    case 2147483647:
                                        var v60 int = -9223372036854775808
                                        _ = v60
                                    case -274:
                                        v61 := []int{}
                                        _ = v61
                                    }
                                }
                                v62 := []int{-2147483648}
                                v41 = v41
                                _ = fi57
                                _ = v62
                            }
                        }
                        _ = fi54
                    }
                case -9223372036854775808:
                    if v41 {
                        v63 := GenS7{F0: -1, F1: 1000, F2: v4}
                        v41 = (!false)
                        _ = v63
                    }
                case 32767:
                    if (-2.25 <= 0.0) {
                        v64 := []int{-460, -293, -9223372036854775808, 42}
                        if v41 {
                            var v65 int = 138
                            _ = v65
                        }
                        _ = v64
                    }
                default:
                    for fi66 := 0; fi66 < 17; fi66++ {
                        v67 := []int{127, fi66, (fi66 + -226), fi66}
                        v4 = v4
                        switch (fi66 * fi66) {
                        case -74:
                            if false {
                                if (1 > 709) {
                                    v68 := GenS9{F0: 1}
                                    _ = v68
                                } else {
                                    var v69 uint64 = 89
                                    fmt.Println(v69, v4, v69)
                                    _ = v69
                                }
                            } else {
                                for fi70 := 0; fi70 < 5; fi70++ {
                                    v41 = (!false)
                                    if (0 == 42) {
                                        switch "" {
                                        case "\n\t":
                                            fmt.Println(v41)
                                        }
                                        v71 := GenS6{F0: -855, F1: 18446744073709551615}
                                        _ = v71
                                    }
                                    _ = fi70
                                }
                                v72 := GenS10{F0: 255, F1: (v4 ^ v4)}
                                _ = v72
                            }
                        }
                        _ = fi66
                        _ = v67
                    }
                }
            }
        case 32768:
            v73 := []int{32768}
            _ = v73
        default:
            fmt.Println(v1, v4, v4)
        }
    }
    for fi74 := 0; fi74 < 17; fi74++ {
        switch (^(fi74)) {
        case 163:
            v4 = (v4 % ((v4) | 1))
        case 32768:
            for fi75 := 0; fi75 < 17; fi75++ {
                v4 -= genFn1(-1, 255)
                if true {
                    var v76 uint64 = 42
                    _ = v76
                } else {
                    var v77 bool = v41
                    v77 = ((-(v1)) != v1)
                    _ = v77
                }
                v78 := []int{}
                var v79 int64 = 0
                _ = fi75
                _ = v78
                _ = v79
            }
        case 42:
            v41 = ((0.1 * 0.0) > 1e10)
        }
        var v80 uint32 = 605
        _ = fi74
        _ = v80
    }
    for fi81 := 0; fi81 < 9; fi81++ {
        v82 := []int{fi81, fi81, 32768, (fi81 ^ (fi81 ^ (fi81 / ((32767) | 1))))}
        _ = fi81
        _ = v82
    }
    v1 = 127
    v83 := GenS7{F0: -128, F1: 455, F2: (v4 >> ((957) & 63))}
    var v84 uint32 = 2147483647
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v84:", v84)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v5
    _ = v83
    _ = v84
}
