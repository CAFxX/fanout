package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int64; F1 uint16; F2 uint16 }
type GenS4 struct { F0 int8; F1 string }
type GenS5 struct { F0 bool; F1 int; F2 int32 }
type GenS6 struct { F0 uint; F1 uint; F2 int16 }
type GenS7 struct { F0 int; F1 uint64; F2 int32 }
type GenS8 struct { F0 int64; F1 int8; F2 int8 }
type GenS9 struct { F0 int; F1 uint32; F2 int8 }
type GenS10 struct { F0 uint8 }
type GenS11 struct { F0 uint8 }
type GenS12 struct { F0 int64 }
type GenS13 struct { F0 int64; F1 int32 }
func genFn0(p0 int16, p1 int8, p2 int8) bool {
    v6 := GenS1{F0: 2147483647}
    for fi7 := 0; fi7 < 8; fi7++ {
        for fi8 := 0; fi8 < 3; fi8++ {
            switch fi7 {
            case 0:
                var v9 int64 = -1
                _ = v9
            }
            if (-728 == 1) {
                v10 := GenS2{F0: 42}
                _ = v10
            } else {
                var v11 float64 = 1e10
                switch (fi7 & 4294967295) {
                case -9223372036854775808:
                    v12 := GenS3{F0: 42, F1: 265, F2: 65535}
                    fmt.Println(p0, p1)
                    var v13 uint8 = 255
                    _ = v12
                    _ = v13
                case 128:
                    v14 := GenS1{F0: -2147483648}
                    _ = v14
                case 32767:
                    v15 := []int{(fi7 >> ((917) & 63)), fi8, fi7, fi8}
                    _ = v15
                }
                _ = v11
            }
            _ = fi8
        }
        for fi16 := 0; fi16 < 12; fi16++ {
            var v17 uint16 = 0
            v18 := []int{fi7}
            if (128 < 150) {
                var v19 int32 = -210
                var v20 uint32 = 408
                var v21 float32 = (((0.0 - float32((0.0 / 1e10))) + 1.5) / 1.0)
                if ((-(v19)) >= (v19 * 42)) {
                    var v22 int16 = p0
                    fmt.Println(p1, v17)
                    _ = v22
                }
                _ = v19
                _ = v20
                _ = v21
            }
            _ = fi16
            _ = v17
            _ = v18
        }
        fmt.Println(p1)
        _ = fi7
    }
    switch "a" {
    case "\n\t":
        var v23 float32 = (((3.14159 - (1e10 / (1e10 + 1e10))) + 3.14159) * 3.14159)
        _ = v23
    case "foo bar":
        v24 := []int{2147483647, -922}
        _ = v24
    }
    if ((((p0 % ((1) | 1)) < 1) && (true || false)) && (0 <= 659)) {
        for fi25 := 0; fi25 < 5; fi25++ {
            if true {
                v26 := GenS4{F0: (p2 % ((100) | 1)), F1: "\n\t"}
                fmt.Println(p2, p0, fi25)
                v27 := []int{}
                _ = v26
                _ = v27
            }
            var v28 uint8 = 1
            _ = fi25
            _ = v28
        }
        var v29 float64 = float64(0.0)
        _ = v29
    }
    for fi30 := 0; fi30 < 9; fi30++ {
        p2 = p1
        switch (fi30 << ((65535) & 63)) {
        case 9223372036854775807:
            if (127 <= 254) {
                p1 += 1
            } else {
                var v31 float64 = 1.5
                if ((p2 <= p2) || false) {
                    for fi32 := 0; fi32 < 10; fi32++ {
                        v33 := GenS4{F0: 127, F1: (((("hello" + "a") + "a") + "a") + ("" + "hello"))}
                        switch fi30 {
                        case 2147483647:
                            switch "hello" {
                            case "hello":
                                if (!false) {
                                    v31 += ((v31 * (v31 / v31)) * (((0.0 + 1.5) * 0.0) + (-2.25 / ((0.1 * v31) + v31))))
                                }
                            default:
                                if (false || true) {
                                    if true {
                                        var v34 uint16 = 32767
                                        _ = v34
                                    } else {
                                        if (148 == 1) {
                                            fmt.Println(fi32, p2, fi30)
                                            fmt.Println(p2)
                                            if false {
                                                v35 := GenS2{F0: p0}
                                                switch fi30 {
                                                case 128:
                                                    var v36 uint64 = 4294967295
                                                    _ = v36
                                                }
                                                _ = v35
                                            } else {
                                                var v37 uint64 = 0
                                                if false {
                                                    p1 = -127
                                                    v38 := GenS5{F0: false, F1: (^(fi30)), F2: -2147483648}
                                                    v39 := []int{(-(fi32)), fi32, (fi30 % ((-1) | 1)), (fi30 | (fi30 | (fi32 << ((4294967295) & 63))))}
                                                    _ = v38
                                                    _ = v39
                                                }
                                                if (875 <= 32767) {
                                                    v40 := []int{(fi30 & fi32), (fi32 & fi32), fi32}
                                                    p0 += p0
                                                    var v41 uint8 = 2
                                                    v42 := GenS6{F0: 577, F1: 4294967295, F2: (^(p0))}
                                                    v43 := GenS4{F0: p1, F1: "foo bar"}
                                                    var v44 uint8 = v41
                                                    _ = v40
                                                    _ = v41
                                                    _ = v42
                                                    _ = v43
                                                    _ = v44
                                                } else {
                                                    p2 += (p1 >> ((1) & 63))
                                                    v37 -= 345
                                                    v45 := []int{178, -980}
                                                    _ = v45
                                                }
                                                _ = v37
                                            }
                                        } else {
                                            switch (fi32 | (-(fi32))) {
                                            case 32768:
                                                p0 += 157
                                            case -258:
                                                if false {
                                                    var v46 uint = 2
                                                    _ = v46
                                                }
                                            default:
                                                if false {
                                                    var v47 float64 = v31
                                                    if false {
                                                        var v48 int8 = p1
                                                        if true {
                                                            if (v31 < (v47 + v31)) {
                                                                if true {
                                                                    v49 := []int{(^(fi32)), (fi32 << ((14) & 63)), fi30, fi30}
                                                                    _ = v49
                                                                }
                                                                fmt.Println(v31, p2, p2)
                                                                var v50 float64 = (v31 + (1e10 * v31))
                                                                _ = v50
                                                            }
                                                        }
                                                        _ = v48
                                                    }
                                                    var v51 int32 = 0
                                                    v52 := []int{(fi32 >> ((148) & 63)), (fi30 ^ (fi32 + 127)), -793}
                                                    v53 := GenS7{F0: 634, F1: 2, F2: 1}
                                                    _ = v47
                                                    _ = v51
                                                    _ = v52
                                                    _ = v53
                                                }
                                            }
                                            v54 := GenS8{F0: 1, F1: p1, F2: (p1 * (p1 - p1))}
                                            if (!(("\n\t" + "foo bar") >= "hello")) {
                                                v55 := GenS1{F0: 2147483647}
                                                switch "" {
                                                case "":
                                                    v56 := GenS9{F0: 32768, F1: 2147483648, F2: 1}
                                                    _ = v56
                                                }
                                                _ = v55
                                            } else {
                                                fmt.Println(fi32, v31, v31)
                                                var v57 bool = (v31 != 3.14159)
                                                v58 := GenS10{F0: 42}
                                                _ = v57
                                                _ = v58
                                            }
                                            _ = v54
                                        }
                                    }
                                    v59 := GenS5{F0: ((p2 / ((p2) | 1)) != p1), F1: 2147483647, F2: -2147483648}
                                    v60 := GenS10{F0: 255}
                                    _ = v59
                                    _ = v60
                                }
                            }
                        default:
                            v61 := GenS11{F0: 0}
                            _ = v61
                        }
                        switch "foo bar" {
                        case "\n\t":
                            var v62 bool = false
                            _ = v62
                        default:
                            var v63 float64 = (3.14159 + ((v31 - 0.1) - float64((0.0 - (float32((3.14159 * ((v31 - -2.25) / v31))) * (float32((v31 / v31)) + ((((1.5 + 3.14159) / 3.14159) + 1e10) - 3.14159)))))))
                            _ = v63
                        }
                        _ = fi32
                        _ = v33
                    }
                }
                _ = v31
            }
        case 4294967295:
            v64 := []int{}
            _ = v64
        default:
            var v65 float32 = 1e10
            _ = v65
        }
        for fi66 := 0; fi66 < 12; fi66++ {
            var v67 uint16 = 42
            _ = fi66
            _ = v67
        }
        _ = fi30
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return true
}
func genFn1(p0 bool) bool {
    switch "foo bar" {
    case "a":
        switch "foo bar" {
        case "":
            p0 = false
        default:
            p0 = true
        }
    case "hello":
        fmt.Println(p0)
    case "\n\t":
        v68 := []int{-9223372036854775808, -237}
        _ = v68
    default:
        fmt.Println(p0)
    }
    var v69 uint64 = 76
    var v70 bool = p0
    var v71 uint8 = 38
    _ = p0
    _ = v69
    _ = v70
    _ = v71
    return (p0 || false)
}

func main() {
    v1 := make(map[uint32]int)
    v2 := []uint32{294}
    var v3 uint8 = 127
    v4 := []int{}
    v5 := make(map[int]bool)
    v72 := GenS12{F0: -656}
    v73 := GenS11{F0: (^(v3))}
    fmt.Println(v3)
    v5[4294967295] = (!false)
    v3 += 2
    for fi74 := 0; fi74 < 13; fi74++ {
        for fi75 := 0; fi75 < 12; fi75++ {
            v4 = append(v4, fi74)
            v3 -= v3
            _ = fi75
        }
        _ = fi74
    }
    if true {
        v76 := GenS5{F0: (2 == 18446744073709551615), F1: 127, F2: 0}
        fmt.Println(v3)
        if (290 <= -197) {
            v3 = (v3 | 42)
            var v77 uint = 0
            _ = v77
        } else {
            switch 9223372036854775807 {
            case -2147483648:
                v3 += v3
            case 601:
                v3 = 128
            case 511:
                fmt.Println(v3)
            default:
                v3 += 127
            }
        }
        _ = v76
    } else {
        v3 = v3
        for fi78 := 0; fi78 < 7; fi78++ {
            if (0 != 0) {
                v3 += 1
            } else {
                _ = len(v5)
                switch 127 {
                case 42:
                    v3 -= v3
                case -978:
                    v2 = append(v2, 42)
                default:
                    var v79 int8 = 42
                    _ = v79
                }
            }
            v3 = v3
            _ = len(v5)
            _ = fi78
        }
    }
    if (-362 >= -916) {
        fmt.Println(v3)
        v80 := GenS13{F0: int64(v3), F1: 42}
        for fi81 := 0; fi81 < 11; fi81++ {
            switch "a" {
            case "":
                _ = len(v5)
            }
            _ = fi81
        }
        _ = v80
    } else {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        var v82 int8 = 103
        _ = v82
    }
    var v83 uint = 2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v83:", v83)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v72
    _ = v73
    _ = v83
}
