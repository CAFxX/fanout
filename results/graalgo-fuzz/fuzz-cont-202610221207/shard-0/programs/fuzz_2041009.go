package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint32; F1 uint16; F2 int }

func main() {
    v1 := []string{("\n\t" + ("foo bar" + ("\n\t" + ("a" + "a"))))}
    v2 := make(map[string]string)
    var v3 int16 = -32768
    var v4 int8 = 81
    v5 := GenS1{F0: 128}
    v4 = -1
    for fi6 := 0; fi6 < 6; fi6++ {
        if true {
            v3 = (v3 - 1000)
        } else {
            for fi7 := 0; fi7 < 4; fi7++ {
                var v8 int16 = v3
                _ = fi7
                _ = v8
            }
            for fi9 := 0; fi9 < 5; fi9++ {
                if true {
                    v4 -= (-(v4))
                    _ = len(v2)
                }
                if len(v1) > 0 {
                    v1[0] = "\n\t"
                }
                _ = fi9
            }
        }
        for fi10 := 0; fi10 < 8; fi10++ {
            fmt.Println(v3)
            _ = fi10
        }
        _ = fi6
    }
    var v11 int32 = 236
    var v12 uint32 = 1
    var v13 string = "hello"
    for fi14 := 0; fi14 < 2; fi14++ {
        v11 = v11
        _ = len(v2)
        if true {
            if (true || (0 == 9223372036854775807)) {
                if ((3.14159 * 0.0) != float64(0.0)) {
                    for fi15 := 0; fi15 < 2; fi15++ {
                        switch (fi15 ^ fi14) {
                        case 764:
                            v16 := GenS1{F0: 128}
                            _ = v16
                        case 2147483647:
                            fmt.Println(v13, v12, v3)
                        case -9223372036854775808:
                            v13 = v13
                        default:
                            _ = len(v1)
                        }
                        _ = fi15
                    }
                    switch (fi14 << ((0) & 63)) {
                    case 32768:
                        switch v13 {
                        case "a":
                            if ((255 <= 255) && true) {
                                v17 := GenS2{F0: v3}
                                switch ((v13 + v13) + v13) {
                                case "\n\t":
                                    for fi18 := 0; fi18 < 6; fi18++ {
                                        _ = len(v2)
                                        _ = fi18
                                    }
                                }
                                var v19 int = (fi14 << ((255) & 63))
                                _ = v17
                                _ = v19
                            }
                        case "foo bar":
                            for fi20 := 0; fi20 < 5; fi20++ {
                                switch v13 {
                                case "":
                                    var v21 uint32 = (v12 & 0)
                                    _ = v21
                                case "foo bar":
                                    fmt.Println(v13, v12, v13)
                                default:
                                    switch (fi14 >> ((2) & 63)) {
                                    case 1:
                                        v12 += v12
                                    case -113:
                                        v11 = v11
                                    default:
                                        fmt.Println(fi20)
                                    }
                                }
                                v3 += -554
                                _ = fi20
                            }
                        default:
                            fmt.Println(v12)
                        }
                    case 1:
                        switch (fi14 / (((fi14 >> ((407) & 63))) | 1)) {
                        case -748:
                            var v22 bool = (false || false)
                            _ = v22
                        case -9223372036854775808:
                            for fi23 := 0; fi23 < 7; fi23++ {
                                var v24 float64 = float64(1e10)
                                _ = fi23
                                _ = v24
                            }
                        case 127:
                            fmt.Println(v4, fi14, v12)
                        default:
                            var v25 uint8 = 125
                            _ = v25
                        }
                    default:
                        if true {
                            v26 := GenS3{F0: v12, F1: 46, F2: fi14}
                            v27 := GenS3{F0: (^(v12)), F1: 795, F2: fi14}
                            var v28 int = (fi14 ^ 127)
                            _ = v26
                            _ = v27
                            _ = v28
                        } else {
                            fmt.Println(v4, v11, v4)
                            switch fi14 {
                            case 127:
                                var v29 float32 = float32(float64(1.5))
                                _ = v29
                            }
                        }
                    }
                    if (false || ((v11 << ((65535) & 63)) < v11)) {
                        v11 += 2147483647
                        fmt.Println(v13, v12, v13)
                    }
                } else {
                    for fi30 := 0; fi30 < 2; fi30++ {
                        v31 := GenS3{F0: v12, F1: 391, F2: 941}
                        _ = fi30
                        _ = v31
                    }
                }
            } else {
                _ = len(v1)
            }
            var v32 int = fi14
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = v32
        } else {
            switch v13 {
            case "foo bar":
                fmt.Println(v13, v3, v11)
            case "\n\t":
                v13 = v13
            case "":
                if false {
                    for fi33 := 0; fi33 < 17; fi33++ {
                        v1 = append(v1, v13)
                        if len(v1) > 0 {
                            v1[0] = (v13 + "foo bar")
                        }
                        _ = fi33
                    }
                }
            default:
                fmt.Println(v4, v3)
            }
        }
        _ = fi14
    }
    for fi34 := 0; fi34 < 18; fi34++ {
        v12 = v12
        _ = len(v1)
        v12 = v12
        _ = fi34
    }
    v13 = ""
    var v35 float64 = 1.5
    v35 = (v35 * (v35 / float64(1e10)))
    v1 = append(v1, v13)
    for fi36 := 0; fi36 < 4; fi36++ {
        var v37 float32 = (0.1 + 0.1)
        v37 -= (v37 - v37)
        _ = fi36
        _ = v37
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v5
}
