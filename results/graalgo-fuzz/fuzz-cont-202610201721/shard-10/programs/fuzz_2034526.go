package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int }
type GenS2 struct { F0 int16; F1 int64 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint; F1 uint8; F2 uint32 }
type GenS5 struct { F0 int16; F1 uint8; F2 uint8 }
type GenS6 struct { F0 uint64; F1 int16; F2 bool }

func main() {
    v1 := make(map[string]int)
    var v2 uint = 710
    var v3 float64 = float64(1.5)
    var v4 uint8 = 1
    v5 := GenS1{F0: 1, F1: 128}
    if (float32(-2.25) > 0.1) {
        for fi6 := 0; fi6 < 2; fi6++ {
            switch "foo bar" {
            case "":
                for fi7 := 0; fi7 < 13; fi7++ {
                    var v8 uint8 = 42
                    fmt.Println(v2)
                    if ((((-2.25 * (-2.25 + (-2.25 - (-2.25 * -2.25)))) == (-2.25 - -2.25)) || false) && true) {
                        v4 = v8
                    } else {
                        switch ("hello" + "hello") {
                        case "":
                            var v9 uint8 = v8
                            _ = v9
                        case "hello":
                            v2 -= v2
                        default:
                            var v10 int64 = 601
                            _ = v10
                        }
                    }
                    _ = fi7
                    _ = v8
                }
            case "hello":
                for fi11 := 0; fi11 < 19; fi11++ {
                    v12 := []int{}
                    _ = fi11
                    _ = v12
                }
            }
            _ = fi6
        }
    }
    var v13 bool = (!true)
    v4 = uint8(v2)
    for fi14 := 0; fi14 < 16; fi14++ {
        fmt.Println(fi14)
        _ = len(v1)
        _ = fi14
    }
    var v15 string = ("\n\t" + "")
    if true {
        for fi16 := 0; fi16 < 15; fi16++ {
            for fi17 := 0; fi17 < 7; fi17++ {
                v18 := GenS2{F0: -1, F1: -9223372036854775808}
                var v19 uint64 = 1
                if v13 {
                    fmt.Println(v3, fi17, v13)
                } else {
                    var v20 int8 = 48
                    fmt.Println(v3, v4)
                    _ = v20
                }
                var v21 uint32 = 65535
                var v22 uint = v2
                if v13 {
                    fmt.Println(v13, v2, v13)
                    if true {
                        if ((-193 >= 42) && true) {
                            v23 := []int{32767, (fi17 * fi17)}
                            switch v15 {
                            case "hello":
                                var v24 int8 = -8
                                _ = v24
                            default:
                                fmt.Println(fi17, v21, v2)
                            }
                            _ = v23
                        }
                        var v25 int16 = -1
                        var v26 float32 = 0.0
                        _ = v25
                        _ = v26
                    } else {
                        v19 = v19
                        v21 += (v21 / ((1) | 1))
                    }
                    if (v4 > 128) {
                        var v27 int8 = -128
                        fmt.Println(v3, v15, v2)
                        if ((0.0 + 0.0) == 3.14159) {
                            var v28 uint8 = v4
                            v1[v15] = (fi17 * 730)
                            v19 += uint64(v21)
                            _ = v28
                        }
                        _ = v27
                    } else {
                        v29 := GenS1{F0: 0, F1: 2147483647}
                        _ = v29
                    }
                } else {
                    if (507 != v22) {
                        v13 = ((v22 >> (((v22 | (v2 & (^(v2))))) & 63)) <= (v22 & uint(v19)))
                        fmt.Println(v2, fi17, v2)
                        fmt.Println(v19)
                    }
                    if v13 {
                        var v30 int8 = -128
                        var v31 uint32 = v21
                        _ = v30
                        _ = v31
                    }
                    if false {
                        v32 := []int{(fi16 << (((v22 / ((v22) | 1))) & 63)), fi17}
                        _ = len(v32)
                        v21 -= v21
                        _ = v32
                    } else {
                        if v13 {
                            v22 = 65535
                            if false {
                                v19 = (^(v19))
                            }
                            if (!(true && v13)) {
                                fmt.Println(v21, v19, v2)
                                var v33 int8 = 95
                                v34 := []int{(fi17 / ((-1) | 1)), -2147483648, int(v19)}
                                v4 -= uint8(v21)
                                v35 := v1[(v15 + v15)]
                                v36 := GenS1{F0: (v33 * v33), F1: (-(fi16))}
                                _ = v33
                                _ = v34
                                _ = v35
                                _ = v36
                            } else {
                                delete(v1, "\n\t")
                                var v37 int16 = -32768
                                _ = v37
                            }
                        } else {
                            if false {
                                v38 := []int{(fi17 >> ((uint(v4)) & 63)), (fi16 * fi17), (fi16 % (((fi16 | fi16)) | 1))}
                                v39 := GenS2{F0: 386, F1: 743}
                                if (-944 < 0) {
                                    var v40 uint = uint(v4)
                                    var v41 int32 = 2147483647
                                    _ = v40
                                    _ = v41
                                }
                                _ = v38
                                _ = v39
                            } else {
                                v1[v15] = 0
                                var v42 string = (v15 + ("foo bar" + ("a" + v15)))
                                v43 := GenS3{F0: (v4 & (v4 - v4))}
                                v42 = ("\n\t" + "\n\t")
                                var v44 int16 = 32767
                                _ = v42
                                _ = v43
                                _ = v44
                            }
                            v45 := GenS4{F0: v22, F1: 157, F2: 2}
                            _ = v45
                        }
                        fmt.Println(fi17)
                    }
                }
                v46 := v1[((v15 + "a") + "foo bar")]
                v46 -= v46
                _ = fi17
                _ = v18
                _ = v19
                _ = v21
                _ = v22
                _ = v46
            }
            _ = fi16
        }
    } else {
        if v13 {
            var v47 int = 128
            v2 -= (^(v2))
            _ = v47
        } else {
            v4 += v4
            v48 := GenS5{F0: 536, F1: (v4 - v4), F2: v4}
            var v49 uint16 = 301
            _ = v48
            _ = v49
        }
        var v50 int32 = -1
        v51 := GenS6{F0: 18446744073709551615, F1: 32767, F2: v13}
        _ = v50
        _ = v51
    }
    v4 = 128
    v52 := v1["\n\t"]
    v53 := []int{}
    var v54 int = (v52 + v52)
    v54 += (v52 % (((v54 * 156)) | 1))
    v15 = ("hello" + "foo bar")
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v52:", v52)
    fmt.Println("slicelen v53:", len(v53), v53)
    fmt.Println("v54:", v54)
    _ = v1
    _ = v13
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v52
    _ = v53
    _ = v54
}
