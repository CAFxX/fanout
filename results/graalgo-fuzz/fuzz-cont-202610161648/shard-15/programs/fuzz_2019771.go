package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8 }
type GenS2 struct { F0 int8; F1 uint64 }
type GenS3 struct { F0 uint; F1 uint32; F2 int }
type GenS4 struct { F0 int64; F1 int8; F2 uint64 }

func main() {
    var v1 uint16 = 255
    var v2 uint16 = v1
    var v3 int64 = -225
    var v4 int64 = v3
    switch (("hello" + "") + "foo bar") {
    case "":
        v5 := GenS1{F0: false, F1: 122}
        _ = v5
    default:
        fmt.Println(v1, v3)
    }
    var v6 float64 = (3.14159 - float64(1e10))
    switch 42 {
    case 2147483647:
        v7 := []int{-749, 1, 32768, 127}
        _ = v7
    default:
        switch "hello" {
        case "":
            v3 -= (v4 ^ (v3 + (v4 % ((0) | 1))))
        case "hello":
            v6 += float64(1e10)
        case "foo bar":
            v8 := []int{-1, 42}
            _ = v8
        default:
            if (18446744073709551615 == 391) {
                v2 = v2
            }
        }
    }
    for fi9 := 0; fi9 < 3; fi9++ {
        var v10 uint64 = 0
        var v11 uint = 566
        for fi12 := 0; fi12 < 2; fi12++ {
            var v13 uint8 = 26
            _ = fi12
            _ = v13
        }
        _ = fi9
        _ = v10
        _ = v11
    }
    var v14 bool = ((-2147483648 != 883) && true)
    v15 := []int{0, -9223372036854775808, -1}
    v16 := GenS2{F0: 42, F1: 515}
    v17 := GenS2{F0: -16, F1: 0}
    v15 = append(v15, -651)
    for fi18 := 0; fi18 < 5; fi18++ {
        switch -77 {
        case -1:
            if len(v15) > 0 {
                v15[0] = 9223372036854775807
            }
        case 849:
            var v19 uint16 = v1
            _ = v19
        default:
            if false {
                for fi20 := 0; fi20 < 1; fi20++ {
                    var v21 bool = (!((uint32(v1) <= 4294967295) || true))
                    if (127 >= 127) {
                        v22 := GenS2{F0: 127, F1: 42}
                        _ = v22
                    }
                    fmt.Println(fi20, v3, v6)
                    _ = fi20
                    _ = v21
                }
                v1 = 2
                for fi23 := 0; fi23 < 18; fi23++ {
                    if (347 >= -32768) {
                        switch "" {
                        case "foo bar":
                            v24 := GenS3{F0: 0, F1: 2147483647, F2: (fi18 / ((-2147483648) | 1))}
                            _ = v24
                        case "hello":
                            if false {
                                fmt.Println(v1, v14)
                                var v25 int = fi23
                                _ = v25
                            } else {
                                v1 = 42
                            }
                        }
                        fmt.Println(v6, v2)
                    } else {
                        v4 += v4
                        var v26 int16 = -32768
                        var v27 uint32 = 0
                        _ = v26
                        _ = v27
                    }
                    var v28 int = fi23
                    if true {
                        v2 = 42
                    } else {
                        v4 += 889
                        if ((127 == 100) && v14) {
                            var v29 uint64 = 4294967295
                            _ = v29
                        } else {
                            v30 := GenS2{F0: 42, F1: 18446744073709551615}
                            switch 9223372036854775807 {
                            case 2147483647:
                                switch -806 {
                                case 4294967295:
                                    var v31 bool = (42 <= 515)
                                    _ = v31
                                case 127:
                                    var v32 int64 = (v4 ^ (v4 ^ (^(v4))))
                                    _ = v32
                                }
                            default:
                                v33 := GenS3{F0: 4294967295, F1: 1, F2: (fi23 | v28)}
                                _ = v33
                            }
                            _ = v30
                        }
                        switch 4294967295 {
                        case 723:
                            var v34 int64 = (-(v3))
                            _ = v34
                        }
                        v35 := GenS4{F0: v4, F1: 127, F2: 724}
                        v15 = append(v15, -482)
                        _ = v35
                    }
                    _ = fi23
                    _ = v28
                }
            }
        }
        var v36 int32 = 766
        v2 = 65535
        _ = fi18
        _ = v36
    }
    for fi37 := 0; fi37 < 8; fi37++ {
        var v38 uint8 = 255
        _ = fi37
        _ = v38
    }
    v1 = (v2 - v1)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v6
}
