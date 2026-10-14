package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16 }
type GenS2 struct { F0 int }

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[uint32]uint32)
    v3 := []uint32{1}
    var v4 uint = 574
    v4 = v4
    v4 += (v4 / (((v4 + v4)) | 1))
    for fi5 := 0; fi5 < 11; fi5++ {
        v4 -= 1
        _ = fi5
    }
    _ = len(v3)
    switch "hello" {
    case "a":
        if true {
            v4 += (v4 % ((v4) | 1))
            var v6 int64 = 42
            _ = v6
        }
    }
    switch "foo bar" {
    case "":
        fmt.Println(v4)
    case "a":
        fmt.Println(v4)
    default:
        v4 = v4
    }
    _ = len(v3)
    fmt.Println(v4)
    switch "" {
    case "":
        var v7 int = 0
        _ = v7
    }
    for fi8 := 0; fi8 < 7; fi8++ {
        if len(v3) > 0 {
            v3[0] = 65535
        }
        switch ("\n\t" + "hello") {
        case "a":
            for fi9 := 0; fi9 < 10; fi9++ {
                if (false || (1 < 65535)) {
                    v3 = append(v3, 283)
                }
                _ = fi9
            }
        case "\n\t":
            for fi10 := 0; fi10 < 4; fi10++ {
                if (65535 < 255) {
                    if true {
                        fmt.Println(fi10, fi8)
                        v4 = (v4 % (((^(v4))) | 1))
                        v4 = v4
                    } else {
                        var v11 int16 = 1
                        var v12 int32 = 0
                        var v13 uint32 = 2147483647
                        _ = v11
                        _ = v12
                        _ = v13
                    }
                    v14 := GenS1{F0: 0, F1: 1}
                    _ = v14
                } else {
                    v4 = 18446744073709551615
                    v15 := GenS1{F0: -274, F1: 736}
                    _ = v15
                }
                if true {
                    v16 := GenS2{F0: fi10}
                    _ = v16
                } else {
                    if (!true) {
                        v4 = 42
                    } else {
                        if true {
                            v4 = v4
                            v4 = (v4 % ((700) | 1))
                            var v17 float32 = 3.14159
                            _ = v17
                        }
                        v4 -= v4
                    }
                }
                _ = fi10
            }
        case "hello":
            v4 += 416
        default:
            for fi18 := 0; fi18 < 8; fi18++ {
                var v19 int64 = 0
                if true {
                    if (!((1.5 - 3.14159) >= -2.25)) {
                        v1[(fi8 + 0)] = false
                        var v20 int16 = 42
                        fmt.Println(fi18, v4)
                        _ = v20
                    } else {
                        fmt.Println(fi18, fi8)
                    }
                    v19 -= (^(v19))
                    if (v4 == (^(v4))) {
                        v19 += v19
                        fmt.Println(v19)
                        if false {
                            fmt.Println(fi18)
                        }
                        fmt.Println(fi8, v4)
                    }
                } else {
                    var v21 uint = v4
                    _ = v21
                }
                if len(v3) > 0 {
                    v3[0] = 491
                }
                _ = fi18
                _ = v19
            }
        }
        _ = fi8
    }
    var v22 uint8 = 73
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v4
}
