package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64; F2 uint8 }
type GenS2 struct { F0 uint32; F1 int64 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint8; F1 uint64; F2 int16 }
type GenS5 struct { F0 string; F1 int64; F2 int16 }
type GenS6 struct { F0 int8; F1 uint64; F2 uint32 }
type GenS7 struct { F0 uint16; F1 uint }
type GenS8 struct { F0 int64; F1 uint16; F2 uint32 }
func genFn0(p0 string, p1 uint, p2 int) int8 {
    fmt.Println(p1, p2)
    v4 := GenS1{F0: false, F1: -305, F2: 0}
    for fi5 := 0; fi5 < 17; fi5++ {
        p1 -= p1
        p0 = (p0 + p0)
        fmt.Println(p2)
        _ = fi5
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return -1
}
func genFn1(p0 string) int64 {
    v6 := GenS1{F0: (true || (2147483647 != 1)), F1: 9223372036854775807, F2: 255}
    p0 = ""
    for fi7 := 0; fi7 < 19; fi7++ {
        for fi8 := 0; fi8 < 15; fi8++ {
            v9 := GenS1{F0: true, F1: -9223372036854775808, F2: 93}
            p0 = ""
            p0 = (p0 + "a")
            _ = fi8
            _ = v9
        }
        fmt.Println(p0, fi7)
        _ = fi7
    }
    v10 := []int{42, 32768}
    if true {
        if (false || (1000 != 431)) {
            if (true && (-1 >= 32767)) {
                v11 := GenS1{F0: false, F1: -510, F2: 2}
                for fi12 := 0; fi12 < 13; fi12++ {
                    switch 32767 {
                    case 763:
                        var v13 float64 = 0.1
                        _ = v13
                    case 127:
                        var v14 uint8 = 42
                        _ = v14
                    case 4294967295:
                        var v15 bool = false
                        _ = v15
                    default:
                        fmt.Println(fi12)
                    }
                    _ = fi12
                }
                for fi16 := 0; fi16 < 1; fi16++ {
                    v17 := GenS1{F0: (1 >= -1), F1: 1, F2: 42}
                    if ((-2147483648 == 519) || true) {
                        if len(v10) > 0 {
                            v10[0] = fi16
                        }
                    }
                    p0 = ("" + p0)
                    _ = fi16
                    _ = v17
                }
                _ = v11
            } else {
                p0 = (p0 + ("hello" + "foo bar"))
                var v18 uint32 = 42
                _ = v18
            }
            var v19 uint64 = 719
            for fi20 := 0; fi20 < 18; fi20++ {
                if true {
                    for fi21 := 0; fi21 < 8; fi21++ {
                        var v22 uint16 = 1
                        _ = fi21
                        _ = v22
                    }
                    if ((true || false) && true) {
                        var v23 uint64 = v19
                        _ = v23
                    } else {
                        fmt.Println(fi20, v19)
                        for fi24 := 0; fi24 < 2; fi24++ {
                            v19 = v19
                            v25 := GenS1{F0: (false || true), F1: 144, F2: 255}
                            var v26 int8 = genFn0(p0, 0, (fi24 | (fi20 ^ 32767)))
                            _ = fi24
                            _ = v25
                            _ = v26
                        }
                    }
                } else {
                    fmt.Println(v19, fi20, v19)
                    p0 = p0
                }
                v19 += 2
                if (false || (!(-1 < 1))) {
                    for fi27 := 0; fi27 < 14; fi27++ {
                        v28 := GenS1{F0: false, F1: -801, F2: 2}
                        v19 = v19
                        var v29 float32 = -2.25
                        _ = fi27
                        _ = v28
                        _ = v29
                    }
                }
                _ = fi20
            }
            _ = v19
        } else {
            v30 := GenS2{F0: 2147483647, F1: 920}
            v10 = append(v10, -9223372036854775808)
            _ = v30
        }
        switch -896 {
        case -9223372036854775808:
            fmt.Println(p0)
        case 269:
            var v31 float32 = 3.14159
            _ = v31
        case -2147483648:
            if len(v10) >= 2 {
                v10 = v10[1:len(v10)]
            }
        }
    } else {
        if len(v10) >= 2 {
            v10 = v10[1:len(v10)]
        }
        for fi32 := 0; fi32 < 14; fi32++ {
            var v33 bool = (true && (65535 >= 65535))
            _ = fi32
            _ = v33
        }
        v34 := GenS3{F0: 0}
        _ = v34
    }
    _ = p0
    _ = v10
    _ = v6
    return -87
}

func main() {
    var v1 uint64 = 1
    var v2 bool = (!(-1 == -1))
    v3 := make(map[int]string)
    if false {
        for fi35 := 0; fi35 < 15; fi35++ {
            for fi36 := 0; fi36 < 15; fi36++ {
                var v37 int8 = 18
                v38 := GenS4{F0: 227, F1: 636, F2: 1000}
                _ = fi36
                _ = v37
                _ = v38
            }
            fmt.Println(fi35, fi35)
            _ = fi35
        }
        for fi39 := 0; fi39 < 11; fi39++ {
            if v2 {
                _ = len(v3)
            } else {
                for fi40 := 0; fi40 < 10; fi40++ {
                    var v41 int64 = 0
                    if (v41 > 0) {
                        delete(v3, (fi40 << ((4294967295) & 63)))
                        var v42 bool = (306 <= 21)
                        _ = v42
                    } else {
                        if v2 {
                            if (false && false) {
                                var v43 int64 = genFn1("\n\t")
                                if v2 {
                                    v44 := GenS1{F0: v2, F1: (v43 ^ 857), F2: 173}
                                    v43 += (v43 - genFn1(("a" + "hello")))
                                    if ((v2 && (!(-1 < 1000))) || true) {
                                        delete(v3, (fi39 % ((fi40) | 1)))
                                        var v45 uint8 = 0
                                        if v2 {
                                            v46 := GenS3{F0: (fi40 & fi39)}
                                            var v47 uint16 = 2
                                            if true {
                                                fmt.Println(fi40, fi40)
                                            }
                                            var v48 float32 = 0.1
                                            _ = v46
                                            _ = v47
                                            _ = v48
                                        }
                                        _ = v45
                                    } else {
                                        var v49 uint = 255
                                        v43 -= genFn1("")
                                        delete(v3, fi40)
                                        v43 = genFn1("a")
                                        switch "" {
                                        case "hello":
                                            var v50 string = "hello"
                                            _ = v50
                                        case "":
                                            fmt.Println(fi40, v43)
                                        default:
                                            v51 := GenS1{F0: ((42 > 17) && ((v2 && (1 <= 976)) && v2)), F1: v43, F2: 0}
                                            if v2 {
                                                var v52 uint = 651
                                                if ((v52 << (((v52 % (((v49 - 255)) | 1))) & 63)) >= (^(v52))) {
                                                    switch (fi39 >> ((v52) & 63)) {
                                                    case 0:
                                                        v53 := GenS5{F0: ("\n\t" + "\n\t"), F1: genFn1("a"), F2: 0}
                                                        _ = v53
                                                    case -713:
                                                        var v54 uint8 = 0
                                                        _ = v54
                                                    default:
                                                        v43 += (-(v43))
                                                    }
                                                } else {
                                                    var v55 int64 = v41
                                                    v2 = true
                                                    fmt.Println(fi39, v52)
                                                    _ = v55
                                                }
                                                if true {
                                                    fmt.Println(v52, v41)
                                                    fmt.Println(v41)
                                                    v3[9223372036854775807] = "hello"
                                                    v43 += (v41 >> ((65535) & 63))
                                                    v56 := GenS6{F0: -99, F1: 2, F2: 34}
                                                    v2 = false
                                                    _ = v56
                                                }
                                                _ = v52
                                            } else {
                                                var v57 uint = v49
                                                var v58 int64 = -34
                                                _ = v57
                                                _ = v58
                                            }
                                            _ = v51
                                        }
                                        _ = v49
                                    }
                                    var v59 uint32 = 0
                                    _ = v44
                                    _ = v59
                                }
                                var v60 uint8 = 128
                                v3[fi40] = "foo bar"
                                _ = v43
                                _ = v60
                            }
                        } else {
                            var v61 bool = (65535 != 255)
                            v62 := GenS7{F0: 1, F1: 491}
                            if (("\n\t" + "\n\t") <= ("hello" + "a")) {
                                var v63 uint = 2
                                v64 := []int{(fi40 * fi40), (^(fi39)), (fi39 ^ fi40)}
                                v41 -= v41
                                _ = v63
                                _ = v64
                            } else {
                                var v65 uint = 654
                                _ = v65
                            }
                            _ = v61
                            _ = v62
                        }
                        v2 = v2
                        v2 = (!v2)
                    }
                    v41 = (v41 % (((-(v41))) | 1))
                    var v66 int8 = 100
                    var v67 uint16 = 1
                    _ = fi40
                    _ = v41
                    _ = v66
                    _ = v67
                }
                switch (fi39 - (fi39 & fi39)) {
                case 32767:
                    switch (("a" + "hello") + (("" + "\n\t") + "")) {
                    case "hello":
                        v68 := GenS8{F0: -580, F1: 32767, F2: 582}
                        _ = v68
                    case "":
                        var v69 uint8 = 52
                        _ = v69
                    case "a":
                        v70 := GenS6{F0: 43, F1: v1, F2: 4294967295}
                        _ = v70
                    default:
                        v1 -= 36
                    }
                }
                v1 -= (v1 >> ((2) & 63))
            }
            _ = fi39
        }
    }
    if (3.14159 == 3.14159) {
        v2 = (!(true || (2 != 32767)))
        if v2 {
            v71 := GenS7{F0: 0, F1: 111}
            _ = v71
        }
        v72 := GenS2{F0: 1, F1: genFn1("")}
        _ = v72
    } else {
        if ((true || (!true)) || (1 > 0)) {
            v1 = (v1 % (((v1 + v1)) | 1))
        }
        v3[-849] = (("\n\t" + (("\n\t" + ("a" + ("a" + "a"))) + "foo bar")) + ("" + ""))
    }
    var v73 int64 = -957
    var v74 uint64 = (v1 / ((18446744073709551615) | 1))
    if (!false) {
        v2 = v2
        if ((!(genFn1("hello") >= 979)) || (65535 >= uint16(v1))) {
            var v75 int32 = 1
            v75 = v75
            _ = v75
        }
        var v76 string = (("hello" + "a") + "")
        _ = v76
    } else {
        v77 := GenS2{F0: 1, F1: 834}
        _ = v77
    }
    _ = len(v3)
    switch 2147483647 {
    case -2147483648:
        if false {
            switch 0 {
            case 512:
                var v78 uint64 = v74
                _ = v78
            case -869:
                fmt.Println(v74, v2, v74)
            case 42:
                fmt.Println(v74, v1, v73)
            }
        } else {
            v79 := GenS3{F0: -1}
            v73 = v73
            _ = v79
        }
    case 0:
        var v80 int32 = 2147483647
        _ = v80
    }
    var v81 int64 = v73
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v73:", v73)
    fmt.Println("v74:", v74)
    fmt.Println("v81:", v81)
    _ = v1
    _ = v2
    _ = v3
    _ = v73
    _ = v74
    _ = v81
}
