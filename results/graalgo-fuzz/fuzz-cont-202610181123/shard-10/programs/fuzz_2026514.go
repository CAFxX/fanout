package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int8; F1 uint16 }

func main() {
    var v1 int32 = 1
    var v2 int64 = -129
    var v3 int = 1
    v4 := make(map[int]int)
    var v5 int8 = 101
    v6 := []int{(v3 - v3), (v3 << ((255) & 63))}
    v2 = (v2 % (((v2 >> ((18446744073709551615) & 63))) | 1))
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    v1 = (v1 / ((42) | 1))
    v5 = (^(v5))
    v3 = 2147483647
    v1 = v1
    var v7 uint64 = 0
    switch v3 {
    case 32767:
        for fi8 := 0; fi8 < 3; fi8++ {
            switch (v3 % (((-(v3))) | 1)) {
            case 32767:
                v1 = (v1 * (v1 << ((836) & 63)))
            case -487:
                switch "\n\t" {
                case "a":
                    fmt.Println(v3, v7, v5)
                case "hello":
                    v9 := GenS1{F0: (v7 >> ((2) & 63))}
                    _ = v9
                }
            case -9223372036854775808:
                switch "" {
                case "foo bar":
                    v10 := GenS2{F0: -128, F1: 1}
                    _ = v10
                case "\n\t":
                    var v11 float64 = 3.14159
                    _ = v11
                case "":
                    if true {
                        v7 += (v7 % (((v7 | (v7 >> ((18446744073709551615) & 63)))) | 1))
                        var v12 int = -2147483648
                        v3 = v12
                        _ = v12
                    }
                default:
                    for fi13 := 0; fi13 < 9; fi13++ {
                        switch ("\n\t" + "a") {
                        case "a":
                            switch "" {
                            case "foo bar":
                                v14 := GenS1{F0: v7}
                                switch "" {
                                case "foo bar":
                                    if len(v6) >= 2 {
                                        v6 = v6[1:len(v6)]
                                    }
                                case "\n\t":
                                    if true {
                                        fmt.Println(v7, fi8, fi13)
                                        var v15 uint32 = 227
                                        var v16 int = v3
                                        _ = v15
                                        _ = v16
                                    }
                                default:
                                    var v17 float32 = 3.14159
                                    _ = v17
                                }
                                _ = v14
                            case "\n\t":
                                var v18 string = ("foo bar" + ("foo bar" + "foo bar"))
                                _ = v18
                            }
                        }
                        var v19 uint16 = 2
                        _ = fi13
                        _ = v19
                    }
                }
            }
            v20 := GenS2{F0: (v5 - v5), F1: 0}
            v7 -= v7
            _ = fi8
            _ = v20
        }
    case -2147483648:
        var v21 float64 = float64(0.0)
        _ = v21
    case 1:
        v7 -= (v7 + v7)
    }
    fmt.Println(v5, v1, v2)
    v22 := GenS2{F0: v5, F1: 44}
    v5 -= v5
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
