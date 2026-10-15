package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8; F2 uint16 }
type GenS2 struct { F0 int8; F1 bool; F2 int16 }
type GenS3 struct { F0 int16 }

func main() {
    var v1 float64 = (0.0 + 0.1)
    var v2 int = -442
    var v3 uint8 = 19
    if (-1 == 785) {
        var v4 uint8 = (v3 << ((158) & 63))
        switch "a" {
        case "foo bar":
            v5 := []int{(v2 * v2), (v2 << ((42) & 63)), v2}
            _ = v5
        }
        switch 179 {
        case -217:
            v1 -= (v1 * (((0.1 / v1) * (v1 * v1)) / v1))
        default:
            for fi6 := 0; fi6 < 5; fi6++ {
                if (!(!true)) {
                    if (339 <= 1) {
                        v2 += v2
                        v7 := GenS1{F0: 970, F1: -1, F2: 255}
                        for fi8 := 0; fi8 < 8; fi8++ {
                            v4 = (v4 << ((4294967295) & 63))
                            fmt.Println(v3)
                            _ = fi8
                        }
                        _ = v7
                    } else {
                        for fi9 := 0; fi9 < 4; fi9++ {
                            v4 = v4
                            _ = fi9
                        }
                        v10 := GenS1{F0: 18446744073709551615, F1: 100, F2: 255}
                        _ = v10
                    }
                    var v11 float64 = (v1 / v1)
                    _ = v11
                }
                fmt.Println(fi6)
                fmt.Println(v4)
                _ = fi6
            }
        }
        _ = v4
    }
    switch (^(v2)) {
    case 647:
        if (93 > v2) {
            if false {
                v3 = v3
                v3 -= (v3 << ((2) & 63))
                if (!(1 != -32768)) {
                    v2 += (v2 << ((2) & 63))
                }
            }
            for fi12 := 0; fi12 < 14; fi12++ {
                var v13 bool = true
                _ = fi12
                _ = v13
            }
            var v14 float32 = 0.1
            _ = v14
        } else {
            v15 := []int{(v2 % ((32768) | 1)), v2, 99, v2}
            _ = v15
        }
    }
    v16 := GenS2{F0: 0, F1: true, F2: -32768}
    v17 := []int{}
    if len(v17) > 0 {
        v17[0] = (v2 << ((1) & 63))
    }
    v2 += 128
    var v18 uint32 = 42
    var v19 uint8 = (v3 * 27)
    var v20 bool = (2 != 42)
    v21 := GenS3{F0: -1}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
}
