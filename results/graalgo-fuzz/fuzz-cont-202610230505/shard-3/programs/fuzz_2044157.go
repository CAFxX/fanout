package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8; F2 uint }
type GenS2 struct { F0 uint16; F1 string }
type GenS3 struct { F0 bool; F1 uint32; F2 uint64 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 bool; F1 uint; F2 int8 }
type GenS6 struct { F0 uint32; F1 int64 }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 string; F1 uint16; F2 uint }
func genFn0(p0 int8, p1 string, p2 string) uint64 {
    p1 = "a"
    var v6 uint64 = 547
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return v6
}
func genFn1(p0 uint16, p1 uint8) int {
    var v7 float64 = 1e10
    v7 -= v7
    if false {
        if true {
            var v8 uint = 989
            switch 0 {
            case 127:
                switch "foo bar" {
                case "\n\t":
                    v9 := []int{219, 4294967295}
                    _ = v9
                case "foo bar":
                    v10 := []int{-1, 512, 486, 1}
                    _ = v10
                default:
                    v11 := GenS1{F0: -9223372036854775808, F1: (p1 >> (((v8 / ((18446744073709551615) | 1))) & 63)), F2: (v8 - (v8 / (((v8 ^ v8)) | 1)))}
                    _ = v11
                }
            case -50:
                v12 := []int{1}
                _ = v12
            case -9223372036854775808:
                switch "" {
                case "foo bar":
                    v13 := GenS1{F0: 1, F1: 128, F2: 1}
                    _ = v13
                case "a":
                    for fi14 := 0; fi14 < 5; fi14++ {
                        switch "\n\t" {
                        case "foo bar":
                            var v15 float64 = v7
                            _ = v15
                        case "a":
                            switch fi14 {
                            case 457:
                                var v16 int8 = -1
                                _ = v16
                            default:
                                for fi17 := 0; fi17 < 14; fi17++ {
                                    p1 -= (p1 - 128)
                                    p0 += 42
                                    _ = fi17
                                }
                            }
                        case "\n\t":
                            switch "a" {
                            case "hello":
                                for fi18 := 0; fi18 < 19; fi18++ {
                                    v19 := GenS2{F0: (p0 << ((v8) & 63)), F1: "\n\t"}
                                    var v20 int32 = 656
                                    _ = fi18
                                    _ = v19
                                    _ = v20
                                }
                            case "a":
                                for fi21 := 0; fi21 < 16; fi21++ {
                                    v22 := GenS2{F0: (p0 * p0), F1: "\n\t"}
                                    v23 := GenS3{F0: ((fi14 / (((fi21 & 32768)) | 1)) == (fi21 & fi14)), F1: 65535, F2: 963}
                                    _ = fi21
                                    _ = v22
                                    _ = v23
                                }
                            case "foo bar":
                                v24 := []int{-2147483648, 829, fi14}
                                _ = v24
                            }
                        default:
                            v25 := []int{(fi14 << (((v8 / (((v8 * v8)) | 1))) & 63)), (fi14 | -173)}
                            _ = v25
                        }
                        fmt.Println(p1, p0)
                        var v26 int16 = 509
                        _ = fi14
                        _ = v26
                    }
                default:
                    fmt.Println(p0, v8, v7)
                }
            }
            var v27 int8 = -14
            _ = v27
            _ = v8
        } else {
            var v28 uint32 = 2147483648
            var v29 int16 = -792
            if (42 == 0) {
                for fi30 := 0; fi30 < 14; fi30++ {
                    p1 = p1
                    v31 := GenS4{F0: 42}
                    _ = fi30
                    _ = v31
                }
                v29 = 805
            }
            _ = v28
            _ = v29
        }
        for fi32 := 0; fi32 < 18; fi32++ {
            for fi33 := 0; fi33 < 12; fi33++ {
                v34 := []int{4294967295, fi33, 32768, fi32}
                p0 = (p0 + (p0 | (p0 >> ((734) & 63))))
                _ = fi33
                _ = v34
            }
            fmt.Println(v7)
            _ = fi32
        }
        for fi35 := 0; fi35 < 8; fi35++ {
            if (-486 == 42) {
                p0 += (p0 + (p0 % ((p0) | 1)))
                v36 := GenS5{F0: true, F1: 2, F2: 71}
                p1 += p1
                _ = v36
            }
            _ = fi35
        }
    }
    v37 := GenS6{F0: 2147483648, F1: 0}
    v38 := GenS6{F0: 2, F1: 1}
    _ = p0
    _ = p1
    _ = v37
    _ = v38
    _ = v7
    return 4294967295
}

func main() {
    var v1 uint = 0
    var v2 float64 = (float64(1.5) / (float64(1e10) + 0.1))
    v3 := []uint32{751}
    var v4 uint = (v1 & 0)
    var v5 float64 = ((v2 / float64(1e10)) - v2)
    v3 = append(v3, 2)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v2 += (v2 * v5)
    for fi39 := 0; fi39 < 20; fi39++ {
        switch (fi39 ^ (fi39 % ((895) | 1))) {
        case -2147483648:
            switch 9223372036854775807 {
            case -9223372036854775808:
                v1 += (v4 * (^(v4)))
            default:
                v4 += (v1 % (((^(v4))) | 1))
            }
        case 32768:
            fmt.Println(v1, fi39, v5)
        case -1:
            if false {
                switch fi39 {
                case 296:
                    var v40 float64 = v2
                    _ = v40
                case -209:
                    fmt.Println(v1, v5, v4)
                default:
                    v5 = (1.5 / float64(((0.1 * -2.25) / 1e10)))
                }
                if true {
                    for fi41 := 0; fi41 < 16; fi41++ {
                        v4 = (v4 - v1)
                        var v42 int16 = -945
                        var v43 uint16 = 42
                        _ = fi41
                        _ = v42
                        _ = v43
                    }
                    switch (fi39 * fi39) {
                    case -877:
                        switch fi39 {
                        case 1:
                            fmt.Println(v5, v5, fi39)
                        }
                    case 2147483647:
                        var v44 uint64 = 0
                        _ = v44
                    case -1:
                        for fi45 := 0; fi45 < 11; fi45++ {
                            v2 = v5
                            var v46 int16 = -1
                            fmt.Println(v4, v2, v1)
                            _ = fi45
                            _ = v46
                        }
                    default:
                        var v47 uint64 = 1
                        _ = v47
                    }
                } else {
                    v2 = v2
                }
                v48 := GenS3{F0: false, F1: 37, F2: 82}
                _ = v48
            }
        }
        _ = fi39
    }
    v49 := GenS7{F0: 2}
    v50 := GenS8{F0: ((("" + (("" + "a") + ("foo bar" + "foo bar"))) + "") + ("\n\t" + ("a" + ""))), F1: 0, F2: (v1 * 783)}
    var v51 uint32 = 42
    var v52 int64 = 0
    var v53 uint16 = 505
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v50
    _ = v51
    _ = v52
    _ = v53
}
