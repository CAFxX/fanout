package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64; F2 uint8 }
type GenS2 struct { F0 int16; F1 int8; F2 uint16 }
type GenS3 struct { F0 uint64; F1 int; F2 uint64 }

func main() {
    var v1 float64 = (-2.25 / (-2.25 + 3.14159))
    v2 := []string{"a"}
    v3 := make(map[int]uint32)
    var v4 float32 = 3.14159
    v5 := []string{}
    var v6 int32 = 0
    if (32767 >= 255) {
        if ((38 < 42) || false) {
            v7 := GenS1{F0: 2, F1: 4294967295, F2: 42}
            if false {
                fmt.Println(v6, v4)
                v1 -= 1.5
                v6 += v6
            }
            _ = v7
        } else {
            fmt.Println(v6)
            for fi8 := 0; fi8 < 7; fi8++ {
                v1 -= v1
                switch fi8 {
                case 128:
                    for fi9 := 0; fi9 < 11; fi9++ {
                        switch "\n\t" {
                        case "foo bar":
                            v10 := GenS2{F0: 1, F1: -128, F2: 86}
                            _ = v10
                        default:
                            v11 := GenS1{F0: 354, F1: 0, F2: 42}
                            fmt.Println(v4, fi9, fi9)
                            v4 -= (0.0 + v4)
                            _ = v11
                        }
                        if (v1 == v1) {
                            _ = len(v5)
                        } else {
                            switch "\n\t" {
                            case "\n\t":
                                v2 = append(v2, "")
                            case "":
                                if (0 <= 42) {
                                    _ = len(v3)
                                    var v12 uint8 = 255
                                    _ = v12
                                }
                            }
                        }
                        _ = fi9
                    }
                case 324:
                    delete(v3, fi8)
                case -879:
                    var v13 float32 = 3.14159
                    _ = v13
                default:
                    var v14 int32 = 1
                    _ = v14
                }
                _ = fi8
            }
            var v15 int16 = 651
            _ = v15
        }
        fmt.Println(v1, v6)
    }
    var v16 float64 = 3.14159
    for fi17 := 0; fi17 < 20; fi17++ {
        for fi18 := 0; fi18 < 16; fi18++ {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            if (2147483648 <= 42) {
                fmt.Println(v16, fi17)
                var v19 float32 = v4
                switch ("" + "") {
                case "foo bar":
                    fmt.Println(fi18, v6)
                default:
                    var v20 int = (-(fi18))
                    _ = v20
                }
                _ = v19
            } else {
                switch (fi17 * fi17) {
                case 332:
                    var v21 uint16 = 65535
                    _ = v21
                }
                v22 := GenS2{F0: 32767, F1: -128, F2: 42}
                if (!(true || (18446744073709551615 < 1))) {
                    var v23 float32 = -2.25
                    _ = v23
                } else {
                    fmt.Println(fi18, v16, v1)
                }
                v24 := GenS2{F0: 0, F1: 100, F2: 65535}
                _ = v22
                _ = v24
            }
            fmt.Println(v16)
            v1 -= ((v16 * v1) / 1.0)
            _ = fi18
        }
        _ = fi17
    }
    v6 -= (v6 * v6)
    var v25 float32 = 1e10
    switch "\n\t" {
    case "\n\t":
        var v26 float64 = 0.1
        _ = v26
    }
    switch 1 {
    case 32768:
        v27 := GenS1{F0: 2, F1: 18446744073709551615, F2: 127}
        _ = v27
    case 721:
        fmt.Println(v4)
    case 1:
        fmt.Println(v4)
    default:
        v28 := GenS3{F0: 0, F1: 1, F2: 4294967295}
        _ = v28
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    _ = len(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
