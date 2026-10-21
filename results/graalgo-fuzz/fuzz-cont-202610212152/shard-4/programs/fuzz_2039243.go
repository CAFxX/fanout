package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 string; F1 uint16 }
type GenS4 struct { F0 int16; F1 int32; F2 int32 }
type GenS5 struct { F0 uint64; F1 int64 }
type GenS6 struct { F0 uint64; F1 string }

func main() {
    v1 := make(map[uint32]bool)
    var v2 int = -9223372036854775808
    var v3 float64 = -2.25
    var v4 uint8 = 42
    v5 := []int{}
    var v6 uint16 = 1
    var v7 bool = (!(-520 != 42))
    if v7 {
        v3 -= v3
        var v8 int64 = 42
        v6 = v6
        _ = v8
    }
    var v9 uint16 = (^(v6))
    switch "a" {
    case "\n\t":
        if false {
            var v10 uint = 458
            _ = v10
        }
    case "hello":
        var v11 float32 = 0.1
        _ = v11
    case "foo bar":
        for fi12 := 0; fi12 < 2; fi12++ {
            v4 -= v4
            var v13 bool = (!true)
            _ = fi12
            _ = v13
        }
    default:
        switch (v2 - (v2 | 4294967295)) {
        case -9223372036854775808:
            if ((v4 > 0) && v7) {
                var v14 uint16 = 221
                _ = v14
            } else {
                fmt.Println(v2, v7)
            }
        default:
            var v15 string = ("" + ("hello" + "foo bar"))
            _ = v15
        }
    }
    fmt.Println(v4, v7)
    if (0 == 625) {
        v16 := GenS1{F0: -346, F1: -19}
        if (!true) {
            for fi17 := 0; fi17 < 15; fi17++ {
                if v7 {
                    v18 := GenS1{F0: -79, F1: -883}
                    _ = v18
                }
                if v7 {
                    v2 -= fi17
                }
                _ = fi17
            }
            switch (v2 * v2) {
            case 32768:
                if v7 {
                    v19 := GenS2{F0: 2}
                    fmt.Println(v3)
                    for fi20 := 0; fi20 < 20; fi20++ {
                        if true {
                            v1[0] = (9 < v4)
                            switch 2147483647 {
                            case 998:
                                v7 = true
                            case 2147483647:
                                var v21 int = v2
                                _ = v21
                            case -23:
                                if v7 {
                                    v3 = (v3 / 3.14159)
                                    fmt.Println(fi20, v2)
                                } else {
                                    v22 := GenS3{F0: "foo bar", F1: v9}
                                    v23 := GenS4{F0: 1000, F1: 0, F2: 0}
                                    _ = v22
                                    _ = v23
                                }
                            default:
                                v24 := GenS3{F0: "\n\t", F1: uint16(v4)}
                                _ = v24
                            }
                            v25 := GenS4{F0: 740, F1: 740, F2: 2147483647}
                            _ = v25
                        }
                        fmt.Println(v2)
                        for fi26 := 0; fi26 < 2; fi26++ {
                            v9 = 42
                            fmt.Println(v6, v2)
                            var v27 uint16 = 42
                            _ = fi26
                            _ = v27
                        }
                        _ = fi20
                    }
                    _ = v19
                }
            case 4294967295:
                for fi28 := 0; fi28 < 4; fi28++ {
                    v4 += v4
                    _ = fi28
                }
            default:
                var v29 uint = uint(v9)
                _ = v29
            }
        }
        for fi30 := 0; fi30 < 13; fi30++ {
            var v31 int = 9223372036854775807
            var v32 float64 = v3
            _ = fi30
            _ = v31
            _ = v32
        }
        _ = v16
    } else {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        if len(v5) > 0 {
            v5[0] = 336
        }
        for fi33 := 0; fi33 < 20; fi33++ {
            if true {
                var v34 int = 0
                _ = v34
            } else {
                v7 = (((v4 / (((v4 * (v4 << ((288) & 63)))) | 1)) == 124) || false)
            }
            _ = fi33
        }
    }
    v6 -= (v9 >> ((42) & 63))
    for fi35 := 0; fi35 < 14; fi35++ {
        v9 = (^(v6))
        v3 = 1e10
        _ = fi35
    }
    delete(v1, 88)
    v7 = v7
    v7 = v7
    if false {
        v2 -= v2
    }
    if v7 {
        v36 := GenS5{F0: uint64(v9), F1: 0}
        v7 = (1000 != 272)
        v37 := GenS6{F0: 4294967295, F1: "a"}
        _ = v36
        _ = v37
    } else {
        fmt.Println(v2, v9, v7)
        var v38 string = "\n\t"
        _ = v38
    }
    var v39 string = "\n\t"
    switch (v2 >> ((2) & 63)) {
    case 32768:
        fmt.Println(v4)
    default:
        if true {
            for fi40 := 0; fi40 < 11; fi40++ {
                v1[uint32(v9)] = v7
                v1[912] = (true || (!false))
                _ = fi40
            }
            if (v7 || (!(!v7))) {
                delete(v1, 4294967295)
                if (!false) {
                    v6 += 237
                } else {
                    _ = len(v5)
                }
            }
        }
    }
    v9 -= (v6 >> ((861) & 63))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
