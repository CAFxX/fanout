package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint; F1 string }
type GenS3 struct { F0 uint16; F1 int; F2 int16 }
type GenS4 struct { F0 int32 }

func main() {
    v1 := []string{"foo bar", "", "hello"}
    var v2 int32 = 42
    var v3 float32 = 0.1
    var v4 int32 = (v2 ^ 0)
    var v5 float32 = 1e10
    if true {
        fmt.Println(v2)
        if (("\n\t" + "hello") < "") {
            v3 = v3
            for fi6 := 0; fi6 < 14; fi6++ {
                var v7 uint = 0
                if false {
                    v4 = v4
                    var v8 uint = v7
                    v5 = v5
                    _ = v8
                }
                for fi9 := 0; fi9 < 4; fi9++ {
                    if true {
                        v4 += (v4 >> ((18446744073709551615) & 63))
                        fmt.Println(fi6, v3, v2)
                    }
                    fmt.Println(v2, v4, v5)
                    v4 = -788
                    _ = fi9
                }
                _ = fi6
                _ = v7
            }
        } else {
            v2 = -1
            v1 = append(v1, "a")
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    } else {
        var v10 int8 = -128
        v4 -= -854
        var v11 uint64 = 166
        _ = v10
        _ = v11
    }
    var v12 int64 = -1
    fmt.Println(v12, v5, v3)
    v13 := GenS1{F0: -688}
    switch "\n\t" {
    case "foo bar":
        v4 = v2
    }
    v2 += v2
    for fi14 := 0; fi14 < 12; fi14++ {
        if ((^(fi14)) != -1) {
            v15 := GenS2{F0: 996, F1: ("foo bar" + ("foo bar" + ("hello" + ("foo bar" + "hello"))))}
            switch 32767 {
            case -1:
                v16 := GenS1{F0: (^(v2))}
                _ = v16
            case 32768:
                var v17 float32 = ((v3 + 1.5) + 1.5)
                _ = v17
            case 9223372036854775807:
                if (!(false && true)) {
                    for fi18 := 0; fi18 < 12; fi18++ {
                        var v19 int8 = 100
                        v20 := GenS1{F0: -1}
                        _ = fi18
                        _ = v19
                        _ = v20
                    }
                    var v21 float64 = (1e10 / ((1.5 - 0.1) - 1.5))
                    switch "a" {
                    case "\n\t":
                        fmt.Println(v3, v21, v12)
                    }
                    _ = v21
                } else {
                    if false {
                        if false {
                            if true {
                                _ = len(v1)
                                var v22 int64 = (v12 << ((1) & 63))
                                _ = v22
                            }
                            v1 = append(v1, "\n\t")
                        } else {
                            var v23 uint16 = 2
                            var v24 uint = 4294967295
                            _ = v23
                            _ = v24
                        }
                    } else {
                        v5 -= v5
                        v3 -= ((1e10 - 3.14159) / (v3 - 1.5))
                        v25 := GenS3{F0: 582, F1: fi14, F2: 562}
                        _ = v25
                    }
                }
            default:
                switch fi14 {
                case -9223372036854775808:
                    for fi26 := 0; fi26 < 20; fi26++ {
                        var v27 uint64 = 42
                        var v28 uint32 = 438
                        _ = fi26
                        _ = v27
                        _ = v28
                    }
                case -1:
                    var v29 uint32 = 1
                    _ = v29
                default:
                    if false {
                        for fi30 := 0; fi30 < 8; fi30++ {
                            if (71 != 744) {
                                if len(v1) > 0 {
                                    v1[0] = "hello"
                                }
                                v3 += (v3 / ((((v3 + ((-2.25 + ((v3 * 0.1) / float32(3.14159))) * float32(((1e10 / 1.0) * float64((v3 * v3)))))) + v5) * (float32(0.1) * 3.14159)) - v5))
                                var v31 bool = true
                                _ = v31
                            }
                            _ = fi30
                        }
                        if (2 < 128) {
                            v5 = (3.14159 / ((((v3 + v5) + v5) / ((0.1 + ((v3 - 0.0) - v5)) + v5)) - v5))
                            v32 := GenS4{F0: -1}
                            var v33 uint8 = 2
                            _ = v32
                            _ = v33
                        } else {
                            v5 += 0.1
                            var v34 int32 = v2
                            _ = v34
                        }
                    }
                }
            }
            _ = v15
        }
        _ = fi14
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
