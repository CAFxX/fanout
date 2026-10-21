package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint32; F1 int8; F2 bool }
type GenS4 struct { F0 uint64; F1 int64 }
type GenS5 struct { F0 int }

func main() {
    var v1 bool = false
    v2 := make(map[string]string)
    v3 := make(map[uint32]int)
    var v4 float64 = (1e10 - ((0.1 / float64(1.5)) / (1.5 + 1e10)))
    var v5 uint64 = 4294967295
    v6 := []int{224, -92, -888}
    v1 = v1
    var v7 int64 = 42
    v4 -= (1.5 / v4)
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    v8 := GenS1{F0: 0}
    if v1 {
        var v9 uint32 = 2147483648
        fmt.Println(v7, v5)
        _ = v9
    } else {
        for fi10 := 0; fi10 < 9; fi10++ {
            _ = len(v3)
            _ = fi10
        }
        var v11 int32 = -2147483648
        v1 = v1
        _ = v11
    }
    v12 := v2["hello"]
    v3[848] = 9223372036854775807
    if ((0.1 == -2.25) && v1) {
        switch v12 {
        case "foo bar":
            for fi13 := 0; fi13 < 14; fi13++ {
                v4 = 0.0
                _ = fi13
            }
        case "hello":
            if (4294967295 >= 146) {
                v7 = 544
                switch 128 {
                case -2147483648:
                    fmt.Println(v5, v4, v12)
                case -9223372036854775808:
                    v14 := v3[42]
                    _ = v14
                }
                switch "a" {
                case "foo bar":
                    for fi15 := 0; fi15 < 9; fi15++ {
                        _ = len(v2)
                        var v16 int8 = -1
                        _ = fi15
                        _ = v16
                    }
                case "\n\t":
                    var v17 uint = 42
                    _ = v17
                case "a":
                    v5 = v5
                }
            } else {
                if (0 != 255) {
                    fmt.Println(v1, v12)
                    v18 := GenS2{F0: 848}
                    var v19 int32 = -215
                    _ = v18
                    _ = v19
                }
            }
        default:
            var v20 float64 = ((v4 / v4) - v4)
            _ = v20
        }
        switch 127 {
        case 127:
            v4 = v4
        default:
            for fi21 := 0; fi21 < 4; fi21++ {
                v22 := GenS3{F0: 359, F1: 127, F2: true}
                var v23 bool = (42 < v5)
                v6 = append(v6, (fi21 * (fi21 % ((fi21) | 1))))
                _ = fi21
                _ = v22
                _ = v23
            }
        }
    }
    v6 = append(v6, 0)
    v24 := GenS4{F0: 42, F1: v7}
    if false {
        fmt.Println(v7, v4)
        fmt.Println(v4)
        v25 := GenS5{F0: 868}
        _ = v25
    } else {
        v26 := GenS4{F0: v5, F1: v7}
        _ = v26
    }
    v27 := GenS3{F0: 806, F1: 0, F2: true}
    v12 = "\n\t"
    v28 := GenS5{F0: -101}
    v4 = v4
    var v29 uint = 18446744073709551615
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v2
    _ = v24
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
