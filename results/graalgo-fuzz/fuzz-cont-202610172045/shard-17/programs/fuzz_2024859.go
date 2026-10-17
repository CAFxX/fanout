package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32 }
type GenS2 struct { F0 bool; F1 string; F2 uint16 }
type GenS3 struct { F0 string; F1 uint8; F2 uint32 }
type GenS4 struct { F0 int; F1 int64; F2 int64 }
type GenS5 struct { F0 uint; F1 int; F2 int32 }
type GenS6 struct { F0 uint; F1 uint32 }
type GenS7 struct { F0 bool; F1 uint8; F2 bool }
type GenS8 struct { F0 uint8 }
func genFn0(p0 uint8) uint {
    v7 := GenS1{F0: -462, F1: uint32(p0)}
    for fi8 := 0; fi8 < 1; fi8++ {
        for fi9 := 0; fi9 < 6; fi9++ {
            p0 += (p0 << ((0) & 63))
            _ = fi9
        }
        _ = fi8
    }
    v10 := GenS2{F0: (32767 <= -9223372036854775808), F1: ("" + "foo bar"), F2: 116}
    var v11 float64 = ((0.0 - 1e10) - 3.14159)
    _ = p0
    _ = v10
    _ = v11
    _ = v7
    return 571
}
func genFn1(p0 int64, p1 int16) uint32 {
    switch 1 {
    case 2147483647:
        var v12 float32 = 0.0
        _ = v12
    case 250:
        fmt.Println(p0, p1)
    }
    p0 = 433
    p1 = (^(p1))
    _ = p0
    _ = p1
    return 385
}

func main() {
    var v1 int16 = 42
    var v2 string = "foo bar"
    var v3 uint16 = 32767
    v4 := []uint32{389, 4294967295, 0, 1}
    v5 := make(map[int]bool)
    v6 := make(map[uint32]uint32)
    v1 += v1
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println(v2, v2)
    if len(v4) > 0 {
        v4[0] = 774
    }
    fmt.Println(v1)
    if false {
        for fi13 := 0; fi13 < 4; fi13++ {
            v14 := GenS2{F0: false, F1: v2, F2: (v3 & (v3 * v3))}
            _ = fi13
            _ = v14
        }
        for fi15 := 0; fi15 < 9; fi15++ {
            if true {
                v1 -= v1
                switch fi15 {
                case -9223372036854775808:
                    v2 = ("hello" + "hello")
                case 751:
                    v3 = (v3 % ((2) | 1))
                default:
                    for fi16 := 0; fi16 < 4; fi16++ {
                        if true {
                            var v17 uint8 = 1
                            switch fi15 {
                            case 1:
                                fmt.Println(v1, v1, v1)
                            case 31:
                                switch fi16 {
                                case -236:
                                    v2 = v2
                                default:
                                    var v18 int16 = (v1 / ((v1) | 1))
                                    _ = v18
                                }
                            }
                            if true {
                                v17 -= v17
                            }
                            _ = v17
                        } else {
                            if len(v4) > 0 {
                                v4[0] = 260
                            }
                        }
                        v19 := GenS1{F0: v1, F1: genFn1(-1, (v1 - 32767))}
                        var v20 uint = 2
                        _ = fi16
                        _ = v19
                        _ = v20
                    }
                }
                var v21 float64 = (1e10 + 1.5)
                _ = v21
            }
            var v22 float64 = 0.1
            for fi23 := 0; fi23 < 8; fi23++ {
                switch v2 {
                case "":
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                case "hello":
                    var v24 uint32 = 65535
                    _ = v24
                case "\n\t":
                    v25 := GenS3{F0: v2, F1: uint8(v3), F2: 2}
                    _ = v25
                }
                switch v2 {
                case "foo bar":
                    v22 = (v22 * (v22 - v22))
                case "\n\t":
                    v26 := GenS4{F0: 127, F1: int64(v1), F2: int64(fi15)}
                    _ = v26
                case "hello":
                    v27 := v6[4294967295]
                    _ = v27
                }
                _ = fi23
            }
            _ = fi15
            _ = v22
        }
        fmt.Println(v2, v2)
    } else {
        if true {
            if false {
                v1 = v1
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                v1 = 42
            }
            if (1.5 == 0.0) {
                var v28 float64 = 1e10
                v29 := GenS1{F0: 42, F1: genFn1(42, (v1 - v1))}
                switch -2147483648 {
                case 2147483647:
                    var v30 int64 = 341
                    _ = v30
                case -2147483648:
                    switch 625 {
                    case -2147483648:
                        var v31 bool = true
                        _ = v31
                    default:
                        v32 := GenS5{F0: 1, F1: int(v3), F2: -1}
                        _ = v32
                    }
                case 4294967295:
                    var v33 uint64 = 0
                    _ = v33
                }
                _ = v28
                _ = v29
            } else {
                if false {
                    var v34 int64 = 1
                    _ = v34
                } else {
                    v1 += v1
                    v3 = 42
                    v5[127] = false
                }
            }
            for fi35 := 0; fi35 < 19; fi35++ {
                if (770 < 0) {
                    _ = len(v4)
                    switch ((v2 + "") + v2) {
                    case "":
                        if ((v3 ^ 1) > 2) {
                            var v36 int32 = 2147483647
                            var v37 int64 = 94
                            var v38 string = v2
                            _ = v36
                            _ = v37
                            _ = v38
                        } else {
                            fmt.Println(v3, v2)
                        }
                    case "foo bar":
                        var v39 int = 211
                        _ = v39
                    case "\n\t":
                        if ((!false) || (true || true)) {
                            v1 += -32768
                        }
                    }
                }
                _ = fi35
            }
        }
    }
    for fi40 := 0; fi40 < 3; fi40++ {
        v2 = v2
        fmt.Println(v3, v2, fi40)
        var v41 float32 = 1.5
        _ = fi40
        _ = v41
    }
    var v42 int64 = 1
    v43 := GenS6{F0: 65535, F1: 674}
    v44 := GenS7{F0: ((^(v1)) < 0), F1: 127, F2: (0 > 127)}
    for fi45 := 0; fi45 < 5; fi45++ {
        var v46 bool = (true || true)
        _ = fi45
        _ = v46
    }
    v47 := v6[194]
    var v48 string = ((v2 + (v2 + ("" + v2))) + v2)
    v49 := GenS8{F0: 42}
    var v50 uint32 = v47
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v42:", v42)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v50:", v50)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v44
    _ = v47
    _ = v48
    _ = v49
    _ = v5
    _ = v50
    _ = v6
}
