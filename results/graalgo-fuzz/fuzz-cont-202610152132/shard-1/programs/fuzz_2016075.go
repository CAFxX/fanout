package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int32; F1 bool }

func main() {
    v1 := make(map[int]int)
    var v2 uint = 255
    v3 := []uint32{}
    v4 := []int{128, 127}
    fmt.Println(v2)
    var v5 uint16 = 786
    for fi6 := 0; fi6 < 5; fi6++ {
        for fi7 := 0; fi7 < 5; fi7++ {
            v8 := GenS1{F0: -1}
            if true {
                var v9 float32 = 3.14159
                var v10 string = "\n\t"
                _ = v10
                _ = v9
            } else {
                if false {
                    switch "\n\t" {
                    case "a":
                        if false {
                            v11 := GenS2{F0: 2147483647}
                            _ = v11
                        } else {
                            v12 := GenS3{F0: 0, F1: true}
                            var v13 int16 = 32767
                            var v14 int32 = -2147483648
                            _ = v12
                            _ = v13
                            _ = v14
                        }
                    case "":
                        var v15 uint32 = 1
                        _ = v15
                    }
                } else {
                    fmt.Println(fi7, v5)
                    switch "" {
                    case "a":
                        if false {
                            v5 = 0
                            var v16 uint16 = v5
                            _ = v16
                        } else {
                            switch "a" {
                            case "hello":
                                var v17 string = "foo bar"
                                _ = v17
                            case "\n\t":
                                var v18 int16 = -446
                                _ = v18
                            case "a":
                                fmt.Println(v5, fi7, v2)
                            }
                            switch "foo bar" {
                            case "\n\t":
                                v2 = 18446744073709551615
                            }
                        }
                    case "foo bar":
                        _ = len(v1)
                    default:
                        var v19 float32 = 1.5
                        _ = v19
                    }
                    var v20 uint32 = uint32(v5)
                    _ = v20
                }
                v2 = (v2 + (v2 << (((v2 | v2)) & 63)))
                v5 -= (^(v5))
            }
            _ = fi7
            _ = v8
        }
        v2 -= v2
        _ = fi6
    }
    if ((v5 & (v5 % ((v5) | 1))) == 65535) {
        fmt.Println(v2, v5)
    }
    var v21 uint16 = v5
    if len(v4) > 0 {
        v4[0] = 0
    }
    var v22 uint16 = v5
    var v23 bool = (!false)
    var v24 uint8 = 44
    switch "a" {
    case "a":
        var v25 float64 = float64(3.14159)
        _ = v25
    default:
        var v26 string = ("\n\t" + "foo bar")
        _ = v26
    }
    var v27 bool = (32767 > -186)
    if ((^(v2)) < v2) {
        if (1.5 != 1.5) {
            switch "hello" {
            case "foo bar":
                if (!false) {
                    fmt.Println(v24)
                }
            }
            var v28 int32 = -1
            fmt.Println(v23, v28)
            _ = v28
        }
        var v29 uint16 = v21
        _ = v29
    } else {
        switch "hello" {
        case "a":
            for fi30 := 0; fi30 < 16; fi30++ {
                v1[(fi30 | (fi30 * 9223372036854775807))] = fi30
                _ = fi30
            }
        }
    }
    for fi31 := 0; fi31 < 14; fi31++ {
        var v32 float64 = float64(1e10)
        _ = fi31
        _ = v32
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
