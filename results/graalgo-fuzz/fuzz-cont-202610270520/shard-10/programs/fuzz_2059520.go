package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint64; F1 int16 }
func genFn0() int64 {
    var v7 uint = 4294967295
    for fi8 := 0; fi8 < 5; fi8++ {
        var v9 bool = (-1 < 1)
        v10 := []int{(fi8 + -1), (-(fi8)), (^(fi8)), int(v7)}
        switch -2147483648 {
        case 128:
            v9 = v9
        case 1:
            v11 := GenS1{F0: 858, F1: (v7 ^ (v7 | v7))}
            _ = v11
        case -1:
            for fi12 := 0; fi12 < 14; fi12++ {
                fmt.Println(v9, fi8)
                _ = fi12
            }
        }
        _ = fi8
        _ = v10
        _ = v9
    }
    for fi13 := 0; fi13 < 5; fi13++ {
        for fi14 := 0; fi14 < 12; fi14++ {
            v15 := GenS2{F0: 0}
            var v16 int16 = -32768
            switch fi13 {
            case 42:
                switch "foo bar" {
                case "":
                    v7 += v7
                    v17 := []int{}
                    if false {
                        if ((2 <= 2147483647) && true) {
                            switch "foo bar" {
                            case "\n\t":
                                if (true && (-478 <= 1)) {
                                    v18 := GenS1{F0: 143, F1: (v7 % (((v7 % (((v7 + 860)) | 1))) | 1))}
                                    var v19 uint32 = 42
                                    _ = v18
                                    _ = v19
                                }
                            case "a":
                                if len(v17) > 0 {
                                    v17[0] = -722
                                }
                            }
                            var v20 string = "foo bar"
                            if len(v17) > 0 {
                                v17[0] = -414
                            }
                            _ = v20
                        }
                        v7 = v7
                    } else {
                        if len(v17) >= 2 {
                            v17 = v17[1:len(v17)]
                        }
                        if (39 == 127) {
                            if len(v17) >= 2 {
                                v17 = v17[1:len(v17)]
                            }
                            v21 := GenS1{F0: 65535, F1: (v7 % (((v7 / (((^(v7))) | 1))) | 1))}
                            v16 = v16
                            if len(v17) > 0 {
                                v17[0] = fi13
                            }
                            fmt.Println(v16, v7)
                            _ = v21
                        }
                        switch (fi14 * fi14) {
                        case -603:
                            var v22 string = ""
                            _ = v22
                        }
                    }
                    _ = v17
                default:
                    var v23 uint32 = 0
                    _ = v23
                }
            case 2147483647:
                if false {
                    switch "a" {
                    case "":
                        v16 = v16
                        var v24 float64 = 1.5
                        var v25 string = "a"
                        _ = v24
                        _ = v25
                    case "foo bar":
                        var v26 int = (fi13 ^ (fi14 ^ 2147483647))
                        _ = v26
                    case "\n\t":
                        if false {
                            var v27 int8 = 42
                            v28 := []int{}
                            var v29 int64 = 0
                            _ = v27
                            _ = v28
                            _ = v29
                        } else {
                            v7 = 0
                            v30 := GenS1{F0: 2, F1: (v7 * v7)}
                            v7 += (v7 / (((v7 / ((v7) | 1))) | 1))
                            _ = v30
                        }
                    default:
                        var v31 float64 = -2.25
                        _ = v31
                    }
                    if false {
                        var v32 int32 = -1
                        v33 := []int{}
                        v34 := GenS3{F0: 4294967295}
                        _ = v32
                        _ = v33
                        _ = v34
                    }
                } else {
                    var v35 int = (fi13 - fi14)
                    switch "hello" {
                    case "foo bar":
                        var v36 uint64 = 4294967295
                        _ = v36
                    case "":
                        v7 += v7
                    case "\n\t":
                        v37 := []int{9223372036854775807, fi14, (fi13 ^ (-(fi14))), (^(v35))}
                        _ = v37
                    default:
                        v38 := GenS2{F0: 687}
                        _ = v38
                    }
                    _ = v35
                }
            default:
                if (255 == 430) {
                    if (!(!false)) {
                        var v39 bool = true
                        var v40 float64 = ((-2.25 + 3.14159) - 0.1)
                        _ = v39
                        _ = v40
                    } else {
                        v41 := []int{(fi14 >> ((v7) & 63)), fi13, (fi13 | (fi14 + 9223372036854775807)), 9223372036854775807}
                        _ = v41
                    }
                } else {
                    var v42 uint32 = 365
                    v42 -= (v42 * (v42 / ((v42) | 1)))
                    v42 += v42
                    _ = v42
                }
            }
            _ = fi14
            _ = v15
            _ = v16
        }
        switch (fi13 / (((fi13 ^ fi13)) | 1)) {
        case 127:
            var v43 float64 = 0.1
            _ = v43
        case 9223372036854775807:
            v44 := GenS3{F0: 2}
            _ = v44
        }
        v7 = (v7 - v7)
        _ = fi13
    }
    if true {
        var v45 string = "foo bar"
        if (0 != 18446744073709551615) {
            var v46 float64 = 1e10
            _ = v46
        }
        v47 := []int{42, 958, 874, 2147483647}
        _ = v45
        _ = v47
    } else {
        v48 := []int{}
        _ = v48
    }
    v7 += (v7 | v7)
    _ = v7
    return -609
}

func main() {
    v1 := make(map[uint32]string)
    var v2 int16 = 0
    var v3 int8 = -56
    v4 := make(map[uint32]string)
    v5 := []string{("" + ""), "foo bar", ("\n\t" + ("a" + ("\n\t" + "\n\t"))), ""}
    v6 := make(map[int]uint32)
    var v49 uint = 18446744073709551615
    if true {
        v49 = v49
        for fi50 := 0; fi50 < 1; fi50++ {
            v3 += (-(v3))
            _ = fi50
        }
        v51 := v1[663]
        _ = v51
    }
    v49 = 18446744073709551615
    fmt.Println(v49, v49, v2)
    var v52 uint16 = 1
    v3 = 0
    var v53 int32 = 361
    for fi54 := 0; fi54 < 5; fi54++ {
        var v55 uint64 = 2
        v56 := GenS4{F0: v55}
        _ = fi54
        _ = v55
        _ = v56
    }
    for fi57 := 0; fi57 < 4; fi57++ {
        delete(v4, 48)
        if (!true) {
            var v58 uint16 = (v52 % (((v52 >> (((v49 | (v49 ^ v49))) & 63))) | 1))
            v4[2] = (("\n\t" + "") + ("hello" + ("a" + ("hello" + "hello"))))
            _ = v58
        } else {
            for fi59 := 0; fi59 < 13; fi59++ {
                v60 := GenS4{F0: 1}
                _ = fi59
                _ = v60
            }
            if false {
                v61 := GenS5{F0: 510, F1: (v2 >> (((v49 & v49)) & 63))}
                _ = v61
            } else {
                v62 := v6[0]
                for fi63 := 0; fi63 < 11; fi63++ {
                    switch fi63 {
                    case 32767:
                        _ = len(v4)
                    default:
                        v49 = (v49 ^ (v49 * 2))
                        delete(v4, (v62 % ((1) | 1)))
                        var v64 int8 = (v3 << ((uint(v52)) & 63))
                        _ = v64
                    }
                    _ = fi63
                }
                _ = v62
            }
        }
        _ = fi57
    }
    v52 = v52
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v49:", v49)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v52
    _ = v53
    _ = v6
}
