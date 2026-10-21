package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int }

func main() {
    v1 := make(map[int]uint32)
    var v2 uint = 255
    var v3 int = 1
    fmt.Println(v2, v2)
    var v4 int = (v3 & (v3 + 703))
    v5 := v1[(v4 / ((-2147483648) | 1))]
    v3 += -2147483648
    if false {
        v6 := GenS1{F0: -730, F1: (v3 ^ v4)}
        var v7 int16 = 1
        for fi8 := 0; fi8 < 13; fi8++ {
            var v9 uint = 42
            if (1 < -2147483648) {
                v2 += uint(v5)
                fmt.Println(fi8)
                var v10 bool = true
                _ = v10
            } else {
                for fi11 := 0; fi11 < 1; fi11++ {
                    switch v3 {
                    case 1:
                        fmt.Println(v9, fi8, fi8)
                    default:
                        v1[(fi8 - fi11)] = (v5 << ((v9) & 63))
                    }
                    switch "\n\t" {
                    case "hello":
                        v12 := v1[-917]
                        _ = v12
                    default:
                        var v13 int64 = -1
                        _ = v13
                    }
                    v2 = v9
                    _ = fi11
                }
                for fi14 := 0; fi14 < 12; fi14++ {
                    var v15 uint8 = 166
                    fmt.Println(v4)
                    switch fi8 {
                    case 127:
                        var v16 int32 = 518
                        _ = v16
                    case -1:
                        fmt.Println(v15, v7, v7)
                    case 4294967295:
                        v5 -= uint32(v15)
                    default:
                        if true {
                            v17 := []int{(-(fi8)), 4294967295, 112, fi14}
                            fmt.Println(v4)
                            v18 := GenS1{F0: -413, F1: 113}
                            _ = v17
                            _ = v18
                        } else {
                            v19 := GenS1{F0: 1, F1: 9223372036854775807}
                            var v20 float64 = 0.0
                            _ = len(v1)
                            _ = v19
                            _ = v20
                        }
                        var v21 int64 = int64(v2)
                        _ = v21
                    }
                    _ = fi14
                    _ = v15
                }
                switch (("\n\t" + "hello") + ("a" + ("a" + "a"))) {
                case "":
                    var v22 int64 = -419
                    _ = v22
                case "a":
                    v5 += 42
                }
            }
            fmt.Println(v7)
            _ = fi8
            _ = v9
        }
        _ = v6
        _ = v7
    }
    fmt.Println(v2)
    var v23 uint8 = uint8(v5)
    v24 := []int{(v4 ^ (v4 - (v3 | v4)))}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
