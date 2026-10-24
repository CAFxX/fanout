package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64; F2 int8 }
type GenS2 struct { F0 uint32; F1 uint }
func genFn0() uint {
    var v7 uint32 = 704
    var v8 float32 = 0.0
    _ = v7
    _ = v8
    return 0
}

func main() {
    v1 := []string{"hello", ("" + "hello"), "\n\t", (("" + "foo bar") + "hello")}
    var v2 uint32 = 844
    v3 := make(map[uint32]int)
    var v4 bool = (true || (false || false))
    var v5 uint8 = 157
    var v6 uint32 = v2
    fmt.Println(v4, v4, v2)
    fmt.Println(v6)
    for fi9 := 0; fi9 < 15; fi9++ {
        fmt.Println(v4, v2, v5)
        v10 := GenS1{F0: v6, F1: 42, F2: 100}
        _ = fi9
        _ = v10
    }
    switch "\n\t" {
    case "foo bar":
        var v11 int8 = 42
        _ = v11
    case "\n\t":
        v1 = append(v1, "\n\t")
    case "":
        var v12 int16 = -1
        _ = v12
    default:
        for fi13 := 0; fi13 < 7; fi13++ {
            if (255 >= 255) {
                v2 = (v6 * (v2 / ((2) | 1)))
            }
            var v14 bool = (1e10 <= float64((1.5 * (float32(-2.25) / (3.14159 / (float32((-2.25 + 0.0)) / -2.25))))))
            _ = fi13
            _ = v14
        }
    }
    v5 = (v5 >> ((genFn0()) & 63))
    _ = len(v3)
    fmt.Println(v2, v6)
    if true {
        v5 = v5
    } else {
        var v15 int32 = 569
        switch "" {
        case "a":
            switch 1 {
            case 0:
                var v16 float32 = 1e10
                _ = v16
            case 32767:
                v17 := GenS1{F0: 1, F1: 1, F2: 0}
                _ = v17
            default:
                v18 := GenS1{F0: v2, F1: 42, F2: 7}
                _ = v18
            }
        case "foo bar":
            var v19 int64 = 16
            _ = v19
        default:
            var v20 int64 = -9223372036854775808
            _ = v20
        }
        _ = v15
    }
    var v21 int16 = -1
    v6 = (v2 & v2)
    var v22 float64 = -2.25
    _ = len(v3)
    v5 += 255
    var v23 uint64 = 0
    var v24 int32 = -3
    v21 += v21
    if false {
        for fi25 := 0; fi25 < 7; fi25++ {
            v24 = 0
            v23 -= (v23 | 1)
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = fi25
        }
    } else {
        fmt.Println(v4, v22)
        var v26 int16 = 260
        v2 += (v2 >> ((715) & 63))
        _ = v26
    }
    for fi27 := 0; fi27 < 12; fi27++ {
        v6 += v2
        v6 = (v6 ^ 462)
        if v4 {
            if v4 {
                fmt.Println(v5, v5)
                v4 = (v23 > (v23 + 1))
            }
            if v4 {
                for fi28 := 0; fi28 < 19; fi28++ {
                    v6 += (^(v2))
                    var v29 float64 = (v22 + 0.1)
                    v5 = (v5 ^ (^(v5)))
                    _ = fi28
                    _ = v29
                }
                if v4 {
                    var v30 int64 = 9223372036854775807
                    _ = v30
                } else {
                    var v31 uint8 = v5
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    _ = v31
                }
            } else {
                fmt.Println(fi27, v21, v2)
            }
            v32 := GenS2{F0: v6, F1: genFn0()}
            _ = v32
        }
        _ = fi27
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
