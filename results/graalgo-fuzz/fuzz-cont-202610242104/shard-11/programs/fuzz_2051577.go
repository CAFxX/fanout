package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16 }
type GenS2 struct { F0 bool; F1 int; F2 int16 }
type GenS3 struct { F0 bool; F1 int64; F2 int8 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 uint32; F1 int; F2 int32 }
type GenS7 struct { F0 int; F1 uint32; F2 uint64 }
type GenS8 struct { F0 int; F1 int8; F2 string }
func genFn0(p0 uint64, p1 int64, p2 uint32) uint64 {
    for fi7 := 0; fi7 < 13; fi7++ {
        p2 = uint32(p0)
        for fi8 := 0; fi8 < 5; fi8++ {
            if (1 > p1) {
                if (564 == 32767) {
                    v9 := GenS1{F0: 120, F1: 255}
                    _ = v9
                } else {
                    var v10 int = 382
                    v10 -= 9223372036854775807
                    _ = v10
                }
                fmt.Println(fi8, fi7)
                p2 += p2
            }
            fmt.Println(p0, p1)
            _ = fi8
        }
        _ = fi7
    }
    var v11 int64 = p1
    v11 = (v11 >> ((42) & 63))
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return (p0 ^ (p0 / (((^(p0))) | 1)))
}

func main() {
    var v1 int = -239
    var v2 uint = 943
    v3 := make(map[int]bool)
    v4 := make(map[int]bool)
    v5 := make(map[string]int)
    var v6 int64 = -1
    switch v1 {
    case 1:
        for fi12 := 0; fi12 < 17; fi12++ {
            v5[("a" + ("a" + "foo bar"))] = fi12
            switch 1 {
            case 140:
                v6 -= 560
            case 1:
                var v13 int32 = 2147483647
                _ = v13
            case -9223372036854775808:
                if (true && true) {
                    v1 = (^(v1))
                    if ((0 > 1) || (438 > 2)) {
                        for fi14 := 0; fi14 < 12; fi14++ {
                            var v15 uint16 = 255
                            v16 := GenS2{F0: ((3.14159 + -2.25) != (0.1 * (-2.25 * ((1e10 / (0.1 + 0.1)) / 1.0)))), F1: v1, F2: 32767}
                            _ = fi14
                            _ = v15
                            _ = v16
                        }
                        if (true && (!(!false))) {
                            delete(v4, fi12)
                            var v17 uint = v2
                            _ = v17
                        } else {
                            v18 := GenS2{F0: ("foo bar" >= "a"), F1: 1, F2: 1000}
                            var v19 uint64 = 18446744073709551615
                            switch 9223372036854775807 {
                            case 127:
                                for fi20 := 0; fi20 < 7; fi20++ {
                                    if false {
                                        var v21 uint16 = 354
                                        fmt.Println(v2, v2, v19)
                                        v2 -= (v2 >> (((^(v2))) & 63))
                                        _ = v21
                                    } else {
                                        v22 := []int{-1}
                                        _ = v22
                                    }
                                    _ = fi20
                                }
                            default:
                                for fi23 := 0; fi23 < 17; fi23++ {
                                    switch "" {
                                    case "\n\t":
                                        if false {
                                            v24 := GenS2{F0: ((("" + "\n\t") + "a") != "hello"), F1: (-(v1)), F2: 1}
                                            fmt.Println(v19, fi12)
                                            var v25 string = ("" + "foo bar")
                                            _ = v24
                                            _ = v25
                                        } else {
                                            v1 -= 349
                                            var v26 int = (-(v1))
                                            _ = v26
                                        }
                                    default:
                                        v19 = v19
                                    }
                                    _ = fi23
                                }
                            }
                            _ = v18
                            _ = v19
                        }
                    }
                }
            default:
                v27 := GenS3{F0: ((0.0 * (0.0 / (float32(((0.1 + ((0.1 + (((float64(-2.25) - (0.1 - -2.25)) / (3.14159 + (3.14159 + (float64(0.1) + 3.14159)))) / 1e10)) * 1.5)) - 0.0)) + 1.5))) <= 0.1), F1: (v6 | 9223372036854775807), F2: 0}
                _ = v27
            }
            if (!true) {
                fmt.Println(v6, v2, v6)
                v1 = (-(fi12))
            } else {
                var v28 int16 = 42
                if true {
                    switch -9223372036854775808 {
                    case -2147483648:
                        v29 := v4[v1]
                        _ = v29
                    case 42:
                        var v30 uint64 = 476
                        _ = v30
                    }
                    v31 := GenS4{F0: false}
                    var v32 int8 = 1
                    _ = v31
                    _ = v32
                } else {
                    v28 = int16(v6)
                    var v33 uint = (v2 + 794)
                    var v34 uint32 = 1
                    _ = v33
                    _ = v34
                }
                fmt.Println(v2, v28)
                _ = v28
            }
            _ = fi12
        }
    default:
        v35 := GenS3{F0: (v6 >= (v6 + v6)), F1: v6, F2: -1}
        _ = v35
    }
    v36 := GenS5{F0: -1}
    fmt.Println(v6, v6, v2)
    if (128 <= 1) {
        var v37 string = "\n\t"
        _ = v37
    }
    v38 := []int{v1}
    fmt.Println(v1, v2)
    var v39 string = "foo bar"
    if false {
        if (!(1e10 < (-2.25 + ((3.14159 / 0.1) - (0.1 + 0.0))))) {
            if ((false || false) && true) {
                v1 -= -158
                fmt.Println(v1, v39, v2)
                v40 := GenS1{F0: -1, F1: 32767}
                _ = v40
            } else {
                var v41 string = ((v39 + (((v39 + "\n\t") + v39) + ("foo bar" + (v39 + (v39 + v39))))) + (((v39 + ((v39 + "foo bar") + v39)) + v39) + "hello"))
                _ = v41
            }
            for fi42 := 0; fi42 < 4; fi42++ {
                v2 += 65535
                v6 = v6
                _ = fi42
            }
            for fi43 := 0; fi43 < 10; fi43++ {
                var v44 float64 = ((3.14159 + 0.0) - (0.1 / 1e10))
                for fi45 := 0; fi45 < 3; fi45++ {
                    var v46 uint8 = 42
                    var v47 uint32 = 47
                    v44 = v44
                    var v48 int32 = -1
                    _ = fi45
                    _ = v46
                    _ = v47
                    _ = v48
                }
                if (2 == 1) {
                    v1 -= (fi43 % (((v1 * v1)) | 1))
                } else {
                    v38 = append(v38, v1)
                    if len(v38) >= 2 {
                        v38 = v38[1:len(v38)]
                    }
                }
                _ = fi43
                _ = v44
            }
        } else {
            var v49 int32 = -724
            _ = v49
        }
        if true {
            var v50 uint = (v2 / (((v2 ^ 18446744073709551615)) | 1))
            for fi51 := 0; fi51 < 15; fi51++ {
                v52 := GenS1{F0: -128, F1: 746}
                _ = fi51
                _ = v52
            }
            if true {
                var v53 int16 = -496
                v54 := v5[v39]
                v55 := GenS6{F0: 2147483647, F1: (^(v54)), F2: 0}
                _ = v53
                _ = v54
                _ = v55
            } else {
                v50 = v2
                for fi56 := 0; fi56 < 13; fi56++ {
                    _ = len(v38)
                    for fi57 := 0; fi57 < 17; fi57++ {
                        v58 := v5[v39]
                        _ = fi57
                        _ = v58
                    }
                    if true {
                        var v59 int64 = (v6 & (-(v6)))
                        _ = v59
                    }
                    _ = fi56
                }
                switch v1 {
                case 32767:
                    if ((v6 >> ((v50) & 63)) != v6) {
                        if true {
                            for fi60 := 0; fi60 < 15; fi60++ {
                                if len(v38) > 0 {
                                    v38[0] = (fi60 >> ((v50) & 63))
                                }
                                _ = fi60
                            }
                            v38 = append(v38, v1)
                            switch "a" {
                            case "\n\t":
                                v3[v1] = true
                            case "foo bar":
                                if ((!(127 > 127)) || true) {
                                    if (323 <= genFn0(2, 42, 383)) {
                                        var v61 int64 = v6
                                        var v62 uint16 = 65535
                                        v63 := GenS7{F0: v1, F1: 4294967295, F2: 4294967295}
                                        _ = v61
                                        _ = v62
                                        _ = v63
                                    }
                                    v64 := GenS4{F0: true}
                                    _ = v64
                                }
                            case "a":
                                for fi65 := 0; fi65 < 12; fi65++ {
                                    var v66 string = (v39 + "hello")
                                    if (v39 > "foo bar") {
                                        var v67 int64 = -829
                                        _ = v67
                                    }
                                    _ = fi65
                                    _ = v66
                                }
                            }
                        } else {
                            v1 -= 432
                        }
                        v6 -= v6
                        v1 = -669
                    }
                case 32768:
                    var v68 int = (-(v1))
                    _ = v68
                case 0:
                    var v69 bool = false
                    _ = v69
                }
            }
            _ = v50
        }
    }
    fmt.Println(v1)
    var v70 int16 = 777
    var v71 int16 = (v70 * 69)
    if true {
        v71 = v71
        v6 -= -318
        var v72 string = v39
        _ = v72
    }
    for fi73 := 0; fi73 < 15; fi73++ {
        v2 += (^(v2))
        var v74 uint64 = 1
        for fi75 := 0; fi75 < 15; fi75++ {
            v76 := GenS8{F0: fi73, F1: 100, F2: "\n\t"}
            if (2 != 612) {
                if false {
                    v77 := GenS1{F0: 42, F1: 42}
                    var v78 int32 = -1
                    _ = v77
                    _ = v78
                }
            }
            _ = len(v4)
            _ = fi75
            _ = v76
        }
        _ = fi73
        _ = v74
    }
    if false {
        v70 += 0
    } else {
        var v79 uint8 = 222
        _ = v79
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v38:", len(v38), v38)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v70:", v70)
    fmt.Println("v71:", v71)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v6
    _ = v70
    _ = v71
}
