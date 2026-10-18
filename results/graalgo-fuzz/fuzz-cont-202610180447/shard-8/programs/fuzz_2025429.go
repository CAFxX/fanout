package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 string; F1 uint }
type GenS4 struct { F0 string; F1 uint16; F2 uint64 }

func main() {
    var v1 int32 = -1
    v2 := []uint32{2147483648, 0, 65535, 1}
    var v3 int32 = 1
    var v4 uint8 = 42
    var v5 string = "a"
    var v6 uint64 = 0
    if (!true) {
        fmt.Println(v3, v4, v4)
        if (255 > 42) {
            if true {
                switch 4294967295 {
                case 2147483647:
                    fmt.Println(v3, v1, v4)
                case 0:
                    v7 := GenS1{F0: v1, F1: 875}
                    _ = v7
                case -9223372036854775808:
                    if (v6 > v6) {
                        var v8 string = v5
                        var v9 bool = true
                        _ = v8
                        _ = v9
                    } else {
                        for fi10 := 0; fi10 < 9; fi10++ {
                            switch (fi10 & fi10) {
                            case 127:
                                fmt.Println(v6, v4)
                            case 2147483647:
                                fmt.Println(v4)
                            }
                            v3 -= v3
                            var v11 int64 = 104
                            _ = fi10
                            _ = v11
                        }
                        if (("\n\t" + ("a" + "foo bar")) != v5) {
                            for fi12 := 0; fi12 < 15; fi12++ {
                                v3 = 0
                                _ = fi12
                            }
                            for fi13 := 0; fi13 < 8; fi13++ {
                                for fi14 := 0; fi14 < 20; fi14++ {
                                    var v15 string = ""
                                    v5 = "hello"
                                    _ = fi14
                                    _ = v15
                                }
                                switch v5 {
                                case "foo bar":
                                    v1 = (v3 | v3)
                                case "\n\t":
                                    v6 -= (v6 - (v6 % (((v6 / ((v6) | 1))) | 1)))
                                case "a":
                                    v16 := GenS2{F0: 72}
                                    _ = v16
                                }
                                _ = fi13
                            }
                            for fi17 := 0; fi17 < 6; fi17++ {
                                v18 := GenS1{F0: v1, F1: 114}
                                var v19 int64 = 0
                                if (!true) {
                                    switch fi17 {
                                    case 128:
                                        v5 = "hello"
                                    case -2147483648:
                                        for fi20 := 0; fi20 < 14; fi20++ {
                                            var v21 uint32 = 65535
                                            switch "hello" {
                                            case "foo bar":
                                                v4 = 127
                                            case "\n\t":
                                                v22 := GenS1{F0: (v3 / (((v3 * 42)) | 1)), F1: (v21 * (^(v21)))}
                                                _ = v22
                                            }
                                            _ = fi20
                                            _ = v21
                                        }
                                    }
                                }
                                _ = fi17
                                _ = v18
                                _ = v19
                            }
                        } else {
                            v4 = 127
                        }
                        for fi23 := 0; fi23 < 2; fi23++ {
                            var v24 bool = ((fi23 & fi23) >= int(v6))
                            v25 := GenS3{F0: (v5 + "a"), F1: 18446744073709551615}
                            var v26 int16 = 0
                            _ = fi23
                            _ = v24
                            _ = v25
                            _ = v26
                        }
                    }
                }
                switch -431 {
                case 4294967295:
                    var v27 int64 = -602
                    _ = v27
                case 9223372036854775807:
                    if (0.1 != (1e10 + 1e10)) {
                        v4 += v4
                    } else {
                        v3 += (v3 + (v3 | v1))
                    }
                default:
                    for fi28 := 0; fi28 < 1; fi28++ {
                        v1 -= (v3 - 828)
                        _ = len(v2)
                        switch v5 {
                        case "hello":
                            switch "foo bar" {
                            case "":
                                for fi29 := 0; fi29 < 12; fi29++ {
                                    v1 -= -2147483648
                                    _ = fi29
                                }
                            default:
                                v2 = append(v2, 171)
                            }
                        case "a":
                            for fi30 := 0; fi30 < 11; fi30++ {
                                v31 := GenS3{F0: v5, F1: 65535}
                                var v32 int64 = 42
                                if ((v6 | v6) <= (^(v6))) {
                                    switch "\n\t" {
                                    case "a":
                                        var v33 float64 = (3.14159 + float64(0.1))
                                        _ = v33
                                    case "":
                                        var v34 int = 9223372036854775807
                                        _ = v34
                                    }
                                } else {
                                    v32 = (v32 | v32)
                                }
                                _ = fi30
                                _ = v31
                                _ = v32
                            }
                        }
                        _ = fi28
                    }
                }
            }
            var v35 int16 = 42
            _ = v35
        } else {
            fmt.Println(v3, v5)
            var v36 int16 = 1000
            if len(v2) > 0 {
                v2[0] = 2147483647
            }
            _ = v36
        }
        var v37 uint8 = v4
        _ = v37
    } else {
        v38 := GenS4{F0: v5, F1: 255, F2: (v6 + (v6 >> ((937) & 63)))}
        var v39 uint = 227
        _ = v38
        _ = v39
    }
    _ = len(v2)
    fmt.Println(v6, v5, v3)
    v4 = (v4 % (((^(v4))) | 1))
    v40 := GenS4{F0: v5, F1: 230, F2: (v6 >> ((0) & 63))}
    fmt.Println(v4)
    if (-892 > 42) {
        fmt.Println(v1, v4, v1)
    } else {
        v6 = v6
    }
    for fi41 := 0; fi41 < 12; fi41++ {
        for fi42 := 0; fi42 < 20; fi42++ {
            switch v5 {
            case "":
                var v43 int32 = v1
                _ = v43
            }
            _ = fi42
        }
        fmt.Println(v1)
        var v44 int32 = int32(v4)
        _ = fi41
        _ = v44
    }
    var v45 int = 793
    v5 = (("\n\t" + v5) + v5)
    fmt.Println(v3, v6)
    v46 := GenS3{F0: v5, F1: 300}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v45
    _ = v46
    _ = v5
    _ = v6
}
