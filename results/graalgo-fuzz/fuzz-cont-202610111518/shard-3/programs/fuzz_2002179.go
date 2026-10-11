package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8; F2 int64 }
type GenS2 struct { F0 uint; F1 int32 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 bool; F1 uint }

func main() {
    var v1 int64 = 437
    var v2 int16 = 42
    var v3 float64 = (1.5 * 1e10)
    v2 = (v2 & v2)
    v1 += v1
    if false {
        v3 -= 1.5
        if true {
            for fi4 := 0; fi4 < 18; fi4++ {
                fmt.Println(v1, fi4, v3)
                if true {
                    v2 -= (v2 | (v2 ^ (v2 >> ((2) & 63))))
                }
                _ = fi4
            }
            v1 = v1
        }
    } else {
        if false {
            if true {
                v5 := GenS1{F0: 0, F1: 1, F2: v1}
                _ = v5
            } else {
                var v6 float32 = (0.1 * 3.14159)
                _ = v6
            }
        }
        v7 := GenS1{F0: 42, F1: 225, F2: v1}
        _ = v7
    }
    if false {
        v3 = v3
    }
    var v8 uint32 = 1
    v1 -= v1
    for fi9 := 0; fi9 < 9; fi9++ {
        if false {
            v10 := []int{fi9, (fi9 * (fi9 >> ((255) & 63)))}
            for fi11 := 0; fi11 < 11; fi11++ {
                if (1000 < v2) {
                    if false {
                        var v12 int = 32767
                        _ = v12
                    } else {
                        v3 = (v3 - 1.5)
                        v1 = -958
                    }
                    var v13 int8 = 100
                    _ = v13
                } else {
                    if (1 < fi11) {
                        v10 = append(v10, fi9)
                    } else {
                        fmt.Println(v3, v8, v1)
                        if true {
                            v14 := GenS2{F0: 2, F1: -116}
                            _ = v14
                        } else {
                            if (-128 > -128) {
                                v10 = append(v10, fi9)
                                switch fi11 {
                                case 4294967295:
                                    _ = len(v10)
                                }
                            } else {
                                v15 := GenS1{F0: -421, F1: 235, F2: -100}
                                v2 += (-(v2))
                                _ = v15
                            }
                            fmt.Println(fi9)
                            v1 -= v1
                        }
                        if false {
                            fmt.Println(v1, fi9, fi11)
                            switch "hello" {
                            case "":
                                var v16 uint8 = 237
                                _ = v16
                            case "a":
                                v8 += (v8 - v8)
                            default:
                                switch (fi9 * (^(fi9))) {
                                case 9223372036854775807:
                                    v8 = (^(v8))
                                case 1:
                                    if false {
                                        v17 := GenS3{F0: fi11}
                                        var v18 int32 = 241
                                        _ = v17
                                        _ = v18
                                    }
                                default:
                                    v10 = append(v10, 32767)
                                }
                            }
                        }
                        v3 += float64(0.0)
                        v19 := GenS4{F0: false, F1: 15}
                        _ = v19
                    }
                    fmt.Println(v1)
                }
                switch fi11 {
                case 2147483647:
                    if (true || (((v1 >> ((0) & 63)) > 758) && (0.1 > 1.5))) {
                        switch fi11 {
                        case 128:
                            v10 = append(v10, -2147483648)
                        case -913:
                            v20 := GenS1{F0: -2147483648, F1: 42, F2: 9223372036854775807}
                            _ = v20
                        default:
                            switch "a" {
                            case "a":
                                var v21 int64 = 962
                                _ = v21
                            case "":
                                if true {
                                    v22 := GenS3{F0: (fi9 << ((868) & 63))}
                                    var v23 int32 = 0
                                    if false {
                                        var v24 uint16 = 32767
                                        if (v23 != 1) {
                                            var v25 uint16 = v24
                                            fmt.Println(fi11, fi9, v24)
                                            _ = v25
                                        } else {
                                            switch ("a" + "foo bar") {
                                            case "\n\t":
                                                v2 -= 818
                                            }
                                            var v26 uint64 = 393
                                            if false {
                                                v27 := GenS3{F0: (-(fi9))}
                                                _ = v27
                                            }
                                            _ = v26
                                        }
                                        _ = v24
                                    }
                                    _ = v22
                                    _ = v23
                                }
                            }
                        }
                        var v28 int64 = (^(v1))
                        _ = v28
                    }
                case 127:
                    v10 = append(v10, -9223372036854775808)
                case -2147483648:
                    fmt.Println(fi9, v8, v2)
                default:
                    var v29 int8 = -1
                    _ = v29
                }
                v8 = (v8 | 91)
                _ = fi11
            }
            if len(v10) >= 2 {
                v10 = v10[1:len(v10)]
            }
            _ = v10
        }
        if true {
            var v30 int32 = 222
            v31 := []int{32768, (fi9 | (fi9 ^ int(v30))), fi9}
            for fi32 := 0; fi32 < 10; fi32++ {
                v1 -= (v1 >> ((0) & 63))
                var v33 uint16 = 275
                var v34 int8 = -97
                _ = fi32
                _ = v33
                _ = v34
            }
            _ = v30
            _ = v31
        }
        _ = fi9
    }
    var v35 float64 = (v3 / v3)
    v2 += v2
    v36 := GenS3{F0: 128}
    for fi37 := 0; fi37 < 16; fi37++ {
        v8 = 4294967295
        _ = fi37
    }
    v38 := []int{1, 1, -428, -2147483648}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v38:", len(v38), v38)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v38
    _ = v8
}
