package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint; F1 uint32 }
type GenS3 struct { F0 string; F1 uint32; F2 uint16 }
type GenS4 struct { F0 int64; F1 uint8; F2 bool }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 uint64 }

func main() {
    v1 := make(map[uint32]string)
    var v2 uint64 = 268
    v3 := []uint32{}
    var v4 int8 = 42
    v2 = (v2 + v2)
    v5 := GenS1{F0: 890}
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v6 float32 = 1.5
    switch ("hello" + "") {
    case "foo bar":
        for fi7 := 0; fi7 < 13; fi7++ {
            var v8 int32 = 0
            if true {
                switch "hello" {
                case "\n\t":
                    v2 += v2
                case "a":
                    switch (((("a" + "hello") + "\n\t") + ("foo bar" + "\n\t")) + "\n\t") {
                    case "":
                        switch 2147483647 {
                        case 127:
                            switch fi7 {
                            case 572:
                                if false {
                                    v9 := GenS2{F0: 2, F1: 4294967295}
                                    v10 := GenS2{F0: 0, F1: 1}
                                    _ = v10
                                    _ = v9
                                } else {
                                    if ((true || true) && ((fi7 << ((396) & 63)) > int(v4))) {
                                        v11 := GenS2{F0: 65535, F1: 1}
                                        _ = v11
                                    }
                                    if true {
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        if (true || false) {
                                            v12 := GenS2{F0: 0, F1: 65535}
                                            _ = v12
                                        }
                                    }
                                    if len(v3) > 0 {
                                        v3[0] = 879
                                    }
                                }
                            case -490:
                                delete(v1, 538)
                            default:
                                if (true && false) {
                                    v13 := GenS1{F0: 4294967295}
                                    _ = v13
                                } else {
                                    fmt.Println(v8, v6, v4)
                                    fmt.Println(fi7, fi7)
                                }
                            }
                        default:
                            for fi14 := 0; fi14 < 13; fi14++ {
                                switch ("hello" + "foo bar") {
                                case "hello":
                                    if ((true && ((-(v8)) != (v8 % ((v8) | 1)))) && false) {
                                        fmt.Println(v8, v4)
                                        if (100 <= (v4 ^ v4)) {
                                            fmt.Println(v2, v8, fi7)
                                            v2 += (v2 + (v2 & (^(v2))))
                                            var v15 uint64 = (v2 << ((609) & 63))
                                            _ = v15
                                        } else {
                                            v16 := GenS1{F0: 220}
                                            _ = v16
                                        }
                                    } else {
                                        v17 := GenS3{F0: "foo bar", F1: 1, F2: 116}
                                        if len(v3) > 0 {
                                            v3[0] = 2147483647
                                        }
                                        _ = v17
                                    }
                                }
                                _ = fi14
                            }
                        }
                    case "foo bar":
                        v2 += v2
                    }
                default:
                    for fi18 := 0; fi18 < 13; fi18++ {
                        v2 = (v2 - (^(v2)))
                        switch "foo bar" {
                        case "\n\t":
                            v19 := GenS4{F0: 42, F1: 255, F2: false}
                            _ = v19
                        }
                        _ = fi18
                    }
                }
            } else {
                v6 = float32(-2.25)
            }
            _ = fi7
            _ = v8
        }
    case "a":
        var v20 float64 = (0.1 * (0.1 + (-2.25 - float64(1.5))))
        _ = v20
    default:
        switch 128 {
        case 4294967295:
            var v21 int64 = 0
            _ = v21
        }
    }
    v4 -= 1
    if (313 != 925) {
        if true {
            switch 42 {
            case 4294967295:
                if len(v3) > 0 {
                    v3[0] = 2147483648
                }
            case 9223372036854775807:
                var v22 string = "foo bar"
                _ = v22
            default:
                if (-292 <= -9223372036854775808) {
                    v23 := GenS5{F0: 252}
                    _ = v23
                }
            }
            var v24 int16 = 542
            switch int(v24) {
            case -1:
                if (333 <= 32768) {
                    v4 += (^(v4))
                    for fi25 := 0; fi25 < 16; fi25++ {
                        for fi26 := 0; fi26 < 5; fi26++ {
                            v27 := GenS6{F0: v24}
                            _ = fi26
                            _ = v27
                        }
                        var v28 int8 = v4
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        _ = fi25
                        _ = v28
                    }
                }
            case 566:
                var v29 int8 = v4
                _ = v29
            default:
                delete(v1, 0)
            }
            _ = v24
        } else {
            delete(v1, 2)
        }
        for fi30 := 0; fi30 < 15; fi30++ {
            fmt.Println(v4, v6, v2)
            if false {
                _ = len(v1)
            }
            v31 := GenS3{F0: "\n\t", F1: 165, F2: 150}
            _ = fi30
            _ = v31
        }
        if ((v2 * 18446744073709551615) <= v2) {
            v32 := GenS7{F0: 42}
            _ = v32
        } else {
            v6 = ((v6 * (float32(1e10) + v6)) - float32(3.14159))
            if (-1 > int16(v4)) {
                v4 -= -76
            } else {
                if (false || (true && (!(1 != -1)))) {
                    v6 += float32(-2.25)
                    switch -689 {
                    case 911:
                        v4 = v4
                    default:
                        var v33 int8 = (v4 - v4)
                        _ = v33
                    }
                    v34 := GenS2{F0: 987, F1: 652}
                    _ = v34
                } else {
                    var v35 float32 = v6
                    _ = v35
                }
                _ = len(v3)
            }
        }
    } else {
        v2 -= (v2 + (v2 / ((v2) | 1)))
        fmt.Println(v4, v2)
        _ = len(v3)
    }
    if len(v3) > 0 {
        v3[0] = 42
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
