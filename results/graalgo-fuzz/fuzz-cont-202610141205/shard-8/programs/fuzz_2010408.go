package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 bool; F1 int32 }

func main() {
    v1 := []uint32{42, 714, 2147483647}
    var v2 float64 = (1.5 - 3.14159)
    v3 := make(map[string]uint32)
    for fi4 := 0; fi4 < 2; fi4++ {
        var v5 uint64 = 18446744073709551615
        v5 += 1
        _ = fi4
        _ = v5
    }
    v1 = append(v1, 0)
    if false {
        var v6 int8 = 100
        var v7 int64 = -326
        v7 += 607
        _ = v6
        _ = v7
    }
    v2 -= v2
    v8 := GenS1{F0: 743}
    for fi9 := 0; fi9 < 7; fi9++ {
        v2 -= v2
        switch "hello" {
        case "hello":
            fmt.Println(v2)
        case "foo bar":
            if true {
                switch ("hello" + "\n\t") {
                case "hello":
                    for fi10 := 0; fi10 < 12; fi10++ {
                        v2 = 0.1
                        var v11 string = "\n\t"
                        if (18 > 100) {
                            v12 := GenS1{F0: 596}
                            v13 := GenS1{F0: 0}
                            if true {
                                v2 += v2
                            } else {
                                switch "\n\t" {
                                case "\n\t":
                                    v11 = v11
                                case "hello":
                                    if (float32((-2.25 * v2)) == 3.14159) {
                                        var v14 int32 = 42
                                        v2 = float64(float32(1e10))
                                        _ = v14
                                    } else {
                                        var v15 uint8 = 255
                                        v11 = v11
                                        _ = v15
                                    }
                                default:
                                    var v16 float64 = 0.0
                                    _ = v16
                                }
                                switch (v11 + "\n\t") {
                                case "hello":
                                    v11 = v11
                                    var v17 uint = 398
                                    var v18 float32 = 3.14159
                                    _ = v17
                                    _ = v18
                                case "":
                                    v2 = v2
                                default:
                                    switch fi9 {
                                    case 560:
                                        if (((fi9 ^ fi10) <= fi9) || (false && false)) {
                                            v2 += (v2 / (v2 + v2))
                                            v19 := GenS2{F0: false, F1: 907}
                                            _ = v19
                                        } else {
                                            v11 = ""
                                        }
                                    case -9223372036854775808:
                                        v20 := GenS1{F0: 255}
                                        _ = v20
                                    }
                                }
                            }
                            _ = v12
                            _ = v13
                        }
                        _ = fi10
                        _ = v11
                    }
                case "":
                    v1 = append(v1, 42)
                }
            } else {
                var v21 int16 = -32768
                fmt.Println(v21)
                _ = v21
            }
        case "":
            switch "a" {
            case "a":
                for fi22 := 0; fi22 < 3; fi22++ {
                    switch (fi9 & 946) {
                    case 42:
                        var v23 uint32 = 206
                        _ = v23
                    case 127:
                        if true {
                            v24 := GenS2{F0: true, F1: -1}
                            _ = v24
                        } else {
                            v1 = append(v1, 4294967295)
                        }
                    }
                    v2 -= 1.5
                    if (!true) {
                        v2 += float64(1e10)
                        v2 -= float64((float32(1e10) / 3.14159))
                        if ((0.0 > 0.1) && false) {
                            v25 := GenS2{F0: (98 <= 226), F1: 355}
                            v2 = float64(((-2.25 * (1e10 / 1.0)) * ((-2.25 / 1.0) + 0.0)))
                            var v26 uint64 = 42
                            _ = v25
                            _ = v26
                        }
                    }
                    _ = fi22
                }
            }
        default:
            if (-32768 < -1) {
                var v27 uint64 = 165
                var v28 int32 = -2147483648
                _ = v27
                _ = v28
            } else {
                var v29 float32 = -2.25
                if false {
                    var v30 uint = 817
                    _ = v30
                } else {
                    var v31 string = (("" + "\n\t") + "a")
                    if false {
                        var v32 uint16 = 32767
                        _ = v32
                    } else {
                        v31 = v31
                        v29 += 1e10
                        switch (v31 + v31) {
                        case "hello":
                            fmt.Println(v29)
                        default:
                            var v33 string = "hello"
                            _ = v33
                        }
                    }
                    if (((0.1 / float64((v29 - v29))) * v2) == (1.5 * (float64(v29) / v2))) {
                        delete(v3, ("\n\t" + v31))
                    }
                    _ = v31
                }
                var v34 int64 = 9223372036854775807
                _ = v29
                _ = v34
            }
        }
        switch "hello" {
        case "\n\t":
            fmt.Println(fi9, fi9)
        }
        _ = fi9
    }
    var v35 string = ""
    fmt.Println(v2, v35)
    v2 += (v2 / v2)
    if len(v1) > 0 {
        v1[0] = 42
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v8
}
