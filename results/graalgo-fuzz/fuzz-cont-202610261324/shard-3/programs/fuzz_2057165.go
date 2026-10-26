package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int32; F1 uint16; F2 uint16 }
type GenS4 struct { F0 bool; F1 int8; F2 uint32 }
type GenS5 struct { F0 int32; F1 uint8 }
type GenS6 struct { F0 int64; F1 int; F2 uint16 }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 string; F1 uint; F2 uint }
type GenS9 struct { F0 int16 }
type GenS10 struct { F0 uint16 }
func genFn0() uint8 {
    for fi4 := 0; fi4 < 11; fi4++ {
        fmt.Println(fi4)
        _ = fi4
    }
    var v5 uint8 = 176
    v6 := []int{}
    for fi7 := 0; fi7 < 16; fi7++ {
        var v8 int64 = 42
        _ = fi7
        _ = v8
    }
    _ = v5
    _ = v6
    return (^(v5))
}
func genFn1(p0 uint16, p1 int) int8 {
    v9 := []int{(p1 % ((p1) | 1)), 444, p1}
    var v10 int16 = 0
    v11 := GenS1{F0: (v10 % ((42) | 1)), F1: -9223372036854775808}
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v9
    return 100
}

func main() {
    v1 := make(map[int]string)
    var v2 uint64 = 18446744073709551615
    var v3 bool = ((-2.25 - (0.1 / 1.0)) <= 0.1)
    v2 = v2
    for fi12 := 0; fi12 < 12; fi12++ {
        var v13 int = (-(fi12))
        v14 := []int{v13, v13}
        switch v13 {
        case 127:
            for fi15 := 0; fi15 < 16; fi15++ {
                v16 := GenS2{F0: 42}
                var v17 bool = true
                _ = fi15
                _ = v16
                _ = v17
            }
        case 4294967295:
            v3 = ("\n\t" >= "hello")
        case 896:
            var v18 uint64 = (v2 >> ((0) & 63))
            _ = v18
        default:
            v19 := GenS1{F0: 74, F1: 1}
            _ = v19
        }
        _ = fi12
        _ = v13
        _ = v14
    }
    for fi20 := 0; fi20 < 14; fi20++ {
        switch -231 {
        case 9223372036854775807:
            if (fi20 < fi20) {
                if v3 {
                    v21 := GenS3{F0: 2147483647, F1: 65535, F2: 65535}
                    if v3 {
                        var v22 uint32 = 454
                        v23 := GenS2{F0: 127}
                        var v24 uint64 = v2
                        _ = v22
                        _ = v23
                        _ = v24
                    }
                    _ = v21
                }
            } else {
                var v25 int64 = -727
                _ = v25
            }
        case -246:
            var v26 uint = 549
            _ = v26
        default:
            var v27 float64 = 0.1
            _ = v27
        }
        var v28 uint8 = 255
        v2 += (^(v2))
        _ = fi20
        _ = v28
    }
    if (!true) {
        switch "" {
        case "\n\t":
            v29 := GenS4{F0: v3, F1: 1, F2: 890}
            _ = v29
        case "a":
            if ((457 > -32768) || ((((("" + "foo bar") + (("" + ("\n\t" + "foo bar")) + "foo bar")) + "\n\t") != "a") && (2147483647 > 2147483647))) {
                var v30 int = 354
                _ = v30
            } else {
                v31 := GenS2{F0: 255}
                for fi32 := 0; fi32 < 16; fi32++ {
                    v3 = ((("\n\t" + "a") > "foo bar") && ((!(1 >= 42)) && (float64(1e10) != 3.14159)))
                    if (-32768 >= 42) {
                        for fi33 := 0; fi33 < 18; fi33++ {
                            v1[1] = "a"
                            if v3 {
                                fmt.Println(v2, v2, fi32)
                            } else {
                                v2 -= 42
                                v2 += 0
                                v2 = (^(v2))
                            }
                            _ = fi33
                        }
                        if (v2 != v2) {
                            v3 = (65535 >= 212)
                            var v34 int8 = 127
                            _ = v34
                        }
                        v2 += (v2 % ((198) | 1))
                    }
                    _ = fi32
                }
                _ = v31
            }
        }
        for fi35 := 0; fi35 < 7; fi35++ {
            v36 := GenS5{F0: 139, F1: 42}
            _ = fi35
            _ = v36
        }
        switch ("" + (("hello" + "") + "a")) {
        case "hello":
            for fi37 := 0; fi37 < 12; fi37++ {
                for fi38 := 0; fi38 < 1; fi38++ {
                    v39 := []int{}
                    _ = fi38
                    _ = v39
                }
                if (!v3) {
                    var v40 string = "a"
                    fmt.Println(v2, fi37)
                    _ = v40
                } else {
                    for fi41 := 0; fi41 < 17; fi41++ {
                        v42 := []int{fi37, (fi37 | 1), (fi37 - (fi41 >> ((1) & 63)))}
                        v3 = (!v3)
                        v2 += (v2 ^ (v2 ^ (^(v2))))
                        _ = fi41
                        _ = v42
                    }
                    v1[(fi37 * -1)] = "foo bar"
                    for fi43 := 0; fi43 < 13; fi43++ {
                        if ((167 == 2) || (v3 || v3)) {
                            v2 = 1
                            var v44 bool = (v2 > (v2 % (((v2 << ((1) & 63))) | 1)))
                            _ = v44
                        }
                        var v45 string = (("" + "hello") + "")
                        v46 := []int{32767, fi37}
                        _ = fi43
                        _ = v45
                        _ = v46
                    }
                }
                if (true && true) {
                    fmt.Println(fi37, v2)
                    var v47 uint32 = 42
                    v2 = uint64(v47)
                    _ = v47
                } else {
                    var v48 uint8 = 128
                    switch "\n\t" {
                    case "":
                        if true {
                            delete(v1, (^(fi37)))
                            v49 := GenS6{F0: -807, F1: (^(fi37)), F2: 42}
                            var v50 float64 = -2.25
                            _ = v49
                            _ = v50
                        }
                    case "a":
                        v51 := GenS2{F0: (v48 % (((v48 >> ((4294967295) & 63))) | 1))}
                        _ = v51
                    case "foo bar":
                        v1[(^(fi37))] = "a"
                    }
                    switch (("" + "") + ("\n\t" + "a")) {
                    case "foo bar":
                        v52 := GenS5{F0: -190, F1: 81}
                        _ = v52
                    }
                    _ = v48
                }
                _ = fi37
            }
        case "foo bar":
            if v3 {
                v1[2147483647] = "foo bar"
                v53 := []int{}
                var v54 float32 = ((0.1 + ((1.5 * ((1.5 * (0.0 - (0.0 - (3.14159 - 1e10)))) / ((-2.25 - 1.5) * float32(3.14159)))) / 1e10)) * 0.0)
                _ = v53
                _ = v54
            } else {
                for fi55 := 0; fi55 < 3; fi55++ {
                    if (0.1 < (1.5 * 0.0)) {
                        switch 127 {
                        case -340:
                            v2 = v2
                        case 32767:
                            switch "\n\t" {
                            case "":
                                v56 := GenS3{F0: 1, F1: 1, F2: 0}
                                _ = v56
                            default:
                                v57 := v1[fi55]
                                _ = v57
                            }
                        }
                    }
                    v58 := []int{(fi55 << ((255) & 63)), fi55}
                    fmt.Println(v3, v3, v2)
                    _ = fi55
                    _ = v58
                }
                v59 := GenS4{F0: v3, F1: -128, F2: 2147483647}
                switch "foo bar" {
                case "\n\t":
                    v3 = v3
                case "hello":
                    fmt.Println(v3, v3)
                default:
                    v60 := GenS5{F0: -2147483648, F1: 2}
                    _ = v60
                }
                _ = v59
            }
        default:
            switch 9223372036854775807 {
            case -9223372036854775808:
                var v61 bool = (!v3)
                _ = v61
            default:
                fmt.Println(v3, v2)
            }
        }
    }
    var v62 uint16 = 255
    var v63 float32 = (0.0 + 1.5)
    var v64 float64 = 0.1
    for fi65 := 0; fi65 < 20; fi65++ {
        switch 0 {
        case -1:
            v66 := GenS6{F0: -871, F1: (fi65 % ((fi65) | 1)), F2: (^(v62))}
            _ = v66
        case -72:
            for fi67 := 0; fi67 < 3; fi67++ {
                v1[int(v2)] = "a"
                if true {
                    if ((v2 & v2) == (v2 % ((v2) | 1))) {
                        switch fi67 {
                        case 9223372036854775807:
                            var v68 int64 = 794
                            switch 42 {
                            case -603:
                                v69 := GenS7{F0: 255}
                                _ = v69
                            }
                            _ = v68
                        case -517:
                            var v70 uint32 = 2
                            _ = v70
                        case -1:
                            var v71 int64 = 1
                            _ = v71
                        }
                    } else {
                        var v72 float32 = 1.5
                        v2 += (^(v2))
                        _ = v72
                    }
                } else {
                    v2 = 4294967295
                    if ((^(fi65)) != (-(fi67))) {
                        if (!((42 > -2147483648) || (v3 || (1 < -2147483648)))) {
                            delete(v1, -632)
                            v73 := []int{}
                            _ = v73
                        }
                    } else {
                        v74 := GenS8{F0: "hello", F1: 0, F2: 350}
                        var v75 int16 = -755
                        v63 = 1.5
                        _ = v74
                        _ = v75
                    }
                }
                switch "" {
                case "":
                    var v76 float64 = 1e10
                    _ = v76
                }
                _ = fi67
            }
        case -2147483648:
            fmt.Println(v63)
        default:
            for fi77 := 0; fi77 < 8; fi77++ {
                if (!true) {
                    v78 := GenS6{F0: -243, F1: (fi65 + fi77), F2: v62}
                    v79 := []int{-2147483648, (-(fi65)), (fi77 * -60), (fi77 << ((593) & 63))}
                    _ = v78
                    _ = v79
                }
                _ = len(v1)
                if true {
                    var v80 float64 = (v64 / (0.1 - ((-2.25 * 0.1) + (1e10 + v64))))
                    var v81 string = "hello"
                    if false {
                        v82 := GenS9{F0: -1}
                        v62 = v62
                        v81 = "a"
                        _ = v82
                    }
                    _ = v80
                    _ = v81
                }
                _ = fi77
            }
        }
        v64 -= v64
        _ = fi65
    }
    v83 := GenS10{F0: (^(v62))}
    v63 += v63
    var v84 uint32 = 557
    v85 := GenS8{F0: (("a" + "a") + ("a" + "")), F1: 2, F2: 2}
    for fi86 := 0; fi86 < 8; fi86++ {
        v3 = (!(genFn0() != 255))
        v87 := GenS9{F0: 872}
        switch "a" {
        case "hello":
            var v88 int32 = 2147483647
            _ = v88
        default:
            var v89 uint8 = uint8(v62)
            _ = v89
        }
        _ = fi86
        _ = v87
    }
    if v3 {
        v84 -= (v84 % ((v84) | 1))
    }
    if v3 {
        fmt.Println(v84)
    }
    var v90 int8 = 100
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v62:", v62)
    fmt.Println("v63:", v63)
    fmt.Println("v64:", v64)
    fmt.Println("v84:", v84)
    fmt.Println("v90:", v90)
    _ = v1
    _ = v2
    _ = v3
    _ = v62
    _ = v63
    _ = v64
    _ = v83
    _ = v84
    _ = v85
    _ = v90
}
