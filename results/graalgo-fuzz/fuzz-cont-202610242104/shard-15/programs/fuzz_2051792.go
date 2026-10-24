package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8 }
type GenS2 struct { F0 int32 }

func main() {
    var v1 uint64 = 4294967295
    var v2 int32 = -884
    var v3 float64 = 1e10
    var v4 float32 = 0.0
    var v5 bool = false
    v6 := []int{-9223372036854775808, 777, 4294967295}
    v4 = (v4 / 1e10)
    fmt.Println(v2, v3, v2)
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    if len(v6) > 0 {
        v6[0] = 993
    }
    fmt.Println(v1)
    switch "foo bar" {
    case "a":
        v7 := GenS1{F0: 42, F1: -61}
        _ = v7
    case "\n\t":
        if ((v1 | 1) > v1) {
            v3 -= (v3 + v3)
        } else {
            if v5 {
                if v5 {
                    var v8 int = -2147483648
                    for fi9 := 0; fi9 < 4; fi9++ {
                        v3 = (((v3 - 3.14159) - (((v3 + v3) - ((-2.25 + float64((v4 - (v4 - ((float32((v3 + (-2.25 + float64(v4)))) - v4) * 0.1))))) * v3)) * v3)) + float64(v4))
                        v3 += (v3 + (-2.25 - -2.25))
                        v10 := GenS2{F0: v2}
                        _ = fi9
                        _ = v10
                    }
                    fmt.Println(v4)
                    _ = v8
                }
                var v11 uint16 = 42
                _ = v11
            }
        }
    default:
        for fi12 := 0; fi12 < 15; fi12++ {
            v6 = append(v6, (fi12 ^ (fi12 >> ((968) & 63))))
            for fi13 := 0; fi13 < 19; fi13++ {
                var v14 int64 = 427
                _ = fi13
                _ = v14
            }
            _ = fi12
        }
    }
    fmt.Println(v5)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
