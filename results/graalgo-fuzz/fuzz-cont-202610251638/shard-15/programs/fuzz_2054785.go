package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 string; F1 int64; F2 int64 }
func genFn0(p0 int64) bool {
    v5 := []int{768, 160, 4294967295, -995}
    var v6 int16 = -1
    fmt.Println(p0, p0)
    _ = p0
    _ = v5
    _ = v6
    return (-242 <= 236)
}

func main() {
    var v1 int64 = -1
    var v2 float64 = 1e10
    v3 := make(map[int]uint32)
    v4 := []string{}
    var v7 uint = 767
    v1 += v1
    v8 := GenS1{F0: (^(v7))}
    for fi9 := 0; fi9 < 16; fi9++ {
        var v10 bool = (!true)
        _ = fi9
        _ = v10
    }
    var v11 int64 = (v1 % (((v1 + v1)) | 1))
    switch 0 {
    case 890:
        v2 += (v2 + v2)
    case 0:
        v7 += (v7 & v7)
    }
    v1 = (^(v1))
    var v12 int32 = -1
    if (1000 != 1) {
        if (!false) {
            v1 -= -9223372036854775808
        }
        if (false || true) {
            if len(v4) > 0 {
                v4[0] = "\n\t"
            }
        }
        switch "a" {
        case "a":
            switch -9223372036854775808 {
            case -767:
                switch 128 {
                case -454:
                    v13 := GenS1{F0: v7}
                    _ = v13
                case 854:
                    var v14 int = 32768
                    _ = v14
                case 997:
                    v12 = (^(v12))
                }
            case -849:
                v1 += (v11 + (v11 / ((9223372036854775807) | 1)))
            }
        case "\n\t":
            switch 2147483647 {
            case -9223372036854775808:
                _ = len(v3)
            case 994:
                v1 += v11
            default:
                for fi15 := 0; fi15 < 17; fi15++ {
                    _ = len(v3)
                    var v16 float64 = (v2 - 0.0)
                    if true {
                        v17 := v3[(fi15 ^ 2147483647)]
                        for fi18 := 0; fi18 < 9; fi18++ {
                            if ((v11 - v1) < v11) {
                                fmt.Println(fi15, v1)
                            }
                            _ = fi18
                        }
                        _ = v17
                    } else {
                        for fi19 := 0; fi19 < 8; fi19++ {
                            var v20 uint = 0
                            var v21 int16 = 1
                            _ = len(v3)
                            v21 = v21
                            fmt.Println(fi19, v20, v12)
                            var v22 uint8 = uint8(v20)
                            _ = fi19
                            _ = v20
                            _ = v21
                            _ = v22
                        }
                    }
                    _ = fi15
                    _ = v16
                }
            }
        }
    }
    fmt.Println(v11)
    v23 := GenS1{F0: 798}
    v24 := GenS1{F0: 42}
    for fi25 := 0; fi25 < 6; fi25++ {
        var v26 uint32 = 0
        _ = fi25
        _ = v26
    }
    var v27 int = 10
    switch "foo bar" {
    case "foo bar":
        v28 := v3[157]
        _ = v28
    case "\n\t":
        v29 := GenS2{F0: "hello", F1: v11, F2: (v11 % ((v11) | 1))}
        _ = v29
    case "hello":
        v30 := GenS2{F0: "hello", F1: -1, F2: v1}
        _ = v30
    default:
        delete(v3, v27)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v23
    _ = v24
    _ = v27
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
