package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16; F2 uint32 }
type GenS2 struct { F0 int; F1 uint32; F2 uint64 }
type GenS3 struct { F0 string; F1 int; F2 int }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint8; F1 int8; F2 bool }
type GenS6 struct { F0 string }
type GenS7 struct { F0 int8; F1 int64 }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 uint8; F1 int32; F2 uint64 }
func genFn0(p0 string, p1 uint64) string {
    for fi6 := 0; fi6 < 20; fi6++ {
        v7 := GenS1{F0: 2, F1: 32767, F2: 2}
        p1 += (p1 % ((p1) | 1))
        _ = fi6
        _ = v7
    }
    if false {
        for fi8 := 0; fi8 < 8; fi8++ {
            p1 += (p1 / ((p1) | 1))
            switch "hello" {
            case "\n\t":
                for fi9 := 0; fi9 < 20; fi9++ {
                    v10 := []int{-185, fi8, fi9}
                    p1 += 1
                    switch (p0 + ((p0 + "") + p0)) {
                    case "a":
                        var v11 string = p0
                        _ = v11
                    case "hello":
                        switch (fi8 >> ((4294967295) & 63)) {
                        case -9223372036854775808:
                            fmt.Println(fi8, p1)
                            v12 := GenS2{F0: fi8, F1: 2147483647, F2: (p1 - 42)}
                            p0 = ("" + (("a" + p0) + "hello"))
                            _ = v12
                        case 4294967295:
                            p0 = (p0 + ("\n\t" + (p0 + (p0 + "a"))))
                        case 28:
                            p0 = "\n\t"
                        }
                    default:
                        var v13 uint16 = 0
                        _ = v13
                    }
                    _ = fi9
                    _ = v10
                }
            case "hello":
                var v14 int64 = -459
                _ = v14
            case "a":
                p0 = p0
            default:
                for fi15 := 0; fi15 < 8; fi15++ {
                    var v16 uint16 = 32767
                    v17 := GenS3{F0: p0, F1: (^(fi8)), F2: fi8}
                    v16 = (v16 ^ (v16 / ((945) | 1)))
                    _ = fi15
                    _ = v16
                    _ = v17
                }
            }
            switch fi8 {
            case 585:
                var v18 int32 = -1
                _ = v18
            default:
                p0 = p0
            }
            _ = fi8
        }
    } else {
        v19 := GenS2{F0: 198, F1: 525, F2: (p1 & (^(p1)))}
        if false {
            switch ("" + ("a" + "")) {
            case "\n\t":
                var v20 uint16 = 962
                _ = v20
            case "a":
                if (!((((0.1 != 1.5) && (!true)) || true) && (!((^(p1)) > p1)))) {
                    fmt.Println(p0, p1)
                }
            case "":
                if (!(53 >= 0)) {
                    v21 := []int{128}
                    _ = v21
                } else {
                    v22 := GenS4{F0: 46}
                    var v23 bool = ((!(((557 > -2147483648) || (false || true)) && (126 > 42))) && (-9223372036854775808 != 107))
                    _ = v22
                    _ = v23
                }
            default:
                if (1000 > 42) {
                    p1 = 4294967295
                    p1 += p1
                }
            }
            switch "hello" {
            case "\n\t":
                for fi24 := 0; fi24 < 19; fi24++ {
                    if true {
                        fmt.Println(p1, fi24, p1)
                        var v25 uint32 = 160
                        _ = v25
                    }
                    _ = fi24
                }
            }
            for fi26 := 0; fi26 < 13; fi26++ {
                var v27 uint32 = 912
                if ((false && ((((((p0 + ("\n\t" + p0)) + ((p0 + ((((p0 + (p0 + ("" + "hello"))) + "\n\t") + "a") + "foo bar")) + p0)) + "hello") + ((("hello" + ("" + (p0 + p0))) + ((p0 + (("hello" + (p0 + p0)) + p0)) + p0)) + ("foo bar" + "a"))) != p0) && (true || true))) || false) {
                    v28 := GenS2{F0: (fi26 & fi26), F1: 42, F2: 42}
                    for fi29 := 0; fi29 < 5; fi29++ {
                        fmt.Println(v27, v27, v27)
                        fmt.Println(p1)
                        fmt.Println(fi26, p1)
                        fmt.Println(fi26)
                        _ = fi29
                    }
                    _ = v28
                } else {
                    var v30 uint = 0
                    for fi31 := 0; fi31 < 3; fi31++ {
                        var v32 int64 = -146
                        v33 := GenS2{F0: (fi26 / (((fi26 | fi26)) | 1)), F1: v27, F2: (p1 / (((p1 * p1)) | 1))}
                        if (uint32(p1) >= (v27 * v27)) {
                            fmt.Println(p1, v27)
                            if (true && true) {
                                fmt.Println(fi31, p1)
                                switch "foo bar" {
                                case "a":
                                    var v34 uint8 = 128
                                    _ = v34
                                case "foo bar":
                                    fmt.Println(p1, p0)
                                case "hello":
                                    v35 := GenS5{F0: 1, F1: 100, F2: ((fi26 % ((fi31) | 1)) <= (fi26 ^ -582))}
                                    _ = v35
                                }
                                if ((0 == 0) || true) {
                                    p1 = (p1 * (p1 & p1))
                                } else {
                                    v36 := GenS6{F0: ("" + "hello")}
                                    p1 = (p1 >> ((v30) & 63))
                                    _ = v36
                                }
                            } else {
                                p0 = "hello"
                                p0 = p0
                            }
                        }
                        _ = fi31
                        _ = v32
                        _ = v33
                    }
                    _ = v30
                }
                v37 := GenS7{F0: -88, F1: 1}
                _ = fi26
                _ = v27
                _ = v37
            }
        }
        _ = v19
    }
    _ = p0
    _ = p1
    return (("a" + "\n\t") + "foo bar")
}
func genFn1(p0 int16, p1 bool, p2 int8) int32 {
    fmt.Println(p1, p2)
    switch 303 {
    case 32767:
        p1 = (((0.1 + 1e10) / 1e10) > 3.14159)
    case 971:
        v38 := GenS8{F0: 0}
        _ = v38
    default:
        var v39 uint64 = 42
        _ = v39
    }
    v40 := []int{2147483647}
    var v41 bool = ((9223372036854775807 != 1) && (p1 && p1))
    fmt.Println(v41, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v40
    _ = v41
    return 2147483647
}

func main() {
    var v1 int = 32768
    var v2 string = "\n\t"
    v3 := make(map[uint32]uint32)
    v4 := make(map[string]uint32)
    var v5 bool = true
    if v5 {
        if (42 >= 42) {
            fmt.Println(v2, v1, v1)
            switch ((((v2 + v2) + v2) + v2) + (v2 + "")) {
            case "":
                v1 -= 80
            case "foo bar":
                _ = len(v3)
            }
            if v5 {
                for fi42 := 0; fi42 < 12; fi42++ {
                    v4[(v2 + v2)] = 2147483647
                    var v43 int8 = 103
                    v44 := []int{fi42, v1}
                    _ = fi42
                    _ = v43
                    _ = v44
                }
                _ = len(v4)
                var v45 int8 = 107
                _ = v45
            } else {
                fmt.Println(v2, v5, v2)
            }
        } else {
            for fi46 := 0; fi46 < 9; fi46++ {
                for fi47 := 0; fi47 < 20; fi47++ {
                    v2 = "a"
                    _ = len(v3)
                    _ = fi47
                }
                v48 := GenS9{F0: 255, F1: 1, F2: 18446744073709551615}
                _ = fi46
                _ = v48
            }
        }
    } else {
        _ = len(v3)
        for fi49 := 0; fi49 < 11; fi49++ {
            v1 -= (v1 & 449)
            _ = fi49
        }
        v50 := GenS1{F0: 127, F1: 940, F2: 4294967295}
        _ = v50
    }
    v51 := GenS7{F0: -1, F1: -327}
    fmt.Println(v2)
    var v52 uint16 = 2
    var v53 uint8 = 255
    for fi54 := 0; fi54 < 2; fi54++ {
        var v55 uint32 = 65535
        var v56 uint16 = (^(v52))
        _ = fi54
        _ = v55
        _ = v56
    }
    var v57 int = (v1 / ((1) | 1))
    var v58 uint32 = 1
    v59 := []int{4294967295, (v1 ^ (v57 * (v1 / ((v1) | 1))))}
    var v60 uint64 = 331
    fmt.Println(v5, v53)
    for fi61 := 0; fi61 < 12; fi61++ {
        if v5 {
            _ = len(v59)
        } else {
            fmt.Println(v5, fi61)
            var v62 uint = 567
            _ = v62
        }
        for fi63 := 0; fi63 < 13; fi63++ {
            if (!v5) {
                var v64 string = v2
                fmt.Println(v57, v64)
                v58 = 4294967295
                _ = v64
            }
            _ = fi63
        }
        v65 := GenS2{F0: (v57 | v57), F1: (^(v58)), F2: (v60 / ((v60) | 1))}
        _ = fi61
        _ = v65
    }
    v1 = v1
    if (!((0 == 983) && true)) {
        if false {
            var v66 int16 = 32767
            _ = v66
        } else {
            for fi67 := 0; fi67 < 16; fi67++ {
                v1 = (v1 + (-(v57)))
                var v68 int = (fi67 * 127)
                _ = fi67
                _ = v68
            }
            v53 = v53
        }
    } else {
        switch (v1 ^ (v1 | (v1 - (v57 >> ((65535) & 63))))) {
        case -1:
            fmt.Println(v58)
        case 619:
            for fi69 := 0; fi69 < 6; fi69++ {
                v70 := GenS4{F0: 44}
                _ = fi69
                _ = v70
            }
        case 966:
            var v71 int = 128
            _ = v71
        default:
            if false {
                var v72 uint8 = 42
                fmt.Println(v58, v53)
                var v73 uint16 = 42
                _ = v72
                _ = v73
            }
        }
    }
    var v74 int32 = 0
    v53 -= 241
    _ = len(v4)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    fmt.Println("v57:", v57)
    fmt.Println("v58:", v58)
    fmt.Println("slicelen v59:", len(v59), v59)
    fmt.Println("v60:", v60)
    fmt.Println("v74:", v74)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v51
    _ = v52
    _ = v53
    _ = v57
    _ = v58
    _ = v59
    _ = v60
    _ = v74
}
