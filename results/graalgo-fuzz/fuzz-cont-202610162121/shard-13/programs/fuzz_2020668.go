package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64; F2 int32 }
type GenS2 struct { F0 int32; F1 int; F2 int8 }

func main() {
    v1 := make(map[string]int)
    var v2 int64 = -1
    v3 := []string{"a"}
    if (2147483647 > 466) {
        fmt.Println(v2)
    } else {
        fmt.Println(v2)
    }
    for fi4 := 0; fi4 < 9; fi4++ {
        var v5 float64 = 0.1
        for fi6 := 0; fi6 < 14; fi6++ {
            v7 := GenS1{F0: 42, F1: 42, F2: 763}
            _ = fi6
            _ = v7
        }
        switch "foo bar" {
        case "a":
            if true {
                switch 4294967295 {
                case 336:
                    for fi8 := 0; fi8 < 11; fi8++ {
                        fmt.Println(v2, v2)
                        var v9 int32 = 0
                        _ = fi8
                        _ = v9
                    }
                case 1:
                    var v10 float64 = (v5 + 0.1)
                    _ = v10
                case -1:
                    for fi11 := 0; fi11 < 11; fi11++ {
                        fmt.Println(fi4, fi4, fi11)
                        _ = len(v1)
                        _ = fi11
                    }
                default:
                    fmt.Println(v2)
                }
                switch "a" {
                case "hello":
                    v12 := GenS1{F0: -32768, F1: 4294967295, F2: 573}
                    _ = v12
                default:
                    for fi13 := 0; fi13 < 4; fi13++ {
                        v5 += ((((float64(1.5) + v5) - (3.14159 - v5)) + 1.5) + (float64(1.5) + v5))
                        _ = fi13
                    }
                }
                fmt.Println(v2)
            }
        case "":
            v2 -= v2
        case "foo bar":
            fmt.Println(fi4, v5, v5)
        default:
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
        _ = fi4
        _ = v5
    }
    v2 = 1
    var v14 uint8 = 128
    v14 = 2
    if true {
        v3 = append(v3, "")
        for fi15 := 0; fi15 < 14; fi15++ {
            switch 127 {
            case -1:
                v16 := v1["hello"]
                _ = v16
            case 128:
                _ = len(v1)
            }
            _ = fi15
        }
        for fi17 := 0; fi17 < 4; fi17++ {
            if (!true) {
                _ = len(v3)
            } else {
                if true {
                    v14 += v14
                    v18 := GenS2{F0: -2147483648, F1: 9223372036854775807, F2: 100}
                    _ = v18
                } else {
                    v19 := GenS1{F0: -1, F1: 2, F2: 42}
                    _ = v19
                }
            }
            var v20 int = fi17
            fmt.Println(v2, v2)
            _ = fi17
            _ = v20
        }
    }
    v2 = v2
    fmt.Println(v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
}
