package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int; F2 int }
type GenS2 struct { F0 uint8; F1 uint32 }

func main() {
    var v1 uint32 = 533
    var v2 int = 127
    v3 := make(map[int]int)
    var v4 uint = 65535
    var v5 float32 = -2.25
    v6 := []uint32{(v1 >> ((v4) & 63)), v1, (v1 % (((v1 + (v1 >> (((v4 / ((v4) | 1))) & 63)))) | 1))}
    if true {
        if false {
            var v7 uint16 = 0
            v1 += v1
            switch v2 {
            case 128:
                for fi8 := 0; fi8 < 8; fi8++ {
                    if (661 >= 2147483647) {
                        switch (^(v2)) {
                        case 42:
                            var v9 uint64 = 0
                            _ = v9
                        case 32768:
                            for fi10 := 0; fi10 < 1; fi10++ {
                                v4 = v4
                                _ = fi10
                            }
                        case 127:
                            switch fi8 {
                            case -9223372036854775808:
                                if false {
                                    v4 += 2
                                    v11 := GenS1{F0: 42, F1: (v2 >> ((uint(v1)) & 63)), F2: 4294967295}
                                    _ = v11
                                } else {
                                    v2 = (fi8 ^ 128)
                                    _ = len(v6)
                                }
                            case 127:
                                v12 := GenS2{F0: 42, F1: (v1 - uint32(v4))}
                                _ = v12
                            case -195:
                                v2 = fi8
                            }
                        default:
                            v13 := GenS2{F0: 42, F1: (v1 ^ (v1 + 2))}
                            _ = v13
                        }
                        _ = len(v3)
                        for fi14 := 0; fi14 < 5; fi14++ {
                            fmt.Println(v7, v4, v1)
                            _ = fi14
                        }
                    } else {
                        for fi15 := 0; fi15 < 11; fi15++ {
                            var v16 uint16 = v7
                            _ = fi15
                            _ = v16
                        }
                        for fi17 := 0; fi17 < 15; fi17++ {
                            var v18 float64 = float64(v5)
                            switch "\n\t" {
                            case "\n\t":
                                v6 = append(v6, v1)
                            case "a":
                                v19 := GenS1{F0: 0, F1: 32768, F2: (fi17 & fi8)}
                                _ = v19
                            }
                            fmt.Println(v4, v1)
                            _ = fi17
                            _ = v18
                        }
                    }
                    _ = fi8
                }
            default:
                v5 = (0.0 / 1.0)
            }
            _ = v7
        } else {
            v5 = v5
            fmt.Println(v1, v2)
        }
        var v20 uint64 = 1
        _ = v20
    } else {
        var v21 uint64 = 42
        _ = v21
    }
    if (-2147483648 >= -837) {
        fmt.Println(v5, v4)
        for fi22 := 0; fi22 < 2; fi22++ {
            v23 := GenS1{F0: 42, F1: v2, F2: fi22}
            v3[fi22] = fi22
            _ = fi22
            _ = v23
        }
    }
    for fi24 := 0; fi24 < 2; fi24++ {
        switch "\n\t" {
        case "\n\t":
            v6 = append(v6, 4294967295)
        case "a":
            var v25 uint = v4
            _ = v25
        }
        var v26 int = 2147483647
        _ = fi24
        _ = v26
    }
    v3[v2] = (v2 ^ -440)
    var v27 uint = (v4 | (v4 >> (((v4 >> ((4294967295) & 63))) & 63)))
    if true {
        v1 = 813
    } else {
        v28 := GenS1{F0: 128, F1: 2147483647, F2: v2}
        _ = v28
    }
    var v29 uint8 = 0
    var v30 int16 = 32767
    var v31 float32 = 1e10
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v27
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
