package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint16; F1 int64; F2 bool }
type GenS3 struct { F0 int16; F1 uint; F2 int }
type GenS4 struct { F0 uint; F1 int64; F2 uint32 }
type GenS5 struct { F0 int16; F1 uint; F2 string }
type GenS6 struct { F0 string; F1 int16; F2 uint64 }
type GenS7 struct { F0 uint; F1 uint; F2 int64 }
type GenS8 struct { F0 int32; F1 int; F2 int32 }
type GenS9 struct { F0 uint64; F1 uint32 }
type GenS10 struct { F0 bool }
type GenS11 struct { F0 uint32 }
type GenS12 struct { F0 uint; F1 bool; F2 int64 }
func genFn0(p0 uint32) uint8 {
    for fi4 := 0; fi4 < 15; fi4++ {
        for fi5 := 0; fi5 < 11; fi5++ {
            if (uint16(p0) >= 32767) {
                if ((32767 < -32768) && true) {
                    switch (fi4 % (((fi5 << ((4) & 63))) | 1)) {
                    case 693:
                        var v6 uint16 = 42
                        _ = v6
                    case -9223372036854775808:
                        var v7 int8 = 100
                        v8 := GenS1{F0: p0}
                        _ = v7
                        _ = v8
                    case -299:
                        if (!(42 < (fi5 << ((481) & 63)))) {
                            if (0.1 == 0.1) {
                                var v9 uint = 2
                                _ = v9
                            } else {
                                var v10 uint32 = p0
                                var v11 int64 = 1
                                _ = v10
                                _ = v11
                            }
                        }
                    default:
                        p0 -= p0
                    }
                    var v12 uint16 = 0
                    var v13 uint16 = (v12 - (v12 - 1))
                    _ = v12
                    _ = v13
                }
                p0 = p0
                p0 = p0
            }
            switch fi5 {
            case 537:
                var v14 uint32 = p0
                _ = v14
            case 45:
                if true {
                    v15 := []int{(fi4 - (fi4 >> ((4294967295) & 63)))}
                    v16 := GenS2{F0: 0, F1: 42, F2: (!((p0 >> ((18446744073709551615) & 63)) != (p0 ^ (p0 % (((p0 % ((p0) | 1))) | 1)))))}
                    _ = v15
                    _ = v16
                } else {
                    v17 := []int{32768, fi4, -394}
                    var v18 uint16 = 255
                    v19 := GenS2{F0: 2, F1: 309, F2: false}
                    v17 = append(v17, fi4)
                    _ = v17
                    _ = v18
                    _ = v19
                }
                p0 -= 65535
            case 32768:
                p0 += (^(p0))
            default:
                v20 := []int{fi4, fi5}
                if false {
                    var v21 int32 = -2147483648
                    fmt.Println(p0)
                    _ = v21
                }
                _ = v20
            }
            _ = fi5
        }
        if true {
            for fi22 := 0; fi22 < 15; fi22++ {
                if ((634 != p0) && false) {
                    v23 := GenS1{F0: 4294967295}
                    p0 = p0
                    _ = v23
                }
                if true {
                    v24 := GenS2{F0: 65535, F1: 359, F2: true}
                    if (141 < 1) {
                        v25 := GenS3{F0: int16(fi4), F1: 0, F2: fi22}
                        _ = v25
                    } else {
                        v26 := GenS4{F0: uint(p0), F1: 1, F2: (p0 & 2147483647)}
                        switch (fi22 / (((-(fi4))) | 1)) {
                        case 127:
                            p0 -= 2147483647
                        case 42:
                            var v27 float32 = float32(1.5)
                            if (true && (0 < 0)) {
                                v28 := []int{fi22, (fi22 & (fi22 ^ fi4)), fi4, (^(fi22))}
                                var v29 int8 = 42
                                var v30 int64 = 0
                                v30 -= (v30 * 9223372036854775807)
                                _ = v28
                                _ = v29
                                _ = v30
                            }
                            switch (fi4 + fi22) {
                            case 32768:
                                fmt.Println(fi4, fi22)
                            case -175:
                                switch (fi4 - (^(fi22))) {
                                case 4294967295:
                                    switch (fi22 << ((18446744073709551615) & 63)) {
                                    case -1:
                                        v27 = (1.5 + (((3.14159 * v27) + 0.0) + v27))
                                    default:
                                        var v31 uint32 = (p0 % ((p0) | 1))
                                        v32 := []int{128}
                                        _ = v31
                                        _ = v32
                                    }
                                default:
                                    v33 := GenS2{F0: 892, F1: 0, F2: true}
                                    _ = v33
                                }
                            }
                            switch fi22 {
                            case -280:
                                switch "hello" {
                                case "foo bar":
                                    switch "\n\t" {
                                    case "foo bar":
                                        v34 := GenS2{F0: 2, F1: 1, F2: (true || (32767 < 862))}
                                        _ = v34
                                    case "a":
                                        switch (fi22 - 4294967295) {
                                        case 32768:
                                            v27 = 0.1
                                        case -1:
                                            var v35 uint64 = 0
                                            _ = v35
                                        }
                                    default:
                                        if (false || true) {
                                            if (557 <= 684) {
                                                fmt.Println(fi22, p0)
                                                fmt.Println(v27)
                                            } else {
                                                var v36 uint16 = 1
                                                var v37 int32 = 108
                                                _ = v36
                                                _ = v37
                                            }
                                            var v38 float32 = v27
                                            _ = v38
                                        }
                                    }
                                case "":
                                    v39 := GenS4{F0: 18446744073709551615, F1: 1, F2: (p0 % (((p0 | (p0 >> ((42) & 63)))) | 1))}
                                    _ = v39
                                case "hello":
                                    v40 := []int{fi22}
                                    _ = v40
                                default:
                                    var v41 uint64 = 532
                                    _ = v41
                                }
                            case 9223372036854775807:
                                p0 += (^(p0))
                            default:
                                if false {
                                    v27 = v27
                                    v42 := GenS4{F0: 2, F1: 9223372036854775807, F2: (^(p0))}
                                    if (false || (1.5 > float32(0.0))) {
                                        p0 = (^(p0))
                                    }
                                    _ = v42
                                }
                            }
                            v43 := []int{-562, fi22}
                            _ = v27
                            _ = v43
                        case 128:
                            var v44 uint32 = 2147483647
                            _ = v44
                        default:
                            var v45 int16 = -859
                            _ = v45
                        }
                        if (true || (42 == -1)) {
                            fmt.Println(fi22)
                            p0 = p0
                            var v46 int32 = 1
                            v47 := GenS5{F0: -327, F1: 18446744073709551615, F2: "\n\t"}
                            _ = v46
                            _ = v47
                        }
                        _ = v26
                    }
                    fmt.Println(p0)
                    _ = v24
                }
                if (true && ((("\n\t" + "foo bar") + (("foo bar" + (("hello" + "hello") + "hello")) + (("foo bar" + "foo bar") + ("foo bar" + "\n\t")))) > ("a" + "hello"))) {
                    v48 := []int{(fi4 % (((-(fi4))) | 1))}
                    _ = v48
                } else {
                    switch fi22 {
                    case 128:
                        v49 := GenS6{F0: ("a" + "a"), F1: -1, F2: 4294967295}
                        _ = v49
                    case -2147483648:
                        var v50 uint8 = 1
                        _ = v50
                    }
                    var v51 int16 = 42
                    if true {
                        if (!(1 < -128)) {
                            fmt.Println(v51, fi22)
                            var v52 int16 = v51
                            v53 := []int{(fi4 * (fi4 % (((fi4 & (fi22 / (((fi22 / (((fi4 ^ (fi4 - fi22))) | 1))) | 1)))) | 1))), (fi4 >> ((18446744073709551615) & 63))}
                            _ = v52
                            _ = v53
                        }
                        v51 = v51
                        switch (fi22 ^ fi4) {
                        case 4294967295:
                            p0 += (p0 >> ((4294967295) & 63))
                        case -515:
                            switch (fi4 ^ int(v51)) {
                            case 199:
                                v54 := []int{}
                                _ = v54
                            default:
                                fmt.Println(fi4, p0)
                            }
                        default:
                            switch (fi4 >> ((1) & 63)) {
                            case 1:
                                v55 := GenS7{F0: 494, F1: 0, F2: int64(v51)}
                                _ = v55
                            case -828:
                                v51 += v51
                            case 764:
                                p0 += p0
                            default:
                                v56 := GenS4{F0: 1, F1: 42, F2: (p0 >> ((18446744073709551615) & 63))}
                                _ = v56
                            }
                        }
                    } else {
                        v57 := []int{127, (fi4 * fi4), 29}
                        var v58 uint8 = 2
                        _ = v57
                        _ = v58
                    }
                    _ = v51
                }
                _ = fi22
            }
            p0 = p0
            p0 = p0
        }
        v59 := []int{(fi4 >> ((42) & 63)), (-(fi4))}
        _ = fi4
        _ = v59
    }
    var v60 int = 32767
    var v61 uint32 = (p0 % ((p0) | 1))
    v62 := []int{(v60 << ((0) & 63)), v60, v60, (v60 >> ((4294967295) & 63))}
    _ = p0
    _ = v60
    _ = v61
    _ = v62
    return 132
}

func main() {
    var v1 uint8 = 1
    v2 := []string{"", ("" + "\n\t"), "a"}
    v3 := make(map[string]uint32)
    v63 := GenS8{F0: 1, F1: 9223372036854775807, F2: 2147483647}
    switch ("" + "a") {
    case "a":
        if false {
            if true {
                for fi64 := 0; fi64 < 2; fi64++ {
                    v65 := GenS8{F0: -1, F1: -2147483648, F2: 0}
                    _ = fi64
                    _ = v65
                }
            }
            fmt.Println(v1)
        }
    case "":
        if (1 != 9223372036854775807) {
            v66 := GenS1{F0: 271}
            _ = v66
        }
    }
    if (!(false || true)) {
        switch 387 {
        case 457:
            var v67 float32 = 1.5
            _ = v67
        case -1:
            v1 = v1
        case 32768:
            if true {
                var v68 float64 = 3.14159
                v69 := GenS9{F0: 2, F1: 2}
                v70 := GenS1{F0: 738}
                _ = v68
                _ = v69
                _ = v70
            } else {
                v71 := GenS10{F0: (4294967295 != 609)}
                if true {
                    v1 = 44
                    v72 := GenS11{F0: 0}
                    v73 := GenS4{F0: 255, F1: -1, F2: 2147483647}
                    _ = v72
                    _ = v73
                }
                _ = v71
            }
        }
    } else {
        var v74 uint8 = 127
        _ = v74
    }
    fmt.Println(v1)
    var v75 uint16 = uint16(v1)
    v75 = v75
    var v76 int = int(v1)
    switch "hello" {
    case "hello":
        if (false || (42 == 42)) {
            var v77 bool = true
            var v78 bool = (18446744073709551615 <= 191)
            for fi79 := 0; fi79 < 13; fi79++ {
                switch 156 {
                case 32767:
                    if len(v2) > 0 {
                        v2[0] = "foo bar"
                    }
                default:
                    v2 = append(v2, ("hello" + "a"))
                }
                if (!false) {
                    var v80 int8 = 100
                    for fi81 := 0; fi81 < 19; fi81++ {
                        fmt.Println(v77, fi81)
                        _ = fi81
                    }
                    for fi82 := 0; fi82 < 14; fi82++ {
                        v83 := GenS12{F0: 1, F1: v78, F2: 0}
                        var v84 int64 = -128
                        v77 = v77
                        _ = fi82
                        _ = v83
                        _ = v84
                    }
                    _ = v80
                } else {
                    v3[""] = uint32(v1)
                    var v85 float32 = (-2.25 + 1e10)
                    _ = v85
                }
                _ = fi79
            }
            _ = v77
            _ = v78
        } else {
            if (787 <= 4294967295) {
                fmt.Println(v1, v75)
                if len(v2) > 0 {
                    v2[0] = "\n\t"
                }
                v75 = (v75 ^ 662)
            }
        }
    default:
        fmt.Println(v1, v1, v1)
    }
    fmt.Println(v1, v75, v76)
    fmt.Println(v1, v1)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v75:", v75)
    fmt.Println("v76:", v76)
    _ = v1
    _ = v2
    _ = v3
    _ = v63
    _ = v75
    _ = v76
}
