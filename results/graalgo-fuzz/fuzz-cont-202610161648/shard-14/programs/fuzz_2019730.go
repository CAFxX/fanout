package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32 }
type GenS2 struct { F0 int8 }

func main() {
    v1 := make(map[int]int)
    var v2 float64 = 0.1
    v3 := []uint32{2147483648, 107, 65535}
    var v4 uint16 = 1
    var v5 int16 = 412
    var v6 string = "hello"
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v7 float32 = -2.25
    for fi8 := 0; fi8 < 10; fi8++ {
        fmt.Println(v2)
        fmt.Println(v6, v2)
        _ = fi8
    }
    var v9 int = -112
    v4 += v4
    var v10 uint16 = (^(v4))
    fmt.Println(v5)
    v2 = 0.1
    if false {
        var v11 uint16 = v10
        var v12 uint64 = 4294967295
        _ = v11
        _ = v12
    }
    v7 -= (v7 / (float32(v2) - v7))
    if (9223372036854775807 == -1) {
        var v13 bool = false
        v14 := GenS1{F0: 2, F1: 1}
        _ = v13
        _ = v14
    }
    if (int8(v5) > 126) {
        for fi15 := 0; fi15 < 19; fi15++ {
            v4 = 42
            var v16 uint64 = 0
            _ = fi15
            _ = v16
        }
        delete(v1, (v9 >> ((102) & 63)))
    } else {
        v17 := GenS1{F0: 255, F1: -2147483648}
        if (0 == v5) {
            if (42 < 1) {
                delete(v1, 0)
                for fi18 := 0; fi18 < 7; fi18++ {
                    switch v9 {
                    case 128:
                        v19 := GenS1{F0: 128, F1: 135}
                        _ = v19
                    case 4294967295:
                        switch v6 {
                        case "a":
                            switch ((v6 + (v6 + "foo bar")) + (v6 + v6)) {
                            case "":
                                for fi20 := 0; fi20 < 10; fi20++ {
                                    v2 += float64(float32(v2))
                                    v21 := GenS1{F0: 42, F1: -888}
                                    _ = fi20
                                    _ = v21
                                }
                            case "\n\t":
                                var v22 uint64 = 42
                                _ = v22
                            }
                        case "":
                            switch (fi18 >> ((42) & 63)) {
                            case 816:
                                fmt.Println(v4)
                            case 127:
                                if (0.1 <= (v2 * v2)) {
                                    for fi23 := 0; fi23 < 4; fi23++ {
                                        _ = len(v1)
                                        _ = fi23
                                    }
                                    if ("foo bar" <= v6) {
                                        var v24 int8 = 42
                                        var v25 float32 = ((v7 - (((v7 - v7) + float32((1.5 * v2))) - 1e10)) / (v7 + 0.0))
                                        _ = v24
                                        _ = v25
                                    }
                                } else {
                                    for fi26 := 0; fi26 < 14; fi26++ {
                                        v4 -= 0
                                        v10 -= v4
                                        var v27 uint8 = 1
                                        _ = fi26
                                        _ = v27
                                    }
                                    v28 := GenS1{F0: 200, F1: -2147483648}
                                    _ = v28
                                }
                            case -2147483648:
                                v6 = v6
                            default:
                                if true {
                                    v29 := v1[fi18]
                                    switch v9 {
                                    case 42:
                                        switch v29 {
                                        case 1:
                                            v6 = ((("\n\t" + v6) + "foo bar") + v6)
                                        case 127:
                                            if (((!false) && (!(!(false && (80 != 42))))) || true) {
                                                switch v29 {
                                                case 9223372036854775807:
                                                    var v30 uint32 = 0
                                                    _ = v30
                                                case 597:
                                                    v31 := GenS2{F0: 42}
                                                    _ = v31
                                                }
                                            } else {
                                                delete(v1, v9)
                                                v6 = v6
                                            }
                                        case 9223372036854775807:
                                            v3 = append(v3, 677)
                                        }
                                    case -321:
                                        v4 -= 0
                                    default:
                                        switch fi18 {
                                        case 30:
                                            v3 = append(v3, 993)
                                        case 9223372036854775807:
                                            v2 = (-2.25 - v2)
                                        default:
                                            v32 := GenS1{F0: 209, F1: 151}
                                            _ = v32
                                        }
                                    }
                                    _ = v29
                                }
                            }
                        default:
                            var v33 uint64 = 1
                            _ = v33
                        }
                    default:
                        switch v6 {
                        case "\n\t":
                            if ((300 != 1) || true) {
                                v9 += (fi18 & (fi18 / ((v9) | 1)))
                                if len(v3) > 0 {
                                    v3[0] = 558
                                }
                                var v34 uint32 = 65535
                                _ = v34
                            }
                        case "foo bar":
                            fmt.Println(v4, fi18, v5)
                        case "hello":
                            fmt.Println(v9, v4, v2)
                        default:
                            v4 += v4
                        }
                    }
                    _ = fi18
                }
                switch v6 {
                case "\n\t":
                    v1[(v9 / ((v9) | 1))] = (-(v9))
                case "a":
                    v7 -= v7
                case "":
                    var v35 int = v9
                    _ = v35
                }
            }
            switch v6 {
            case "":
                var v36 int16 = (v5 + 1)
                _ = v36
            }
        }
        v37 := GenS2{F0: -1}
        _ = v17
        _ = v37
    }
    var v38 int16 = 32767
    v39 := GenS1{F0: 255, F1: -2147483648}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
