package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32 }
type GenS2 struct { F0 uint32; F1 uint64; F2 uint16 }
type GenS3 struct { F0 uint8; F1 int8 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint16; F1 int; F2 int16 }
type GenS6 struct { F0 bool; F1 uint32 }
type GenS7 struct { F0 int; F1 int32 }
type GenS8 struct { F0 uint; F1 int8 }
type GenS9 struct { F0 int8; F1 uint64 }
func genFn0(p0 uint32, p1 string) int {
    var v6 uint8 = 243
    v7 := []int{}
    for fi8 := 0; fi8 < 2; fi8++ {
        v9 := GenS1{F0: 42, F1: 2147483647}
        var v10 string = p1
        var v11 string = (v10 + p1)
        _ = fi8
        _ = v10
        _ = v11
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return -1
}
func genFn1(p0 uint8) uint64 {
    var v12 uint32 = 4294967295
    if true {
        for fi13 := 0; fi13 < 16; fi13++ {
            v12 -= (v12 % (((^(v12))) | 1))
            if false {
                if (!false) {
                    p0 = p0
                    var v14 string = "a"
                    _ = v14
                }
                var v15 uint32 = v12
                _ = v15
            }
            v12 = (v12 * (^(v12)))
            _ = fi13
        }
        var v16 int32 = -1
        v17 := GenS2{F0: (v12 % (((^(v12))) | 1)), F1: 1, F2: 812}
        _ = v16
        _ = v17
    }
    var v18 uint = 42
    if true {
        fmt.Println(v12, v12)
    } else {
        v19 := []int{4294967295}
        var v20 uint16 = 0
        for fi21 := 0; fi21 < 11; fi21++ {
            var v22 float32 = 1e10
            switch int(v12) {
            case 0:
                if false {
                    v12 = (v12 | 812)
                    p0 += 255
                } else {
                    for fi23 := 0; fi23 < 9; fi23++ {
                        v22 += 0.0
                        if (uint16(p0) == (v20 ^ v20)) {
                            var v24 int = genFn0(702, "hello")
                            v20 = (^(v20))
                            _ = v24
                        }
                        var v25 uint64 = 4294967295
                        _ = fi23
                        _ = v25
                    }
                }
            }
            _ = fi21
            _ = v22
        }
        _ = v19
        _ = v20
    }
    if (true || (!((^(v18)) == 65535))) {
        if (1.5 == -2.25) {
            fmt.Println(v12, v18, v12)
            switch "hello" {
            case "":
                var v26 uint = (v18 * (v18 + v18))
                _ = v26
            default:
                v27 := []int{genFn0(v12, "\n\t"), -299, 128}
                _ = v27
            }
            fmt.Println(v12, v18)
        } else {
            if false {
                if true {
                    var v28 string = "\n\t"
                    if false {
                        if ((false && (!false)) || false) {
                            if (((v18 << ((v18) & 63)) != (v18 * (v18 + (v18 + 654)))) && (false || ((^(v12)) < 811))) {
                                if (false && (v18 == 42)) {
                                    var v29 int16 = -293
                                    v30 := GenS2{F0: (v12 * v12), F1: 0, F2: 0}
                                    _ = v29
                                    _ = v30
                                } else {
                                    v31 := GenS3{F0: p0, F1: 1}
                                    var v32 int64 = -9223372036854775808
                                    _ = v31
                                    _ = v32
                                }
                                var v33 int = genFn0((v12 | (v12 >> (((v18 % ((v18) | 1))) & 63))), v28)
                                switch v28 {
                                case "a":
                                    v33 = (v33 ^ v33)
                                case "":
                                    v34 := []int{v33, v33, -878}
                                    _ = v34
                                }
                                _ = v33
                            }
                            v35 := []int{42}
                            _ = v35
                        }
                    }
                    v36 := []int{genFn0((v12 ^ uint32(v18)), ("a" + "foo bar")), -9223372036854775808}
                    _ = v28
                    _ = v36
                } else {
                    var v37 uint32 = (v12 >> ((uint(p0)) & 63))
                    var v38 uint16 = uint16(v18)
                    _ = v37
                    _ = v38
                }
            } else {
                v12 = (^(v12))
            }
            var v39 float32 = 1.5
            v40 := GenS4{F0: 42}
            _ = v39
            _ = v40
        }
        var v41 int8 = 127
        _ = v41
    }
    _ = p0
    _ = v12
    _ = v18
    return 18446744073709551615
}

func main() {
    v1 := make(map[int]int)
    var v2 string = "foo bar"
    var v3 uint8 = 127
    v4 := make(map[string]int)
    v5 := make(map[int]uint32)
    if true {
        fmt.Println(v2, v3)
        for fi42 := 0; fi42 < 13; fi42++ {
            v43 := GenS4{F0: -32768}
            _ = fi42
            _ = v43
        }
        switch 755 {
        case 128:
            for fi44 := 0; fi44 < 12; fi44++ {
                fmt.Println(v2)
                if (!((!false) || (42 >= 297))) {
                    var v45 uint16 = 0
                    for fi46 := 0; fi46 < 12; fi46++ {
                        switch v2 {
                        case "foo bar":
                            v45 = 946
                        }
                        _ = fi46
                    }
                    var v47 int32 = 820
                    _ = v45
                    _ = v47
                }
                var v48 string = v2
                _ = fi44
                _ = v48
            }
        case -1:
            fmt.Println(v2)
        case -2147483648:
            v49 := GenS5{F0: 0, F1: 128, F2: 1}
            _ = v49
        default:
            delete(v1, 32767)
        }
    }
    v3 += v3
    v2 = v2
    _ = len(v4)
    switch 119 {
    case 10:
        v50 := GenS3{F0: v3, F1: 1}
        _ = v50
    case -1:
        if true {
            if ((genFn1(v3) < 4294967295) || (false && (!false))) {
                if true {
                    if false {
                        v2 = (v2 + (v2 + "foo bar"))
                        var v51 int64 = 9223372036854775807
                        switch v2 {
                        case "a":
                            fmt.Println(v51)
                        case "foo bar":
                            var v52 uint64 = 42
                            _ = v52
                        }
                        _ = v51
                    } else {
                        fmt.Println(v2, v2)
                        v2 = (((v2 + v2) + (v2 + v2)) + v2)
                    }
                    var v53 bool = (-198 == -556)
                    _ = v53
                }
            }
            v2 = ("" + v2)
        } else {
            if false {
                if true {
                    var v54 int = 4294967295
                    fmt.Println(v3)
                    v54 = 2147483647
                    _ = v54
                } else {
                    switch "a" {
                    case "a":
                        v55 := GenS6{F0: (!((true && ((2 > 216) || true)) && true)), F1: 0}
                        _ = v55
                    case "":
                        v56 := GenS4{F0: 1}
                        _ = v56
                    default:
                        for fi57 := 0; fi57 < 2; fi57++ {
                            v58 := GenS4{F0: -840}
                            switch v2 {
                            case "foo bar":
                                if true {
                                    v3 -= 127
                                }
                            }
                            _ = fi57
                            _ = v58
                        }
                    }
                    var v59 uint = 2
                    _ = v59
                }
                for fi60 := 0; fi60 < 4; fi60++ {
                    v61 := GenS3{F0: (v3 / (((v3 % (((v3 | (v3 >> ((2) & 63)))) | 1))) | 1)), F1: -1}
                    v2 = v2
                    _ = fi60
                    _ = v61
                }
            } else {
                var v62 uint64 = 0
                v63 := GenS4{F0: 366}
                if true {
                    if (14 == 42) {
                        switch v2 {
                        case "":
                            if (false && (86 != 0)) {
                                switch ("\n\t" + v2) {
                                case "\n\t":
                                    v2 = v2
                                }
                            }
                        }
                        v64 := GenS7{F0: 591, F1: 2147483647}
                        _ = v64
                    } else {
                        if (128 > v3) {
                            var v65 uint8 = (v3 & 128)
                            delete(v4, v2)
                            v66 := GenS8{F0: 65535, F1: -8}
                            _ = v65
                            _ = v66
                        }
                        fmt.Println(v2, v2)
                        v67 := []int{-1}
                        _ = v67
                    }
                    v3 += 156
                    v3 = v3
                }
                _ = v62
                _ = v63
            }
            var v68 uint = 65535
            var v69 uint16 = 42
            _ = v68
            _ = v69
        }
    case -388:
        if (3.14159 == 0.1) {
            switch "hello" {
            case "\n\t":
                fmt.Println(v3)
            }
        } else {
            v5[-607] = 65535
            for fi70 := 0; fi70 < 19; fi70++ {
                v3 += (v3 ^ (v3 / (((v3 << ((0) & 63))) | 1)))
                fmt.Println(v3)
                var v71 string = "\n\t"
                _ = fi70
                _ = v71
            }
            v72 := GenS9{F0: 0, F1: 0}
            _ = v72
        }
    default:
        v73 := []int{}
        _ = v73
    }
    if (!(true && false)) {
        switch -732 {
        case 730:
            v3 = 2
        case -1:
            _ = len(v4)
        case -17:
            if true {
                v2 = "\n\t"
            }
        }
        for fi74 := 0; fi74 < 7; fi74++ {
            for fi75 := 0; fi75 < 4; fi75++ {
                switch genFn0(0, (v2 + "\n\t")) {
                case 127:
                    var v76 float32 = 0.0
                    _ = v76
                case -3:
                    var v77 string = "hello"
                    _ = v77
                }
                _ = fi75
            }
            _ = fi74
        }
    } else {
        var v78 float32 = 3.14159
        _ = v78
    }
    for fi79 := 0; fi79 < 11; fi79++ {
        v80 := GenS5{F0: 2, F1: genFn0(842, v2), F2: 258}
        delete(v1, (fi79 / (((fi79 >> ((18446744073709551615) & 63))) | 1)))
        v2 = ""
        _ = fi79
        _ = v80
    }
    var v81 int8 = 88
    switch -418 {
    case 1:
        v82 := GenS1{F0: -64, F1: -1}
        _ = v82
    }
    var v83 int64 = 104
    fmt.Println(v2, v81, v3)
    v84 := v4[v2]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v81:", v81)
    fmt.Println("v83:", v83)
    fmt.Println("v84:", v84)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v81
    _ = v83
    _ = v84
}
