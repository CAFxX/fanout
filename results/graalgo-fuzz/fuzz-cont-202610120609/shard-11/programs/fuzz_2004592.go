package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 string; F1 uint16 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int32; F1 string; F2 string }

func main() {
    v1 := make(map[int]int)
    v2 := make(map[int]string)
    v3 := []int{}
    var v4 int = 42
    var v5 uint32 = 767
    var v6 uint64 = 129
    var v7 uint32 = v5
    if false {
        var v8 uint64 = 1
        if true {
            for fi9 := 0; fi9 < 1; fi9++ {
                v7 = (v5 + v7)
                _ = fi9
            }
            v10 := GenS1{F0: -631}
            _ = len(v3)
            _ = v10
        } else {
            if false {
                v4 -= 127
            }
            switch "" {
            case "hello":
                if false {
                    v11 := GenS1{F0: -171}
                    var v12 string = "foo bar"
                    _ = v11
                    _ = v12
                } else {
                    v13 := GenS1{F0: -32768}
                    _ = v13
                }
            case "foo bar":
                for fi14 := 0; fi14 < 15; fi14++ {
                    v15 := GenS2{F0: 85}
                    _ = fi14
                    _ = v15
                }
            default:
                v7 += v5
            }
            _ = len(v3)
        }
        for fi16 := 0; fi16 < 4; fi16++ {
            for fi17 := 0; fi17 < 12; fi17++ {
                switch "foo bar" {
                case "\n\t":
                    if true {
                        fmt.Println(v7)
                        var v18 string = "foo bar"
                        switch "\n\t" {
                        case "hello":
                            switch v18 {
                            case "foo bar":
                                v19 := GenS3{F0: "\n\t", F1: 32767}
                                _ = v19
                            case "hello":
                                fmt.Println(v7)
                            default:
                                v20 := v2[-2147483648]
                                _ = v20
                            }
                        case "":
                            fmt.Println(v18, v4, v8)
                        default:
                            fmt.Println(v6, v18, v18)
                        }
                        _ = v18
                    }
                case "hello":
                    v21 := GenS2{F0: 127}
                    _ = v21
                }
                var v22 int8 = 42
                if false {
                    switch 42 {
                    case 127:
                        if false {
                            if false {
                                v1[(-(fi17))] = fi16
                                v23 := GenS4{F0: (v6 | (v6 - 1))}
                                delete(v1, 1)
                                v4 = (fi16 ^ 663)
                                v24 := GenS3{F0: "hello", F1: 65535}
                                _ = v23
                                _ = v24
                            } else {
                                v8 = v6
                                v4 = v4
                                v7 += v5
                            }
                        }
                    case 2147483647:
                        v1[(-(fi17))] = 127
                    default:
                        switch -1 {
                        case 127:
                            if len(v3) > 0 {
                                v3[0] = -318
                            }
                        case 32768:
                            var v25 string = "\n\t"
                            _ = v25
                        case 700:
                            fmt.Println(v4, fi16, v5)
                        }
                    }
                    v7 += v5
                    if false {
                        v4 -= (v4 & (fi17 % (((v4 & 2147483647)) | 1)))
                        if (false || (-2147483648 != -2147483648)) {
                            fmt.Println(v4, fi17, v5)
                            v26 := GenS5{F0: 2147483647, F1: "a", F2: "\n\t"}
                            v27 := GenS5{F0: -282, F1: ("" + "foo bar"), F2: (("a" + ("a" + "\n\t")) + "hello")}
                            _ = v26
                            _ = v27
                        } else {
                            if (((((1e10 - (1e10 + (3.14159 - 3.14159))) * 3.14159) * (3.14159 * 1e10)) + 3.14159) < 0.0) {
                                if true {
                                    var v28 int = 694
                                    if (165 <= 255) {
                                        switch v28 {
                                        case 4294967295:
                                            var v29 int32 = 0
                                            if true {
                                                _ = len(v1)
                                                var v30 bool = ((((((0.1 / ((1e10 / (0.1 / 1.0)) - 0.0)) - 0.0) * (0.1 * 1.5)) - (0.1 * 1e10)) * 0.0) > (1.5 * 0.0))
                                                _ = v30
                                            }
                                            v5 -= 935
                                            _ = v29
                                        default:
                                            v4 = -9223372036854775808
                                        }
                                        v31 := GenS4{F0: 1}
                                        _ = v31
                                    }
                                    _ = v28
                                }
                                v7 = (v7 + (^(v5)))
                                var v32 uint8 = 72
                                _ = v32
                            }
                            _ = len(v3)
                            v7 += 65535
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        }
                    }
                } else {
                    v22 = v22
                    fmt.Println(v5, v8)
                }
                var v33 int16 = 0
                _ = fi17
                _ = v22
                _ = v33
            }
            _ = fi16
        }
        _ = v8
    }
    switch v4 {
    case 301:
        var v34 uint32 = (v5 % ((v7) | 1))
        _ = v34
    case 32768:
        if (523 < -172) {
            var v35 float32 = (float32(float64((((1.5 + 0.0) + 1e10) - (-2.25 + 0.0)))) * float32(0.1))
            _ = v35
        } else {
            for fi36 := 0; fi36 < 19; fi36++ {
                v6 += uint64(v5)
                var v37 uint32 = v7
                for fi38 := 0; fi38 < 16; fi38++ {
                    var v39 float32 = 1e10
                    v4 -= (fi38 - v4)
                    v4 = (^(fi36))
                    fmt.Println(v7, v7)
                    _ = fi38
                    _ = v39
                }
                _ = fi36
                _ = v37
            }
        }
    }
    fmt.Println(v7, v4)
    var v40 uint16 = 427
    v7 = 0
    for fi41 := 0; fi41 < 10; fi41++ {
        if false {
            fmt.Println(v5)
            if true {
                v5 -= (v7 >> ((uint(v5)) & 63))
                v42 := GenS3{F0: (("foo bar" + "hello") + "foo bar"), F1: (v40 / (((v40 & 42)) | 1))}
                _ = v42
            }
        } else {
            fmt.Println(v7)
            v43 := GenS3{F0: "a", F1: (^(v40))}
            _ = v43
        }
        _ = fi41
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v5
    _ = v6
    _ = v7
}
