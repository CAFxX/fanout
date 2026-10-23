package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint }
type GenS2 struct { F0 int; F1 uint8; F2 uint32 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint8; F1 int64; F2 uint32 }
type GenS5 struct { F0 uint8; F1 uint8; F2 uint64 }
type GenS6 struct { F0 string; F1 uint32 }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 bool; F1 uint; F2 string }
type GenS9 struct { F0 int16; F1 uint }

func main() {
    var v1 uint32 = 2
    var v2 uint64 = 4294967295
    v3 := make(map[uint32]uint32)
    v4 := []int{}
    v5 := []string{"", ("a" + "foo bar"), "hello", "foo bar"}
    var v6 uint64 = v2
    v2 += (v2 * v2)
    v1 = 109
    if (v6 <= (^(v6))) {
        for fi7 := 0; fi7 < 18; fi7++ {
            v3[(v1 * v1)] = (v1 << ((1) & 63))
            if (fi7 >= 0) {
                v3[626] = (v1 >> ((4294967295) & 63))
                v8 := GenS1{F0: uint16(v2), F1: 0}
                _ = v8
            } else {
                if (539 == -1) {
                    _ = len(v5)
                }
                for fi9 := 0; fi9 < 13; fi9++ {
                    var v10 float32 = 1e10
                    if true {
                        var v11 int8 = int8(v1)
                        if (!(-1 != 1000)) {
                            v10 = (v10 * (0.1 / v10))
                            v11 = v11
                        }
                        if false {
                            v12 := GenS2{F0: fi9, F1: 128, F2: v1}
                            if false {
                                if len(v5) > 0 {
                                    v5[0] = "hello"
                                }
                            } else {
                                var v13 uint16 = 124
                                v14 := GenS2{F0: (fi9 | fi9), F1: 0, F2: v1}
                                _ = v13
                                _ = v14
                            }
                            v10 = v10
                            v15 := GenS3{F0: 1}
                            _ = v12
                            _ = v15
                        } else {
                            v16 := GenS4{F0: 128, F1: 1, F2: (v1 / ((v1) | 1))}
                            v17 := GenS5{F0: 61, F1: 42, F2: v2}
                            v2 -= 460
                            _ = v16
                            _ = v17
                        }
                        _ = v11
                    } else {
                        v2 = v2
                        var v18 float32 = (0.0 / v10)
                        if (true || (!false)) {
                            if (53 == 128) {
                                fmt.Println(v18, v18)
                            }
                            var v19 uint = 0
                            _ = v19
                        }
                        if false {
                            fmt.Println(fi9, v2, v10)
                        }
                        _ = v18
                    }
                    _ = fi9
                    _ = v10
                }
            }
            _ = fi7
        }
        if false {
            if true {
                var v20 bool = true
                switch int(v6) {
                case -9223372036854775808:
                    var v21 int64 = 42
                    _ = v21
                }
                _ = v20
            } else {
                v22 := v3[v1]
                _ = v22
            }
            for fi23 := 0; fi23 < 10; fi23++ {
                var v24 int8 = 100
                v25 := GenS6{F0: "hello", F1: v1}
                for fi26 := 0; fi26 < 14; fi26++ {
                    fmt.Println(v2, v6)
                    if len(v4) > 0 {
                        v4[0] = fi26
                    }
                    _ = fi26
                }
                _ = fi23
                _ = v24
                _ = v25
            }
        }
        v6 += 765
    }
    if true {
        delete(v3, v1)
    }
    v27 := GenS3{F0: 1000}
    var v28 uint16 = uint16(v1)
    for fi29 := 0; fi29 < 14; fi29++ {
        if false {
            v1 += (v1 << ((355) & 63))
            switch (^(fi29)) {
            case 4294967295:
                var v30 float32 = (float32(0.1) + ((1.5 + 0.0) * 0.0))
                _ = v30
            case 32768:
                var v31 uint32 = v1
                _ = v31
            case 1:
                for fi32 := 0; fi32 < 19; fi32++ {
                    var v33 int32 = 1
                    var v34 uint16 = (^(v28))
                    if len(v4) > 0 {
                        v4[0] = fi32
                    }
                    _ = fi32
                    _ = v33
                    _ = v34
                }
            default:
                v2 = v2
            }
            switch ("foo bar" + "\n\t") {
            case "foo bar":
                switch 127 {
                case 721:
                    var v35 uint64 = (v6 / (((v2 >> ((4294967295) & 63))) | 1))
                    _ = v35
                case 127:
                    v36 := GenS5{F0: 87, F1: 118, F2: (v6 + (v6 + v2))}
                    _ = v36
                default:
                    var v37 int32 = 1
                    _ = v37
                }
            case "\n\t":
                v38 := GenS7{F0: 150}
                _ = v38
            }
        } else {
            for fi39 := 0; fi39 < 7; fi39++ {
                var v40 string = ""
                v2 -= v6
                v41 := GenS8{F0: true, F1: 255, F2: (((((v40 + ("" + (v40 + v40))) + "\n\t") + "foo bar") + (v40 + ("hello" + (((("" + (v40 + "foo bar")) + v40) + "") + ((v40 + "a") + v40))))) + v40)}
                _ = fi39
                _ = v40
                _ = v41
            }
            for fi42 := 0; fi42 < 3; fi42++ {
                if false {
                    v28 = (v28 / ((v28) | 1))
                    v1 += v1
                    var v43 int16 = -32768
                    _ = v43
                }
                v44 := GenS2{F0: (-(fi29)), F1: 128, F2: v1}
                _ = fi42
                _ = v44
            }
            _ = len(v3)
        }
        v6 = v2
        _ = fi29
    }
    var v45 string = "\n\t"
    v3[65535] = 221
    _ = len(v3)
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v4 = append(v4, -384)
    v46 := GenS9{F0: 636, F1: 776}
    for fi47 := 0; fi47 < 16; fi47++ {
        v6 -= v6
        switch (((v45 + v45) + "\n\t") + (v45 + (((v45 + v45) + "foo bar") + v45))) {
        case "\n\t":
            v6 = (v6 ^ 814)
        case "hello":
            var v48 uint32 = v1
            _ = v48
        case "a":
            v3[v1] = v1
        default:
            for fi49 := 0; fi49 < 10; fi49++ {
                var v50 int16 = -27
                _ = fi49
                _ = v50
            }
        }
        var v51 float64 = 1.5
        _ = fi47
        _ = v51
    }
    var v52 uint64 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v45:", v45)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v52:", v52)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v5
    _ = v52
    _ = v6
}
