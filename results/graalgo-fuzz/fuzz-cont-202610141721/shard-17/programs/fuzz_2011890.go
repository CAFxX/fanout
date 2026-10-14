package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64 }
type GenS2 struct { F0 int }
func genFn0(p0 bool) uint8 {
    v7 := GenS1{F0: 873, F1: -434}
    p0 = (127 <= 100)
    _ = p0
    _ = v7
    return 0
}
func genFn1(p0 uint8, p1 uint16, p2 int) int8 {
    fmt.Println(p0)
    fmt.Println(p2)
    var v8 int = p2
    fmt.Println(p0, p1, p2)
    fmt.Println(p2, v8, p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return -128
}

func main() {
    var v1 uint16 = 0
    var v2 uint = 255
    var v3 uint16 = v1
    v4 := make(map[string]uint32)
    v5 := make(map[int]bool)
    v6 := make(map[string]int)
    var v9 string = "\n\t"
    var v10 uint32 = 809
    switch 2147483647 {
    case 378:
        v11 := []int{-777, 469, 42, 195}
        _ = v11
    case 127:
        switch 4294967295 {
        case 0:
            v3 += v3
        case 32768:
            switch 889 {
            case -9223372036854775808:
                switch 4294967295 {
                case -870:
                    var v12 uint = v2
                    _ = v12
                case 127:
                    v1 = (^(v1))
                case 32768:
                    if (false || (!(42 != 2))) {
                        var v13 int32 = 2147483647
                        _ = v13
                    } else {
                        if true {
                            fmt.Println(v10, v9, v2)
                        }
                    }
                }
            case -532:
                switch v9 {
                case "hello":
                    v2 += v2
                case "\n\t":
                    var v14 int32 = -733
                    _ = v14
                default:
                    if true {
                        v10 = (^(v10))
                        if true {
                            v2 += (v2 % (((v2 + 0)) | 1))
                            v10 += (v10 & (^(v10)))
                            var v15 int16 = -32768
                            _ = v15
                        } else {
                            fmt.Println(v2)
                            v16 := GenS2{F0: 32767}
                            for fi17 := 0; fi17 < 4; fi17++ {
                                var v18 float64 = float64(0.1)
                                _ = fi17
                                _ = v18
                            }
                            _ = v16
                        }
                        var v19 float32 = (0.1 + -2.25)
                        _ = v19
                    }
                }
            case 0:
                fmt.Println(v9)
            default:
                var v20 int32 = -1
                _ = v20
            }
        }
    }
    v21 := v5[42]
    fmt.Println(v21, v9)
    var v22 uint32 = (v10 * (v10 << (((v2 << ((v2) & 63))) & 63)))
    v23 := GenS2{F0: 32768}
    _ = len(v4)
    switch -294 {
    case 2147483647:
        var v24 int = 4294967295
        _ = v24
    case -396:
        for fi25 := 0; fi25 < 4; fi25++ {
            v3 = 2
            _ = fi25
        }
    default:
        fmt.Println(v9, v21)
    }
    v26 := v5[9223372036854775807]
    v27 := []int{-683, 1, -9223372036854775808}
    for fi28 := 0; fi28 < 3; fi28++ {
        delete(v5, 746)
        var v29 uint64 = 997
        _ = fi28
        _ = v29
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
