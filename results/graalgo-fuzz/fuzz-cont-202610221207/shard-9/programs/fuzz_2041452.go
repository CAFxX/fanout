package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 bool; F1 uint32 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint64; F1 string; F2 uint16 }
type GenS5 struct { F0 int8; F1 uint64 }
type GenS6 struct { F0 int64; F1 int16 }
type GenS7 struct { F0 int32; F1 int8 }
type GenS8 struct { F0 int16 }

func main() {
    v1 := make(map[string]int)
    var v2 int8 = 100
    var v3 int64 = 0
    var v4 int = 42
    var v5 bool = false
    var v6 uint8 = 127
    var v7 float32 = -2.25
    switch "foo bar" {
    case "a":
        switch "" {
        case "":
            v8 := GenS1{F0: 571}
            _ = v8
        case "\n\t":
            v9 := GenS2{F0: (4294967295 == 42), F1: 2147483648}
            _ = v9
        }
    case "hello":
        switch ("a" + ("hello" + "hello")) {
        case "hello":
            var v10 string = "foo bar"
            _ = v10
        case "\n\t":
            switch "a" {
            case "hello":
                if (42 > -32768) {
                    v11 := []int{v4}
                    for fi12 := 0; fi12 < 8; fi12++ {
                        var v13 uint = 874
                        _ = fi12
                        _ = v13
                    }
                    v14 := GenS3{F0: 2}
                    _ = v11
                    _ = v14
                } else {
                    v2 = 0
                    if false {
                        var v15 float64 = 0.1
                        if ((^(v3)) <= (v3 << ((595) & 63))) {
                            if false {
                                v3 -= v3
                                switch "\n\t" {
                                case "a":
                                    v7 += 0.1
                                case "\n\t":
                                    v16 := GenS3{F0: 0}
                                    _ = v16
                                case "":
                                    switch (v4 << ((389) & 63)) {
                                    case 1:
                                        v17 := []int{(v4 % ((v4) | 1)), v4}
                                        _ = v17
                                    case -9223372036854775808:
                                        for fi18 := 0; fi18 < 2; fi18++ {
                                            for fi19 := 0; fi19 < 15; fi19++ {
                                                switch fi18 {
                                                case -9223372036854775808:
                                                    v3 = v3
                                                case 9223372036854775807:
                                                    fmt.Println(fi19, fi19)
                                                    var v20 uint16 = 65535
                                                    v21 := []int{(fi18 >> ((18446744073709551615) & 63))}
                                                    _ = v20
                                                    _ = v21
                                                default:
                                                    v6 -= (v6 >> ((4294967295) & 63))
                                                    var v22 uint8 = (^(v6))
                                                    _ = v22
                                                }
                                                switch "\n\t" {
                                                case "foo bar":
                                                    var v23 bool = (float32((3.14159 * v15)) != (0.1 - (v7 - v7)))
                                                    _ = v23
                                                case "a":
                                                    v6 += 177
                                                }
                                                v3 -= (v3 % ((v3) | 1))
                                                _ = fi19
                                            }
                                            _ = fi18
                                        }
                                    default:
                                        v15 = (0.1 - (0.1 * (3.14159 / v15)))
                                    }
                                }
                            } else {
                                switch v4 {
                                case 4294967295:
                                    v24 := v1["\n\t"]
                                    _ = v24
                                case 42:
                                    v25 := v1["foo bar"]
                                    _ = v25
                                }
                                for fi26 := 0; fi26 < 6; fi26++ {
                                    v15 -= float64(1e10)
                                    _ = fi26
                                }
                                var v27 uint32 = 2147483648
                                _ = v27
                            }
                        } else {
                            v3 = 1
                        }
                        _ = v15
                    } else {
                        if ("hello" <= "foo bar") {
                            switch ("foo bar" + "") {
                            case "a":
                                for fi28 := 0; fi28 < 13; fi28++ {
                                    if ((fi28 >> ((65535) & 63)) > v4) {
                                        v2 = (v2 - 0)
                                        var v29 uint64 = 302
                                        _ = v29
                                    } else {
                                        v30 := GenS1{F0: 1}
                                        if true {
                                            for fi31 := 0; fi31 < 19; fi31++ {
                                                v6 = (^(v6))
                                                _ = fi31
                                            }
                                        } else {
                                            v32 := []int{1, (v4 / (((fi28 * fi28)) | 1)), (-(v4)), (v4 / ((v4) | 1))}
                                            v3 = v3
                                            v3 += v3
                                            _ = v32
                                        }
                                        for fi33 := 0; fi33 < 16; fi33++ {
                                            v34 := []int{(fi28 + v4), 42}
                                            if false {
                                                var v35 uint64 = 4294967295
                                                delete(v1, ((("\n\t" + "\n\t") + "\n\t") + "hello"))
                                                _ = v35
                                            }
                                            switch 9223372036854775807 {
                                            case -935:
                                                if (!(!v5)) {
                                                    v36 := GenS4{F0: 42, F1: ((("a" + (("" + "\n\t") + "a")) + ("" + "a")) + "foo bar"), F2: 912}
                                                    v1["a"] = (fi28 * (fi33 >> ((65535) & 63)))
                                                    var v37 int32 = -37
                                                    _ = v36
                                                    _ = v37
                                                } else {
                                                    switch "" {
                                                    case "foo bar":
                                                        fmt.Println(v4)
                                                    case "a":
                                                        var v38 uint64 = 1
                                                        _ = v38
                                                    default:
                                                        if v5 {
                                                            v7 += v7
                                                            v2 = v2
                                                        }
                                                    }
                                                    v2 = v2
                                                }
                                            }
                                            _ = fi33
                                            _ = v34
                                        }
                                        _ = v30
                                    }
                                    v1[""] = (-(v4))
                                    _ = fi28
                                }
                            default:
                                v39 := GenS1{F0: 158}
                                _ = v39
                            }
                            v2 = (v2 & (^(v2)))
                            for fi40 := 0; fi40 < 12; fi40++ {
                                v41 := GenS5{F0: (-(v2)), F1: 468}
                                _ = fi40
                                _ = v41
                            }
                        }
                    }
                    if false {
                        v7 -= 1.5
                        v5 = (v5 || (float64((0.1 - (v7 * v7))) == 0.1))
                        if (!(637 != 4294967295)) {
                            v2 = v2
                            for fi42 := 0; fi42 < 4; fi42++ {
                                var v43 float64 = 0.1
                                var v44 float32 = (v7 / v7)
                                _ = fi42
                                _ = v43
                                _ = v44
                            }
                        }
                    }
                }
            default:
                if (1.5 < ((-2.25 + 0.1) - 1e10)) {
                    fmt.Println(v6, v3, v3)
                    if ((float32((1e10 * (3.14159 - (3.14159 / 1.0)))) - (v7 / 0.1)) > 1.5) {
                        fmt.Println(v2, v4, v3)
                        var v45 bool = (v6 != (^(v6)))
                        var v46 string = "a"
                        _ = v45
                        _ = v46
                    }
                } else {
                    switch "foo bar" {
                    case "hello":
                        v47 := GenS2{F0: (93 == 2), F1: 613}
                        _ = v47
                    }
                    var v48 uint32 = 189
                    var v49 float32 = (v7 - (v7 - v7))
                    _ = v48
                    _ = v49
                }
            }
        }
    case "":
        if (!(("a" + "\n\t") <= "foo bar")) {
            for fi50 := 0; fi50 < 19; fi50++ {
                var v51 string = "a"
                if v5 {
                    v7 = v7
                    switch v4 {
                    case 4294967295:
                        fmt.Println(v51, v2, v2)
                    case -76:
                        var v52 int = 0
                        _ = v52
                    case 128:
                        for fi53 := 0; fi53 < 17; fi53++ {
                            if v5 {
                                v4 += v4
                                var v54 int64 = v3
                                var v55 int64 = 42
                                _ = v54
                                _ = v55
                            }
                            _ = fi53
                        }
                    default:
                        v5 = v5
                    }
                } else {
                    if v5 {
                        v2 = (v2 / ((0) | 1))
                        switch -2147483648 {
                        case -978:
                            if v5 {
                                var v56 uint16 = 889
                                _ = v56
                            } else {
                                v4 += 0
                                var v57 uint8 = 0
                                _ = v57
                            }
                        case -168:
                            switch v4 {
                            case 897:
                                v7 = v7
                            }
                        case -1:
                            for fi58 := 0; fi58 < 1; fi58++ {
                                v59 := GenS1{F0: 826}
                                switch v51 {
                                case "a":
                                    v6 = 0
                                case "":
                                    v2 = (v2 - -1)
                                case "foo bar":
                                    v60 := []int{}
                                    _ = v60
                                }
                                var v61 bool = (1 < uint64(v6))
                                _ = fi58
                                _ = v59
                                _ = v61
                            }
                        default:
                            var v62 int64 = (v3 / (((v3 % ((v3) | 1))) | 1))
                            _ = v62
                        }
                    }
                    v63 := GenS6{F0: (v3 + int64(v6)), F1: 42}
                    v64 := v1["\n\t"]
                    _ = v63
                    _ = v64
                }
                _ = fi50
                _ = v51
            }
            var v65 int = (v4 + v4)
            var v66 int8 = int8(v65)
            _ = v65
            _ = v66
        }
    default:
        var v67 uint32 = 4294967295
        _ = v67
    }
    var v68 float32 = (v7 * v7)
    fmt.Println(v4, v6, v2)
    var v69 float64 = 1e10
    if v5 {
        fmt.Println(v3, v4, v5)
        switch "foo bar" {
        case "hello":
            if (1 > 42) {
                v6 -= 42
                for fi70 := 0; fi70 < 7; fi70++ {
                    for fi71 := 0; fi71 < 20; fi71++ {
                        delete(v1, ("a" + "\n\t"))
                        _ = fi71
                    }
                    _ = fi70
                }
                v3 = v3
            }
        }
        var v72 uint16 = 478
        _ = v72
    }
    for fi73 := 0; fi73 < 4; fi73++ {
        if (("" + "hello") < ("\n\t" + "a")) {
            if ((^(v6)) != 249) {
                v4 -= 1
                for fi74 := 0; fi74 < 7; fi74++ {
                    fmt.Println(v68)
                    v6 += 128
                    _ = fi74
                }
                if (v6 < 127) {
                    v6 -= v6
                    var v75 float32 = v7
                    for fi76 := 0; fi76 < 13; fi76++ {
                        v2 = (v2 / (((v2 & (v2 ^ (v2 * v2)))) | 1))
                        v77 := GenS5{F0: (v2 % (((v2 << ((951) & 63))) | 1)), F1: 2}
                        _ = fi76
                        _ = v77
                    }
                    _ = v75
                } else {
                    switch ("foo bar" + "\n\t") {
                    case "a":
                        var v78 int32 = 0
                        _ = v78
                    case "foo bar":
                        if ((((3.14159 - (1e10 + float32(-2.25))) + -2.25) > 0.0) || true) {
                            var v79 uint = 65535
                            _ = v79
                        } else {
                            var v80 uint8 = 2
                            _ = v80
                        }
                    default:
                        v81 := GenS7{F0: -864, F1: (v2 ^ 100)}
                        _ = v81
                    }
                }
            }
            v4 -= -1
            v82 := GenS8{F0: 420}
            _ = v82
        } else {
            for fi83 := 0; fi83 < 11; fi83++ {
                var v84 int8 = (v2 ^ (v2 * (v2 * (v2 * (v2 ^ v2)))))
                _ = fi83
                _ = v84
            }
            fmt.Println(v7, v5)
        }
        _ = fi73
    }
    fmt.Println(v69)
    fmt.Println(v69, v7)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v68:", v68)
    fmt.Println("v69:", v69)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v68
    _ = v69
    _ = v7
}
