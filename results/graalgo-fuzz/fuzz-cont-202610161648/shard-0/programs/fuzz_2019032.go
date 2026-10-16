package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 bool; F1 uint }

func main() {
    v1 := make(map[int]bool)
    var v2 float32 = -2.25
    v3 := make(map[string]int)
    v4 := []int{0, 32767, -1}
    var v5 int16 = 838
    for fi6 := 0; fi6 < 20; fi6++ {
        v2 = (-2.25 - (float32(float64(-2.25)) - float32(0.1)))
        v7 := GenS1{F0: -647}
        switch (fi6 * (fi6 | (fi6 << ((65535) & 63)))) {
        case -242:
            var v8 uint64 = 0
            _ = v8
        case 564:
            var v9 uint32 = 2
            _ = v9
        case -2147483648:
            if false {
                for fi10 := 0; fi10 < 5; fi10++ {
                    fmt.Println(v2, fi6, v2)
                    _ = fi10
                }
                v11 := GenS1{F0: 0}
                var v12 uint32 = 789
                _ = v11
                _ = v12
            }
        default:
            var v13 float64 = (-2.25 + 1.5)
            _ = v13
        }
        _ = fi6
        _ = v7
    }
    switch 9223372036854775807 {
    case 127:
        _ = len(v3)
    }
    for fi14 := 0; fi14 < 15; fi14++ {
        for fi15 := 0; fi15 < 17; fi15++ {
            var v16 uint64 = 4294967295
            v16 += (v16 & (v16 + 658))
            if true {
                v2 = ((1e10 / v2) * (3.14159 * ((v2 + 1.5) - 0.1)))
                switch (fi15 + 835) {
                case 127:
                    v17 := GenS2{F0: (!false), F1: 705}
                    _ = v17
                case 584:
                    v16 = (^(v16))
                case 4294967295:
                    fmt.Println(v16, fi15, v5)
                default:
                    switch fi14 {
                    case 2147483647:
                        v2 = v2
                    case -1:
                        var v18 int8 = -1
                        _ = v18
                    default:
                        v19 := GenS1{F0: 42}
                        _ = v19
                    }
                }
                _ = len(v4)
            }
            _ = fi15
            _ = v16
        }
        if len(v4) > 0 {
            v4[0] = (^(fi14))
        }
        var v20 int32 = 0
        _ = fi14
        _ = v20
    }
    var v21 uint16 = 0
    switch "" {
    case "hello":
        if len(v4) > 0 {
            v4[0] = 128
        }
    case "a":
        var v22 uint = 42
        _ = v22
    default:
        for fi23 := 0; fi23 < 14; fi23++ {
            v21 = v21
            var v24 float64 = 0.1
            _ = fi23
            _ = v24
        }
    }
    delete(v1, -9223372036854775808)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
