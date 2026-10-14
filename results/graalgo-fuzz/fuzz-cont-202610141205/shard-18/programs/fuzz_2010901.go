package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int8; F1 int8 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint8; F1 uint16; F2 uint32 }
type GenS6 struct { F0 uint16; F1 bool }
type GenS7 struct { F0 int32; F1 uint16 }
type GenS8 struct { F0 int16; F1 string; F2 int8 }
type GenS9 struct { F0 uint16 }
type GenS10 struct { F0 uint16; F1 uint8 }

func main() {
    var v1 int64 = 1
    v2 := []uint32{197, 288}
    var v3 uint = 18446744073709551615
    v4 := make(map[uint32]int)
    var v5 uint = (^(v3))
    fmt.Println(v3, v1, v1)
    fmt.Println(v5, v1)
    var v6 uint32 = 0
    if false {
        v7 := v4[(v6 * 833)]
        v3 += uint(v6)
        var v8 bool = ((-2.25 <= -2.25) && false)
        _ = v7
        _ = v8
    } else {
        _ = len(v2)
        fmt.Println(v1, v3, v6)
        for fi9 := 0; fi9 < 15; fi9++ {
            v6 = (^(v6))
            v1 = v1
            for fi10 := 0; fi10 < 4; fi10++ {
                var v11 int64 = v1
                _ = fi10
                _ = v11
            }
            _ = fi9
        }
    }
    v3 = v3
    v12 := GenS1{F0: "hello", F1: v1}
    v13 := GenS1{F0: (("" + ("\n\t" + "foo bar")) + ""), F1: v1}
    _ = len(v2)
    if true {
        v1 = v1
        var v14 int32 = 2147483647
        for fi15 := 0; fi15 < 10; fi15++ {
            if false {
                v3 = uint(v6)
                v3 = 379
                switch ("a" + "a") {
                case "hello":
                    v1 = int64(v5)
                case "\n\t":
                    if (0 < 490) {
                        v16 := GenS2{F0: "hello"}
                        if (-12 > 0) {
                            for fi17 := 0; fi17 < 1; fi17++ {
                                if false {
                                    if ((fi15 >> (((v5 >> (((v3 - 42)) & 63))) & 63)) != (fi15 << ((888) & 63))) {
                                        var v18 uint32 = v6
                                        _ = v18
                                    }
                                    v19 := GenS3{F0: 87, F1: -128}
                                    _ = v19
                                }
                                v20 := GenS3{F0: 48, F1: 42}
                                v21 := GenS4{F0: v14}
                                _ = len(v4)
                                _ = fi17
                                _ = v20
                                _ = v21
                            }
                        } else {
                            var v22 uint8 = 153
                            for fi23 := 0; fi23 < 3; fi23++ {
                                fmt.Println(v5, v6)
                                switch ("foo bar" + "\n\t") {
                                case "hello":
                                    v24 := GenS1{F0: ((("\n\t" + "foo bar") + "foo bar") + "hello"), F1: (v1 >> (((v3 / (((v3 / ((942) | 1))) | 1))) & 63))}
                                    _ = v24
                                case "foo bar":
                                    v22 -= uint8(v6)
                                case "a":
                                    if false {
                                        v25 := GenS2{F0: "\n\t"}
                                        v6 = (v6 + v6)
                                        v26 := GenS1{F0: "foo bar", F1: v1}
                                        _ = v25
                                        _ = v26
                                    } else {
                                        switch int(v14) {
                                        case 127:
                                            v1 += v1
                                        case 128:
                                            v1 -= -1
                                        }
                                    }
                                    var v27 int32 = 0
                                    _ = v27
                                default:
                                    v4[(v6 ^ (v6 | v6))] = (fi15 | (fi15 * (^(fi15))))
                                }
                                v6 = (v6 >> (((v5 >> ((4294967295) & 63))) & 63))
                                v14 = (-(v14))
                                var v28 int64 = (v1 & v1)
                                _ = fi23
                                _ = v28
                            }
                            _ = v22
                        }
                        _ = v16
                    } else {
                        v29 := GenS5{F0: 127, F1: 0, F2: v6}
                        _ = v29
                    }
                default:
                    if false {
                        if true {
                            switch "" {
                            case "a":
                                v2 = append(v2, (v6 / ((uint32(v3)) | 1)))
                            case "":
                                var v30 float64 = float64(float32((0.0 - (1e10 / 1.0))))
                                _ = v30
                            default:
                                var v31 uint = (v5 + v3)
                                _ = v31
                            }
                            v32 := GenS4{F0: v14}
                            switch ("\n\t" + ("\n\t" + "\n\t")) {
                            case "foo bar":
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            case "hello":
                                var v33 uint32 = 4294967295
                                _ = v33
                            case "a":
                                v5 = (v3 + v3)
                            default:
                                fmt.Println(v3)
                            }
                            _ = v32
                        } else {
                            switch fi15 {
                            case -404:
                                v3 -= (v5 + (v5 & (v3 >> ((v5) & 63))))
                            case 966:
                                var v34 int64 = (v1 | -828)
                                _ = v34
                            }
                            if len(v2) > 0 {
                                v2[0] = (v6 % (((v6 + (^(v6)))) | 1))
                            }
                        }
                        if (true || false) {
                            if true {
                                if (true && (-2.25 <= 1.5)) {
                                    var v35 float32 = (0.1 / 1.0)
                                    for fi36 := 0; fi36 < 3; fi36++ {
                                        v6 -= v6
                                        if (!true) {
                                            var v37 int8 = 56
                                            _ = v37
                                        }
                                        _ = fi36
                                    }
                                    _ = v35
                                } else {
                                    switch ("" + (("hello" + "a") + "\n\t")) {
                                    case "":
                                        for fi38 := 0; fi38 < 18; fi38++ {
                                            if true {
                                                if (!false) {
                                                    var v39 float32 = float32(float64((1.5 + (0.1 * 1e10))))
                                                    v40 := GenS6{F0: 623, F1: (uint16(v3) < 683)}
                                                    v6 = v6
                                                    _ = v39
                                                    _ = v40
                                                }
                                                v3 += v3
                                            }
                                            _ = fi38
                                        }
                                    default:
                                        v41 := GenS7{F0: v14, F1: 556}
                                        _ = v41
                                    }
                                    for fi42 := 0; fi42 < 3; fi42++ {
                                        v2 = append(v2, (v6 * v6))
                                        if false {
                                            _ = len(v2)
                                        } else {
                                            v43 := GenS8{F0: -1, F1: "", F2: 42}
                                            delete(v4, 65535)
                                            _ = v43
                                        }
                                        _ = fi42
                                    }
                                }
                                if false {
                                    if (true || true) {
                                        v4[v6] = 42
                                        v6 = (v6 % (((v6 << (((v3 & (v3 ^ v3))) & 63))) | 1))
                                    } else {
                                        switch fi15 {
                                        case 1:
                                            var v44 uint64 = 42
                                            _ = v44
                                        case 979:
                                            fmt.Println(v14, fi15)
                                        }
                                    }
                                    var v45 int8 = -19
                                    v46 := GenS6{F0: 2, F1: false}
                                    _ = v45
                                    _ = v46
                                } else {
                                    v14 -= v14
                                }
                            }
                            fmt.Println(v6, v1)
                        }
                    } else {
                        v47 := GenS8{F0: int16(v6), F1: "hello", F2: -113}
                        fmt.Println(v1)
                        v1 += (^(v1))
                        _ = v47
                    }
                }
            }
            _ = fi15
        }
        _ = v14
    } else {
        if ((1 > -113) || (int(v6) > 0)) {
            var v48 uint16 = 1
            v49 := GenS9{F0: (v48 / (((v48 & (^(v48)))) | 1))}
            v3 -= 18446744073709551615
            _ = v48
            _ = v49
        } else {
            v50 := GenS6{F0: 32767, F1: false}
            v51 := GenS10{F0: 158, F1: 1}
            v5 -= 42
            _ = v50
            _ = v51
        }
        for fi52 := 0; fi52 < 1; fi52++ {
            if (!false) {
                if (v6 == (v6 - (v6 * (^(v6))))) {
                    v3 = (^(v3))
                    var v53 int64 = -9223372036854775808
                    if (v6 != v6) {
                        v6 -= (v6 << (((v3 % ((v3) | 1))) & 63))
                    } else {
                        switch "" {
                        case "foo bar":
                            var v54 uint8 = 255
                            _ = v54
                        case "hello":
                            var v55 bool = true
                            _ = v55
                        case "a":
                            if (0 != 1) {
                                v3 = v3
                                var v56 uint16 = 32767
                                v53 += (v1 % ((v1) | 1))
                                _ = v56
                            } else {
                                v1 = v1
                            }
                        default:
                            for fi57 := 0; fi57 < 5; fi57++ {
                                fmt.Println(v1, fi52)
                                _ = fi57
                            }
                        }
                        var v58 int16 = 1
                        _ = v58
                    }
                    _ = v53
                } else {
                    fmt.Println(v3, v6)
                }
            }
            _ = fi52
        }
        if len(v2) > 0 {
            v2[0] = (v6 >> ((984) & 63))
        }
    }
    var v59 string = "hello"
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v59:", v59)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v59
    _ = v6
}
