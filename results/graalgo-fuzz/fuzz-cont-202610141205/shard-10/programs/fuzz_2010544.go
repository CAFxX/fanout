package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int }
type GenS2 struct { F0 uint; F1 uint8; F2 int32 }

func main() {
    v1 := []uint32{2147483648, 2147483648}
    var v2 bool = true
    v3 := []uint32{4294967295, 1, 482, 819}
    v4 := make(map[string]string)
    v5 := make(map[int]bool)
    var v6 int16 = 42
    v7 := GenS1{F0: 308, F1: 128}
    if v2 {
        var v8 uint32 = 0
        switch 2147483647 {
        case 1:
            fmt.Println(v6)
        default:
            switch -2147483648 {
            case 32768:
                if false {
                    if (42 == (v8 >> ((158) & 63))) {
                        fmt.Println(v8)
                    } else {
                        v2 = ((!v2) && (!(v2 || (460 > (v6 ^ v6)))))
                        if (v2 || ((1.5 * 1e10) > 1.5)) {
                            for fi9 := 0; fi9 < 3; fi9++ {
                                var v10 uint8 = 255
                                _ = fi9
                                _ = v10
                            }
                            delete(v5, 127)
                        }
                    }
                    if (2 > 221) {
                        var v11 int32 = 747
                        _ = v11
                    } else {
                        if ((2 != 32767) || v2) {
                            fmt.Println(v6, v2)
                            if len(v3) > 0 {
                                v3[0] = (v8 % ((759) | 1))
                            }
                            v12 := GenS2{F0: 798, F1: 255, F2: 420}
                            _ = v12
                        }
                        var v13 int = 521
                        _ = v13
                    }
                    if len(v1) > 0 {
                        v1[0] = v8
                    }
                }
            default:
                fmt.Println(v6, v2)
            }
        }
        v8 -= (^(v8))
        _ = v8
    }
    var v14 bool = true
    v6 += (v6 | (v6 ^ v6))
    for fi15 := 0; fi15 < 4; fi15++ {
        var v16 int8 = 1
        _ = fi15
        _ = v16
    }
    fmt.Println(v14, v2)
    var v17 float64 = 1.5
    if false {
        fmt.Println(v17)
        v14 = (v14 || (false || (!v14)))
        v17 -= (v17 + v17)
    }
    v6 -= (v6 >> ((1) & 63))
    if (-729 <= 2147483647) {
        switch "a" {
        case "":
            var v18 float64 = v17
            _ = v18
        case "\n\t":
            if len(v3) > 0 {
                v3[0] = 2147483647
            }
        }
        for fi19 := 0; fi19 < 3; fi19++ {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            v6 += (^(v6))
            _ = fi19
        }
        var v20 float64 = v17
        _ = v20
    }
    if len(v1) > 0 {
        v1[0] = 1
    }
    for fi21 := 0; fi21 < 7; fi21++ {
        _ = len(v3)
        var v22 uint16 = 2
        v23 := GenS2{F0: 255, F1: 1, F2: 0}
        _ = fi21
        _ = v22
        _ = v23
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
