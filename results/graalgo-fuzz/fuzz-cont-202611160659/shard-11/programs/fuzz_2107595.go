package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint }
type GenS2 struct { F0 int32 }

func main() {
    var v1 uint16 = 386
    v2 := make(map[int]string)
    var v3 bool = true
    if (93 == 127) {
        var v4 int = 531
        _ = v4
    }
    v5 := []int{0, -2147483648, 4294967295}
    v5 = append(v5, -31)
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v6 := GenS1{F0: 1, F1: 42}
    var v7 uint = 4294967295
    for fi8 := 0; fi8 < 7; fi8++ {
        fmt.Println(v1, fi8)
        var v9 string = "hello"
        _ = fi8
        _ = v9
    }
    var v10 int = -246
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    if len(v5) > 0 {
        v5[0] = (-(v10))
    }
    var v11 string = "hello"
    switch v10 {
    case 37:
        var v12 uint16 = (v1 ^ 42)
        _ = v12
    case -1:
        if (!(true || false)) {
            var v13 int8 = 1
            fmt.Println(v10, v11)
            _ = len(v2)
            _ = v13
        } else {
            fmt.Println(v3, v10, v11)
        }
    case 374:
        if (0 < 100) {
            fmt.Println(v7, v3, v1)
        }
    }
    for fi14 := 0; fi14 < 13; fi14++ {
        if (-1 < -363) {
            v3 = v3
            v2[127] = "hello"
            if (-1 >= -1) {
                for fi15 := 0; fi15 < 16; fi15++ {
                    switch v10 {
                    case 2147483647:
                        v16 := GenS1{F0: (v1 | v1), F1: (v7 & 0)}
                        _ = v16
                    case -2147483648:
                        var v17 int8 = 1
                        _ = v17
                    case 1:
                        switch (^(v10)) {
                        case 32767:
                            var v18 string = (v11 + "")
                            _ = v18
                        case 2147483647:
                            v19 := v2[v10]
                            var v20 int8 = 0
                            v1 = (v1 % (((v1 % ((v1) | 1))) | 1))
                            var v21 int64 = 9223372036854775807
                            var v22 int32 = 798
                            _ = v19
                            _ = v20
                            _ = v21
                            _ = v22
                        }
                    }
                    v11 = "a"
                    _ = fi15
                }
                switch "\n\t" {
                case "hello":
                    for fi23 := 0; fi23 < 11; fi23++ {
                        v24 := GenS2{F0: 42}
                        v7 -= uint(v1)
                        switch "" {
                        case "\n\t":
                            var v25 uint64 = 2
                            _ = v25
                        case "foo bar":
                            if v3 {
                                v1 = 255
                            }
                        case "hello":
                            if (!false) {
                                v3 = v3
                                v26 := GenS1{F0: 65535, F1: (^(v7))}
                                _ = v26
                            } else {
                                var v27 float64 = 0.0
                                v2[fi14] = ("\n\t" + ("\n\t" + (((("\n\t" + v11) + v11) + "a") + "hello")))
                                _ = v27
                            }
                        default:
                            v28 := GenS2{F0: 42}
                            _ = v28
                        }
                        _ = fi23
                        _ = v24
                    }
                case "foo bar":
                    v1 -= 550
                }
            }
        }
        _ = fi14
    }
    v1 -= v1
    fmt.Println(v3, v7)
    v11 = "\n\t"
    _ = len(v2)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v5
    _ = v6
    _ = v7
}
