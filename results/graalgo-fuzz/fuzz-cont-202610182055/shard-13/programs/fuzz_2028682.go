package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8; F2 uint64 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint64 }

func main() {
    var v1 int = 32768
    var v2 bool = (0 >= 0)
    var v3 float32 = (1e10 * (3.14159 / 1.0))
    v4 := []int{-393, -1}
    v2 = (!(v2 && v2))
    for fi5 := 0; fi5 < 6; fi5++ {
        var v6 int32 = -807
        var v7 uint = 0
        v4 = append(v4, (^(fi5)))
        _ = fi5
        _ = v6
        _ = v7
    }
    var v8 float32 = 0.1
    fmt.Println(v1, v2, v8)
    if v2 {
        if v2 {
            v4 = append(v4, v1)
        } else {
            v3 -= v3
            v9 := GenS1{F0: 35, F1: 100, F2: 42}
            _ = v9
        }
        if v2 {
            v3 = (v3 - 0.1)
        }
        if ((0.1 - v3) <= v8) {
            var v10 int64 = 323
            switch -244 {
            case -739:
                if (!(true || v2)) {
                    _ = len(v4)
                    for fi11 := 0; fi11 < 19; fi11++ {
                        for fi12 := 0; fi12 < 1; fi12++ {
                            v4 = append(v4, (v1 - (fi11 + (fi11 << ((0) & 63)))))
                            switch fi11 {
                            case 9223372036854775807:
                                var v13 int8 = -103
                                _ = v13
                            default:
                                v4 = append(v4, 32768)
                            }
                            v2 = (-1 < 17)
                            _ = fi12
                        }
                        switch "hello" {
                        case "\n\t":
                            for fi14 := 0; fi14 < 18; fi14++ {
                                v4 = append(v4, 32767)
                                switch "\n\t" {
                                case "foo bar":
                                    switch fi14 {
                                    case 32767:
                                        if (v2 || (!((1 < -128) && (!v2)))) {
                                            v15 := GenS2{F0: 314}
                                            switch fi11 {
                                            case -1:
                                                var v16 float32 = (v8 / ((v3 + v3) + 1e10))
                                                _ = v16
                                            default:
                                                v4 = append(v4, v1)
                                            }
                                            _ = v15
                                        }
                                    case -9223372036854775808:
                                        var v17 string = "hello"
                                        _ = v17
                                    default:
                                        switch (v1 | (fi14 ^ (^(fi14)))) {
                                        case 32768:
                                            switch -311 {
                                            case 128:
                                                var v18 int = -2147483648
                                                _ = v18
                                            default:
                                                if len(v4) > 0 {
                                                    v4[0] = v1
                                                }
                                            }
                                        default:
                                            var v19 int64 = v10
                                            _ = v19
                                        }
                                    }
                                case "":
                                    v8 = 3.14159
                                case "\n\t":
                                    fmt.Println(fi14)
                                default:
                                    v20 := GenS1{F0: 4294967295, F1: -1, F2: 42}
                                    _ = v20
                                }
                                v21 := GenS2{F0: 1}
                                _ = fi14
                                _ = v21
                            }
                        case "":
                            for fi22 := 0; fi22 < 10; fi22++ {
                                switch 314 {
                                case 128:
                                    v23 := GenS3{F0: 4294967295}
                                    _ = v23
                                }
                                v2 = (v2 && false)
                                _ = fi22
                            }
                        default:
                            fmt.Println(v1)
                        }
                        _ = fi11
                    }
                    if (!v2) {
                        v3 = (1e10 - 0.1)
                        v4 = append(v4, 0)
                    } else {
                        v24 := GenS3{F0: 4294967295}
                        var v25 uint16 = 2
                        _ = v24
                        _ = v25
                    }
                }
            case 2147483647:
                var v26 uint16 = 65535
                _ = v26
            }
            switch "" {
            case "foo bar":
                fmt.Println(v8, v8)
            case "a":
                var v27 uint = 731
                _ = v27
            case "\n\t":
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            }
            _ = v10
        }
    } else {
        v8 += v3
        var v28 int64 = 1
        if (48 == 255) {
            v29 := GenS1{F0: 255, F1: -1, F2: 286}
            _ = v29
        }
        _ = v28
    }
    v4 = append(v4, v1)
    var v30 float32 = (float32(float64(1.5)) * v3)
    if v2 {
        switch (^(v1)) {
        case 128:
            var v31 float32 = (0.0 / (((-2.25 + 0.0) * 1.5) / ((((v8 - v3) + v3) * 0.0) / v8)))
            _ = v31
        case 1:
            var v32 uint64 = 118
            _ = v32
        default:
            switch 538 {
            case 9223372036854775807:
                v2 = (false || false)
            case -1:
                v33 := GenS2{F0: 255}
                _ = v33
            case 587:
                fmt.Println(v8, v1)
            default:
                switch (v1 << ((2) & 63)) {
                case -9223372036854775808:
                    var v34 uint32 = 186
                    _ = v34
                }
            }
        }
        v4 = append(v4, 4294967295)
    }
    if (42 >= 42) {
        var v35 int64 = 42
        _ = v35
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v8
}
