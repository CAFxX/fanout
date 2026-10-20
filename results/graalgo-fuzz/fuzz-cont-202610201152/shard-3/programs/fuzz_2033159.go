package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int8; F1 int32 }
type GenS4 struct { F0 int64; F1 int16 }

func main() {
    var v1 int = 913
    v2 := []int{(v1 >> ((1) & 63)), 9223372036854775807, v1, v1}
    var v3 float64 = ((3.14159 + (0.0 * ((0.0 + float64(3.14159)) + float64(0.1)))) + (-2.25 / 1.5))
    var v4 float32 = (1e10 - (float32(v3) * 1e10))
    var v5 float32 = -2.25
    var v6 bool = (1 >= 0)
    _ = len(v2)
    v7 := GenS1{F0: -1, F1: 1}
    if (v3 == v3) {
        fmt.Println(v3, v5, v1)
        fmt.Println(v1)
    } else {
        v8 := GenS1{F0: 42, F1: -174}
        v9 := GenS2{F0: 215}
        if v6 {
            switch "foo bar" {
            case "a":
                var v10 uint32 = 42
                _ = v10
            case "\n\t":
                v3 += v3
            case "hello":
                if (((float64((v5 / -2.25)) * float64(1e10)) / 1.5) >= ((0.1 * v3) + (v3 - 0.1))) {
                    var v11 float32 = v4
                    v12 := GenS2{F0: 1}
                    _ = v11
                    _ = v12
                }
            }
            var v13 int16 = -32768
            _ = v13
        }
        _ = v8
        _ = v9
    }
    var v14 bool = v6
    for fi15 := 0; fi15 < 10; fi15++ {
        if v6 {
            switch 32768 {
            case -1:
                v16 := GenS3{F0: 127, F1: -1}
                _ = v16
            case 832:
                fmt.Println(v3, v4, v5)
            default:
                v1 += -2147483648
            }
        }
        _ = fi15
    }
    v4 = 3.14159
    if (v14 && v14) {
        fmt.Println(v1, v6, v6)
        for fi17 := 0; fi17 < 2; fi17++ {
            if (18446744073709551615 < 42) {
                for fi18 := 0; fi18 < 15; fi18++ {
                    v5 = ((1e10 - (1.5 * 0.0)) / v4)
                    if (!(42 != -875)) {
                        if ((0.0 * v3) == v3) {
                            v1 = 0
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            v14 = (!v14)
                            if v6 {
                                fmt.Println(v14, v1, v3)
                            } else {
                                if (v6 || (1 >= 1)) {
                                    v14 = v6
                                }
                                var v19 string = ("\n\t" + (("a" + "a") + "hello"))
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                _ = v19
                            }
                        } else {
                            v2 = append(v2, fi17)
                            if (!v14) {
                                v2 = append(v2, v1)
                                v20 := GenS4{F0: -1, F1: 32767}
                                _ = v20
                            }
                        }
                        if ((1e10 * (v4 / (v5 + (v5 - 1.5)))) >= (3.14159 + v5)) {
                            fmt.Println(v6)
                            _ = len(v2)
                            if len(v2) > 0 {
                                v2[0] = v1
                            }
                        }
                    }
                    _ = fi18
                }
            } else {
                var v21 int8 = -15
                switch (-(v1)) {
                case 2147483647:
                    var v22 int8 = v21
                    _ = v22
                default:
                    fmt.Println(v4, v6)
                }
                var v23 int16 = 0
                _ = v21
                _ = v23
            }
            v24 := GenS1{F0: 0, F1: -1}
            _ = fi17
            _ = v24
        }
    }
    v6 = (308 <= 18446744073709551615)
    v1 -= v1
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
