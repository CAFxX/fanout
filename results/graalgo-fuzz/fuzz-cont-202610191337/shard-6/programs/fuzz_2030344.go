package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16 }

func main() {
    v1 := make(map[string]int)
    var v2 float64 = float64(1.5)
    var v3 float64 = v2
    v2 -= ((1e10 / 3.14159) + float64(float32(v3)))
    if true {
        v3 = 0.0
    }
    v2 = (1.5 * ((1e10 + 3.14159) / ((v3 * (v3 * float64(0.1))) + (0.0 + 0.1))))
    v4 := GenS1{F0: -1, F1: 1000}
    v3 += float64(0.0)
    v2 = float64(1e10)
    var v5 int = -884
    fmt.Println(v5, v3, v5)
    v6 := GenS1{F0: -1, F1: 959}
    v3 = -2.25
    v2 -= (1.5 * v2)
    fmt.Println(v5, v5, v5)
    fmt.Println(v3, v5, v5)
    v7 := []int{(^(v5)), v5, v5}
    for fi8 := 0; fi8 < 5; fi8++ {
        for fi9 := 0; fi9 < 6; fi9++ {
            v5 = -437
            switch (v5 >> ((53) & 63)) {
            case -2147483648:
                if true {
                    switch "hello" {
                    case "a":
                        v2 += float64(1e10)
                    case "foo bar":
                        v5 -= fi8
                    }
                    var v10 uint = 178
                    _ = v10
                } else {
                    v11 := v1[""]
                    _ = v11
                }
            case 4294967295:
                v12 := GenS1{F0: 0, F1: -466}
                _ = v12
            case 32767:
                v1["\n\t"] = fi9
                var v13 int16 = -544
                v1[""] = (fi8 * fi8)
                if true {
                    v5 = v5
                }
                fmt.Println(v3)
                v5 += 9223372036854775807
                _ = v13
            default:
                if (115 == -32768) {
                    var v14 bool = (42 < 856)
                    _ = len(v1)
                    v2 = (0.1 / v2)
                    _ = v14
                }
            }
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
            _ = fi9
        }
        for fi15 := 0; fi15 < 16; fi15++ {
            var v16 float32 = 0.0
            _ = fi15
            _ = v16
        }
        _ = fi8
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
