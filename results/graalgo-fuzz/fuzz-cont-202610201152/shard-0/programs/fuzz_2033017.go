package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint16; F1 uint32; F2 int8 }
type GenS3 struct { F0 bool; F1 int8 }
type GenS4 struct { F0 int8; F1 string; F2 int16 }

func main() {
    v1 := []uint32{104, 2147483648, 0}
    var v2 uint16 = 311
    var v3 uint8 = uint8(v2)
    var v4 uint32 = 248
    v5 := make(map[string]int)
    if (-1 > 42) {
        for fi6 := 0; fi6 < 14; fi6++ {
            var v7 int8 = 0
            _ = fi6
            _ = v7
        }
    }
    v5["hello"] = -570
    for fi8 := 0; fi8 < 13; fi8++ {
        for fi9 := 0; fi9 < 20; fi9++ {
            var v10 uint = 227
            switch "" {
            case "":
                switch 9223372036854775807 {
                case 0:
                    v4 -= v4
                case 9223372036854775807:
                    v1 = append(v1, 2)
                case 745:
                    var v11 float64 = float64((0.1 - 0.0))
                    _ = v11
                }
            case "hello":
                fmt.Println(v3)
            default:
                var v12 uint = v10
                _ = v12
            }
            _ = fi9
            _ = v10
        }
        _ = fi8
    }
    if (false || true) {
        v13 := v5["foo bar"]
        v14 := v5["hello"]
        fmt.Println(v14)
        _ = v13
        _ = v14
    } else {
        if ((true && true) && (0 > 42)) {
            v15 := GenS1{F0: false}
            _ = v15
        } else {
            v4 += 649
            v4 = v4
        }
    }
    var v16 int16 = 1
    switch 128 {
    case 42:
        switch ("" + ("foo bar" + ("hello" + (("foo bar" + "foo bar") + "hello")))) {
        case "hello":
            if false {
                v3 += (v3 / ((v3) | 1))
                fmt.Println(v4, v4)
                _ = len(v5)
            }
        case "":
            v17 := GenS2{F0: (v2 >> ((0) & 63)), F1: v4, F2: -4}
            _ = v17
        }
    case 111:
        var v18 uint32 = 2147483647
        _ = v18
    }
    if (9223372036854775807 <= -656) {
        fmt.Println(v3, v3, v3)
        v3 -= v3
        fmt.Println(v2, v3)
    }
    if len(v1) > 0 {
        v1[0] = (v4 * (v4 ^ (v4 - 0)))
    }
    delete(v5, "hello")
    if false {
        if len(v1) > 0 {
            v1[0] = v4
        }
        fmt.Println(v4)
        for fi19 := 0; fi19 < 19; fi19++ {
            for fi20 := 0; fi20 < 1; fi20++ {
                if (!true) {
                    v16 = v16
                }
                fmt.Println(v16, v3)
                _ = fi20
            }
            _ = fi19
        }
    } else {
        v21 := GenS1{F0: false}
        for fi22 := 0; fi22 < 3; fi22++ {
            switch 127 {
            case 9223372036854775807:
                var v23 uint8 = 2
                _ = v23
            case 0:
                if false {
                    v24 := GenS3{F0: (true && false), F1: -128}
                    var v25 int8 = 0
                    for fi26 := 0; fi26 < 17; fi26++ {
                        v25 += -3
                        var v27 int = fi22
                        if ((v16 >> ((uint(v4)) & 63)) < int16(v27)) {
                            v2 = (v2 & (v2 % ((549) | 1)))
                            var v28 float32 = 0.0
                            if (!false) {
                                v29 := v5[("hello" + "hello")]
                                v28 -= (1.5 / 1.0)
                                _ = v29
                            }
                            var v30 uint32 = uint32(v3)
                            v25 += (-(v25))
                            _ = v28
                            _ = v30
                        }
                        _ = fi26
                        _ = v27
                    }
                    _ = v24
                    _ = v25
                } else {
                    v2 = (v2 ^ v2)
                }
            default:
                if false {
                    for fi31 := 0; fi31 < 14; fi31++ {
                        var v32 uint16 = (^(v2))
                        _ = fi31
                        _ = v32
                    }
                } else {
                    if false {
                        for fi33 := 0; fi33 < 19; fi33++ {
                            var v34 int16 = v16
                            _ = fi33
                            _ = v34
                        }
                    } else {
                        v35 := GenS2{F0: (v2 * (^(v2))), F1: v4, F2: int8(v4)}
                        v36 := GenS1{F0: true}
                        _ = v35
                        _ = v36
                    }
                    delete(v5, ("a" + ""))
                    switch (fi22 ^ (fi22 | (fi22 >> ((520) & 63)))) {
                    case -2147483648:
                        switch "hello" {
                        case "\n\t":
                            var v37 uint64 = 953
                            _ = v37
                        case "foo bar":
                            var v38 bool = true
                            _ = v38
                        default:
                            v39 := GenS3{F0: (true || (!true)), F1: 90}
                            _ = v39
                        }
                    case 32768:
                        v3 += v3
                    case 0:
                        for fi40 := 0; fi40 < 8; fi40++ {
                            var v41 int64 = -13
                            v42 := GenS1{F0: false}
                            if false {
                                v43 := GenS4{F0: -78, F1: "hello", F2: (v16 | v16)}
                                if (false && (v16 < v16)) {
                                    switch fi22 {
                                    case 128:
                                        var v44 int8 = 42
                                        _ = v44
                                    case -1:
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                    default:
                                        var v45 int64 = v41
                                        _ = v45
                                    }
                                }
                                _ = v43
                            }
                            var v46 uint8 = uint8(v4)
                            _ = fi40
                            _ = v41
                            _ = v42
                            _ = v46
                        }
                    }
                }
            }
            v47 := v5["foo bar"]
            delete(v5, "hello")
            _ = fi22
            _ = v47
        }
        _ = v21
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
