package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint; F2 int16 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint; F1 uint64 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int }
type GenS6 struct { F0 string; F1 uint64 }

func main() {
    var v1 float64 = 1e10
    v2 := []int{0, 128}
    var v3 float32 = 1e10
    switch -9223372036854775808 {
    case -872:
        var v4 int = 112
        _ = v4
    }
    var v5 uint32 = 2
    v1 = v1
    for fi6 := 0; fi6 < 20; fi6++ {
        if len(v2) > 0 {
            v2[0] = -1
        }
        v1 -= v1
        v2 = append(v2, -2147483648)
        _ = fi6
    }
    for fi7 := 0; fi7 < 6; fi7++ {
        if len(v2) > 0 {
            v2[0] = 2147483647
        }
        fmt.Println(v5, v1, v5)
        for fi8 := 0; fi8 < 7; fi8++ {
            switch (^(fi7)) {
            case 563:
                v1 -= (0.1 - v1)
            case 2147483647:
                v5 += 2
            default:
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            }
            v9 := GenS1{F0: 96, F1: 2, F2: 1000}
            var v10 int = fi8
            _ = fi8
            _ = v10
            _ = v9
        }
        _ = fi7
    }
    v2 = append(v2, -2147483648)
    v3 -= float32(float64(v3))
    if ((^(v5)) < (^(v5))) {
        if (true || false) {
            v11 := GenS2{F0: 691}
            v12 := GenS3{F0: 2, F1: 42}
            if ((-287 == -1) && true) {
                switch "a" {
                case "":
                    var v13 int64 = 0
                    _ = v13
                }
            } else {
                switch 9223372036854775807 {
                case 9223372036854775807:
                    var v14 string = "hello"
                    _ = v14
                case 0:
                    switch 905 {
                    case -231:
                        v15 := GenS4{F0: -910}
                        _ = v15
                    case 32767:
                        fmt.Println(v5, v1, v1)
                    case 9223372036854775807:
                        var v16 int16 = 42
                        _ = v16
                    }
                case 2147483647:
                    var v17 string = "\n\t"
                    _ = v17
                default:
                    var v18 int8 = 103
                    _ = v18
                }
                switch -2147483648 {
                case -549:
                    var v19 uint32 = v5
                    _ = v19
                case 128:
                    var v20 uint32 = v5
                    _ = v20
                case 32768:
                    switch -200 {
                    case 678:
                        fmt.Println(v3)
                    case 32767:
                        v3 = v3
                    case 32768:
                        switch 9223372036854775807 {
                        case 128:
                            fmt.Println(v5)
                        case -640:
                            var v21 uint16 = 32767
                            _ = v21
                        case 0:
                            switch "a" {
                            case "a":
                                if len(v2) > 0 {
                                    v2[0] = 128
                                }
                            default:
                                v22 := GenS5{F0: -961}
                                _ = v22
                            }
                        default:
                            var v23 uint64 = 651
                            _ = v23
                        }
                    default:
                        if (2147483647 <= -978) {
                            v24 := GenS6{F0: ("" + "hello"), F1: 48}
                            _ = v24
                        } else {
                            v25 := GenS1{F0: 18446744073709551615, F1: 4294967295, F2: 324}
                            _ = v25
                        }
                    }
                }
            }
            _ = v11
            _ = v12
        }
        _ = len(v2)
    }
    if true {
        var v26 int = -2147483648
        _ = v26
    }
    v1 += (float64(3.14159) + float64(v3))
    v2 = append(v2, -2147483648)
    fmt.Println(v1, v1)
    var v27 int = 9223372036854775807
    v27 += v27
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v5
}
