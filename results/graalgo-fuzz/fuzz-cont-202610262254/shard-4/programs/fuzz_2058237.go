package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64; F2 uint64 }

func main() {
    v1 := []string{""}
    var v2 int8 = 1
    var v3 bool = (18446744073709551615 > 2)
    v4 := []int{538, 48, -2147483648, 42}
    var v5 int32 = 1
    for fi6 := 0; fi6 < 10; fi6++ {
        var v7 float32 = 1e10
        _ = fi6
        _ = v7
    }
    for fi8 := 0; fi8 < 9; fi8++ {
        for fi9 := 0; fi9 < 5; fi9++ {
            v2 -= (v2 * v2)
            _ = fi9
        }
        _ = fi8
    }
    for fi10 := 0; fi10 < 7; fi10++ {
        v2 = v2
        v2 -= 116
        for fi11 := 0; fi11 < 20; fi11++ {
            v12 := GenS1{F0: 42, F1: 18446744073709551615, F2: 2}
            var v13 int16 = -32768
            _ = fi11
            _ = v12
            _ = v13
        }
        _ = fi10
    }
    var v14 uint64 = 0
    v15 := GenS1{F0: 2147483648, F1: 1, F2: 18446744073709551615}
    fmt.Println(v3, v2, v3)
    fmt.Println(v14, v5, v14)
    v3 = (true && v3)
    for fi16 := 0; fi16 < 8; fi16++ {
        var v17 int8 = (v2 + (v2 ^ v2))
        _ = fi16
        _ = v17
    }
    var v18 uint8 = 127
    v18 = v18
    if v3 {
        _ = len(v1)
        if v3 {
            var v19 int16 = -989
            v4 = append(v4, 32767)
            _ = v19
        } else {
            v2 = (v2 + 42)
        }
        v18 = v18
    }
    v5 = (v5 * (v5 + -2147483648))
    var v20 bool = (v3 || false)
    var v21 uint = 18446744073709551615
    for fi22 := 0; fi22 < 3; fi22++ {
        for fi23 := 0; fi23 < 19; fi23++ {
            if v20 {
                if (1 != 1) {
                    var v24 int32 = 1
                    if ((fi22 ^ fi22) >= 4294967295) {
                        var v25 string = "\n\t"
                        _ = v25
                    } else {
                        v3 = ((0.1 / (float64(3.14159) + -2.25)) <= float64((0.1 * (0.1 * float32((0.1 - 3.14159))))))
                    }
                    var v26 float64 = 1.5
                    _ = v24
                    _ = v26
                } else {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    v5 += (v5 << (((v21 >> ((v21) & 63))) & 63))
                }
                if v3 {
                    var v27 int32 = v5
                    switch "foo bar" {
                    case "":
                        fmt.Println(v27, v14, v27)
                    case "\n\t":
                        v28 := GenS1{F0: 4294967295, F1: v14, F2: v14}
                        _ = v28
                    case "a":
                        if ("\n\t" == "foo bar") {
                            v29 := GenS1{F0: 1, F1: 419, F2: v14}
                            switch (fi22 * fi23) {
                            case -1:
                                v30 := GenS1{F0: 2, F1: (^(v14)), F2: v14}
                                _ = v30
                            case 42:
                                var v31 string = ("a" + "hello")
                                _ = v31
                            case 127:
                                v32 := GenS1{F0: 65535, F1: (^(v14)), F2: 630}
                                _ = v32
                            default:
                                var v33 int8 = (v2 / (((v2 ^ (v2 >> ((v21) & 63)))) | 1))
                                _ = v33
                            }
                            _ = v29
                        } else {
                            v5 = (v27 ^ (v27 ^ -2147483648))
                        }
                    }
                    _ = v27
                } else {
                    fmt.Println(v20, v2)
                }
            }
            fmt.Println(v3)
            _ = fi23
        }
        _ = len(v1)
        _ = fi22
    }
    switch "foo bar" {
    case "\n\t":
        fmt.Println(v14, v14)
    default:
        _ = len(v1)
    }
    if (v20 || true) {
        v18 = (v18 | (v18 | (v18 >> (((v21 << ((65535) & 63))) & 63))))
        if v20 {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            v3 = v3
        }
    } else {
        v20 = v3
        v14 = 551
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v18
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
