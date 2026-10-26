package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8 }
type GenS2 struct { F0 uint; F1 int16; F2 uint32 }
type GenS3 struct { F0 uint32; F1 uint64; F2 bool }
type GenS4 struct { F0 int8; F1 int }

func main() {
    var v1 int16 = -553
    var v2 float32 = 1e10
    var v3 int64 = -9223372036854775808
    var v4 uint = 647
    v3 = v3
    fmt.Println(v1, v4, v3)
    v5 := GenS1{F0: 802, F1: -128}
    v6 := []int{-908, -769, 32768}
    fmt.Println(v3, v3, v4)
    for fi7 := 0; fi7 < 19; fi7++ {
        for fi8 := 0; fi8 < 5; fi8++ {
            v4 += v4
            _ = len(v6)
            v4 -= 325
            if (100 > -1) {
                var v9 uint32 = 2
                var v10 float64 = (float64((v2 / v2)) + float64(v2))
                if false {
                    if true {
                        if (!false) {
                            v11 := GenS2{F0: v4, F1: v1, F2: 786}
                            _ = v11
                        }
                        switch "\n\t" {
                        case "foo bar":
                            var v12 int = 32767
                            _ = v12
                        case "\n\t":
                            if true {
                                var v13 uint64 = uint64(v9)
                                _ = v13
                            } else {
                                v1 += v1
                            }
                            v14 := GenS3{F0: v9, F1: 535, F2: true}
                            _ = v14
                        default:
                            v3 += v3
                        }
                        switch fi8 {
                        case -1:
                            if true {
                                if true {
                                    v15 := GenS3{F0: (v9 >> ((v4) & 63)), F1: 42, F2: (((v2 * float32((-2.25 * v10))) - v2) != v2)}
                                    fmt.Println(v1)
                                    _ = v15
                                }
                                var v16 float32 = v2
                                v2 = 3.14159
                                _ = v16
                            }
                        case 0:
                            _ = len(v6)
                        case 128:
                            var v17 int32 = -1
                            _ = v17
                        default:
                            v2 += v2
                        }
                    } else {
                        v10 = (v10 + ((v10 * ((v10 * -2.25) / v10)) + v10))
                        if len(v6) > 0 {
                            v6[0] = -239
                        }
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                    }
                } else {
                    v3 = v3
                    var v18 string = "hello"
                    v10 = 0.0
                    _ = v18
                }
                v19 := GenS1{F0: 90, F1: 127}
                _ = v10
                _ = v19
                _ = v9
            } else {
                v20 := GenS3{F0: 1, F1: 1, F2: true}
                _ = v20
            }
            _ = fi8
        }
        _ = fi7
    }
    if (v4 > (v4 << ((2) & 63))) {
        if false {
            var v21 int16 = (v1 << (((v4 & 652)) & 63))
            _ = v21
        }
        v6 = append(v6, 42)
    } else {
        for fi22 := 0; fi22 < 19; fi22++ {
            var v23 uint16 = 239
            v24 := GenS4{F0: -48, F1: fi22}
            fmt.Println(v4, fi22, v2)
            _ = fi22
            _ = v23
            _ = v24
        }
    }
    fmt.Println(v1, v3)
    for fi25 := 0; fi25 < 18; fi25++ {
        v26 := GenS3{F0: 2147483648, F1: 395, F2: (true || ((853 >= 559) || false))}
        var v27 uint64 = 42
        _ = fi25
        _ = v26
        _ = v27
    }
    v4 = (v4 << (((v4 - v4)) & 63))
    switch -489 {
    case 1:
        _ = len(v6)
    case -385:
        _ = len(v6)
    default:
        switch 32767 {
        case 32767:
            if true {
                for fi28 := 0; fi28 < 17; fi28++ {
                    v6 = append(v6, (fi28 << ((v4) & 63)))
                    var v29 uint16 = 42
                    v2 = 3.14159
                    _ = fi28
                    _ = v29
                }
                for fi30 := 0; fi30 < 17; fi30++ {
                    v4 += (v4 ^ (v4 / (((v4 / (((v4 * 757)) | 1))) | 1)))
                    _ = len(v6)
                    _ = fi30
                }
                fmt.Println(v4)
            }
        case 1:
            if len(v6) > 0 {
                v6[0] = 32768
            }
        default:
            v31 := GenS4{F0: -68, F1: -1}
            _ = v31
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
