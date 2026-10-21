package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16; F2 bool }
type GenS2 struct { F0 uint32; F1 int32 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 int8; F1 uint64; F2 int16 }
func genFn0() int16 {
    if true {
        v7 := GenS1{F0: -9223372036854775808, F1: -1, F2: false}
        var v8 float64 = 3.14159
        _ = v7
        _ = v8
    } else {
        var v9 int = 32767
        if false {
            v9 -= -107
            switch ("\n\t" + "hello") {
            case "a":
                fmt.Println(v9)
            default:
                if false {
                    v10 := GenS1{F0: 1, F1: -32768, F2: false}
                    v11 := []int{(v9 + (v9 / ((v9) | 1)))}
                    _ = v10
                    _ = v11
                } else {
                    fmt.Println(v9)
                }
            }
            switch v9 {
            case 858:
                switch v9 {
                case 0:
                    v12 := []int{}
                    _ = v12
                default:
                    var v13 float64 = -2.25
                    _ = v13
                }
            case 127:
                v14 := GenS2{F0: 42, F1: -1}
                _ = v14
            default:
                fmt.Println(v9)
            }
        }
        _ = v9
    }
    var v15 int64 = 167
    v15 += v15
    var v16 bool = false
    _ = v15
    _ = v16
    return -32768
}
func genFn1() int16 {
    for fi17 := 0; fi17 < 1; fi17++ {
        var v18 int64 = 9223372036854775807
        if (0.1 >= 1.5) {
            if (!false) {
                if false {
                    switch ("hello" + "") {
                    case "a":
                        if (((0.0 / 1.0) - 1.5) > (1.5 / 0.1)) {
                            var v19 uint64 = 0
                            v18 += v18
                            _ = v19
                        }
                    case "hello":
                        if (31 > 2) {
                            v20 := GenS2{F0: 42, F1: 0}
                            v21 := []int{int(v18), -1, fi17}
                            _ = v20
                            _ = v21
                        } else {
                            v18 = v18
                        }
                    case "foo bar":
                        v18 = v18
                    }
                } else {
                    v18 = v18
                    for fi22 := 0; fi22 < 17; fi22++ {
                        fmt.Println(fi22, v18)
                        _ = fi22
                    }
                    for fi23 := 0; fi23 < 18; fi23++ {
                        switch "a" {
                        case "a":
                            switch (fi23 ^ fi23) {
                            case 32768:
                                var v24 uint32 = 2147483648
                                v25 := []int{fi17, fi17}
                                _ = v24
                                _ = v25
                            }
                        }
                        _ = fi23
                    }
                }
                v18 = -865
                v18 += (v18 >> ((1) & 63))
            } else {
                for fi26 := 0; fi26 < 4; fi26++ {
                    v18 -= (v18 | v18)
                    v18 = (v18 << ((89) & 63))
                    v27 := []int{(fi26 & fi17), -9223372036854775808}
                    _ = fi26
                    _ = v27
                }
            }
        }
        _ = fi17
        _ = v18
    }
    v28 := GenS2{F0: 0, F1: -277}
    var v29 int8 = -1
    _ = v28
    _ = v29
    return 32767
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := []uint32{0, 581, 543}
    var v3 int = -340
    v4 := []int{}
    var v5 uint8 = 255
    v6 := []int{v3, v3}
    v5 = v5
    for fi30 := 0; fi30 < 19; fi30++ {
        var v31 uint16 = 255
        if (0 <= 59) {
            var v32 float32 = 3.14159
            switch (("hello" + "") + "\n\t") {
            case "foo bar":
                v5 -= v5
            case "hello":
                v6 = append(v6, v3)
            case "\n\t":
                switch "hello" {
                case "hello":
                    var v33 uint = 2
                    _ = v33
                case "\n\t":
                    v34 := GenS3{F0: 32767}
                    _ = v34
                case "a":
                    if true {
                        _ = len(v6)
                        var v35 uint64 = 1
                        if (-123 == 657) {
                            v5 -= v5
                        } else {
                            var v36 int64 = -1
                            _ = v36
                        }
                        _ = v35
                    }
                default:
                    var v37 uint16 = 475
                    _ = v37
                }
            default:
                _ = len(v6)
            }
            _ = v32
        }
        _ = fi30
        _ = v31
    }
    var v38 int8 = 100
    var v39 int8 = -128
    switch ("a" + ("\n\t" + "a")) {
    case "a":
        for fi40 := 0; fi40 < 13; fi40++ {
            var v41 int32 = -359
            var v42 bool = false
            for fi43 := 0; fi43 < 9; fi43++ {
                if false {
                    v44 := GenS3{F0: 1}
                    fmt.Println(v38, fi40)
                    var v45 int8 = (v39 / ((v38) | 1))
                    _ = v44
                    _ = v45
                }
                var v46 int = (fi40 << ((1) & 63))
                switch v46 {
                case -982:
                    var v47 uint8 = (v5 + (v5 ^ (v5 >> ((4294967295) & 63))))
                    _ = v47
                case 32767:
                    v48 := GenS4{F0: -708}
                    _ = v48
                case -201:
                    var v49 int16 = -982
                    _ = v49
                default:
                    v3 = (v46 / (((v46 + (v46 * -1))) | 1))
                }
                _ = fi43
                _ = v46
            }
            _ = fi40
            _ = v41
            _ = v42
        }
    case "":
        fmt.Println(v5, v39, v38)
    case "hello":
        v50 := GenS5{F0: 1}
        _ = v50
    default:
        var v51 uint = 42
        _ = v51
    }
    v2 = append(v2, 42)
    var v52 float64 = 3.14159
    v53 := v1[2147483648]
    if (((v52 - 3.14159) / 1e10) < v52) {
        var v54 float32 = 3.14159
        if ("" > "foo bar") {
            for fi55 := 0; fi55 < 8; fi55++ {
                v38 = v38
                v52 = 1e10
                for fi56 := 0; fi56 < 7; fi56++ {
                    var v57 int32 = int32(v3)
                    if (v54 > (v54 - v54)) {
                        fmt.Println(v3)
                        if v53 {
                            var v58 uint32 = 4294967295
                            _ = v58
                        }
                    } else {
                        if ((!v53) || true) {
                            v54 = ((v54 * (v54 - (float32((((v52 + 0.1) + (v52 * 1e10)) - float64((float32((0.1 - v52)) + (v54 - (0.0 + ((1e10 / ((v54 + (((v54 - v54) * v54) - v54)) + v54)) + (v54 - float32((1.5 - v52)))))))))) - 3.14159))) / v54)
                            v38 = v39
                            v5 -= 98
                        }
                        fmt.Println(v5, v39)
                    }
                    _ = fi56
                    _ = v57
                }
                _ = fi55
            }
        } else {
            var v59 uint16 = 42
            v60 := GenS6{F0: v38, F1: 389, F2: 42}
            for fi61 := 0; fi61 < 18; fi61++ {
                var v62 int = 0
                if len(v2) > 0 {
                    v2[0] = 634
                }
                var v63 int8 = 24
                _ = fi61
                _ = v62
                _ = v63
            }
            _ = v59
            _ = v60
        }
        v39 -= v38
        _ = v54
    }
    v3 = v3
    v5 -= (^(v5))
    var v64 bool = ("" >= "foo bar")
    var v65 int32 = -2147483648
    fmt.Println(v39, v5, v38)
    v66 := GenS1{F0: 0, F1: 577, F2: (!true)}
    if true {
        for fi67 := 0; fi67 < 13; fi67++ {
            v4 = append(v4, 4294967295)
            var v68 int32 = v65
            fmt.Println(v38, v68)
            _ = fi67
            _ = v68
        }
        if len(v4) > 0 {
            v4[0] = 896
        }
        _ = len(v1)
    } else {
        fmt.Println(v64, v3, v53)
    }
    v1[0] = v64
    for fi69 := 0; fi69 < 14; fi69++ {
        if v53 {
            if v53 {
                fmt.Println(v65, v64)
                var v70 uint32 = 927
                v71 := v1[v70]
                _ = v70
                _ = v71
            } else {
                if len(v4) > 0 {
                    v4[0] = v3
                }
            }
        } else {
            var v72 int8 = (^(v38))
            v53 = v64
            _ = v72
        }
        v73 := v1[360]
        _ = fi69
        _ = v73
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v64:", v64)
    fmt.Println("v65:", v65)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v52
    _ = v53
    _ = v6
    _ = v64
    _ = v65
    _ = v66
}
