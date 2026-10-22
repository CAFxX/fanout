package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint64 }
func genFn0(p0 int32, p1 bool, p2 uint8) int16 {
    fmt.Println(p1, p1)
    if (true && false) {
        var v4 uint8 = 1
        v4 -= (v4 & 191)
        v4 = (p2 + (p2 << ((148) & 63)))
        _ = v4
    }
    _ = p0
    _ = p1
    _ = p2
    return 32767
}

func main() {
    var v1 uint = 42
    v2 := make(map[string]bool)
    var v3 float32 = -2.25
    v5 := GenS1{F0: (!(122 > 107))}
    for fi6 := 0; fi6 < 11; fi6++ {
        v3 = ((v3 / (v3 / -2.25)) * v3)
        _ = fi6
    }
    var v7 uint = 2
    switch "hello" {
    case "":
        for fi8 := 0; fi8 < 14; fi8++ {
            v3 = (float32(0.0) + (float32(0.0) * 0.1))
            switch "a" {
            case "a":
                v9 := GenS2{F0: v1}
                _ = v9
            }
            _ = fi8
        }
    case "a":
        var v10 bool = true
        _ = v10
    default:
        var v11 string = ""
        _ = v11
    }
    v12 := GenS3{F0: 127}
    switch 32767 {
    case -9223372036854775808:
        v13 := GenS4{F0: 38}
        _ = v13
    case 127:
        v7 += (v7 ^ v7)
    case 1:
        v1 = (v1 + 460)
    default:
        switch "" {
        case "hello":
            if (!false) {
                var v14 int64 = -121
                if (!false) {
                    for fi15 := 0; fi15 < 7; fi15++ {
                        for fi16 := 0; fi16 < 18; fi16++ {
                            fmt.Println(v1, fi15, v7)
                            if (42 <= 1) {
                                var v17 int = -1
                                _ = v17
                            } else {
                                v18 := []int{}
                                var v19 int16 = 42
                                var v20 uint64 = 42
                                _ = v18
                                _ = v19
                                _ = v20
                            }
                            v7 = (v1 * 760)
                            _ = fi16
                        }
                        _ = fi15
                    }
                } else {
                    if (!(false && false)) {
                        if ((1.5 < 0.1) || (!(44 <= uint8(v1)))) {
                            v21 := GenS2{F0: 65535}
                            fmt.Println(v3)
                            var v22 uint16 = 69
                            _ = v21
                            _ = v22
                        }
                    }
                }
                _ = v14
            }
        case "":
            v23 := GenS4{F0: 392}
            _ = v23
        case "\n\t":
            v1 = 2
        }
    }
    v24 := []int{}
    for fi25 := 0; fi25 < 12; fi25++ {
        var v26 int16 = -32768
        _ = fi25
        _ = v26
    }
    var v27 int16 = -32768
    v3 = 3.14159
    fmt.Println(v3)
    v24 = append(v24, 224)
    v3 -= 0.0
    for fi28 := 0; fi28 < 15; fi28++ {
        v27 += (-(v27))
        v7 += 666
        _ = fi28
    }
    var v29 float32 = v3
    switch 2147483647 {
    case 128:
        v27 = -839
    case -920:
        fmt.Println(v27, v29, v27)
    default:
        for fi30 := 0; fi30 < 13; fi30++ {
            var v31 uint16 = 1
            _ = len(v24)
            _ = fi30
            _ = v31
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v2
    _ = v24
    _ = v27
    _ = v29
    _ = v3
    _ = v5
    _ = v7
}
