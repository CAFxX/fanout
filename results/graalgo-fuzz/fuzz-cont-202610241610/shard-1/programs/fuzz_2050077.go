package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32 }
type GenS2 struct { F0 bool; F1 int16; F2 int64 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int8; F1 uint }
type GenS5 struct { F0 string }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 int16; F1 string }
type GenS8 struct { F0 bool; F1 uint64; F2 string }
type GenS9 struct { F0 uint16; F1 int64 }
type GenS10 struct { F0 int8; F1 int16; F2 uint64 }
func genFn0(p0 uint32, p1 int32, p2 int) uint64 {
    p0 = p0
    var v4 uint64 = 881
    for fi5 := 0; fi5 < 9; fi5++ {
        if (true || true) {
            v6 := []int{(p2 & fi5), int(p1), (fi5 | 127), fi5}
            fmt.Println(v4, p2)
            _ = v6
        }
        _ = fi5
    }
    if false {
        v7 := []int{1}
        if (v4 <= uint64(p0)) {
            if len(v7) > 0 {
                v7[0] = 55
            }
            var v8 int64 = -9223372036854775808
            if (true || (true && (-1 != int32(v4)))) {
                var v9 bool = (((88 <= 759) && false) && true)
                v4 = (v4 >> ((255) & 63))
                if v9 {
                    if len(v7) > 0 {
                        v7[0] = (^(p2))
                    }
                    _ = len(v7)
                } else {
                    var v10 int16 = -422
                    v7 = append(v7, (^(p2)))
                    _ = v10
                }
                _ = v9
            } else {
                for fi11 := 0; fi11 < 8; fi11++ {
                    v12 := GenS1{F0: 62, F1: (p0 | (p0 % (((p0 >> ((2) & 63))) | 1)))}
                    _ = fi11
                    _ = v12
                }
            }
            _ = v8
        } else {
            var v13 float64 = 1e10
            _ = v13
        }
        switch "foo bar" {
        case "foo bar":
            switch -1 {
            case 32767:
                for fi14 := 0; fi14 < 16; fi14++ {
                    v15 := GenS2{F0: ((false || (18446744073709551615 > 820)) || (4294967295 != v4)), F1: 32767, F2: -1}
                    _ = fi14
                    _ = v15
                }
            case 1:
                var v16 float32 = 3.14159
                _ = v16
            case 4294967295:
                var v17 float64 = 3.14159
                _ = v17
            default:
                switch p2 {
                case 195:
                    if len(v7) >= 2 {
                        v7 = v7[1:len(v7)]
                    }
                default:
                    if len(v7) >= 2 {
                        v7 = v7[1:len(v7)]
                    }
                }
            }
        case "a":
            if (-1 < p1) {
                v18 := GenS1{F0: (v4 >> ((2) & 63)), F1: 1}
                v7 = append(v7, p2)
                _ = v18
            }
        default:
            p2 += (p2 << ((uint(v4)) & 63))
        }
        _ = v7
    }
    if (true || true) {
        var v19 bool = (42 != 32767)
        switch ("foo bar" + "hello") {
        case "":
            p0 += (^(p0))
        default:
            v20 := []int{(-(p2)), p2, (p2 * p2)}
            _ = v20
        }
        _ = v19
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return (v4 ^ 4294967295)
}

func main() {
    var v1 uint8 = 0
    var v2 int8 = 1
    var v3 int = 9223372036854775807
    if ((42 == 0) || false) {
        v21 := GenS1{F0: 0, F1: 0}
        var v22 float64 = -2.25
        v3 += (v3 ^ v3)
        _ = v21
        _ = v22
    }
    if (27 >= 58) {
        for fi23 := 0; fi23 < 5; fi23++ {
            v24 := GenS3{F0: (560 == 2147483647)}
            _ = fi23
            _ = v24
        }
        v25 := GenS3{F0: false}
        _ = v25
    } else {
        switch (v3 ^ (v3 | -808)) {
        case 9223372036854775807:
            v26 := GenS2{F0: true, F1: 32767, F2: -312}
            _ = v26
        case -638:
            v1 += (v1 << ((0) & 63))
        case 4294967295:
            v27 := GenS1{F0: 97, F1: 42}
            _ = v27
        }
        v28 := []int{v3, (v3 & (v3 & (v3 - 127))), (v3 << ((18446744073709551615) & 63)), (v3 % ((-2147483648) | 1))}
        for fi29 := 0; fi29 < 4; fi29++ {
            switch (-(fi29)) {
            case 449:
                for fi30 := 0; fi30 < 10; fi30++ {
                    var v31 float32 = (1.5 + (1.5 / 1.0))
                    if ((42 == -1) && false) {
                        v28 = append(v28, 1)
                        if (463 != 310) {
                            switch "" {
                            case "":
                                var v32 uint64 = 0
                                _ = v32
                            case "\n\t":
                                _ = len(v28)
                            default:
                                v1 = v1
                                v2 = v2
                            }
                        } else {
                            if true {
                                var v33 int8 = 100
                                fmt.Println(fi30, v33)
                                _ = v33
                            } else {
                                switch "\n\t" {
                                case "foo bar":
                                    fmt.Println(fi30)
                                default:
                                    fmt.Println(v31, v1)
                                }
                                var v34 float32 = 0.1
                                _ = v34
                            }
                            var v35 float64 = 0.1
                            v2 = 1
                            _ = v35
                        }
                        switch v3 {
                        case 128:
                            var v36 uint = 4294967295
                            _ = v36
                        case 32768:
                            if len(v28) >= 2 {
                                v28 = v28[1:len(v28)]
                            }
                        }
                    }
                    _ = fi30
                    _ = v31
                }
            case 7:
                if false {
                    var v37 uint = 2
                    _ = v37
                } else {
                    v38 := GenS2{F0: false, F1: 97, F2: -9223372036854775808}
                    v2 = (v2 / ((v2) | 1))
                    var v39 uint = 42
                    _ = v38
                    _ = v39
                }
            }
            v2 -= (v2 >> ((194) & 63))
            _ = fi29
        }
        _ = v28
    }
    fmt.Println(v1, v2, v2)
    var v40 int64 = 1
    v2 -= v2
    v41 := GenS2{F0: true, F1: 32767, F2: 9223372036854775807}
    var v42 uint32 = 542
    v2 = -29
    var v43 int = v3
    switch "hello" {
    case "hello":
        if (v1 > v1) {
            for fi44 := 0; fi44 < 10; fi44++ {
                for fi45 := 0; fi45 < 17; fi45++ {
                    if (791 != 1000) {
                        if (!(784 <= 2147483647)) {
                            var v46 uint = 557
                            if false {
                                if (42 < 325) {
                                    var v47 uint64 = genFn0(v42, 42, int(v42))
                                    v1 = 2
                                    _ = v47
                                }
                                v48 := GenS4{F0: v2, F1: 0}
                                _ = v48
                            } else {
                                var v49 uint64 = 4
                                _ = v49
                            }
                            v50 := GenS5{F0: ("" + "foo bar")}
                            _ = v46
                            _ = v50
                        }
                    }
                    v51 := GenS5{F0: "a"}
                    switch ("a" + "foo bar") {
                    case "foo bar":
                        v2 = v2
                    case "a":
                        v40 += v40
                    case "hello":
                        var v52 int64 = 0
                        _ = v52
                    default:
                        switch (("\n\t" + "a") + ("a" + ("foo bar" + ("foo bar" + "\n\t")))) {
                        case "":
                            v3 += v3
                        case "\n\t":
                            switch (fi45 + fi45) {
                            case 42:
                                v53 := []int{(-(v43)), (^(fi45)), (fi44 / (((v3 % ((fi45) | 1))) | 1)), fi45}
                                if ((v42 ^ v42) > (v42 << ((4294967295) & 63))) {
                                    var v54 uint8 = (v1 | v1)
                                    v1 = (v54 >> ((293) & 63))
                                    _ = v54
                                }
                                _ = v53
                            case -2147483648:
                                v2 = v2
                            case 642:
                                var v55 int = -470
                                _ = v55
                            }
                        }
                    }
                    _ = fi45
                    _ = v51
                }
                if false {
                    v40 += int64(v42)
                    if true {
                        switch (fi44 & -9223372036854775808) {
                        case 0:
                            v42 -= (v42 << ((765) & 63))
                        case -1:
                            v56 := GenS3{F0: false}
                            _ = v56
                        default:
                            v3 = (v43 ^ v43)
                        }
                        v57 := []int{int(v2), 127, v43, (fi44 << ((4294967295) & 63))}
                        _ = v57
                    } else {
                        v40 = (-(v40))
                    }
                    fmt.Println(v43)
                }
                var v58 int16 = -1
                _ = fi44
                _ = v58
            }
            if (-32768 > 1) {
                var v59 int = v43
                v60 := GenS6{F0: false}
                _ = v59
                _ = v60
            } else {
                fmt.Println(v1, v43)
                var v61 uint16 = 2
                _ = v61
            }
            v62 := GenS7{F0: 1, F1: "foo bar"}
            _ = v62
        }
    case "a":
        v63 := GenS1{F0: 1, F1: v42}
        _ = v63
    case "":
        v42 += (^(v42))
    default:
        for fi64 := 0; fi64 < 6; fi64++ {
            var v65 float64 = 1.5
            _ = fi64
            _ = v65
        }
    }
    v2 = (v2 % ((-1) | 1))
    for fi66 := 0; fi66 < 18; fi66++ {
        switch "foo bar" {
        case "":
            fmt.Println(v42)
        case "\n\t":
            for fi67 := 0; fi67 < 6; fi67++ {
                v68 := []int{v3, 1, (v3 & v3)}
                v43 = 127
                v69 := GenS8{F0: (!(1.5 > -2.25)), F1: 753, F2: "hello"}
                _ = fi67
                _ = v68
                _ = v69
            }
        case "foo bar":
            var v70 uint64 = 2
            _ = v70
        default:
            v43 = (fi66 ^ (v3 * (v3 / (((fi66 / (((fi66 / ((v43) | 1))) | 1))) | 1))))
        }
        _ = fi66
    }
    var v71 int = v43
    v72 := GenS9{F0: 499, F1: v40}
    v1 += v1
    for fi73 := 0; fi73 < 6; fi73++ {
        v74 := []int{(v71 | (v3 % ((v71) | 1))), (fi73 / (((v71 + int(v42))) | 1)), -884}
        for fi75 := 0; fi75 < 13; fi75++ {
            var v76 int16 = int16(v2)
            _ = fi75
            _ = v76
        }
        for fi77 := 0; fi77 < 5; fi77++ {
            switch "\n\t" {
            case "a":
                if (!((!(2 <= 718)) && (!false))) {
                    v78 := GenS10{F0: -128, F1: -571, F2: 1}
                    if false {
                        var v79 int8 = v2
                        _ = v79
                    } else {
                        v3 += (-(v43))
                    }
                    v80 := GenS2{F0: false, F1: 1000, F2: -1}
                    _ = v78
                    _ = v80
                }
            case "hello":
                if true {
                    var v81 int16 = 42
                    v74 = append(v74, (^(v71)))
                    _ = v81
                }
            }
            _ = fi77
        }
        _ = fi73
        _ = v74
    }
    v82 := []int{v71, v3, v43, v71}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v40:", v40)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v71:", v71)
    fmt.Println("slicelen v82:", len(v82), v82)
    _ = v1
    _ = v2
    _ = v3
    _ = v40
    _ = v41
    _ = v42
    _ = v43
    _ = v71
    _ = v72
    _ = v82
}
