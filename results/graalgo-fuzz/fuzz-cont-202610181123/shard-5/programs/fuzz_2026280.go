package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int32; F1 int8 }
type GenS4 struct { F0 int64 }

func main() {
    var v1 float32 = (3.14159 * 1.5)
    v2 := make(map[string]bool)
    var v3 int = -812
    var v4 bool = false
    var v5 int = (v3 | (v3 ^ (v3 % (((v3 - (v3 | -2147483648))) | 1))))
    for fi6 := 0; fi6 < 8; fi6++ {
        if v4 {
            var v7 int32 = 2147483647
            v8 := []int{(v3 << ((18446744073709551615) & 63)), 128, fi6, v5}
            var v9 int32 = v7
            _ = v7
            _ = v8
            _ = v9
        }
        switch v5 {
        case 807:
            v10 := GenS1{F0: 378, F1: 2}
            _ = v10
        case 748:
            v4 = true
        case 2147483647:
            var v11 int16 = 0
            _ = v11
        default:
            v12 := v2["hello"]
            _ = v12
        }
        _ = fi6
    }
    v5 += v5
    v2["\n\t"] = true
    var v13 bool = true
    v4 = false
    for fi14 := 0; fi14 < 16; fi14++ {
        switch "hello" {
        case "hello":
            v5 -= v5
        case "foo bar":
            var v15 int32 = 539
            _ = v15
        default:
            for fi16 := 0; fi16 < 11; fi16++ {
                if false {
                    if ((-1 < -2147483648) && false) {
                        v17 := []int{0, (fi16 % (((fi14 / (((v5 << ((342) & 63))) | 1))) | 1)), 0, (v3 - (fi14 - -604))}
                        _ = v17
                    } else {
                        var v18 uint8 = 42
                        if (65535 <= 42) {
                            var v19 int64 = -9223372036854775808
                            switch "" {
                            case "foo bar":
                                fmt.Println(v5, v3)
                                v5 += fi16
                                v1 = v1
                            case "a":
                                if (2147483647 < uint32(v18)) {
                                    if (!(2 >= 2147483647)) {
                                        v20 := GenS2{F0: 32767}
                                        _ = v20
                                    }
                                    v3 = (fi16 ^ v3)
                                    v19 -= (v19 << ((49) & 63))
                                }
                                var v21 string = "a"
                                _ = v21
                            case "":
                                var v22 int16 = 0
                                _ = v22
                            default:
                                v23 := GenS1{F0: 1, F1: 2}
                                _ = v23
                            }
                            _ = v19
                        }
                        var v24 uint8 = (v18 * (v18 % ((v18) | 1)))
                        _ = v18
                        _ = v24
                    }
                    switch -9223372036854775808 {
                    case -9223372036854775808:
                        v3 = 1
                    }
                } else {
                    fmt.Println(v1, v1, fi14)
                    switch "\n\t" {
                    case "a":
                        fmt.Println(v4, v4)
                    case "hello":
                        v1 = v1
                    }
                }
                if (42 <= 424) {
                    var v25 int = fi16
                    var v26 int8 = 100
                    _ = len(v2)
                    _ = v25
                    _ = v26
                } else {
                    fmt.Println(fi14, fi14, v5)
                    if false {
                        v3 += v5
                        switch "\n\t" {
                        case "\n\t":
                            var v27 float64 = 0.0
                            _ = v27
                        case "hello":
                            var v28 uint16 = 317
                            _ = v28
                        case "a":
                            if v4 {
                                delete(v2, "\n\t")
                                v29 := GenS3{F0: -2147483648, F1: -39}
                                _ = len(v2)
                                v1 += (1e10 / -2.25)
                                _ = v29
                            } else {
                                delete(v2, ("hello" + ("foo bar" + "hello")))
                                v30 := []int{}
                                _ = v30
                            }
                            v3 = fi16
                        default:
                            if false {
                                fmt.Println(v5, v5, v5)
                                var v31 string = ""
                                var v32 uint16 = 255
                                _ = v31
                                _ = v32
                            } else {
                                v33 := GenS4{F0: 0}
                                _ = v33
                            }
                        }
                        v3 += (v3 + 379)
                    } else {
                        if v13 {
                            var v34 bool = v4
                            _ = len(v2)
                            var v35 float64 = 1e10
                            _ = v34
                            _ = v35
                        }
                        v1 = (0.0 - (v1 - (3.14159 + v1)))
                    }
                }
                _ = fi16
            }
        }
        for fi36 := 0; fi36 < 16; fi36++ {
            var v37 float64 = 1.5
            var v38 float64 = v37
            _ = fi36
            _ = v37
            _ = v38
        }
        if v4 {
            var v39 uint32 = 0
            _ = v39
        }
        _ = fi14
    }
    fmt.Println(v5)
    var v40 float32 = -2.25
    var v41 int8 = 42
    v42 := GenS2{F0: -1}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v5
}
