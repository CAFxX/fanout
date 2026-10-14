package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool; F2 uint16 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint32 }

func main() {
    var v1 int64 = 9223372036854775807
    var v2 bool = (622 > 65535)
    var v3 uint64 = 4294967295
    var v4 float32 = 1.5
    v4 -= float32((0.1 - 3.14159))
    fmt.Println(v3, v3, v3)
    v1 -= v1
    var v5 int8 = 100
    fmt.Println(v4)
    v6 := []int{9223372036854775807, 464}
    if v2 {
        fmt.Println(v2, v4)
        var v7 uint32 = 143
        _ = len(v6)
        _ = v7
    } else {
        if (("foo bar" + "foo bar") >= "hello") {
            if (!(((v4 <= (v4 - float32((3.14159 + 3.14159)))) && v2) || true)) {
                switch "\n\t" {
                case "a":
                    v1 = 524
                default:
                    var v8 int16 = 0
                    _ = v8
                }
                for fi9 := 0; fi9 < 7; fi9++ {
                    fmt.Println(v4)
                    var v10 float32 = v4
                    _ = fi9
                    _ = v10
                }
                var v11 uint16 = 667
                _ = v11
            }
            v1 = 42
        }
    }
    v12 := GenS1{F0: 65535, F1: (1.5 >= 1e10), F2: 32767}
    if ((v1 << ((0) & 63)) >= (v1 + v1)) {
        for fi13 := 0; fi13 < 3; fi13++ {
            switch (fi13 << ((18446744073709551615) & 63)) {
            case -267:
                fmt.Println(fi13, v3)
            }
            _ = fi13
        }
    }
    var v14 uint = 4294967295
    if (!v2) {
        fmt.Println(v14, v14)
        var v15 uint32 = 1
        _ = v15
    } else {
        if (263 > 2147483647) {
            var v16 int8 = (v5 * 1)
            fmt.Println(v5, v3)
            _ = v16
        }
        if (!v2) {
            switch "hello" {
            case "a":
                var v17 float64 = (1e10 - 3.14159)
                _ = v17
            default:
                switch -119 {
                case -133:
                    v5 -= (v5 << (((^(v14))) & 63))
                case 4294967295:
                    if ((v3 >> (((^(v14))) & 63)) > (^(v3))) {
                        for fi18 := 0; fi18 < 13; fi18++ {
                            v19 := GenS1{F0: v14, F1: (uint8(v14) < 255), F2: 1}
                            var v20 float64 = 0.0
                            _ = fi18
                            _ = v19
                            _ = v20
                        }
                    } else {
                        switch -446 {
                        case -2147483648:
                            v21 := GenS2{F0: 32767}
                            _ = v21
                        default:
                            _ = len(v6)
                        }
                        for fi22 := 0; fi22 < 7; fi22++ {
                            v23 := GenS1{F0: v14, F1: (65535 == 0), F2: 251}
                            var v24 string = "a"
                            v25 := GenS2{F0: -1}
                            _ = fi22
                            _ = v23
                            _ = v24
                            _ = v25
                        }
                    }
                case -530:
                    v2 = v2
                default:
                    if v2 {
                        if (v14 >= v14) {
                            var v26 int32 = -2147483648
                            if (v5 > (v5 + -1)) {
                                fmt.Println(v1)
                                fmt.Println(v3)
                                v5 = (v5 - v5)
                            }
                            var v27 uint = 0
                            _ = v26
                            _ = v27
                        } else {
                            if v2 {
                                var v28 int = 9223372036854775807
                                for fi29 := 0; fi29 < 13; fi29++ {
                                    var v30 int64 = (v1 % (((v1 >> (((v14 | v14)) & 63))) | 1))
                                    if len(v6) >= 2 {
                                        v6 = v6[1:len(v6)]
                                    }
                                    _ = fi29
                                    _ = v30
                                }
                                v5 += (v5 * (v5 % (((v5 * v5)) | 1)))
                                _ = v28
                            } else {
                                var v31 bool = (2 == 124)
                                switch "\n\t" {
                                case "hello":
                                    v5 += (v5 - (v5 >> ((v14) & 63)))
                                }
                                for fi32 := 0; fi32 < 15; fi32++ {
                                    v3 += v3
                                    for fi33 := 0; fi33 < 1; fi33++ {
                                        if false {
                                            if false {
                                                v1 = -1
                                                if (!v31) {
                                                    v31 = (0 >= (^(fi33)))
                                                }
                                            } else {
                                                v31 = v31
                                            }
                                        }
                                        switch (-(fi32)) {
                                        case -467:
                                            if v31 {
                                                fmt.Println(v2)
                                                var v34 uint64 = (v3 + v3)
                                                _ = v34
                                            }
                                        case -9223372036854775808:
                                            v35 := GenS3{F0: 843}
                                            _ = v35
                                        default:
                                            v4 = 0.1
                                        }
                                        fmt.Println(fi33)
                                        v36 := GenS3{F0: 42}
                                        if (v3 != v3) {
                                            var v37 bool = v2
                                            _ = v37
                                        }
                                        _ = fi33
                                        _ = v36
                                    }
                                    _ = fi32
                                }
                                _ = v31
                            }
                            var v38 uint32 = 147
                            if len(v6) > 0 {
                                v6[0] = 128
                            }
                            _ = v38
                        }
                    }
                }
            }
        } else {
            fmt.Println(v2, v14)
        }
    }
    if true {
        switch -1 {
        case 535:
            var v39 int16 = -1
            _ = v39
        case 9223372036854775807:
            v3 += 719
        case 42:
            _ = len(v6)
        default:
            var v40 float64 = (float64(v4) - 0.1)
            _ = v40
        }
        if len(v6) > 0 {
            v6[0] = 42
        }
    } else {
        for fi41 := 0; fi41 < 14; fi41++ {
            if (!(0 < 1000)) {
                v6 = append(v6, (fi41 + fi41))
            }
            for fi42 := 0; fi42 < 11; fi42++ {
                fmt.Println(fi41)
                _ = fi42
            }
            _ = fi41
        }
        switch 998 {
        case -247:
            fmt.Println(v14, v4, v1)
        default:
            if v2 {
                v3 = 2
                v2 = v2
            }
        }
        var v43 float64 = (-2.25 - float64(0.0))
        _ = v43
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v12
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
