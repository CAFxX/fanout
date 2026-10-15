package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint16; F1 int64; F2 int64 }
type GenS3 struct { F0 int32; F1 uint8 }
type GenS4 struct { F0 uint16; F1 string }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 uint64; F1 string; F2 int64 }
type GenS7 struct { F0 bool; F1 string; F2 int8 }
type GenS8 struct { F0 string; F1 uint32; F2 bool }
type GenS9 struct { F0 int; F1 int32; F2 int32 }
type GenS10 struct { F0 int64 }
type GenS11 struct { F0 bool; F1 uint8 }
type GenS12 struct { F0 int }
func genFn0() string {
    var v5 float32 = 0.0
    v5 = v5
    switch 4294967295 {
    case 4294967295:
        switch 9223372036854775807 {
        case 0:
            v5 += (0.1 + v5)
        default:
            var v6 uint32 = 42
            _ = v6
        }
    case 32768:
        for fi7 := 0; fi7 < 8; fi7++ {
            for fi8 := 0; fi8 < 9; fi8++ {
                v9 := GenS1{F0: 2}
                v10 := GenS1{F0: 713}
                _ = fi8
                _ = v10
                _ = v9
            }
            _ = fi7
        }
    case -9223372036854775808:
        fmt.Println(v5)
    }
    switch (("a" + ("hello" + "\n\t")) + "\n\t") {
    case "hello":
        if true {
            for fi11 := 0; fi11 < 13; fi11++ {
                fmt.Println(fi11, v5)
                v12 := []int{}
                _ = fi11
                _ = v12
            }
            var v13 float32 = v5
            _ = v13
        } else {
            var v14 int16 = 42
            if (("\n\t" + "\n\t") < "") {
                v14 += v14
                switch "" {
                case "\n\t":
                    var v15 uint64 = 42
                    _ = v15
                case "a":
                    v16 := GenS1{F0: 580}
                    _ = v16
                case "hello":
                    if false {
                        v17 := GenS1{F0: 65535}
                        _ = v17
                    } else {
                        var v18 int32 = 1
                        _ = v18
                    }
                default:
                    v19 := []int{4294967295, 32767}
                    _ = v19
                }
                v20 := []int{int(v14)}
                _ = v20
            }
            v21 := GenS2{F0: 0, F1: 790, F2: 9223372036854775807}
            _ = v14
            _ = v21
        }
    default:
        var v22 int32 = 1
        _ = v22
    }
    _ = v5
    return "\n\t"
}
func genFn1(p0 uint, p1 int8, p2 uint32) int8 {
    var v23 uint8 = 2
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v23
    return p1
}

func main() {
    var v1 int8 = 1
    var v2 uint32 = 4294967295
    var v3 int8 = v1
    var v4 int8 = (v1 | v1)
    var v24 uint32 = v2
    for fi25 := 0; fi25 < 4; fi25++ {
        var v26 uint64 = 438
        v27 := GenS3{F0: 230, F1: uint8(v26)}
        v28 := GenS3{F0: -142, F1: 42}
        _ = fi25
        _ = v26
        _ = v27
        _ = v28
    }
    if (42 <= -2147483648) {
        v3 -= (v1 % ((v4) | 1))
        fmt.Println(v1, v24, v4)
    }
    v4 = genFn1(4294967295, genFn1(0, (-(v4)), v2), 0)
    switch 127 {
    case 9223372036854775807:
        var v29 uint64 = 0
        _ = v29
    }
    for fi30 := 0; fi30 < 8; fi30++ {
        v31 := []int{42, (fi30 * fi30), fi30}
        for fi32 := 0; fi32 < 15; fi32++ {
            switch 4294967295 {
            case -2147483648:
                if (genFn1(0, genFn1(977, (v3 ^ v4), (v24 ^ v2)), v24) != 1) {
                    v1 = (^(v1))
                    if (32767 >= 706) {
                        v2 = 0
                        if (42 < 47) {
                            switch (("hello" + "") + ("a" + "\n\t")) {
                            case "":
                                if len(v31) >= 2 {
                                    v31 = v31[1:len(v31)]
                                }
                            case "a":
                                v31 = append(v31, fi32)
                            case "hello":
                                if (0 == 75) {
                                    switch fi30 {
                                    case 32768:
                                        _ = len(v31)
                                    case 128:
                                        _ = len(v31)
                                    case -521:
                                        if (967 == 1) {
                                            var v33 string = "foo bar"
                                            v34 := GenS4{F0: 945, F1: v33}
                                            if len(v31) > 0 {
                                                v31[0] = fi30
                                            }
                                            _ = v33
                                            _ = v34
                                        } else {
                                            v35 := GenS5{F0: 714}
                                            v4 = v4
                                            _ = v35
                                        }
                                    }
                                    if (("\n\t" >= "foo bar") && false) {
                                        v36 := GenS3{F0: 485, F1: 39}
                                        var v37 float64 = (float64(float32(3.14159)) / 3.14159)
                                        v24 = 614
                                        _ = v36
                                        _ = v37
                                    }
                                }
                                var v38 float64 = 0.0
                                v3 -= genFn1(391, v1, (v24 >> ((4294967295) & 63)))
                                _ = v38
                            }
                        } else {
                            fmt.Println(v3, fi32, v2)
                            v39 := GenS4{F0: 321, F1: "a"}
                            var v40 bool = true
                            _ = v39
                            _ = v40
                        }
                        v3 -= (v4 * 96)
                    } else {
                        var v41 int16 = 784
                        var v42 float32 = 1e10
                        if (-1 <= (v41 * v41)) {
                            var v43 int = (fi32 << ((255) & 63))
                            v3 = v4
                            _ = v43
                        } else {
                            v24 += (v2 / (((v24 << ((2) & 63))) | 1))
                        }
                        v3 -= v3
                        _ = v41
                        _ = v42
                    }
                    fmt.Println(v2, v4)
                    fmt.Println(v1, v3, v1)
                    v44 := GenS6{F0: 279, F1: "hello", F2: -9223372036854775808}
                    if true {
                        if false {
                            v31 = append(v31, (fi32 % ((9223372036854775807) | 1)))
                            v2 = (^(v2))
                        } else {
                            v2 -= 42
                            var v45 float64 = 0.0
                            _ = v45
                        }
                        if (127 == 78) {
                            fmt.Println(v1, v24, fi30)
                            var v46 bool = (!(0.1 != 0.1))
                            switch (("\n\t" + ((("" + "foo bar") + ("\n\t" + ("" + ("" + "a")))) + "hello")) + ((("\n\t" + "") + "a") + "")) {
                            case "\n\t":
                                v2 = 506
                            default:
                                switch ("hello" + ("" + (("a" + ((("\n\t" + ("hello" + (((("a" + "a") + "") + "") + ("\n\t" + ((("hello" + "hello") + ("hello" + ("\n\t" + (("hello" + ("\n\t" + ("a" + ""))) + "hello")))) + "foo bar"))))) + "foo bar") + ("" + ("hello" + "\n\t")))) + "a"))) {
                                case "hello":
                                    switch "\n\t" {
                                    case "foo bar":
                                        if (-1 != 141) {
                                            v3 = (v1 / (((^(v4))) | 1))
                                            switch "hello" {
                                            case "a":
                                                var v47 uint8 = 152
                                                _ = v47
                                            default:
                                                v48 := GenS2{F0: 42, F1: 9223372036854775807, F2: 1}
                                                _ = v48
                                            }
                                        } else {
                                            v49 := GenS4{F0: 134, F1: ""}
                                            var v50 int32 = 2147483647
                                            if (true || v46) {
                                                if (true || (!v46)) {
                                                    v51 := GenS7{F0: v46, F1: ("foo bar" + ""), F2: v3}
                                                    _ = v51
                                                } else {
                                                    if v46 {
                                                        switch 478 {
                                                        case 4294967295:
                                                            var v52 float32 = (-2.25 / 1.0)
                                                            _ = v52
                                                        }
                                                        if (!(fi32 == fi30)) {
                                                            var v53 uint64 = 1
                                                            _ = v53
                                                        }
                                                    }
                                                    fmt.Println(v24, v3, v1)
                                                }
                                            }
                                            _ = v49
                                            _ = v50
                                        }
                                    }
                                case "a":
                                    v31 = append(v31, 2147483647)
                                    var v54 uint16 = 42
                                    _ = v54
                                case "\n\t":
                                    v55 := GenS8{F0: ("\n\t" + ("foo bar" + "a")), F1: (v24 >> ((157) & 63)), F2: v46}
                                    _ = v55
                                }
                            }
                            _ = v46
                        }
                        v1 -= (v4 >> ((4294967295) & 63))
                    } else {
                        var v56 uint32 = (v2 % ((4294967295) | 1))
                        if (v1 > genFn1(4294967295, (v3 ^ (v3 & v4)), v56)) {
                            v31 = append(v31, (^(fi30)))
                        }
                        v4 -= v1
                        _ = v56
                    }
                    var v57 int16 = -463
                    _ = v44
                    _ = v57
                } else {
                    if false {
                        if false {
                            v24 -= 2147483648
                            if len(v31) > 0 {
                                v31[0] = -828
                            }
                            v4 = v1
                        } else {
                            if (true && (true && (229 == 127))) {
                                switch (("\n\t" + (("" + "a") + "hello")) + (("hello" + "") + "a")) {
                                case "":
                                    _ = len(v31)
                                default:
                                    v2 = v2
                                }
                                if (-2147483648 >= -1) {
                                    fmt.Println(fi30, v1)
                                    var v58 string = (("" + "") + "foo bar")
                                    _ = v58
                                } else {
                                    v59 := GenS9{F0: -9223372036854775808, F1: 2147483647, F2: 1}
                                    _ = v59
                                }
                            }
                            var v60 uint16 = 558
                            _ = v60
                        }
                        if len(v31) > 0 {
                            v31[0] = 201
                        }
                    } else {
                        if (-2.25 >= float64(((1.5 + 0.1) - float32(3.14159)))) {
                            if (!(42 < 255)) {
                                if len(v31) > 0 {
                                    v31[0] = (fi30 / (((fi32 * (fi32 / ((fi30) | 1)))) | 1))
                                }
                            } else {
                                if len(v31) > 0 {
                                    v31[0] = fi30
                                }
                            }
                        } else {
                            v4 += (v1 - (v4 << ((4294967295) & 63)))
                            var v61 int8 = v1
                            v31 = append(v31, -2147483648)
                            var v62 uint8 = 178
                            var v63 uint8 = v62
                            var v64 int64 = -1
                            fmt.Println(v24, v24, v61)
                            _ = v61
                            _ = v62
                            _ = v63
                            _ = v64
                        }
                        switch ("a" + (("a" + "foo bar") + (("a" + "a") + ""))) {
                        case "a":
                            v4 = (v3 * -44)
                        default:
                            var v65 int64 = 42
                            _ = v65
                        }
                    }
                }
            default:
                var v66 uint32 = (v2 << ((971) & 63))
                _ = v66
            }
            v3 += -59
            _ = fi32
        }
        _ = fi30
        _ = v31
    }
    var v67 int64 = 9223372036854775807
    fmt.Println(v67, v3, v4)
    for fi68 := 0; fi68 < 14; fi68++ {
        if false {
            v4 += (v1 % ((100) | 1))
            for fi69 := 0; fi69 < 9; fi69++ {
                v24 -= (v2 << ((409) & 63))
                fmt.Println(v3)
                var v70 bool = true
                _ = fi69
                _ = v70
            }
            var v71 uint32 = v2
            _ = v71
        } else {
            v1 += int8(v24)
            v24 += 0
            for fi72 := 0; fi72 < 16; fi72++ {
                v4 += (v3 & v3)
                v73 := GenS10{F0: v67}
                _ = fi72
                _ = v73
            }
        }
        _ = fi68
    }
    v4 = (v1 / ((100) | 1))
    if false {
        switch ("hello" + (("\n\t" + "foo bar") + ("a" + "hello"))) {
        case "":
            v74 := GenS11{F0: false, F1: 127}
            _ = v74
        case "foo bar":
            var v75 int64 = v67
            _ = v75
        case "\n\t":
            for fi76 := 0; fi76 < 18; fi76++ {
                var v77 uint32 = v2
                _ = fi76
                _ = v77
            }
        default:
            var v78 uint32 = 0
            _ = v78
        }
    }
    for fi79 := 0; fi79 < 2; fi79++ {
        switch ("\n\t" + "") {
        case "":
            v80 := []int{}
            _ = v80
        case "hello":
            for fi81 := 0; fi81 < 17; fi81++ {
                v67 -= 537
                fmt.Println(v67)
                _ = fi81
            }
        }
        _ = fi79
    }
    if true {
        v82 := []int{-415}
        _ = v82
    }
    v3 += (v1 >> ((uint(v2)) & 63))
    if (42 <= 32767) {
        v2 += (v24 * v24)
    }
    v83 := GenS12{F0: 4294967295}
    var v84 int8 = (^(v3))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v67:", v67)
    fmt.Println("v84:", v84)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v67
    _ = v83
    _ = v84
}
