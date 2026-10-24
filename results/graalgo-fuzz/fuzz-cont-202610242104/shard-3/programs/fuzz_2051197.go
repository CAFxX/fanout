package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16 }

func main() {
    v1 := []int{-9223372036854775808, 128}
    var v2 float64 = float64(1e10)
    var v3 uint = 0
    v2 += (v2 * float64((0.1 - -2.25)))
    v4 := GenS1{F0: 42, F1: 267}
    v3 = (v3 + v3)
    v1 = append(v1, -1)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v2 = 0.0
    var v5 int8 = -1
    for fi6 := 0; fi6 < 2; fi6++ {
        v3 = v3
        for fi7 := 0; fi7 < 17; fi7++ {
            switch -2147483648 {
            case 2147483647:
                v8 := GenS1{F0: 0, F1: 32767}
                _ = v8
            }
            v9 := GenS1{F0: 1, F1: 779}
            var v10 string = "\n\t"
            _ = fi7
            _ = v10
            _ = v9
        }
        if (-233 == -221) {
            if false {
                if len(v1) > 0 {
                    v1[0] = fi6
                }
                v5 -= (v5 >> ((v3) & 63))
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
        }
        _ = fi6
    }
    v2 += v2
    v3 = 296
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
