package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64; F2 bool }
type GenS2 struct { F0 int; F1 int16; F2 int64 }

func main() {
    var v1 int = -1
    v2 := []int{(v1 * 32768), (v1 + v1)}
    var v3 int64 = 0
    v4 := make(map[uint32]bool)
    var v5 bool = (0 > -128)
    v6 := []int{(^(v1)), v1, v1}
    for fi7 := 0; fi7 < 1; fi7++ {
        var v8 int8 = 1
        v9 := GenS1{F0: 0, F1: (v3 * (v3 + (v3 & 42))), F2: v5}
        if false {
            switch "foo bar" {
            case "hello":
                switch (v1 + 0) {
                case -452:
                    for fi10 := 0; fi10 < 1; fi10++ {
                        var v11 bool = v5
                        delete(v4, 65535)
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        _ = fi10
                        _ = v11
                    }
                default:
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                }
            default:
                switch ("\n\t" + "") {
                case "a":
                    var v12 int8 = 127
                    _ = v12
                }
            }
            for fi13 := 0; fi13 < 6; fi13++ {
                switch "" {
                case "foo bar":
                    v14 := GenS1{F0: v3, F1: -9223372036854775808, F2: (((18446744073709551615 == 0) && (-1 < 1)) || v5)}
                    _ = v14
                case "":
                    v1 = fi7
                case "a":
                    v6 = append(v6, (fi13 % (((fi7 / (((^(fi7))) | 1))) | 1)))
                    var v15 uint8 = 127
                    _ = len(v6)
                    _ = v15
                default:
                    switch "foo bar" {
                    case "hello":
                        fmt.Println(v3)
                    default:
                        fmt.Println(fi13, v8)
                    }
                }
                v16 := GenS1{F0: -1, F1: (^(v3)), F2: false}
                _ = fi13
                _ = v16
            }
        }
        _ = fi7
        _ = v8
        _ = v9
    }
    v17 := GenS1{F0: v3, F1: (v3 * (v3 ^ -522)), F2: v5}
    v18 := GenS2{F0: (v1 - v1), F1: 42, F2: (v3 >> ((4294967295) & 63))}
    v3 += (v3 * v3)
    v1 += v1
    var v19 uint64 = 6
    var v20 int32 = 0
    if ((v1 / ((v1) | 1)) > 9223372036854775807) {
        v21 := GenS2{F0: v1, F1: 0, F2: v3}
        var v22 int64 = int64(v1)
        var v23 int16 = -32768
        _ = v21
        _ = v22
        _ = v23
    } else {
        for fi24 := 0; fi24 < 12; fi24++ {
            var v25 int16 = 217
            fmt.Println(fi24)
            _ = fi24
            _ = v25
        }
    }
    var v26 uint8 = 12
    v19 -= (v19 | (v19 | v19))
    _ = len(v4)
    fmt.Println(v1, v26, v20)
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
