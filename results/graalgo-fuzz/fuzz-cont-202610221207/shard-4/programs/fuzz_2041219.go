package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8 }
type GenS2 struct { F0 uint16; F1 uint32 }
type GenS3 struct { F0 int16; F1 uint }
type GenS4 struct { F0 uint }

func main() {
    v1 := make(map[string]int)
    v2 := make(map[string]uint32)
    v3 := make(map[string]bool)
    var v4 uint32 = 249
    var v5 uint = 18446744073709551615
    v4 -= 2
    var v6 int = -2147483648
    var v7 int8 = -128
    for fi8 := 0; fi8 < 14; fi8++ {
        switch "" {
        case "\n\t":
            delete(v1, "a")
        default:
            if (128 > 247) {
                for fi9 := 0; fi9 < 16; fi9++ {
                    fmt.Println(v5)
                    _ = fi9
                }
                switch fi8 {
                case 127:
                    for fi10 := 0; fi10 < 13; fi10++ {
                        var v11 int32 = -2147483648
                        _ = fi10
                        _ = v11
                    }
                case 317:
                    var v12 int16 = 459
                    _ = v12
                }
                if (90 < 2) {
                    for fi13 := 0; fi13 < 14; fi13++ {
                        if false {
                            v14 := GenS1{F0: (v5 / ((v5) | 1)), F1: v7}
                            v15 := []int{(fi8 | -915), fi13, 128, (fi8 | (-(fi8)))}
                            switch "a" {
                            case "\n\t":
                                v6 = (v6 + v6)
                            case "foo bar":
                                var v16 uint8 = 42
                                _ = v16
                            case "hello":
                                v6 = 32768
                            }
                            _ = v14
                            _ = v15
                        }
                        _ = fi13
                    }
                    switch "" {
                    case "foo bar":
                        if ((v5 != v5) && ((v7 - v7) > (v7 << ((v5) & 63)))) {
                            v17 := GenS1{F0: v5, F1: 74}
                            for fi18 := 0; fi18 < 12; fi18++ {
                                v19 := GenS2{F0: 255, F1: v4}
                                var v20 int = (fi8 & (v6 >> ((v5) & 63)))
                                _ = fi18
                                _ = v19
                                _ = v20
                            }
                            _ = v17
                        } else {
                            for fi21 := 0; fi21 < 17; fi21++ {
                                v22 := GenS1{F0: (v5 | (v5 % ((v5) | 1))), F1: 1}
                                _ = fi21
                                _ = v22
                            }
                            var v23 float32 = (3.14159 * (0.0 + (1e10 - (-2.25 - 1e10))))
                            var v24 uint8 = 1
                            _ = v23
                            _ = v24
                        }
                    case "":
                        _ = len(v1)
                    }
                    switch "hello" {
                    case "hello":
                        v25 := GenS1{F0: (v5 & (v5 << ((979) & 63))), F1: -128}
                        _ = v25
                    case "":
                        fmt.Println(fi8, v7, v5)
                    case "\n\t":
                        if (0 <= 0) {
                            v26 := GenS1{F0: 255, F1: (v7 ^ 100)}
                            _ = v26
                        } else {
                            for fi27 := 0; fi27 < 4; fi27++ {
                                var v28 int16 = -541
                                var v29 int = v6
                                if (!(!(740 < 0))) {
                                    v29 = (fi8 / ((v29) | 1))
                                    v30 := v3["\n\t"]
                                    var v31 string = "\n\t"
                                    _ = v30
                                    _ = v31
                                }
                                _ = fi27
                                _ = v28
                                _ = v29
                            }
                            v1["a"] = v6
                        }
                    default:
                        fmt.Println(fi8, fi8, v4)
                    }
                }
            } else {
                switch ("foo bar" + ("hello" + "")) {
                case "\n\t":
                    var v32 int64 = 1
                    _ = v32
                default:
                    switch "" {
                    case "a":
                        v3[("\n\t" + "")] = true
                    case "hello":
                        v6 -= v6
                    default:
                        if true {
                            v33 := GenS3{F0: 42, F1: v5}
                            if (128 < 200) {
                                fmt.Println(v4, fi8)
                                v34 := v2[("hello" + ("" + "hello"))]
                                _ = v34
                            }
                            _ = v33
                        } else {
                            if false {
                                v35 := GenS1{F0: (v5 | v5), F1: 0}
                                _ = v35
                            } else {
                                var v36 float32 = float32(0.1)
                                _ = v36
                            }
                        }
                    }
                }
                for fi37 := 0; fi37 < 16; fi37++ {
                    v38 := []int{(v6 >> ((v5) & 63))}
                    var v39 int64 = 9223372036854775807
                    v4 += uint32(v5)
                    var v40 uint = 4294967295
                    fmt.Println(v40, v39, fi8)
                    _ = fi37
                    _ = v38
                    _ = v39
                    _ = v40
                }
            }
        }
        _ = fi8
    }
    switch v6 {
    case 2147483647:
        v2[""] = v4
    case -9223372036854775808:
        v5 += v5
    default:
        var v41 int64 = -519
        _ = v41
    }
    for fi42 := 0; fi42 < 12; fi42++ {
        v43 := GenS1{F0: v5, F1: v7}
        _ = fi42
        _ = v43
    }
    v5 = (v5 / (((v5 * v5)) | 1))
    fmt.Println(v4)
    v1[(("a" + "\n\t") + "")] = (v6 << (((v5 | v5)) & 63))
    v44 := GenS4{F0: (v5 ^ 359)}
    if true {
        v6 -= -2147483648
        for fi45 := 0; fi45 < 2; fi45++ {
            fmt.Println(v6, v6, v7)
            _ = fi45
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v5
    _ = v6
    _ = v7
}
