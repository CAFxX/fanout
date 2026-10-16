package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8 }

func main() {
    v1 := make(map[int]bool)
    var v2 uint64 = 42
    var v3 int = -869
    v4 := []uint32{2147483648, 0}
    var v5 uint = 341
    v6 := GenS1{F0: 0, F1: 1}
    _ = len(v1)
    v5 = v5
    fmt.Println(v5, v5)
    switch "a" {
    case "\n\t":
        _ = len(v4)
    case "foo bar":
        var v7 int32 = 844
        _ = v7
    case "a":
        fmt.Println(v3, v3)
    }
    v2 = v2
    v2 = v2
    v5 -= v5
    var v8 float64 = 3.14159
    v2 -= 2
    for fi9 := 0; fi9 < 3; fi9++ {
        var v10 uint16 = 42
        _ = fi9
        _ = v10
    }
    var v11 bool = false
    if v11 {
        for fi12 := 0; fi12 < 14; fi12++ {
            v5 += v5
            var v13 uint64 = v2
            if v11 {
                v3 = fi12
                v14 := v1[(v3 << (((v5 >> ((v5) & 63))) & 63))]
                _ = v14
            }
            _ = fi12
            _ = v13
        }
    } else {
        var v15 uint8 = 39
        if v11 {
            if v11 {
                var v16 int8 = 1
                switch ("foo bar" + "") {
                case "hello":
                    v17 := GenS1{F0: 0, F1: (v15 << (((v5 & v5)) & 63))}
                    _ = v17
                default:
                    var v18 int16 = 32767
                    _ = v18
                }
                _ = v16
            } else {
                switch "a" {
                case "":
                    v15 = v15
                case "a":
                    v19 := GenS1{F0: 9223372036854775807, F1: v15}
                    _ = v19
                case "\n\t":
                    for fi20 := 0; fi20 < 12; fi20++ {
                        switch fi20 {
                        case -923:
                            var v21 uint32 = 65535
                            _ = v21
                        default:
                            var v22 int = (fi20 % ((v3) | 1))
                            _ = v22
                        }
                        fmt.Println(fi20)
                        _ = fi20
                    }
                default:
                    if v11 {
                        switch ("foo bar" + "\n\t") {
                        case "foo bar":
                            var v23 int64 = 0
                            _ = v23
                        }
                        for fi24 := 0; fi24 < 12; fi24++ {
                            for fi25 := 0; fi25 < 15; fi25++ {
                                v26 := GenS1{F0: -9223372036854775808, F1: uint8(v2)}
                                _ = fi25
                                _ = v26
                            }
                            _ = fi24
                        }
                    } else {
                        v27 := GenS1{F0: 9223372036854775807, F1: 128}
                        _ = v27
                    }
                }
                if (-9223372036854775808 == -51) {
                    v15 += (v15 % (((^(v15))) | 1))
                    fmt.Println(v5, v11, v11)
                }
            }
            var v28 int8 = 42
            v29 := GenS1{F0: -1, F1: v15}
            _ = v28
            _ = v29
        } else {
            v30 := v1[(v3 | v3)]
            _ = v30
        }
        switch ("foo bar" + "foo bar") {
        case "a":
            var v31 float32 = 1.5
            _ = v31
        default:
            fmt.Println(v15, v8, v2)
        }
        _ = v15
    }
    var v32 uint8 = 127
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
