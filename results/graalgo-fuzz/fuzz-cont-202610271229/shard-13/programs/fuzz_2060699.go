package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string; F2 uint32 }
type GenS2 struct { F0 uint; F1 bool; F2 uint32 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int; F1 uint32; F2 uint }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 string; F1 uint64; F2 int8 }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 uint64; F1 int8; F2 uint16 }
func genFn0() uint32 {
    v5 := []int{1, 2147483647, 4294967295, 4294967295}
    var v6 uint32 = 4294967295
    var v7 float32 = -2.25
    v8 := GenS1{F0: 127, F1: ("a" + "foo bar"), F2: v6}
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return v6
}
func genFn1(p0 int, p1 uint32, p2 int16) uint32 {
    switch (p0 + (p0 ^ p0)) {
    case 2147483647:
        for fi9 := 0; fi9 < 7; fi9++ {
            v10 := []int{}
            for fi11 := 0; fi11 < 12; fi11++ {
                var v12 int64 = 9223372036854775807
                _ = fi11
                _ = v12
            }
            for fi13 := 0; fi13 < 6; fi13++ {
                var v14 uint16 = 255
                p2 = (-(p2))
                _ = fi13
                _ = v14
            }
            _ = fi9
            _ = v10
        }
    case -9223372036854775808:
        var v15 int8 = -6
        _ = v15
    }
    v16 := []int{p0, -88, (p0 - p0)}
    v16 = append(v16, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    return 957
}

func main() {
    var v1 int16 = -32768
    var v2 uint64 = 998
    var v3 uint64 = v2
    var v4 int16 = v1
    v1 += v1
    v3 += 450
    for fi17 := 0; fi17 < 16; fi17++ {
        v18 := []int{(fi17 << ((2) & 63)), -954}
        if len(v18) > 0 {
            v18[0] = fi17
        }
        switch 0 {
        case 4294967295:
            if len(v18) >= 2 {
                v18 = v18[1:len(v18)]
            }
        case 32767:
            v4 = (-(v1))
        case -187:
            v19 := GenS2{F0: 0, F1: (v2 == (v2 + v2)), F2: 1}
            _ = v19
        }
        _ = fi17
        _ = v18
    }
    v20 := []int{-1}
    if false {
        fmt.Println(v3, v2, v4)
    }
    var v21 string = "\n\t"
    for fi22 := 0; fi22 < 8; fi22++ {
        var v23 uint32 = 0
        var v24 float32 = 0.1
        _ = fi22
        _ = v23
        _ = v24
    }
    _ = len(v20)
    v2 -= v2
    fmt.Println(v4, v21, v2)
    if len(v20) > 0 {
        v20[0] = -449
    }
    switch -675 {
    case 0:
        if false {
            if (100 < 0) {
                switch 803 {
                case 489:
                    switch ("a" + "a") {
                    case "\n\t":
                        var v25 float64 = 3.14159
                        _ = v25
                    case "hello":
                        switch 0 {
                        case -9223372036854775808:
                            var v26 int = 9223372036854775807
                            _ = v26
                        case 32768:
                            v27 := GenS2{F0: 2, F1: (4294967295 <= (^(v3))), F2: 2}
                            _ = v27
                        case 196:
                            v4 -= (v1 + (v1 / (((v4 * (^(v4)))) | 1)))
                        default:
                            var v28 uint16 = 1
                            _ = v28
                        }
                    }
                case 127:
                    var v29 uint16 = 1
                    _ = v29
                case 32767:
                    for fi30 := 0; fi30 < 11; fi30++ {
                        switch v21 {
                        case "foo bar":
                            v31 := GenS1{F0: -33, F1: v21, F2: 106}
                            _ = v31
                        case "":
                            v32 := GenS1{F0: 100, F1: (v21 + v21), F2: 42}
                            _ = v32
                        }
                        v21 = "a"
                        var v33 uint16 = 0
                        _ = fi30
                        _ = v33
                    }
                default:
                    v21 = ""
                }
                if len(v20) >= 2 {
                    v20 = v20[1:len(v20)]
                }
            } else {
                if len(v20) >= 2 {
                    v20 = v20[1:len(v20)]
                }
            }
        }
    case 2147483647:
        for fi34 := 0; fi34 < 5; fi34++ {
            v35 := GenS3{F0: 255}
            for fi36 := 0; fi36 < 17; fi36++ {
                if true {
                    if true {
                        var v37 uint8 = 2
                        var v38 float64 = (0.0 + 1.5)
                        _ = v37
                        _ = v38
                    } else {
                        v2 += (v3 - v2)
                        if true {
                            var v39 int16 = (-(v1))
                            if false {
                                var v40 uint32 = 254
                                switch (fi36 - fi34) {
                                case 32767:
                                    fmt.Println(v21, v4)
                                }
                                if true {
                                    v20 = append(v20, fi34)
                                } else {
                                    v39 -= -30
                                }
                                _ = v40
                            } else {
                                v3 += v2
                                var v41 int64 = -9223372036854775808
                                v4 = v1
                                _ = v41
                            }
                            _ = v39
                        }
                    }
                    v20 = append(v20, -2147483648)
                    var v42 int16 = -288
                    _ = v42
                }
                if ((-753 < 2147483647) || (!((v1 > (v1 | -107)) && false))) {
                    v43 := GenS4{F0: 0}
                    _ = v43
                }
                switch 1 {
                case -1:
                    var v44 uint32 = genFn1(2147483647, genFn1((fi36 % ((fi36) | 1)), 2147483648, (v4 << ((0) & 63))), v1)
                    _ = v44
                case 0:
                    v45 := GenS1{F0: 127, F1: v21, F2: 2147483647}
                    _ = v45
                default:
                    var v46 uint64 = v2
                    _ = v46
                }
                _ = fi36
            }
            if true {
                if false {
                    var v47 bool = false
                    v4 = (v4 / ((v1) | 1))
                    _ = v47
                }
                if false {
                    v48 := GenS2{F0: 18446744073709551615, F1: (1.5 >= (0.0 * -2.25)), F2: 2}
                    for fi49 := 0; fi49 < 19; fi49++ {
                        var v50 int8 = 0
                        v51 := GenS4{F0: 42}
                        _ = fi49
                        _ = v50
                        _ = v51
                    }
                    switch "a" {
                    case "hello":
                        fmt.Println(v3, v3)
                    }
                    _ = v48
                }
            }
            _ = fi34
            _ = v35
        }
    default:
        fmt.Println(v1)
    }
    v52 := GenS5{F0: -2147483648, F1: genFn0(), F2: 0}
    for fi53 := 0; fi53 < 20; fi53++ {
        v3 += v2
        for fi54 := 0; fi54 < 13; fi54++ {
            if len(v20) > 0 {
                v20[0] = fi54
            }
            var v55 uint = 4294967295
            var v56 float32 = 1.5
            v3 += v3
            var v57 float64 = 1.5
            _ = fi54
            _ = v55
            _ = v56
            _ = v57
        }
        switch v21 {
        case "\n\t":
            v4 += 42
        case "":
            v58 := GenS3{F0: 0}
            _ = v58
        }
        _ = fi53
    }
    for fi59 := 0; fi59 < 8; fi59++ {
        for fi60 := 0; fi60 < 18; fi60++ {
            var v61 int64 = 42
            if len(v20) >= 2 {
                v20 = v20[1:len(v20)]
            }
            switch (("\n\t" + (v21 + v21)) + ((v21 + v21) + "hello")) {
            case "\n\t":
                v62 := GenS6{F0: 42}
                _ = v62
            case "":
                v3 -= (v2 | (v2 % ((v2) | 1)))
            case "foo bar":
                if false {
                    switch (fi60 * 1) {
                    case -247:
                        fmt.Println(v2, v4)
                    }
                    if false {
                        switch v21 {
                        case "a":
                            v4 -= (-(v1))
                        case "foo bar":
                            if (1 <= 42) {
                                switch 1 {
                                case -997:
                                    _ = len(v20)
                                case 0:
                                    var v63 string = v21
                                    _ = v63
                                case 9223372036854775807:
                                    v61 = (v61 - (v61 + -488))
                                    switch (fi60 ^ (fi60 * fi59)) {
                                    case 32767:
                                        v64 := GenS1{F0: 100, F1: "a", F2: 4294967295}
                                        _ = v64
                                    case 283:
                                        if len(v20) >= 2 {
                                            v20 = v20[1:len(v20)]
                                        }
                                    }
                                    v4 -= (^(v4))
                                    fmt.Println(v21, v61, v3)
                                    v4 -= (v4 ^ 32767)
                                default:
                                    v61 -= v61
                                    var v65 uint8 = 42
                                    if len(v20) > 0 {
                                        v20[0] = (fi59 | fi60)
                                    }
                                    _ = v65
                                }
                            } else {
                                switch "" {
                                case "":
                                    fmt.Println(v2, fi60)
                                case "foo bar":
                                    fmt.Println(v4, v1)
                                    if true {
                                        v20 = append(v20, fi59)
                                        fmt.Println(fi59, fi59, fi59)
                                    }
                                    v20 = append(v20, (fi59 % (((fi59 & fi59)) | 1)))
                                case "hello":
                                    switch "a" {
                                    case "\n\t":
                                        var v66 float64 = 1e10
                                        _ = v66
                                    case "hello":
                                        var v67 int64 = v61
                                        _ = v67
                                    case "foo bar":
                                        switch fi60 {
                                        case -9223372036854775808:
                                            v68 := GenS7{F0: v21, F1: 42, F2: int8(fi59)}
                                            _ = v68
                                        case 356:
                                            if (true && false) {
                                                var v69 int = fi59
                                                _ = v69
                                            }
                                        default:
                                            var v70 int64 = (-(v61))
                                            _ = v70
                                        }
                                    }
                                }
                                switch 82 {
                                case 1:
                                    if false {
                                        var v71 uint64 = 4294967295
                                        _ = v71
                                    } else {
                                        var v72 int8 = 42
                                        fmt.Println(v2, fi59, v4)
                                        var v73 int32 = -634
                                        _ = v72
                                        _ = v73
                                    }
                                case -260:
                                    v74 := GenS5{F0: fi59, F1: 65535, F2: 371}
                                    _ = v74
                                case -199:
                                    v21 = v21
                                }
                                v21 = (v21 + ((v21 + ((v21 + "foo bar") + "foo bar")) + ("a" + "\n\t")))
                            }
                        default:
                            var v75 uint32 = 182
                            _ = v75
                        }
                        if (((v61 | v61) <= (-(v61))) && false) {
                            v76 := GenS8{F0: -1}
                            _ = v76
                        } else {
                            var v77 float32 = 3.14159
                            var v78 string = v21
                            var v79 int32 = -2147483648
                            _ = v77
                            _ = v78
                            _ = v79
                        }
                    } else {
                        v2 += (v2 ^ v2)
                    }
                    if (255 <= 42) {
                        v80 := GenS8{F0: 0}
                        var v81 uint64 = v3
                        _ = v80
                        _ = v81
                    }
                }
            default:
                var v82 uint8 = 0
                _ = v82
            }
            _ = fi60
            _ = v61
        }
        _ = len(v20)
        v83 := GenS9{F0: 1, F1: 1, F2: 255}
        _ = fi59
        _ = v83
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v52
}
