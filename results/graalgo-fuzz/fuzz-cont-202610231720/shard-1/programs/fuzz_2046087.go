package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint64; F1 uint32; F2 uint32 }
type GenS4 struct { F0 int16 }

func main() {
    v1 := []uint32{1, 65535, 1}
    var v2 bool = true
    v3 := make(map[uint32]uint32)
    var v4 float32 = 1.5
    v5 := make(map[string]uint32)
    var v6 int = 32767
    v6 += (v6 << ((2) & 63))
    for fi7 := 0; fi7 < 15; fi7++ {
        v3[4294967295] = 0
        _ = fi7
    }
    for fi8 := 0; fi8 < 20; fi8++ {
        for fi9 := 0; fi9 < 13; fi9++ {
            if v2 {
                v10 := GenS1{F0: 1, F1: 0}
                fmt.Println(fi9, fi8)
                fmt.Println(v4)
                _ = v10
            }
            fmt.Println(v4, fi8)
            if v2 {
                v11 := GenS1{F0: 18446744073709551615, F1: 58}
                switch v6 {
                case 4294967295:
                    v6 = 127
                case -894:
                    var v12 uint32 = 656
                    _ = v12
                default:
                    var v13 int32 = 42
                    _ = v13
                }
                if (-2.25 >= 1e10) {
                    v6 -= (fi8 % ((-2147483648) | 1))
                } else {
                    _ = len(v3)
                    v4 = v4
                    if (42 > 18446744073709551615) {
                        switch (fi8 & v6) {
                        case -9223372036854775808:
                            v4 = (v4 - (float32(0.1) + 0.1))
                        case -356:
                            var v14 float32 = ((v4 * (1e10 * (float32(0.1) / (float32(1.5) - (v4 * float32(1.5)))))) - (1e10 - 0.1))
                            _ = v14
                        default:
                            v4 = (v4 + v4)
                        }
                        if v2 {
                            v15 := GenS1{F0: 0, F1: 0}
                            switch "hello" {
                            case "a":
                                fmt.Println(v6, v4)
                            case "\n\t":
                                fmt.Println(v4, v2, v2)
                                var v16 int64 = 9223372036854775807
                                _ = v16
                            }
                            v4 = (v4 + float32(-2.25))
                            _ = v15
                        } else {
                            v17 := GenS1{F0: 216, F1: -957}
                            _ = v17
                        }
                        v18 := GenS1{F0: 0, F1: 0}
                        _ = v18
                    } else {
                        switch fi9 {
                        case -1:
                            switch "hello" {
                            case "":
                                var v19 string = "a"
                                _ = v19
                            }
                        default:
                            fmt.Println(v4)
                        }
                    }
                }
                _ = v11
            }
            _ = fi9
        }
        _ = fi8
    }
    switch (v6 / ((1) | 1)) {
    case 108:
        var v20 uint32 = 969
        _ = v20
    case -1:
        for fi21 := 0; fi21 < 10; fi21++ {
            v22 := GenS2{F0: (!v2)}
            v23 := GenS2{F0: (128 >= 1)}
            v4 -= v4
            _ = fi21
            _ = v22
            _ = v23
        }
    case 0:
        fmt.Println(v2, v4)
    }
    v24 := GenS3{F0: 193, F1: 85, F2: 608}
    v4 += v4
    v25 := GenS4{F0: -32768}
    v2 = false
    v2 = (1 != 204)
    for fi26 := 0; fi26 < 4; fi26++ {
        v2 = (!(1.5 > float64(((float32((1.5 + 0.1)) * (0.1 - (v4 - (1e10 * v4)))) * (float32(0.0) - v4)))))
        var v27 uint16 = 32767
        v1 = append(v1, 1)
        _ = fi26
        _ = v27
    }
    v28 := GenS4{F0: 722}
    var v29 int32 = 1
    v30 := GenS4{F0: 1}
    var v31 float64 = float64(float32(3.14159))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
