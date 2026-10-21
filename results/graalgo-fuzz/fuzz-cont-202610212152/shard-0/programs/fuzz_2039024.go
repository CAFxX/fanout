package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 bool; F1 int32 }
type GenS3 struct { F0 uint; F1 int8 }

func main() {
    var v1 uint32 = 4294967295
    var v2 float32 = 1.5
    v3 := []int{32768, 0, 128, 736}
    var v4 int = -9223372036854775808
    var v5 int8 = 42
    switch "\n\t" {
    case "a":
        fmt.Println(v4, v4)
    default:
        fmt.Println(v1)
    }
    var v6 string = "\n\t"
    var v7 uint32 = (v1 & v1)
    switch v4 {
    case 94:
        var v8 int32 = 0
        _ = v8
    case -2147483648:
        switch v6 {
        case "foo bar":
            if (true && false) {
                var v9 float32 = (((0.0 - ((1e10 / v2) * -2.25)) * -2.25) * ((-2.25 * float32(((0.0 * -2.25) + 1.5))) - v2))
                v5 += (v5 % ((100) | 1))
                var v10 uint8 = 47
                _ = v10
                _ = v9
            } else {
                var v11 uint16 = 1
                v12 := GenS1{F0: v6}
                v11 += v11
                _ = v11
                _ = v12
            }
        default:
            v6 = v6
        }
    }
    v3 = append(v3, 494)
    if true {
        if (float64(v2) >= (0.1 / 3.14159)) {
            v13 := GenS1{F0: ("a" + v6)}
            _ = v13
        } else {
            fmt.Println(v1)
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            switch "" {
            case "foo bar":
                for fi14 := 0; fi14 < 14; fi14++ {
                    v7 = 1
                    if ((42 > 0) || false) {
                        for fi15 := 0; fi15 < 11; fi15++ {
                            if (255 <= 0) {
                                fmt.Println(v5, fi14, v6)
                            } else {
                                var v16 float64 = 3.14159
                                var v17 int16 = -32768
                                var v18 bool = (v6 > (v6 + (v6 + v6)))
                                _ = v16
                                _ = v17
                                _ = v18
                            }
                            switch (fi15 - 545) {
                            case 127:
                                switch fi15 {
                                case 42:
                                    v19 := GenS2{F0: false, F1: -2147483648}
                                    _ = v19
                                case 626:
                                    var v20 float32 = v2
                                    _ = v20
                                case 9223372036854775807:
                                    v7 += 2147483648
                                default:
                                    v1 -= (v7 / ((v1) | 1))
                                }
                            default:
                                var v21 int8 = v5
                                _ = v21
                            }
                            v5 -= v5
                            _ = fi15
                        }
                    } else {
                        switch "\n\t" {
                        case "foo bar":
                            v2 += float32(3.14159)
                        case "\n\t":
                            var v22 int64 = 285
                            _ = v22
                        }
                        if true {
                            v6 = (v6 + "foo bar")
                        } else {
                            v23 := GenS2{F0: false, F1: 960}
                            _ = v23
                        }
                    }
                    _ = fi14
                }
            case "a":
                var v24 uint64 = 18446744073709551615
                _ = v24
            default:
                v1 = (v7 << ((531) & 63))
            }
        }
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    switch 1 {
    case 785:
        v5 = (-(v5))
    default:
        v25 := GenS3{F0: 392, F1: v5}
        _ = v25
    }
    var v26 uint8 = 76
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
