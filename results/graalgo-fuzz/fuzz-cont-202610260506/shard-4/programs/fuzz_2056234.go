package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int32; F1 uint }

func main() {
    var v1 uint8 = 1
    var v2 int = 4294967295
    var v3 float64 = (1.5 + 1e10)
    var v4 uint64 = 4294967295
    var v5 string = ("foo bar" + "hello")
    if false {
        var v6 uint64 = 396
        if ((-603 == 1) || (!(18446744073709551615 < 4294967295))) {
            for fi7 := 0; fi7 < 20; fi7++ {
                for fi8 := 0; fi8 < 4; fi8++ {
                    switch v5 {
                    case "hello":
                        v9 := []int{v2, v2}
                        _ = v9
                    case "":
                        fmt.Println(v1)
                    case "\n\t":
                        switch fi7 {
                        case 643:
                            v10 := GenS1{F0: v6, F1: 32767}
                            _ = v10
                        }
                    }
                    if (1 < -326) {
                        if (v6 == (v4 | v6)) {
                            var v11 float64 = v3
                            var v12 float64 = (1.5 + v11)
                            v13 := []int{fi7, v2, (v2 % ((fi8) | 1)), fi7}
                            _ = v11
                            _ = v12
                            _ = v13
                        } else {
                            v1 -= (^(v1))
                            var v14 int = (fi7 % ((562) | 1))
                            switch (-(fi7)) {
                            case -283:
                                var v15 bool = ("a" < ("foo bar" + v5))
                                _ = v15
                            case -413:
                                var v16 uint32 = 42
                                if false {
                                    var v17 uint = 2
                                    v18 := GenS1{F0: (v6 - 0), F1: 920}
                                    _ = v17
                                    _ = v18
                                }
                                var v19 uint64 = (v6 >> ((42) & 63))
                                v20 := []int{}
                                _ = v16
                                _ = v19
                                _ = v20
                            case 1:
                                v3 += v3
                            default:
                                v21 := GenS2{F0: -424}
                                _ = v21
                            }
                            _ = v14
                        }
                        if false {
                            if (!true) {
                                fmt.Println(v2)
                                fmt.Println(v1, v6, v4)
                            }
                            if (481 == -209) {
                                var v22 uint64 = 70
                                _ = v22
                            }
                        } else {
                            fmt.Println(v1, v5)
                            var v23 int64 = 1
                            _ = v23
                        }
                    } else {
                        if false {
                            v3 = ((v3 + v3) + (((0.1 + v3) - (v3 / 1.0)) - v3))
                            switch (fi7 * (v2 * fi8)) {
                            case 1:
                                v4 = 4294967295
                            case 42:
                                var v24 int8 = -128
                                _ = v24
                            default:
                                var v25 uint = 2
                                _ = v25
                            }
                        } else {
                            v26 := []int{fi7, (v2 / (((-(fi8))) | 1)), 4294967295}
                            switch v2 {
                            case -1:
                                switch ("" + v5) {
                                case "":
                                    switch (fi7 >> ((1) & 63)) {
                                    case 9223372036854775807:
                                        v3 = 0.1
                                    }
                                case "a":
                                    var v27 string = v5
                                    _ = v27
                                }
                                var v28 float64 = v3
                                v2 = 9223372036854775807
                                _ = v28
                            case -2147483648:
                                fmt.Println(v4, v3)
                            default:
                                fmt.Println(v1, v1, v1)
                            }
                            _ = v26
                        }
                        fmt.Println(v3, v1, v3)
                    }
                    _ = fi8
                }
                v29 := []int{fi7, v2}
                if ((65535 == 336) && (42 > 0)) {
                    var v30 int = (fi7 ^ (fi7 & (-(fi7))))
                    v6 = v6
                    switch (v5 + v5) {
                    case "\n\t":
                        if true {
                            v2 = (v30 + 162)
                            v2 = v2
                        }
                    }
                    _ = v30
                } else {
                    for fi31 := 0; fi31 < 20; fi31++ {
                        v32 := GenS1{F0: 0, F1: 1000}
                        v3 = v3
                        var v33 int16 = 280
                        var v34 uint64 = 1
                        if (65535 == 2147483647) {
                            var v35 int64 = 42
                            _ = v35
                        } else {
                            fmt.Println(fi7, v1, v6)
                            v36 := GenS2{F0: 0}
                            var v37 int = (v2 * 1)
                            _ = v36
                            _ = v37
                        }
                        _ = fi31
                        _ = v32
                        _ = v33
                        _ = v34
                    }
                    if len(v29) >= 2 {
                        v29 = v29[1:len(v29)]
                    }
                }
                _ = fi7
                _ = v29
            }
        } else {
            v38 := GenS1{F0: (v6 * v6), F1: int16(v6)}
            fmt.Println(v3)
            _ = v38
        }
        v39 := GenS1{F0: (v6 >> ((65535) & 63)), F1: 792}
        _ = v39
        _ = v6
    }
    v1 = (^(v1))
    switch "foo bar" {
    case "hello":
        v40 := GenS3{F0: 1, F1: 65535}
        _ = v40
    case "\n\t":
        var v41 uint = 2
        _ = v41
    case "foo bar":
        var v42 int = -2147483648
        _ = v42
    default:
        for fi43 := 0; fi43 < 17; fi43++ {
            var v44 float64 = (v3 * (v3 + v3))
            _ = fi43
            _ = v44
        }
    }
    v45 := []int{v2, (v2 % ((127) | 1)), v2, v2}
    fmt.Println(v5, v4)
    v4 += v4
    _ = len(v45)
    fmt.Println(v4, v3)
    v2 -= 4294967295
    v3 = float64(float32(float64(0.0)))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v45:", len(v45), v45)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v5
}
