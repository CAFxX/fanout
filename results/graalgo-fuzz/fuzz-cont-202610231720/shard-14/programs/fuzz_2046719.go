package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32 }
type GenS2 struct { F0 uint32; F1 uint16 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int8; F1 int8 }
type GenS5 struct { F0 uint32; F1 string; F2 bool }
type GenS6 struct { F0 string; F1 uint16 }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 int; F1 uint8 }
type GenS9 struct { F0 int64; F1 uint32; F2 int }

func main() {
    var v1 int32 = 196
    var v2 int64 = -304
    v3 := make(map[string]uint32)
    _ = len(v3)
    for fi4 := 0; fi4 < 3; fi4++ {
        switch (fi4 >> ((18446744073709551615) & 63)) {
        case 327:
            v5 := GenS1{F0: (fi4 % ((-9223372036854775808) | 1)), F1: 1}
            _ = v5
        case 32767:
            var v6 uint = 65535
            _ = v6
        }
        switch "hello" {
        case "":
            switch "" {
            case "foo bar":
                if true {
                    if (fi4 != fi4) {
                        v7 := GenS2{F0: 65535, F1: 32767}
                        _ = v7
                    }
                    var v8 int = (-(fi4))
                    v9 := []int{32767, 4294967295, fi4}
                    _ = v8
                    _ = v9
                }
            case "\n\t":
                v1 = v1
            case "a":
                v1 = 900
            default:
                v2 += (^(v2))
            }
        case "hello":
            if false {
                switch fi4 {
                case 1:
                    for fi10 := 0; fi10 < 7; fi10++ {
                        switch (("foo bar" + "\n\t") + ("hello" + (("\n\t" + "\n\t") + ("foo bar" + ("a" + ("\n\t" + "")))))) {
                        case "\n\t":
                            switch (fi4 >> ((18446744073709551615) & 63)) {
                            case -1:
                                v2 -= v2
                            case 562:
                                var v11 bool = true
                                _ = v11
                            case 0:
                                fmt.Println(v1, v1, fi10)
                            default:
                                switch ("a" + (("a" + "") + "a")) {
                                case "":
                                    v1 = v1
                                    v1 = -2147483648
                                    v1 = (v1 >> ((0) & 63))
                                case "\n\t":
                                    fmt.Println(v2)
                                case "hello":
                                    v2 = v2
                                }
                                v12 := GenS1{F0: fi4, F1: 389}
                                switch "a" {
                                case "a":
                                    if true {
                                        v13 := GenS3{F0: 2}
                                        v2 -= 1
                                        v2 = v2
                                        _ = v13
                                    } else {
                                        if false {
                                            var v14 int32 = 42
                                            var v15 uint64 = 0
                                            var v16 int32 = v14
                                            _ = v14
                                            _ = v15
                                            _ = v16
                                        }
                                        switch 371 {
                                        case -284:
                                            if (((fi4 + fi10) < 1) || (!false)) {
                                                var v17 string = "a"
                                                v17 = v17
                                                v18 := GenS1{F0: fi10, F1: 65535}
                                                _ = v17
                                                _ = v18
                                            }
                                        case 9223372036854775807:
                                            var v19 uint8 = 127
                                            _ = v19
                                        default:
                                            v2 -= int64(fi10)
                                        }
                                        v20 := GenS2{F0: 4294967295, F1: 996}
                                        _ = v20
                                    }
                                case "\n\t":
                                    var v21 int = (-(fi10))
                                    _ = v21
                                default:
                                    v22 := []int{-9223372036854775808, int(v1), 701}
                                    _ = v22
                                }
                                _ = v12
                            }
                        }
                        _ = fi10
                    }
                default:
                    v2 = (v2 * v2)
                }
                fmt.Println(v2, fi4)
            } else {
                v23 := []int{}
                v23 = append(v23, fi4)
                _ = v23
            }
        case "\n\t":
            v24 := []int{fi4}
            _ = v24
        default:
            v25 := GenS4{F0: -128, F1: -128}
            _ = v25
        }
        _ = fi4
    }
    var v26 float32 = 3.14159
    v3[("a" + ("" + ("hello" + "a")))] = 65535
    switch 32767 {
    case -111:
        switch 32767 {
        case 128:
            v27 := GenS5{F0: 999, F1: "\n\t", F2: true}
            _ = v27
        case 9223372036854775807:
            var v28 int32 = -1
            _ = v28
        case -534:
            switch 985 {
            case 2147483647:
                switch 427 {
                case 128:
                    v29 := GenS3{F0: 0}
                    _ = v29
                }
            }
        default:
            if (false || true) {
                _ = len(v3)
                fmt.Println(v26)
            }
        }
    }
    for fi30 := 0; fi30 < 20; fi30++ {
        v2 += v2
        v2 += (v2 - (-(v2)))
        _ = fi30
    }
    if (18446744073709551615 == 209) {
        v26 = float32((0.1 + -2.25))
    } else {
        v26 = (-2.25 * v26)
        var v31 uint8 = 42
        _ = v31
    }
    switch 0 {
    case -9223372036854775808:
        if false {
            for fi32 := 0; fi32 < 5; fi32++ {
                if true {
                    v33 := GenS5{F0: 4294967295, F1: "", F2: ((("hello" + "") + ((("\n\t" + ("a" + ("\n\t" + ("a" + "a")))) + "foo bar") + ("\n\t" + "\n\t"))) == "\n\t")}
                    var v34 int32 = v1
                    v35 := []int{-1, 32768, int(v2), fi32}
                    _ = v33
                    _ = v34
                    _ = v35
                }
                fmt.Println(v2)
                _ = fi32
            }
            if ((-(v1)) < 1) {
                for fi36 := 0; fi36 < 15; fi36++ {
                    _ = len(v3)
                    fmt.Println(v2, v26)
                    v37 := GenS6{F0: ("hello" + "hello"), F1: 65535}
                    _ = fi36
                    _ = v37
                }
            } else {
                if true {
                    var v38 uint64 = 0
                    v2 = v2
                    v1 -= (^(v1))
                    _ = v38
                } else {
                    v1 -= v1
                }
                v3[("a" + "a")] = 1
            }
            for fi39 := 0; fi39 < 19; fi39++ {
                v40 := GenS1{F0: (fi39 | (-(fi39))), F1: 359}
                if (fi39 != fi39) {
                    fmt.Println(v26)
                    v1 += -1
                    if false {
                        switch ("" + "foo bar") {
                        case "hello":
                            switch 1 {
                            case -2147483648:
                                for fi41 := 0; fi41 < 7; fi41++ {
                                    var v42 bool = true
                                    var v43 uint8 = 165
                                    _ = fi41
                                    _ = v42
                                    _ = v43
                                }
                            case -897:
                                v44 := GenS4{F0: 0, F1: 127}
                                _ = v44
                            }
                        case "foo bar":
                            if ((!true) && true) {
                                switch (fi39 % (((fi39 ^ (fi39 ^ 42))) | 1)) {
                                case 0:
                                    for fi45 := 0; fi45 < 1; fi45++ {
                                        var v46 uint = 938
                                        v3["foo bar"] = 42
                                        _ = fi45
                                        _ = v46
                                    }
                                }
                                for fi47 := 0; fi47 < 8; fi47++ {
                                    if ((!((9223372036854775807 > (fi39 ^ fi47)) || false)) || true) {
                                        var v48 int16 = 900
                                        var v49 bool = false
                                        _ = v48
                                        _ = v49
                                    } else {
                                        var v50 int32 = v1
                                        var v51 uint32 = 2
                                        v52 := v3["hello"]
                                        _ = v50
                                        _ = v51
                                        _ = v52
                                    }
                                    _ = fi47
                                }
                            } else {
                                if ((767 > 2) && true) {
                                    var v53 string = "hello"
                                    _ = v53
                                }
                            }
                        case "a":
                            v54 := GenS3{F0: 4294967295}
                            _ = v54
                        default:
                            v55 := []int{(fi39 << ((0) & 63)), (fi39 % ((128) | 1)), fi39}
                            _ = v55
                        }
                        var v56 float32 = v26
                        _ = v56
                    } else {
                        v57 := GenS6{F0: ("\n\t" + "a"), F1: 0}
                        fmt.Println(v2, v1)
                        _ = v57
                    }
                }
                v58 := GenS6{F0: "hello", F1: 0}
                _ = fi39
                _ = v40
                _ = v58
            }
        } else {
            v26 = v26
        }
    case -746:
        switch 954 {
        case -9223372036854775808:
            var v59 int16 = -1
            _ = v59
        case 9223372036854775807:
            v26 -= (v26 + v26)
        case 127:
            for fi60 := 0; fi60 < 7; fi60++ {
                var v61 uint32 = 2147483647
                switch fi60 {
                case 42:
                    switch 2147483647 {
                    case 9223372036854775807:
                        fmt.Println(v2, v26, v2)
                    case 128:
                        fmt.Println(fi60, v2, fi60)
                    case 1:
                        v61 -= (v61 * (v61 | (^(v61))))
                    }
                case -1:
                    v62 := v3["hello"]
                    _ = v62
                default:
                    for fi63 := 0; fi63 < 12; fi63++ {
                        if false {
                            v64 := []int{(fi63 % ((fi60) | 1)), (fi63 - fi63), -2147483648, fi60}
                            v65 := GenS5{F0: 318, F1: "\n\t", F2: true}
                            var v66 int32 = (v1 ^ int32(fi63))
                            v66 -= (v1 & (v66 << ((65535) & 63)))
                            _ = v64
                            _ = v65
                            _ = v66
                        } else {
                            if false {
                                v67 := GenS2{F0: v61, F1: 255}
                                _ = v67
                            }
                            var v68 uint64 = 167
                            _ = v68
                        }
                        var v69 uint8 = 42
                        v26 += ((1e10 / (float32(-2.25) / v26)) / v26)
                        _ = fi63
                        _ = v69
                    }
                }
                _ = fi60
                _ = v61
            }
        default:
            var v70 int32 = 2147483647
            _ = v70
        }
    case 777:
        fmt.Println(v1)
    default:
        fmt.Println(v26, v1, v1)
    }
    v1 = v1
    switch 9223372036854775807 {
    case -230:
        if (v2 >= 1) {
            if ((false || true) && false) {
                for fi71 := 0; fi71 < 9; fi71++ {
                    v1 = int32(v2)
                    if ((false || true) && false) {
                        if (false && true) {
                            fmt.Println(v2, v26)
                        }
                    } else {
                        v2 = v2
                    }
                    var v72 int = 42
                    _ = fi71
                    _ = v72
                }
            }
        }
    case -512:
        for fi73 := 0; fi73 < 18; fi73++ {
            v74 := GenS3{F0: 2}
            for fi75 := 0; fi75 < 3; fi75++ {
                v26 = (v26 + -2.25)
                if true {
                    if ((128 != 255) || (589 >= 331)) {
                        fmt.Println(fi75, v26, fi75)
                    }
                }
                _ = fi75
            }
            v1 = int32(fi73)
            _ = fi73
            _ = v74
        }
    case -491:
        switch -1 {
        case -754:
            var v76 int32 = v1
            _ = v76
        case 1:
            for fi77 := 0; fi77 < 20; fi77++ {
                switch "a" {
                case "a":
                    for fi78 := 0; fi78 < 14; fi78++ {
                        if true {
                            v26 += v26
                            v79 := GenS7{F0: int8(v2)}
                            var v80 int16 = -1
                            var v81 uint32 = 42
                            v2 = 42
                            _ = v79
                            _ = v80
                            _ = v81
                        }
                        _ = fi78
                    }
                default:
                    fmt.Println(fi77, fi77)
                }
                var v82 float64 = 1.5
                _ = fi77
                _ = v82
            }
        }
    }
    v83 := GenS8{F0: 22, F1: 42}
    for fi84 := 0; fi84 < 10; fi84++ {
        var v85 uint16 = 333
        _ = fi84
        _ = v85
    }
    v86 := GenS6{F0: "hello", F1: 42}
    v1 -= v1
    v87 := GenS9{F0: v2, F1: 1, F2: 32767}
    var v88 uint64 = 2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v88:", v88)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v83
    _ = v86
    _ = v87
    _ = v88
}
