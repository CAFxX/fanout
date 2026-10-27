package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int16; F1 bool; F2 int64 }
type GenS4 struct { F0 uint64; F1 int16 }
type GenS5 struct { F0 uint }
func genFn0(p0 uint8, p1 uint64, p2 uint16) string {
    var v4 int32 = -1
    v5 := []int{9223372036854775807, 841, -210, 128}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return ("hello" + "\n\t")
}
func genFn1() int8 {
    var v6 uint32 = 2147483648
    var v7 float64 = 3.14159
    if (13 >= 2) {
        v7 = v7
        if false {
            for fi8 := 0; fi8 < 6; fi8++ {
                for fi9 := 0; fi9 < 20; fi9++ {
                    v10 := []int{(fi9 * 127), (fi9 ^ (-(fi8))), (fi8 % ((32768) | 1)), (^(fi9))}
                    _ = fi9
                    _ = v10
                }
                var v11 uint16 = 0
                _ = fi8
                _ = v11
            }
            v6 = (v6 & 1)
            v7 -= v7
        } else {
            v12 := GenS1{F0: v6}
            switch ("" + (("a" + "a") + "hello")) {
            case "a":
                v13 := GenS1{F0: v6}
                _ = v13
            case "hello":
                v14 := []int{-2147483648, 723, 42}
                _ = v14
            case "":
                for fi15 := 0; fi15 < 13; fi15++ {
                    for fi16 := 0; fi16 < 9; fi16++ {
                        if false {
                            if true {
                                switch fi15 {
                                case 213:
                                    v17 := GenS1{F0: (v6 ^ (v6 ^ v6))}
                                    _ = v17
                                default:
                                    var v18 int64 = -9223372036854775808
                                    _ = v18
                                }
                            }
                        } else {
                            if (305 > 1) {
                                v19 := GenS1{F0: 752}
                                v20 := GenS2{F0: 42}
                                _ = v19
                                _ = v20
                            }
                            if false {
                                v7 -= 3.14159
                            }
                        }
                        fmt.Println(v6, fi16, v7)
                        v21 := GenS2{F0: 1}
                        _ = fi16
                        _ = v21
                    }
                    if (true || false) {
                        v22 := GenS3{F0: 42, F1: false, F2: 0}
                        for fi23 := 0; fi23 < 2; fi23++ {
                            v24 := []int{(-(fi15)), fi15, 2147483647}
                            var v25 uint16 = 42
                            var v26 int32 = -2147483648
                            _ = fi23
                            _ = v24
                            _ = v25
                            _ = v26
                        }
                        _ = v22
                    } else {
                        v27 := []int{fi15, fi15}
                        var v28 bool = (v6 == (^(v6)))
                        _ = v27
                        _ = v28
                    }
                    _ = fi15
                }
            default:
                v6 += (^(v6))
            }
            _ = v12
        }
    } else {
        fmt.Println(v7, v6)
    }
    _ = v6
    _ = v7
    return 14
}

func main() {
    v1 := make(map[string]int)
    var v2 uint = 65535
    var v3 uint64 = 0
    for fi29 := 0; fi29 < 7; fi29++ {
        for fi30 := 0; fi30 < 15; fi30++ {
            var v31 uint32 = 1
            _ = fi30
            _ = v31
        }
        if true {
            if (int16(fi29) == 1000) {
                var v32 float64 = 1.5
                v33 := GenS4{F0: (^(v3)), F1: -582}
                _ = v32
                _ = v33
            } else {
                v2 = uint(v3)
                v34 := GenS4{F0: 0, F1: -558}
                for fi35 := 0; fi35 < 2; fi35++ {
                    delete(v1, "a")
                    if true {
                        if ((("a" + "foo bar") + "") == "\n\t") {
                            fmt.Println(v3, fi35, fi35)
                        } else {
                            var v36 int32 = 0
                            var v37 int = (fi35 | (fi29 ^ (fi35 / (((fi29 + (-(fi29)))) | 1))))
                            v37 -= (fi29 + (fi35 * (v37 >> ((0) & 63))))
                            _ = v36
                            _ = v37
                        }
                    }
                    var v38 float64 = 0.1
                    _ = fi35
                    _ = v38
                }
                _ = v34
            }
            v3 = v3
        }
        _ = fi29
    }
    var v39 int8 = -22
    var v40 int32 = -284
    switch "a" {
    case "":
        var v41 int64 = 0
        _ = v41
    case "foo bar":
        if true {
            v2 = 65535
            if false {
                for fi42 := 0; fi42 < 11; fi42++ {
                    var v43 int8 = (^(v39))
                    v3 = v3
                    _ = fi42
                    _ = v43
                }
                v44 := GenS5{F0: 18446744073709551615}
                fmt.Println(v39, v39)
                _ = v44
            }
        }
    case "hello":
        var v45 uint8 = 43
        _ = v45
    }
    fmt.Println(v3, v3)
    var v46 int32 = v40
    v3 = 327
    switch "a" {
    case "foo bar":
        var v47 uint8 = 255
        _ = v47
    default:
        v48 := []int{-37, -681}
        _ = v48
    }
    var v49 int16 = 287
    v50 := GenS4{F0: 4294967295, F1: v49}
    var v51 int16 = v49
    var v52 uint8 = 114
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v40:", v40)
    fmt.Println("v46:", v46)
    fmt.Println("v49:", v49)
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v40
    _ = v46
    _ = v49
    _ = v50
    _ = v51
    _ = v52
}
