package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32; F2 int32 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int64; F1 int32 }
type GenS4 struct { F0 int64; F1 int8; F2 uint16 }
type GenS5 struct { F0 uint8; F1 int16 }
type GenS6 struct { F0 int; F1 int8; F2 uint32 }
type GenS7 struct { F0 uint; F1 int }
type GenS8 struct { F0 string; F1 uint64 }
func genFn0(p0 string, p1 int32) uint {
    var v5 uint16 = 0
    v6 := []int{-9223372036854775808, 128, -9223372036854775808, 4294967295}
    fmt.Println(v5, v5)
    v7 := GenS1{F0: 255, F1: 2147483647, F2: (p1 + int32(v5))}
    var v8 uint16 = 216
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return 2
}

func main() {
    var v1 uint16 = 65535
    v2 := make(map[string]int)
    var v3 uint = 4294967295
    v4 := make(map[int]int)
    if ((true && true) && true) {
        v9 := GenS2{F0: 32767}
        v10 := GenS3{F0: 0, F1: 2147483647}
        var v11 int8 = 0
        _ = v10
        _ = v11
        _ = v9
    } else {
        delete(v2, "\n\t")
        var v12 uint64 = 4294967295
        if ((v3 >> ((630) & 63)) < (^(v3))) {
            v2["hello"] = 0
            if (!false) {
                var v13 uint = genFn0(("hello" + ("a" + "\n\t")), 770)
                for fi14 := 0; fi14 < 5; fi14++ {
                    v13 += v3
                    for fi15 := 0; fi15 < 11; fi15++ {
                        v16 := GenS2{F0: 516}
                        _ = fi15
                        _ = v16
                    }
                    _ = fi14
                }
                var v17 int8 = 77
                _ = v13
                _ = v17
            } else {
                for fi18 := 0; fi18 < 9; fi18++ {
                    v19 := []int{(fi18 ^ (fi18 & fi18)), (fi18 ^ 114), 9223372036854775807, (fi18 + (fi18 * fi18))}
                    v20 := GenS1{F0: 143, F1: 2147483648, F2: -1}
                    if false {
                        var v21 int8 = -68
                        v12 += v12
                        _ = v21
                    } else {
                        var v22 int8 = 42
                        _ = v22
                    }
                    _ = fi18
                    _ = v19
                    _ = v20
                }
            }
        }
        _ = v12
    }
    fmt.Println(v3, v1)
    switch -2147483648 {
    case 607:
        v1 += v1
    case -1:
        v3 = (v3 ^ v3)
    default:
        var v23 bool = (false || false)
        _ = v23
    }
    for fi24 := 0; fi24 < 9; fi24++ {
        _ = len(v2)
        _ = fi24
    }
    v3 -= v3
    switch 9223372036854775807 {
    case -2147483648:
        var v25 uint = (v3 >> (((^(v3))) & 63))
        _ = v25
    case 2147483647:
        for fi26 := 0; fi26 < 17; fi26++ {
            if ((!true) && ("a" != (("a" + ("a" + ("\n\t" + ("\n\t" + ("" + ""))))) + "a"))) {
                v2["hello"] = (fi26 * (fi26 >> ((v3) & 63)))
                v27 := GenS2{F0: 150}
                switch "foo bar" {
                case "\n\t":
                    for fi28 := 0; fi28 < 5; fi28++ {
                        var v29 uint32 = 1
                        var v30 int8 = -85
                        _ = fi28
                        _ = v29
                        _ = v30
                    }
                case "hello":
                    for fi31 := 0; fi31 < 11; fi31++ {
                        v1 = (v1 ^ (v1 + v1))
                        v1 = v1
                        _ = fi31
                    }
                case "foo bar":
                    v3 = v3
                }
                _ = v27
            }
            v3 = v3
            _ = fi26
        }
    default:
        var v32 uint = v3
        _ = v32
    }
    if false {
        v3 -= v3
        v3 -= v3
        var v33 int8 = 1
        _ = v33
    }
    v34 := []int{349, 1}
    if (!(-6 < 0)) {
        switch 606 {
        case 32767:
            v3 -= v3
        default:
            switch -540 {
            case 1:
                switch 304 {
                case 32768:
                    var v35 uint = (v3 | (v3 * uint(v1)))
                    _ = v35
                default:
                    v34 = append(v34, 1)
                }
            case 0:
                fmt.Println(v3, v1)
            case 42:
                v1 -= v1
            }
        }
        var v36 string = (("foo bar" + "foo bar") + "a")
        switch (v36 + v36) {
        case "a":
            if (!true) {
                var v37 int32 = 2147483647
                v2[(v36 + "a")] = 9223372036854775807
                var v38 int8 = -1
                _ = v37
                _ = v38
            }
        case "":
            v1 = 42
        case "foo bar":
            switch v36 {
            case "a":
                if (!(1 != 2147483647)) {
                    if ((-1 != -346) && (!(v1 != (v1 | (v1 / ((v1) | 1)))))) {
                        for fi39 := 0; fi39 < 9; fi39++ {
                            fmt.Println(fi39, fi39, v1)
                            _ = fi39
                        }
                    }
                } else {
                    v1 = (v1 << ((4294967295) & 63))
                    v40 := GenS4{F0: 610, F1: 101, F2: 255}
                    _ = v40
                }
            }
        default:
            v36 = "foo bar"
        }
        _ = v36
    } else {
        var v41 int64 = 1
        for fi42 := 0; fi42 < 10; fi42++ {
            v43 := GenS2{F0: 572}
            var v44 uint16 = (v1 | v1)
            switch (fi42 & fi42) {
            case -1:
                v45 := GenS1{F0: 1, F1: 65535, F2: 42}
                _ = v45
            default:
                var v46 float64 = 3.14159
                _ = v46
            }
            _ = fi42
            _ = v43
            _ = v44
        }
        v41 = v41
        _ = v41
    }
    if false {
        if ((((127 >= 100) && false) || (!(true && (false || false)))) || true) {
            v1 -= v1
            v1 = v1
        }
        var v47 int8 = 113
        _ = v47
    }
    for fi48 := 0; fi48 < 10; fi48++ {
        switch "hello" {
        case "foo bar":
            for fi49 := 0; fi49 < 9; fi49++ {
                v1 = v1
                _ = fi49
            }
        }
        if len(v34) > 0 {
            v34[0] = fi48
        }
        for fi50 := 0; fi50 < 6; fi50++ {
            if true {
                switch fi50 {
                case 32768:
                    delete(v4, fi50)
                    if false {
                        switch fi48 {
                        case 128:
                            v51 := GenS4{F0: 9223372036854775807, F1: -120, F2: 193}
                            _ = v51
                        case 9223372036854775807:
                            var v52 string = ("hello" + "\n\t")
                            _ = v52
                        }
                        v1 = (v1 / ((v1) | 1))
                        var v53 float32 = 3.14159
                        switch fi50 {
                        case 9223372036854775807:
                            if (float64(v53) <= -2.25) {
                                v34 = append(v34, fi48)
                                var v54 int = (^(fi50))
                                _ = v54
                            }
                        default:
                            v53 = (1.5 + (float32((-2.25 - float64(v53))) - -2.25))
                        }
                        _ = v53
                    } else {
                        v3 = (v3 % ((v3) | 1))
                        fmt.Println(fi50)
                        var v55 float32 = 1.5
                        _ = v55
                    }
                    if false {
                        delete(v4, (fi48 & (-(fi50))))
                        if true {
                            fmt.Println(v1, v1)
                        }
                        v1 += (^(v1))
                    } else {
                        v56 := GenS5{F0: 2, F1: -784}
                        _ = v56
                    }
                default:
                    fmt.Println(fi50)
                    var v57 int8 = -106
                    _ = len(v2)
                    _ = v57
                }
                _ = len(v4)
            } else {
                switch ((("" + "foo bar") + "") + "hello") {
                case "foo bar":
                    var v58 string = "\n\t"
                    v59 := GenS6{F0: fi48, F1: 0, F2: 132}
                    v3 += genFn0(v58, -2147483648)
                    v1 -= (^(v1))
                    _ = v58
                    _ = v59
                case "":
                    if (true && (!true)) {
                        v34 = append(v34, fi48)
                        v3 += (v3 / (((v3 ^ (v3 << (((v3 >> ((v3) & 63))) & 63)))) | 1))
                        if (95 < 2147483647) {
                            v1 = v1
                            if false {
                                v60 := GenS7{F0: v3, F1: 971}
                                v1 -= v1
                                var v61 string = "\n\t"
                                var v62 int16 = -936
                                _ = v60
                                _ = v61
                                _ = v62
                            } else {
                                switch int(v1) {
                                case -873:
                                    v3 += v3
                                case 32767:
                                    var v63 uint16 = (v1 << (((v3 >> (((v3 >> ((v3) & 63))) & 63))) & 63))
                                    if true {
                                        if true {
                                            if true {
                                                v3 = 255
                                                if (!((-2147483648 > (-(fi48))) || true)) {
                                                    var v64 float64 = (0.0 / 3.14159)
                                                    if false {
                                                        if (-126 >= 2147483647) {
                                                            if (18446744073709551615 <= 47) {
                                                                var v65 int = (fi50 >> (((v3 % ((genFn0("hello", -928)) | 1))) & 63))
                                                                if true {
                                                                    if len(v34) >= 2 {
                                                                        v34 = v34[1:len(v34)]
                                                                    }
                                                                }
                                                                _ = v65
                                                            } else {
                                                                if false {
                                                                    v66 := GenS7{F0: v3, F1: fi50}
                                                                    var v67 uint32 = 2
                                                                    _ = v66
                                                                    _ = v67
                                                                } else {
                                                                    v64 = v64
                                                                    var v68 int32 = 365
                                                                    v69 := GenS2{F0: 32767}
                                                                    _ = v68
                                                                    _ = v69
                                                                }
                                                                fmt.Println(fi50, v64)
                                                            }
                                                            switch "" {
                                                            case "hello":
                                                                var v70 int16 = 1
                                                                _ = v70
                                                            default:
                                                                switch (fi50 * 748) {
                                                                case 480:
                                                                    if len(v34) > 0 {
                                                                        v34[0] = 4294967295
                                                                    }
                                                                default:
                                                                    if (!false) {
                                                                        v71 := GenS7{F0: (v3 << ((475) & 63)), F1: fi50}
                                                                        v34 = append(v34, fi48)
                                                                        _ = v71
                                                                    } else {
                                                                        v63 += (v1 * 42)
                                                                    }
                                                                }
                                                            }
                                                            fmt.Println(fi50, v3)
                                                        } else {
                                                            var v72 uint64 = 948
                                                            switch "\n\t" {
                                                            case "a":
                                                                var v73 uint16 = v63
                                                                _ = v73
                                                            default:
                                                                fmt.Println(fi48)
                                                            }
                                                            _ = v72
                                                        }
                                                        if len(v34) >= 2 {
                                                            v34 = v34[1:len(v34)]
                                                        }
                                                    }
                                                    fmt.Println(fi48)
                                                    _ = v64
                                                }
                                            }
                                            v3 = (^(v3))
                                            fmt.Println(fi48, v3, fi48)
                                        }
                                        v1 += v63
                                    } else {
                                        if len(v34) >= 2 {
                                            v34 = v34[1:len(v34)]
                                        }
                                        var v74 int64 = int64(v3)
                                        v75 := GenS8{F0: "", F1: 409}
                                        _ = v74
                                        _ = v75
                                    }
                                    v3 = 2
                                    _ = v63
                                case -42:
                                    var v76 uint16 = v1
                                    _ = v76
                                default:
                                    fmt.Println(fi48, fi50)
                                }
                                v1 += v1
                                v77 := v4[fi48]
                                _ = v77
                            }
                            var v78 uint8 = 1
                            _ = v78
                        }
                    } else {
                        delete(v2, "\n\t")
                        var v79 int64 = 9223372036854775807
                        _ = v79
                    }
                }
            }
            var v80 float64 = (float64(3.14159) / 0.1)
            v3 -= (^(v3))
            _ = fi50
            _ = v80
        }
        _ = fi48
    }
    var v81 uint16 = (v1 * v1)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v81:", v81)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v81
}
