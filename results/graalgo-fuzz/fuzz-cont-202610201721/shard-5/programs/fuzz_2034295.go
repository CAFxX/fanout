package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int8; F1 uint32; F2 uint16 }
type GenS4 struct { F0 int64; F1 uint32 }
type GenS5 struct { F0 uint8; F1 int64 }
func genFn0(p0 uint64) int32 {
    var v4 uint32 = 42
    for fi5 := 0; fi5 < 12; fi5++ {
        switch (^(fi5)) {
        case 127:
            fmt.Println(p0, p0)
        default:
            v6 := []int{-201, (fi5 * fi5), (fi5 - fi5), (^(fi5))}
            _ = v6
        }
        var v7 int64 = -9223372036854775808
        for fi8 := 0; fi8 < 6; fi8++ {
            v4 = v4
            _ = fi8
        }
        _ = fi5
        _ = v7
    }
    _ = p0
    _ = v4
    return 1
}
func genFn1(p0 uint32) int64 {
    v9 := GenS1{F0: "foo bar", F1: p0}
    fmt.Println(p0)
    _ = p0
    _ = v9
    return -317
}

func main() {
    v1 := []uint32{2147483648, 2147483648}
    var v2 int = -707
    v3 := make(map[int]string)
    v10 := GenS2{F0: 47}
    v1 = append(v1, 0)
    for fi11 := 0; fi11 < 15; fi11++ {
        var v12 float64 = ((-2.25 + (1e10 / 1.0)) / float64(1e10))
        var v13 uint16 = 163
        for fi14 := 0; fi14 < 12; fi14++ {
            if (true && false) {
                fmt.Println(v13)
            } else {
                var v15 uint64 = 2
                v13 = v13
                _ = v15
            }
            fmt.Println(fi14, v13)
            _ = fi14
        }
        _ = fi11
        _ = v12
        _ = v13
    }
    fmt.Println(v2)
    v16 := GenS2{F0: 2}
    v2 = v2
    var v17 uint8 = 1
    var v18 float32 = -2.25
    var v19 float64 = 1e10
    switch "" {
    case "":
        for fi20 := 0; fi20 < 8; fi20++ {
            var v21 uint8 = v17
            _ = fi20
            _ = v21
        }
    default:
        for fi22 := 0; fi22 < 3; fi22++ {
            switch "\n\t" {
            case "foo bar":
                switch (fi22 & -9223372036854775808) {
                case 127:
                    var v23 int32 = 42
                    _ = v23
                }
            case "a":
                if (((false || true) && (!(1 < 18446744073709551615))) || (((float64(-2.25) / (v19 * v19)) >= 0.0) && false)) {
                    var v24 int8 = -1
                    _ = v24
                } else {
                    fmt.Println(fi22, fi22)
                    fmt.Println(v19, v17, fi22)
                    if false {
                        for fi25 := 0; fi25 < 20; fi25++ {
                            if false {
                                if (42 > 0) {
                                    _ = len(v3)
                                    var v26 int64 = 9223372036854775807
                                    var v27 uint8 = 128
                                    if true {
                                        v1 = append(v1, 65535)
                                    }
                                    _ = v26
                                    _ = v27
                                } else {
                                    fmt.Println(fi25)
                                    v28 := GenS3{F0: 42, F1: 289, F2: 42}
                                    _ = v28
                                }
                                v1 = append(v1, 538)
                                if (978 >= 1) {
                                    if true {
                                        if true {
                                            switch -258 {
                                            case -2147483648:
                                                delete(v3, (fi22 % ((v2) | 1)))
                                            case -319:
                                                v29 := GenS3{F0: 127, F1: 2147483647, F2: 459}
                                                _ = v29
                                            default:
                                                v19 = float64((v18 - v18))
                                            }
                                            v30 := GenS3{F0: -128, F1: 2147483647, F2: 295}
                                            _ = v30
                                        } else {
                                            switch (v2 ^ fi22) {
                                            case 332:
                                                switch ("a" + ("a" + (("\n\t" + "foo bar") + "hello"))) {
                                                case "\n\t":
                                                    var v31 int16 = 741
                                                    _ = v31
                                                case "foo bar":
                                                    v1 = append(v1, 65535)
                                                case "hello":
                                                    var v32 int32 = genFn0(1)
                                                    _ = v32
                                                default:
                                                    v17 = (^(v17))
                                                }
                                            }
                                            v2 += (fi25 / ((0) | 1))
                                            v18 = (0.1 + 3.14159)
                                        }
                                    } else {
                                        var v33 uint8 = v17
                                        v19 += 1e10
                                        var v34 int64 = genFn1(576)
                                        _ = v33
                                        _ = v34
                                    }
                                    var v35 uint8 = v17
                                    var v36 string = ""
                                    _ = v35
                                    _ = v36
                                } else {
                                    if (265 < 808) {
                                        var v37 float32 = v18
                                        if true {
                                            v38 := GenS4{F0: -539, F1: 65535}
                                            _ = v38
                                        }
                                        _ = v37
                                    } else {
                                        v17 = v17
                                    }
                                    v39 := GenS5{F0: 127, F1: -9223372036854775808}
                                    v40 := GenS5{F0: v17, F1: 42}
                                    fmt.Println(fi22)
                                    _ = v39
                                    _ = v40
                                }
                            }
                            _ = fi25
                        }
                        v2 += (-(fi22))
                    } else {
                        fmt.Println(v2)
                        var v41 bool = (857 >= 0)
                        _ = v41
                    }
                }
            default:
                v17 = 127
            }
            var v42 float32 = 0.0
            if (379 > 78) {
                var v43 string = ("a" + "foo bar")
                v44 := GenS2{F0: 2}
                _ = len(v3)
                _ = v43
                _ = v44
            } else {
                v42 = v42
                _ = len(v3)
            }
            _ = fi22
            _ = v42
        }
    }
    v17 = v17
    v19 = float64((float32(0.1) / v18))
    v17 = 48
    if (!true) {
        fmt.Println(v17, v18, v18)
    }
    v1 = append(v1, 898)
    fmt.Println(v18, v18)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v10
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
}
