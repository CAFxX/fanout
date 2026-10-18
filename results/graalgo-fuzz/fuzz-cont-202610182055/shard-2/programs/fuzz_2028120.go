package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint16; F1 int16; F2 int }
type GenS4 struct { F0 int32; F1 uint32 }
type GenS5 struct { F0 int32; F1 uint }
type GenS6 struct { F0 int16; F1 bool; F2 uint16 }
type GenS7 struct { F0 uint16; F1 int16 }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 uint; F1 int }
type GenS10 struct { F0 bool; F1 uint64 }
type GenS11 struct { F0 int32 }
type GenS12 struct { F0 int32; F1 int64 }
func genFn0(p0 string, p1 int64, p2 uint32) int32 {
    v4 := GenS1{F0: 249}
    v5 := GenS2{F0: 255}
    p1 = -9223372036854775808
    var v6 int16 = 1000
    v7 := GenS2{F0: 255}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return int32(p2)
}

func main() {
    var v1 uint = 0
    v2 := make(map[string]int)
    var v3 uint32 = 428
    for fi8 := 0; fi8 < 11; fi8++ {
        if false {
            switch ("" + "hello") {
            case "foo bar":
                if ((1e10 + 3.14159) >= 1e10) {
                    if true {
                        v9 := GenS3{F0: 876, F1: -356, F2: (fi8 >> (((v1 - (^(v1)))) & 63))}
                        if false {
                            v3 -= v3
                        } else {
                            _ = len(v2)
                        }
                        if true {
                            if (118 >= 1) {
                                var v10 uint = v1
                                _ = v10
                            }
                            if true {
                                for fi11 := 0; fi11 < 16; fi11++ {
                                    switch fi8 {
                                    case 0:
                                        var v12 float64 = 1.5
                                        _ = v12
                                    default:
                                        v13 := v2["\n\t"]
                                        _ = v13
                                    }
                                    v2["foo bar"] = -50
                                    v14 := []int{2147483647}
                                    _ = fi11
                                    _ = v14
                                }
                                for fi15 := 0; fi15 < 16; fi15++ {
                                    v16 := GenS2{F0: v1}
                                    if false {
                                        fmt.Println(v1, fi8, fi15)
                                    }
                                    fmt.Println(v3, v1, v3)
                                    _ = fi15
                                    _ = v16
                                }
                            }
                            if (("a" + ("a" + ("" + ""))) == ("hello" + "")) {
                                if true {
                                    var v17 int64 = 42
                                    _ = v17
                                }
                                for fi18 := 0; fi18 < 2; fi18++ {
                                    _ = len(v2)
                                    var v19 int16 = 1
                                    _ = fi18
                                    _ = v19
                                }
                            }
                        }
                        _ = v9
                    }
                    v20 := GenS2{F0: 1}
                    for fi21 := 0; fi21 < 4; fi21++ {
                        if true {
                            if false {
                                switch fi8 {
                                case 4294967295:
                                    if true {
                                        _ = len(v2)
                                        v22 := []int{-1, -397, (fi21 / ((fi8) | 1)), fi21}
                                        var v23 float64 = 3.14159
                                        _ = v22
                                        _ = v23
                                    }
                                case 32767:
                                    fmt.Println(fi21)
                                default:
                                    fmt.Println(fi8)
                                }
                                var v24 uint16 = 65535
                                v3 = v3
                                v25 := GenS4{F0: genFn0(("" + "\n\t"), 1, (v3 ^ v3)), F1: uint32(v1)}
                                fmt.Println(fi21)
                                _ = v24
                                _ = v25
                            }
                            v1 += (v1 ^ v1)
                            var v26 uint16 = 32767
                            v27 := GenS1{F0: v1}
                            var v28 uint16 = (v26 % (((v26 ^ (^(v26)))) | 1))
                            if true {
                                v1 = (v1 & (v1 * v1))
                                var v29 string = "hello"
                                var v30 int8 = 42
                                _ = v29
                                _ = v30
                            }
                            var v31 uint32 = uint32(v1)
                            if true {
                                switch (fi21 >> (((v1 & (v1 - v1))) & 63)) {
                                case 32768:
                                    var v32 string = ""
                                    _ = v32
                                default:
                                    fmt.Println(v28)
                                }
                                v28 = 42
                                var v33 int8 = 1
                                _ = v33
                            } else {
                                v3 = v31
                                _ = len(v2)
                            }
                            _ = v26
                            _ = v27
                            _ = v28
                            _ = v31
                        }
                        _ = fi21
                    }
                    _ = v20
                } else {
                    v34 := []int{-9223372036854775808, (fi8 % ((-433) | 1))}
                    v35 := GenS5{F0: genFn0(("hello" + "\n\t"), -9223372036854775808, (v3 << ((v1) & 63))), F1: v1}
                    _ = v34
                    _ = v35
                }
            case "hello":
                v3 = (v3 >> (((v1 & v1)) & 63))
            default:
                if (66 <= 42) {
                    var v36 int8 = int8(fi8)
                    for fi37 := 0; fi37 < 15; fi37++ {
                        switch "hello" {
                        case "hello":
                            if ((9223372036854775807 < 42) && true) {
                                v2[("hello" + ("" + ("\n\t" + "foo bar")))] = 0
                            }
                        }
                        v36 = v36
                        _ = fi37
                    }
                    _ = v36
                } else {
                    v1 -= (v1 - (v1 << ((v1) & 63)))
                    v1 -= v1
                    if (-19 >= 1) {
                        var v38 float64 = 0.0
                        var v39 int = 2147483647
                        _ = v38
                        _ = v39
                    }
                }
            }
            _ = len(v2)
            for fi40 := 0; fi40 < 20; fi40++ {
                switch -1 {
                case 128:
                    v3 += (v3 << ((18446744073709551615) & 63))
                default:
                    if false {
                        v3 += uint32(v1)
                    } else {
                        v41 := []int{(fi40 << ((602) & 63)), -9223372036854775808, -2147483648, (fi40 * fi40)}
                        v3 -= (v3 >> ((v1) & 63))
                        var v42 uint = v1
                        fmt.Println(v1)
                        v2["a"] = fi8
                        _ = v41
                        _ = v42
                    }
                }
                _ = fi40
            }
        } else {
            for fi43 := 0; fi43 < 15; fi43++ {
                v1 -= v1
                fmt.Println(fi8)
                if false {
                    v44 := []int{}
                    var v45 uint32 = v3
                    var v46 uint8 = 128
                    _ = v44
                    _ = v45
                    _ = v46
                } else {
                    if (!(genFn0("", -9223372036854775808, (^(v3))) <= int32(fi8))) {
                        switch (fi8 ^ fi8) {
                        case -1:
                            switch ("hello" + "\n\t") {
                            case "a":
                                v3 = v3
                            default:
                                switch "hello" {
                                case "":
                                    _ = len(v2)
                                default:
                                    v47 := GenS3{F0: 909, F1: 102, F2: (fi8 | -1)}
                                    _ = v47
                                }
                            }
                        case 32767:
                            switch 9223372036854775807 {
                            case -2147483648:
                                if true {
                                    v2[("a" + "\n\t")] = fi8
                                    fmt.Println(v3, fi43)
                                } else {
                                    if (127 != 98) {
                                        var v48 string = "a"
                                        fmt.Println(fi43)
                                        _ = v48
                                    } else {
                                        if (float64((-2.25 + 3.14159)) < 1e10) {
                                            v49 := GenS6{F0: 180, F1: ((!true) && false), F2: 819}
                                            _ = v49
                                        } else {
                                            v50 := []int{fi8, fi8, fi8}
                                            v1 -= (^(v1))
                                            _ = v50
                                        }
                                    }
                                }
                                fmt.Println(v1, fi8)
                                if false {
                                    var v51 float32 = (((1.5 / float32(float64(-2.25))) - 0.0) + (float32((0.1 * 0.1)) / 1.0))
                                    v52 := []int{fi8, (fi43 | (^(fi43)))}
                                    v3 -= 2147483648
                                    _ = v51
                                    _ = v52
                                } else {
                                    fmt.Println(fi8, v3)
                                }
                            case 1:
                                var v53 uint64 = 4294967295
                                _ = v53
                            default:
                                v1 += v1
                            }
                        case 2147483647:
                            switch -261 {
                            case 127:
                                v3 = (^(v3))
                                v3 -= (v3 & (v3 ^ 0))
                            case 9223372036854775807:
                                if true {
                                    if true {
                                        switch ("\n\t" + "foo bar") {
                                        case "\n\t":
                                            if (!(true && false)) {
                                                v54 := GenS1{F0: v1}
                                                v3 += (^(v3))
                                                v1 -= v1
                                                if (fi43 > (fi8 + (fi43 % (((fi43 << ((v1) & 63))) | 1)))) {
                                                    v1 = 2
                                                } else {
                                                    v3 = 769
                                                    v3 = v3
                                                    if false {
                                                        var v55 float64 = (0.0 - 1.5)
                                                        v56 := v2[""]
                                                        _ = v55
                                                        _ = v56
                                                    }
                                                }
                                                var v57 int64 = 9223372036854775807
                                                _ = v54
                                                _ = v57
                                            }
                                        case "a":
                                            v1 = v1
                                        }
                                        fmt.Println(v1)
                                        if false {
                                            v1 = (v1 * (^(v1)))
                                            v1 -= v1
                                            switch int(v3) {
                                            case 84:
                                                v58 := GenS7{F0: 42, F1: 0}
                                                _ = v58
                                            default:
                                                v3 = v3
                                            }
                                        } else {
                                            fmt.Println(v3, fi43, v3)
                                        }
                                    }
                                }
                            case 1:
                                v59 := GenS8{F0: -1}
                                _ = v59
                            default:
                                var v60 bool = true
                                _ = v60
                            }
                        }
                        v61 := GenS9{F0: 0, F1: (fi43 / ((fi8) | 1))}
                        v1 = (v1 % ((v1) | 1))
                        _ = v61
                    }
                    if true {
                        v62 := GenS6{F0: 42, F1: (674 >= 65535), F2: 0}
                        fmt.Println(v3, v1, fi8)
                        fmt.Println(v3, v3, fi43)
                        _ = v62
                    } else {
                        switch (fi43 << ((971) & 63)) {
                        case 865:
                            fmt.Println(v1, fi8, fi8)
                            v63 := GenS10{F0: true, F1: 2}
                            if true {
                                v64 := []int{int(v3)}
                                v3 = 2147483648
                                _ = v64
                            }
                            fmt.Println(v1)
                            v65 := []int{fi43, fi8, fi43}
                            _ = v63
                            _ = v65
                        }
                    }
                    var v66 uint16 = 2
                    _ = v66
                }
                _ = fi43
            }
            for fi67 := 0; fi67 < 1; fi67++ {
                var v68 int16 = 0
                var v69 uint64 = 627
                switch ("" + (("foo bar" + ("\n\t" + "foo bar")) + "\n\t")) {
                case "\n\t":
                    switch "hello" {
                    case "hello":
                        var v70 bool = (-128 <= 127)
                        _ = v70
                    case "":
                        fmt.Println(v68, fi8, v3)
                    default:
                        var v71 float64 = float64(float32(3.14159))
                        _ = v71
                    }
                }
                v72 := []int{}
                v1 += v1
                var v73 uint16 = 2
                var v74 uint8 = 1
                v75 := GenS5{F0: 42, F1: 255}
                v1 += uint(v74)
                _ = fi67
                _ = v68
                _ = v69
                _ = v72
                _ = v73
                _ = v74
                _ = v75
            }
        }
        var v76 float32 = 0.0
        _ = fi8
        _ = v76
    }
    v77 := GenS11{F0: int32(v1)}
    v78 := []int{32768, -1, 128, 377}
    if (!((false && false) && (!(!true)))) {
        var v79 float32 = 0.0
        fmt.Println(v79, v3, v79)
        switch -683 {
        case 2147483647:
            v80 := GenS7{F0: 0, F1: 1}
            _ = v80
        default:
            var v81 uint32 = (^(v3))
            _ = v81
        }
        _ = v79
    }
    var v82 int32 = -2147483648
    var v83 uint64 = 0
    for fi84 := 0; fi84 < 8; fi84++ {
        v85 := GenS7{F0: 1, F1: -768}
        fmt.Println(v1, fi84, v3)
        _ = fi84
        _ = v85
    }
    v3 = 685
    var v86 int = -582
    fmt.Println(v86, v83, v1)
    fmt.Println(v1)
    for fi87 := 0; fi87 < 12; fi87++ {
        var v88 float64 = 0.1
        _ = fi87
        _ = v88
    }
    var v89 uint16 = 151
    var v90 uint8 = 0
    v91 := GenS12{F0: genFn0("hello", -532, (v3 % ((168) | 1))), F1: int64(v3)}
    v86 -= -719
    _ = len(v78)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v78:", len(v78), v78)
    fmt.Println("v82:", v82)
    fmt.Println("v83:", v83)
    fmt.Println("v86:", v86)
    fmt.Println("v89:", v89)
    fmt.Println("v90:", v90)
    _ = v1
    _ = v2
    _ = v3
    _ = v77
    _ = v78
    _ = v82
    _ = v83
    _ = v86
    _ = v89
    _ = v90
    _ = v91
}
