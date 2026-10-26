package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8; F2 int8 }

func main() {
    var v1 int64 = -1
    v2 := make(map[int]bool)
    var v3 float64 = ((((0.0 + (1e10 / (float64(0.1) / ((float64((((1.5 - (-2.25 - 1e10)) * ((0.0 * 0.0) * -2.25)) / 1e10)) * (0.0 / 3.14159)) - (-2.25 + (float64(1e10) * float64(1e10))))))) + (1.5 - -2.25)) * 0.1) / (-2.25 / float64(float32(0.1))))
    v4 := make(map[int]int)
    v1 += v1
    for fi5 := 0; fi5 < 15; fi5++ {
        for fi6 := 0; fi6 < 10; fi6++ {
            if ((1.5 < (1e10 * 3.14159)) || (42 >= 56)) {
                switch "foo bar" {
                case "\n\t":
                    var v7 int32 = 42
                    _ = v7
                case "":
                    v3 = v3
                case "hello":
                    var v8 uint16 = 42
                    _ = v8
                }
                v4[(-(fi5))] = fi5
            } else {
                switch "a" {
                case "\n\t":
                    v1 = (v1 & (^(v1)))
                case "":
                    var v9 uint16 = 32767
                    _ = v9
                default:
                    v10 := []int{32767, 1, (fi5 - fi6), 672}
                    _ = v10
                }
                if (0 < 698) {
                    fmt.Println(fi6, v1)
                    var v11 uint8 = 0
                    _ = v11
                }
            }
            _ = fi6
        }
        _ = fi5
    }
    v12 := GenS1{F0: 0, F1: 1, F2: 100}
    for fi13 := 0; fi13 < 7; fi13++ {
        delete(v2, fi13)
        switch "a" {
        case "a":
            if false {
                v3 = (-2.25 - v3)
            } else {
                switch "" {
                case "":
                    for fi14 := 0; fi14 < 7; fi14++ {
                        fmt.Println(fi14)
                        v15 := v4[fi13]
                        _ = fi14
                        _ = v15
                    }
                default:
                    v16 := GenS1{F0: 2147483647, F1: 1, F2: -73}
                    _ = v16
                }
            }
        case "hello":
            v3 = v3
        }
        v17 := GenS1{F0: 0, F1: 72, F2: 1}
        _ = fi13
        _ = v17
    }
    v1 -= v1
    for fi18 := 0; fi18 < 10; fi18++ {
        v1 = (v1 | v1)
        v1 = 42
        switch (fi18 | 2147483647) {
        case 2147483647:
            v1 = v1
        case 32767:
            v19 := GenS1{F0: 2147483647, F1: 42, F2: 46}
            _ = v19
        case 76:
            if (false && (9223372036854775807 == 128)) {
                v20 := []int{fi18}
                _ = v20
            }
        }
        _ = fi18
    }
    fmt.Println(v1, v3)
    switch 2147483647 {
    case 127:
        v21 := []int{128, 42}
        _ = v21
    }
    v1 += v1
    var v22 uint64 = 2
    v22 = (v22 + (v22 >> ((883) & 63)))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v12
    _ = v2
    _ = v22
    _ = v3
    _ = v4
}
