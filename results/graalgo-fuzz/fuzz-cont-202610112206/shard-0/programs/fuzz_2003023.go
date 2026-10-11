package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint32; F1 int8 }
type GenS3 struct { F0 uint8; F1 uint64; F2 uint32 }
type GenS4 struct { F0 bool; F1 uint32; F2 int }

func main() {
    var v1 float64 = 1e10
    var v2 uint16 = 0
    v3 := make(map[uint32]uint32)
    var v4 int = -1
    fmt.Println(v2, v4)
    var v5 uint32 = 318
    v6 := []int{v4}
    v7 := GenS1{F0: 414}
    delete(v3, v5)
    for fi8 := 0; fi8 < 5; fi8++ {
        var v9 uint8 = 255
        var v10 uint = 65535
        switch 128 {
        case -2147483648:
            v3[(v5 * v5)] = 389
        case 658:
            if (1 > 0) {
                v1 -= (float64(float32((v1 * ((((v1 - 3.14159) - ((1e10 / v1) / v1)) / v1) * 0.0)))) + ((-2.25 * float64((float32((v1 * -2.25)) * (1.5 - (3.14159 - (3.14159 + -2.25)))))) / (0.0 * v1)))
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
            } else {
                var v11 int8 = -1
                _ = v11
            }
        case -321:
            v3[(v5 | (v5 % ((v5) | 1)))] = (v5 * v5)
        }
        _ = fi8
        _ = v10
        _ = v9
    }
    v12 := GenS2{F0: (^(v5)), F1: 0}
    v5 -= v5
    var v13 int16 = 32767
    v5 -= v5
    var v14 bool = (42 != 610)
    v2 = v2
    var v15 int64 = 42
    for fi16 := 0; fi16 < 15; fi16++ {
        v17 := GenS3{F0: 0, F1: 0, F2: (v5 ^ v5)}
        _ = fi16
        _ = v17
    }
    var v18 uint = 255
    for fi19 := 0; fi19 < 5; fi19++ {
        for fi20 := 0; fi20 < 7; fi20++ {
            if v14 {
                switch -296 {
                case 32768:
                    fmt.Println(v5, v5, v14)
                default:
                    var v21 uint64 = 42
                    _ = v21
                }
                switch -797 {
                case -951:
                    if (uint64(v5) >= 2) {
                        switch (fi20 << ((v18) & 63)) {
                        case 9223372036854775807:
                            v4 -= -2147483648
                        case 4294967295:
                            var v22 uint = (v18 * (v18 / ((v18) | 1)))
                            _ = v22
                        }
                        var v23 int32 = -2147483648
                        if true {
                            switch "" {
                            case "\n\t":
                                v15 = v15
                            default:
                                if v14 {
                                    fmt.Println(v5, v15, v18)
                                    v2 -= (^(v2))
                                }
                            }
                            var v24 uint64 = 42
                            fmt.Println(v18, v13)
                            _ = v24
                        } else {
                            v2 = (v2 % ((uint16(v5)) | 1))
                            var v25 int32 = v23
                            switch "a" {
                            case "foo bar":
                                var v26 uint = 42
                                _ = v26
                            default:
                                var v27 string = "foo bar"
                                _ = v27
                            }
                            switch v4 {
                            case 42:
                                var v28 uint64 = 42
                                _ = v28
                            case -2147483648:
                                var v29 int = (fi19 ^ (v4 * v4))
                                _ = v29
                            case 9223372036854775807:
                                v14 = v14
                            default:
                                if ((!(v2 >= v2)) && v14) {
                                    v6 = append(v6, 127)
                                    if (!v14) {
                                        var v30 string = "\n\t"
                                        _ = v30
                                    } else {
                                        var v31 uint = v18
                                        _ = v31
                                    }
                                }
                            }
                            _ = v25
                        }
                        _ = v23
                    } else {
                        if v14 {
                            fmt.Println(v18)
                        }
                    }
                }
            }
            _ = fi20
        }
        v32 := GenS4{F0: false, F1: v5, F2: v4}
        v5 = (v5 * v5)
        _ = fi19
        _ = v32
    }
    if v14 {
        var v33 int = (v4 & int(v5))
        _ = len(v6)
        _ = v33
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
