package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int64; F1 uint; F2 uint }
type GenS3 struct { F0 uint16; F1 uint; F2 uint8 }
type GenS4 struct { F0 uint8; F1 string }

func main() {
    v1 := []int{-483, 973}
    var v2 int = -125
    v3 := []int{(v2 * (v2 / ((v2) | 1))), v2, v2}
    v4 := make(map[int]int)
    v5 := v4[(v2 ^ (^(v2)))]
    switch (v2 + v5) {
    case -1:
        var v6 float32 = -2.25
        _ = v6
    case -9223372036854775808:
        for fi7 := 0; fi7 < 13; fi7++ {
            fmt.Println(v5, v5)
            v8 := GenS1{F0: 42}
            _ = fi7
            _ = v8
        }
    }
    for fi9 := 0; fi9 < 13; fi9++ {
        fmt.Println(fi9)
        _ = fi9
    }
    var v10 uint8 = 0
    var v11 int = v5
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    for fi12 := 0; fi12 < 14; fi12++ {
        for fi13 := 0; fi13 < 4; fi13++ {
            fmt.Println(v2)
            v14 := GenS2{F0: 725, F1: 118, F2: 4294967295}
            fmt.Println(v11, v10)
            v2 = (v2 / ((127) | 1))
            _ = fi13
            _ = v14
        }
        switch "" {
        case "a":
            var v15 string = (("foo bar" + "a") + (("a" + ("" + "hello")) + ""))
            _ = v15
        case "hello":
            v2 += 42
        }
        _ = fi12
    }
    if (!false) {
        v4[(v5 / ((v2) | 1))] = v5
        v16 := GenS3{F0: 2, F1: 18446744073709551615, F2: 28}
        switch -813 {
        case -2147483648:
            for fi17 := 0; fi17 < 20; fi17++ {
                for fi18 := 0; fi18 < 4; fi18++ {
                    switch "hello" {
                    case "\n\t":
                        fmt.Println(v2)
                    case "":
                        var v19 uint32 = 1
                        _ = v19
                    case "foo bar":
                        switch v2 {
                        case 695:
                            v5 = (v2 ^ (v11 * 128))
                        case 32767:
                            v5 += v5
                        default:
                            v20 := GenS2{F0: 42, F1: 4294967295, F2: 255}
                            _ = v20
                        }
                    }
                    v21 := GenS4{F0: v10, F1: "\n\t"}
                    v4[v5] = (v5 * (fi17 >> ((65535) & 63)))
                    _ = fi18
                    _ = v21
                }
                v10 += (v10 / ((2) | 1))
                var v22 bool = (v10 >= (^(v10)))
                _ = fi17
                _ = v22
            }
        case 128:
            var v23 uint32 = 2147483647
            _ = v23
        case -705:
            v10 = (v10 | 128)
        default:
            switch "" {
            case "":
                for fi24 := 0; fi24 < 8; fi24++ {
                    switch v5 {
                    case 127:
                        _ = len(v4)
                    default:
                        var v25 uint8 = 128
                        _ = v25
                    }
                    _ = fi24
                }
            default:
                v5 += (v11 >> ((301) & 63))
            }
        }
        _ = v16
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
