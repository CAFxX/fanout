package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }

func main() {
    var v1 uint8 = 0
    var v2 uint32 = 42
    v3 := []int{-1, -2147483648}
    var v4 int64 = -9223372036854775808
    v5 := make(map[string]uint32)
    v6 := []uint32{v2, v2}
    v7 := GenS1{F0: (int8(v2) == int8(v2))}
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    var v8 uint32 = (v2 + (^(v2)))
    if (!false) {
        v4 = v4
    }
    v2 = 489
    v8 = (v2 & (v8 + v2))
    fmt.Println(v1, v4)
    var v9 int = -9223372036854775808
    v1 += v1
    for fi10 := 0; fi10 < 2; fi10++ {
        if false {
            v11 := GenS1{F0: true}
            switch ("foo bar" + "foo bar") {
            case "":
                var v12 int64 = (v4 - (v4 * (v4 | -1)))
                _ = v12
            default:
                switch (^(v9)) {
                case -199:
                    v13 := GenS1{F0: true}
                    _ = v13
                default:
                    switch v9 {
                    case 128:
                        if len(v6) > 0 {
                            v6[0] = (v2 * (v8 / (((v8 << ((1000) & 63))) | 1)))
                        }
                    case 32768:
                        fmt.Println(v9, v2, v8)
                    default:
                        var v14 int16 = -32768
                        _ = v14
                    }
                }
            }
            var v15 uint64 = 797
            _ = v11
            _ = v15
        }
        switch "foo bar" {
        case "foo bar":
            delete(v5, "")
        case "a":
            if len(v3) > 0 {
                v3[0] = 42
            }
        case "\n\t":
            v16 := v5[("\n\t" + "hello")]
            _ = v16
        default:
            switch "hello" {
            case "a":
                var v17 int16 = 1
                _ = v17
            default:
                var v18 bool = (!(255 == 669))
                _ = v18
            }
        }
        switch (v9 & 9223372036854775807) {
        case 42:
            v3 = append(v3, (-(v9)))
        default:
            switch -1 {
            case 9223372036854775807:
                var v19 float32 = (-2.25 * (1e10 / 0.1))
                _ = v19
            default:
                if (0 > (v4 >> ((164) & 63))) {
                    fmt.Println(v4, v9, v4)
                    var v20 uint8 = (^(v1))
                    _ = v20
                } else {
                    switch "" {
                    case "a":
                        v21 := GenS1{F0: true}
                        _ = v21
                    default:
                        _ = len(v5)
                    }
                    for fi22 := 0; fi22 < 4; fi22++ {
                        v23 := GenS1{F0: (false || (1 < 0))}
                        var v24 uint64 = 4294967295
                        var v25 int = fi10
                        fmt.Println(fi22, fi10)
                        var v26 uint = 0
                        var v27 uint8 = 127
                        _ = fi22
                        _ = v23
                        _ = v24
                        _ = v25
                        _ = v26
                        _ = v27
                    }
                    _ = len(v3)
                }
            }
        }
        _ = fi10
    }
    v28 := GenS1{F0: (uint32(v1) >= 42)}
    fmt.Println(v9, v9)
    var v29 float64 = ((-2.25 + (1.5 + 0.1)) - 3.14159)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
