package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int; F2 uint }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint8; F1 int8; F2 int8 }
type GenS5 struct { F0 int; F1 int32; F2 int16 }
type GenS6 struct { F0 uint8; F1 bool }
type GenS7 struct { F0 uint32; F1 int8; F2 int16 }
type GenS8 struct { F0 int }
type GenS9 struct { F0 int16; F1 bool; F2 uint }
func genFn0(p0 string, p1 int, p2 int64) int32 {
    switch 4294967295 {
    case 42:
        switch p1 {
        case 0:
            v4 := GenS1{F0: (p2 << ((42) & 63)), F1: 0, F2: 977}
            _ = v4
        case -4:
            v5 := GenS2{F0: 0}
            _ = v5
        case 1:
            for fi6 := 0; fi6 < 17; fi6++ {
                p0 = "foo bar"
                _ = fi6
            }
        }
    case 131:
        for fi7 := 0; fi7 < 19; fi7++ {
            for fi8 := 0; fi8 < 20; fi8++ {
                v9 := GenS1{F0: (p2 ^ (-(p2))), F1: -946, F2: 42}
                var v10 uint16 = 897
                v11 := []int{222}
                _ = fi8
                _ = v10
                _ = v11
                _ = v9
            }
            _ = fi7
        }
    case -2147483648:
        var v12 int16 = -1
        _ = v12
    }
    if ("hello" > p0) {
        var v13 float64 = (((3.14159 + 1e10) + 0.0) * 1e10)
        var v14 uint = 4294967295
        var v15 uint32 = 556
        _ = v13
        _ = v14
        _ = v15
    }
    switch "" {
    case "":
        if (true && ((605 == 681) && true)) {
            v16 := []int{95, -1}
            _ = v16
        }
    case "hello":
        if true {
            p2 -= (-(p2))
            if (false && (((0 == (p2 - (p2 % (((-(p2))) | 1)))) || false) || true)) {
                p0 = (p0 + p0)
            }
        }
    case "a":
        if false {
            v17 := GenS2{F0: 1}
            if true {
                p1 = (^(p1))
                fmt.Println(p0, p0)
                for fi18 := 0; fi18 < 8; fi18++ {
                    p2 = (p2 | (^(p2)))
                    _ = fi18
                }
            } else {
                switch (p1 << ((1) & 63)) {
                case -9223372036854775808:
                    p0 = ("foo bar" + p0)
                case 4294967295:
                    for fi19 := 0; fi19 < 13; fi19++ {
                        for fi20 := 0; fi20 < 8; fi20++ {
                            p1 = fi20
                            p2 = 42
                            p2 = p2
                            _ = fi20
                        }
                        p1 = fi19
                        var v21 uint8 = 223
                        _ = fi19
                        _ = v21
                    }
                }
            }
            _ = v17
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return -315
}

func main() {
    v1 := []int{-254, 42}
    var v2 int32 = -803
    v3 := make(map[int]string)
    _ = len(v3)
    v22 := GenS1{F0: -684, F1: 32767, F2: 615}
    _ = len(v3)
    for fi23 := 0; fi23 < 5; fi23++ {
        v24 := GenS2{F0: 0}
        var v25 bool = false
        var v26 uint = 123
        _ = fi23
        _ = v24
        _ = v25
        _ = v26
    }
    if true {
        if len(v1) > 0 {
            v1[0] = -2147483648
        }
        v2 -= (v2 << ((413) & 63))
    }
    switch "hello" {
    case "hello":
        for fi27 := 0; fi27 < 6; fi27++ {
            for fi28 := 0; fi28 < 16; fi28++ {
                var v29 int8 = -76
                v29 -= v29
                if (1 > 319) {
                    v30 := GenS1{F0: -9223372036854775808, F1: (^(fi28)), F2: 13}
                    v2 = v2
                    var v31 uint16 = 32767
                    _ = v30
                    _ = v31
                }
                _ = fi28
                _ = v29
            }
            if false {
                for fi32 := 0; fi32 < 10; fi32++ {
                    var v33 int16 = 598
                    _ = fi32
                    _ = v33
                }
            }
            v34 := v3[fi27]
            _ = fi27
            _ = v34
        }
    case "foo bar":
        var v35 uint8 = 2
        _ = v35
    case "\n\t":
        var v36 float64 = (0.0 - -2.25)
        _ = v36
    default:
        switch "foo bar" {
        case "\n\t":
            v37 := GenS2{F0: -128}
            _ = v37
        case "":
            fmt.Println(v2)
        case "hello":
            if true {
                v3[323] = ""
                _ = len(v3)
                v2 -= (v2 >> ((4294967295) & 63))
            } else {
                for fi38 := 0; fi38 < 19; fi38++ {
                    v39 := GenS1{F0: 1, F1: (fi38 >> ((0) & 63)), F2: 466}
                    v40 := GenS3{F0: 65535}
                    _ = fi38
                    _ = v39
                    _ = v40
                }
                fmt.Println(v2)
                var v41 int = -656
                _ = v41
            }
        default:
            var v42 int16 = -580
            _ = v42
        }
    }
    fmt.Println(v2)
    var v43 uint32 = 2
    var v44 uint64 = 869
    fmt.Println(v2, v44, v2)
    var v45 bool = false
    if false {
        v43 += v43
        v2 += genFn0("hello", 32767, 9223372036854775807)
    } else {
        if v45 {
            if (v45 || v45) {
                switch "hello" {
                case "hello":
                    var v46 int = 127
                    _ = v46
                case "":
                    switch -9223372036854775808 {
                    case -265:
                        for fi47 := 0; fi47 < 6; fi47++ {
                            v48 := GenS4{F0: 205, F1: 1, F2: -111}
                            if (-78 != -84) {
                                v49 := v3[(fi47 << ((143) & 63))]
                                _ = v49
                            }
                            _ = fi47
                            _ = v48
                        }
                    case 32768:
                        fmt.Println(v43, v44)
                    case -9223372036854775808:
                        var v50 uint8 = 127
                        _ = v50
                    default:
                        var v51 string = ("" + "")
                        _ = v51
                    }
                case "foo bar":
                    if (!((v43 + v43) >= 0)) {
                        v3[-298] = ("" + "\n\t")
                        var v52 string = "a"
                        _ = v52
                    } else {
                        switch (("a" + "foo bar") + (("hello" + "a") + "\n\t")) {
                        case "\n\t":
                            for fi53 := 0; fi53 < 14; fi53++ {
                                v54 := GenS5{F0: (fi53 | (fi53 ^ fi53)), F1: (v2 * int32(v44)), F2: 1000}
                                v55 := v3[fi53]
                                v3[(fi53 | (fi53 >> ((0) & 63)))] = v55
                                _ = fi53
                                _ = v54
                                _ = v55
                            }
                        case "a":
                            if v45 {
                                v43 -= 2147483648
                                var v56 bool = v45
                                _ = v56
                            }
                        case "":
                            v45 = v45
                        }
                        _ = len(v1)
                        v1 = append(v1, 42)
                    }
                default:
                    for fi57 := 0; fi57 < 8; fi57++ {
                        var v58 float64 = 3.14159
                        _ = fi57
                        _ = v58
                    }
                }
                v3[804] = ("foo bar" + "a")
            }
        } else {
            v2 -= v2
            var v59 int32 = -2147483648
            _ = v59
        }
    }
    var v60 int8 = 1
    if v45 {
        if (127 > 127) {
            if (-18 == 42) {
                if (!true) {
                    var v61 int = 279
                    _ = v61
                } else {
                    v62 := GenS5{F0: 32767, F1: (v2 * (-(v2))), F2: 1000}
                    for fi63 := 0; fi63 < 4; fi63++ {
                        v64 := GenS4{F0: uint8(v43), F1: v60, F2: -1}
                        fmt.Println(v60, v43, v45)
                        var v65 uint64 = (v44 / ((v44) | 1))
                        _ = fi63
                        _ = v64
                        _ = v65
                    }
                    _ = v62
                }
                fmt.Println(v60)
            } else {
                fmt.Println(v60)
            }
            v45 = v45
        } else {
            v3[42] = (("" + "") + ("hello" + "hello"))
            if v45 {
                if (0 > 0) {
                    v1 = append(v1, 32767)
                } else {
                    var v66 int64 = 1
                    v43 += (v43 ^ (^(v43)))
                    _ = v66
                }
                for fi67 := 0; fi67 < 13; fi67++ {
                    var v68 string = ""
                    switch ((("hello" + v68) + v68) + v68) {
                    case "hello":
                        if false {
                            for fi69 := 0; fi69 < 6; fi69++ {
                                v43 = 42
                                if v45 {
                                    v3[127] = (v68 + (("hello" + "hello") + (("foo bar" + "") + (((v68 + "a") + v68) + (v68 + v68)))))
                                }
                                _ = len(v1)
                                delete(v3, -1)
                                _ = fi69
                            }
                        } else {
                            for fi70 := 0; fi70 < 16; fi70++ {
                                if true {
                                    var v71 float64 = -2.25
                                    switch (fi70 | fi70) {
                                    case 257:
                                        v68 = v68
                                    default:
                                        if v45 {
                                            if true {
                                                fmt.Println(fi67, fi67)
                                                if true {
                                                    switch 4294967295 {
                                                    case -113:
                                                        switch -1 {
                                                        case 0:
                                                            var v72 int16 = 1
                                                            _ = v72
                                                        case 32767:
                                                            v73 := GenS5{F0: (^(fi67)), F1: v2, F2: int16(v44)}
                                                            var v74 int = (^(fi70))
                                                            _ = v73
                                                            _ = v74
                                                        }
                                                    }
                                                    v75 := GenS3{F0: v43}
                                                    fmt.Println(v60, v71)
                                                    _ = v75
                                                } else {
                                                    switch (v68 + v68) {
                                                    case "hello":
                                                        switch v68 {
                                                        case "hello":
                                                            _ = len(v3)
                                                        default:
                                                            v60 -= (^(v60))
                                                        }
                                                    case "a":
                                                        fmt.Println(v43, v60)
                                                    case "\n\t":
                                                        v3[(fi67 >> ((446) & 63))] = v68
                                                    default:
                                                        v2 -= 2147483647
                                                    }
                                                }
                                                if true {
                                                    v76 := GenS4{F0: 42, F1: (v60 ^ -1), F2: v60}
                                                    _ = v76
                                                }
                                            }
                                        } else {
                                            v77 := GenS6{F0: 195, F1: v45}
                                            _ = v77
                                        }
                                    }
                                    v78 := GenS7{F0: 4294967295, F1: (^(v60)), F2: 735}
                                    fmt.Println(fi67)
                                    v43 = (^(v43))
                                    v79 := GenS8{F0: 32767}
                                    _ = len(v1)
                                    _ = v71
                                    _ = v78
                                    _ = v79
                                }
                                var v80 float64 = -2.25
                                delete(v3, int(v44))
                                _ = fi70
                                _ = v80
                            }
                        }
                    case "":
                        if (!v45) {
                            if v45 {
                                v60 += v60
                                if v45 {
                                    var v81 int16 = -32768
                                    var v82 int32 = genFn0((((v68 + v68) + ("\n\t" + (((v68 + "") + v68) + v68))) + (v68 + v68)), 32768, -45)
                                    v2 += v2
                                    _ = v81
                                    _ = v82
                                } else {
                                    _ = len(v1)
                                    if ((667 > 996) && v45) {
                                        var v83 bool = true
                                        _ = v83
                                    }
                                }
                                v84 := GenS9{F0: 655, F1: (("" < "\n\t") && (v45 || (v68 >= "foo bar"))), F2: 1}
                                _ = v84
                            } else {
                                v60 = (v60 / ((-49) | 1))
                                var v85 int = (-(fi67))
                                v68 = v68
                                _ = v85
                            }
                            v86 := GenS4{F0: 2, F1: -128, F2: v60}
                            _ = v86
                        } else {
                            v87 := GenS7{F0: (v43 % ((v43) | 1)), F1: 86, F2: 0}
                            _ = v87
                        }
                    }
                    v88 := GenS7{F0: (^(v43)), F1: v60, F2: 1}
                    _ = fi67
                    _ = v68
                    _ = v88
                }
            }
        }
        v89 := GenS4{F0: 1, F1: int8(v2), F2: v60}
        _ = v89
    } else {
        v45 = false
        v2 += genFn0("\n\t", 2147483647, 1)
    }
    v45 = ((42 == 18446744073709551615) || ((9223372036854775807 == 347) && false))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v43
    _ = v44
    _ = v45
    _ = v60
}
