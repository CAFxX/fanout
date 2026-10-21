package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8; F2 uint }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int64; F1 uint32; F2 bool }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint8; F1 int16; F2 int64 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 string = "\n\t"
    v3 := []string{"hello", "foo bar", "\n\t"}
    if (-1 > -128) {
        v2 = (("foo bar" + (((v2 + v2) + ((v2 + v2) + (((v2 + ("a" + "foo bar")) + (v2 + v2)) + v2))) + v2)) + "")
    }
    if true {
        fmt.Println(v2)
        var v4 bool = true
        switch v2 {
        case "foo bar":
            var v5 int64 = 0
            _ = v5
        case "":
            v6 := v1[0]
            _ = v6
        }
        _ = v4
    }
    var v7 int64 = 1
    if false {
        if false {
            v7 -= v7
            v7 += v7
            var v8 bool = false
            _ = v8
        }
        for fi9 := 0; fi9 < 10; fi9++ {
            var v10 uint16 = 2
            _ = fi9
            _ = v10
        }
        for fi11 := 0; fi11 < 9; fi11++ {
            switch v2 {
            case "a":
                fmt.Println(v2, fi11)
            case "\n\t":
                v7 += v7
            case "foo bar":
                if (!false) {
                    if (4 < 42) {
                        v12 := GenS1{F0: -825, F1: 113, F2: 159}
                        for fi13 := 0; fi13 < 6; fi13++ {
                            v7 += (v7 ^ (v7 % ((v7) | 1)))
                            var v14 uint8 = 1
                            v14 = v14
                            _ = fi13
                            _ = v14
                        }
                        _ = v12
                    }
                }
            default:
                var v15 int32 = 2147483647
                _ = v15
            }
            if true {
                var v16 float64 = (3.14159 - 0.0)
                if len(v3) > 0 {
                    v3[0] = v2
                }
                var v17 uint32 = 601
                _ = v16
                _ = v17
            } else {
                if ((32767 == 42) && false) {
                    v18 := GenS2{F0: 1000}
                    var v19 string = (("\n\t" + v2) + v2)
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    _ = v18
                    _ = v19
                } else {
                    for fi20 := 0; fi20 < 16; fi20++ {
                        switch (fi11 >> ((255) & 63)) {
                        case 9223372036854775807:
                            if len(v3) > 0 {
                                v3[0] = v2
                            }
                        case 715:
                            var v21 int8 = 110
                            _ = v21
                        default:
                            var v22 int64 = (v7 << ((712) & 63))
                            _ = v22
                        }
                        var v23 int64 = v7
                        var v24 int32 = -1
                        if ((false && false) && false) {
                            var v25 uint32 = 625
                            v1[v25] = (!(((0.1 / ((0.0 - float64(0.0)) - 0.1)) / (1e10 * 0.1)) <= ((0.1 * 0.1) / (1.5 * 0.1))))
                            _ = v25
                        } else {
                            v2 = v2
                            fmt.Println(v23, v24)
                        }
                        _ = fi20
                        _ = v23
                        _ = v24
                    }
                    switch v2 {
                    case "foo bar":
                        _ = len(v1)
                    case "\n\t":
                        v2 = ("foo bar" + "\n\t")
                    }
                    for fi26 := 0; fi26 < 2; fi26++ {
                        v2 = (("\n\t" + "hello") + ((v2 + v2) + (v2 + ("\n\t" + "a"))))
                        _ = fi26
                    }
                }
                v27 := GenS3{F0: v7, F1: 237, F2: true}
                switch "foo bar" {
                case "foo bar":
                    delete(v1, 176)
                }
                _ = v27
            }
            v7 = 1
            _ = fi11
        }
    }
    for fi28 := 0; fi28 < 14; fi28++ {
        var v29 uint8 = 35
        _ = fi28
        _ = v29
    }
    for fi30 := 0; fi30 < 11; fi30++ {
        v7 += v7
        switch v2 {
        case "foo bar":
            var v31 float64 = ((1e10 + 0.0) * 3.14159)
            _ = v31
        case "hello":
            fmt.Println(v7)
        case "":
            v7 = (^(v7))
        default:
            v2 = "foo bar"
        }
        _ = fi30
    }
    var v32 int32 = 42
    v7 -= (v7 - v7)
    v33 := GenS3{F0: 1, F1: 65535, F2: false}
    var v34 float64 = (3.14159 * 1e10)
    var v35 int16 = 860
    if (2 > 65535) {
        if false {
            switch 32768 {
            case -2147483648:
                var v36 int32 = (-(v32))
                _ = v36
            case 710:
                _ = len(v1)
            case 127:
                delete(v1, 2147483647)
            default:
                fmt.Println(v35, v35, v35)
            }
        } else {
            var v37 bool = false
            v35 += (v35 + v35)
            var v38 uint32 = 42
            _ = v37
            _ = v38
        }
    }
    if (!((65535 != 42) || true)) {
        _ = len(v3)
    }
    var v39 float32 = 3.14159
    if (v35 >= -310) {
        fmt.Println(v34, v2, v32)
    } else {
        for fi40 := 0; fi40 < 5; fi40++ {
            v39 = (0.1 + ((((v39 - (float32(v34) * ((1e10 * 1e10) / v39))) + v39) - (((-2.25 * 1.5) * 3.14159) - v39)) - v39))
            for fi41 := 0; fi41 < 2; fi41++ {
                fmt.Println(v34, fi40, fi41)
                v42 := GenS4{F0: 1}
                var v43 string = "\n\t"
                v44 := GenS1{F0: 0, F1: -101, F2: 2}
                switch v43 {
                case "a":
                    v45 := GenS5{F0: 203, F1: v35, F2: v7}
                    _ = v45
                case "hello":
                    v43 = "foo bar"
                case "\n\t":
                    switch v2 {
                    case "foo bar":
                        v2 = v2
                    case "":
                        v32 = 863
                    default:
                        if (int32(fi40) < (^(v32))) {
                            v7 += -1
                        } else {
                            if (v34 != float64((float32((v34 - (v34 - v34))) + 0.0))) {
                                v32 -= (v32 << ((65535) & 63))
                            }
                            if true {
                                _ = len(v3)
                                switch fi41 {
                                case -353:
                                    if true {
                                        var v46 int32 = (v32 & -522)
                                        v47 := GenS1{F0: 888, F1: -77, F2: 255}
                                        _ = v46
                                        _ = v47
                                    }
                                default:
                                    v35 = -1
                                }
                                switch fi41 {
                                case 0:
                                    switch (fi41 ^ (fi40 + (fi40 / ((int(v35)) | 1)))) {
                                    case 32767:
                                        var v48 uint16 = 32767
                                        _ = v48
                                    case -9223372036854775808:
                                        v43 = ("\n\t" + v43)
                                    case 127:
                                        var v49 int32 = (v32 ^ 2147483647)
                                        _ = v49
                                    default:
                                        switch (v2 + "a") {
                                        case "a":
                                            delete(v1, 637)
                                        default:
                                            var v50 uint32 = 4294967295
                                            _ = v50
                                        }
                                    }
                                }
                            }
                            v51 := GenS1{F0: 32767, F1: -1, F2: 4294967295}
                            _ = v51
                        }
                    }
                }
                _ = fi41
                _ = v42
                _ = v43
                _ = v44
            }
            _ = fi40
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v39:", v39)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v39
    _ = v7
}
