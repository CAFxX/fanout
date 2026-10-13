package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32; F2 uint8 }
type GenS2 struct { F0 string; F1 uint64 }
type GenS3 struct { F0 uint64 }

func main() {
    var v1 string = "foo bar"
    var v2 uint64 = 0
    v3 := make(map[int]uint32)
    var v4 uint = 606
    v5 := make(map[uint32]string)
    v6 := []int{-477, 890, 869, 127}
    for fi7 := 0; fi7 < 7; fi7++ {
        v8 := GenS1{F0: v1, F1: 667, F2: 255}
        for fi9 := 0; fi9 < 6; fi9++ {
            switch ("foo bar" + v1) {
            case "foo bar":
                fmt.Println(v2)
            case "\n\t":
                switch fi7 {
                case 0:
                    var v10 int64 = 1
                    _ = v10
                case 4294967295:
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                    fmt.Println(fi7, fi7, v1)
                    var v11 uint16 = 0
                    _ = v11
                }
            case "":
                if (false || false) {
                    _ = len(v6)
                    if (!false) {
                        if false {
                            if true {
                                v3[(fi9 << (((v4 & (v4 >> ((4294967295) & 63)))) & 63))] = 965
                                v2 = v2
                                delete(v3, (fi7 % (((fi9 << (((v4 ^ (v4 ^ (v4 & (v4 / ((255) | 1)))))) & 63))) | 1)))
                                _ = len(v6)
                                var v12 uint64 = v2
                                _ = v12
                            }
                        }
                        var v13 int = fi9
                        v14 := GenS2{F0: v1, F1: 42}
                        _ = v13
                        _ = v14
                    }
                    v2 = 2
                    fmt.Println(v2, v2, v1)
                }
            }
            _ = fi9
        }
        fmt.Println(v4)
        _ = fi7
        _ = v8
    }
    var v15 int32 = 2147483647
    var v16 float32 = (0.1 - (0.0 / (-2.25 / 0.1)))
    v15 -= (v15 ^ v15)
    if false {
        if false {
            switch v1 {
            case "foo bar":
                var v17 int64 = 9223372036854775807
                _ = v17
            }
            fmt.Println(v4)
            delete(v3, -8)
        }
    }
    var v18 uint8 = 255
    v19 := GenS1{F0: v1, F1: 804, F2: (^(v18))}
    v20 := GenS3{F0: v2}
    switch -1 {
    case -9223372036854775808:
        v4 = v4
    case 989:
        _ = len(v6)
    case 0:
        var v21 string = (v1 + v1)
        _ = v21
    default:
        _ = len(v5)
    }
    var v22 uint64 = v2
    var v23 uint64 = 4294967295
    var v24 bool = (true && (!false))
    if v24 {
        v1 = (v1 + "hello")
    } else {
        v4 -= (v4 << (((v4 * (v4 % (((v4 ^ v4)) | 1)))) & 63))
        v16 = (1.5 + (((0.0 / float32(-2.25)) - (v16 / ((v16 * 1.5) / ((1e10 / 1e10) + float32((1e10 - 3.14159)))))) + 1e10))
        fmt.Println(v18, v22)
    }
    if v24 {
        if true {
            v25 := GenS2{F0: v1, F1: 120}
            var v26 int8 = -77
            _ = v25
            _ = v26
        }
        _ = len(v6)
        for fi27 := 0; fi27 < 18; fi27++ {
            v28 := v3[840]
            _ = len(v6)
            if len(v6) > 0 {
                v6[0] = (fi27 | 9223372036854775807)
            }
            _ = fi27
            _ = v28
        }
    }
    var v29 int8 = 23
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v22
    _ = v23
    _ = v24
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
