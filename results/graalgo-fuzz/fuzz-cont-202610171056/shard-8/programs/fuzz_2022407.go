package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32 }
type GenS2 struct { F0 int32; F1 uint16 }
type GenS3 struct { F0 uint32; F1 int }

func main() {
    v1 := []uint32{}
    var v2 uint32 = 494
    var v3 string = "a"
    v4 := make(map[string]int)
    for fi5 := 0; fi5 < 11; fi5++ {
        var v6 uint = 42
        _ = len(v4)
        _ = fi5
        _ = v6
    }
    if (1 != 4294967295) {
        var v7 int16 = 32767
        _ = v7
    }
    if ((0 > 559) && false) {
        var v8 uint64 = 18446744073709551615
        v8 -= v8
        _ = v8
    } else {
        v9 := GenS1{F0: 98, F1: (v2 / ((v2) | 1))}
        if (-280 > 9223372036854775807) {
            if false {
                var v10 uint = 2
                _ = v10
            } else {
                var v11 uint16 = 255
                if ((v2 >> ((255) & 63)) < 597) {
                    v1 = append(v1, (v2 / (((v2 % ((v2) | 1))) | 1)))
                    fmt.Println(v3, v11)
                }
                for fi12 := 0; fi12 < 1; fi12++ {
                    if false {
                        v11 += v11
                        var v13 float64 = 0.0
                        if (1 != 42) {
                            var v14 float64 = 3.14159
                            fmt.Println(v2, fi12)
                            _ = v14
                        } else {
                            switch v3 {
                            case "a":
                                v1 = append(v1, 65535)
                            default:
                                for fi15 := 0; fi15 < 14; fi15++ {
                                    switch fi15 {
                                    case 127:
                                        _ = len(v4)
                                    case 2147483647:
                                        var v16 int32 = -2147483648
                                        _ = v16
                                    }
                                    v3 = v3
                                    _ = fi15
                                }
                            }
                            var v17 uint = 255
                            _ = v17
                        }
                        _ = v13
                    } else {
                        v4[v3] = -1
                        _ = len(v4)
                    }
                    _ = fi12
                }
                _ = v11
            }
        } else {
            if true {
                var v18 int32 = 0
                v19 := GenS1{F0: -95, F1: (v2 - 4294967295)}
                _ = len(v1)
                _ = v18
                _ = v19
            }
            v20 := GenS1{F0: 0, F1: 0}
            _ = v20
        }
        v3 = v3
        _ = v9
    }
    for fi21 := 0; fi21 < 2; fi21++ {
        switch v3 {
        case "hello":
            if true {
                var v22 int = fi21
                var v23 int = (fi21 | 1)
                _ = v22
                _ = v23
            }
        default:
            for fi24 := 0; fi24 < 9; fi24++ {
                var v25 float32 = 1.5
                v25 = -2.25
                _ = fi24
                _ = v25
            }
        }
        v2 = (v2 | v2)
        if true {
            v26 := GenS1{F0: -1, F1: v2}
            if false {
                delete(v4, ("foo bar" + ""))
                switch v3 {
                case "":
                    v3 = v3
                case "foo bar":
                    fmt.Println(v3)
                }
                if false {
                    var v27 uint32 = 4294967295
                    _ = v27
                }
            }
            _ = v26
        } else {
            v2 += v2
            v3 = "\n\t"
            v28 := GenS2{F0: 798, F1: 255}
            _ = v28
        }
        _ = fi21
    }
    var v29 bool = (!((v2 | (v2 >> ((1) & 63))) == 65535))
    fmt.Println(v3)
    for fi30 := 0; fi30 < 11; fi30++ {
        for fi31 := 0; fi31 < 19; fi31++ {
            v29 = (32767 > 42)
            _ = fi31
        }
        var v32 int64 = 813
        _ = fi30
        _ = v32
    }
    for fi33 := 0; fi33 < 3; fi33++ {
        _ = len(v4)
        switch -1 {
        case -39:
            if ("foo bar" != v3) {
                fmt.Println(v3)
                for fi34 := 0; fi34 < 7; fi34++ {
                    v3 = v3
                    if (((((1.5 - 1.5) - 0.0) > 1.5) && (!false)) && (-426 >= 600)) {
                        var v35 bool = v29
                        v3 = (v3 + v3)
                        switch 2147483647 {
                        case -9223372036854775808:
                            v29 = true
                        case 681:
                            v2 = v2
                        case 0:
                            if v35 {
                                v29 = (!((v2 % (((v2 << ((1) & 63))) | 1)) <= v2))
                            } else {
                                _ = len(v1)
                                var v36 float32 = 0.0
                                if (!v35) {
                                    if (!(!(v29 || v29))) {
                                        v37 := GenS3{F0: v2, F1: fi34}
                                        _ = v37
                                    }
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    if ((42 < 65535) && (v35 || true)) {
                                        _ = len(v1)
                                        v3 = "hello"
                                    }
                                }
                                _ = v36
                            }
                        default:
                            v38 := v4[(("hello" + (("foo bar" + "") + v3)) + v3)]
                            _ = v38
                        }
                        _ = v35
                    } else {
                        v29 = ((3.14159 * 0.1) >= (0.1 + -2.25))
                        var v39 int16 = 445
                        _ = v39
                    }
                    v29 = false
                    _ = fi34
                }
                switch "hello" {
                case "":
                    var v40 float64 = 3.14159
                    _ = v40
                case "a":
                    var v41 int64 = 0
                    _ = v41
                case "hello":
                    v29 = (false || true)
                default:
                    for fi42 := 0; fi42 < 3; fi42++ {
                        var v43 int8 = -109
                        v3 = v3
                        _ = len(v4)
                        _ = fi42
                        _ = v43
                    }
                }
            }
        case 4294967295:
            v2 += 4294967295
        case -9223372036854775808:
            v29 = v29
        }
        v4[v3] = fi33
        _ = fi33
    }
    v29 = (42 > -9223372036854775808)
    var v44 int8 = 42
    for fi45 := 0; fi45 < 11; fi45++ {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        for fi46 := 0; fi46 < 16; fi46++ {
            v47 := GenS1{F0: (v44 % (((v44 - int8(fi45))) | 1)), F1: (v2 & 914)}
            var v48 int64 = -835
            _ = fi46
            _ = v47
            _ = v48
        }
        _ = fi45
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v44:", v44)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v44
}
