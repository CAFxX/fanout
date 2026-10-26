package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64 }

func main() {
    var v1 float32 = (float32((-2.25 / 1e10)) + 1e10)
    var v2 uint64 = 4294967295
    var v3 uint16 = 65535
    var v4 float64 = float64(0.0)
    var v5 int32 = -696
    for fi6 := 0; fi6 < 8; fi6++ {
        v1 = 0.1
        _ = fi6
    }
    v7 := []int{699, 92, 2147483647}
    for fi8 := 0; fi8 < 1; fi8++ {
        switch 9223372036854775807 {
        case -1:
            for fi9 := 0; fi9 < 6; fi9++ {
                if false {
                    var v10 float64 = (v4 / v4)
                    fmt.Println(v2, v2, fi9)
                    _ = v10
                } else {
                    if (0 > 42) {
                        fmt.Println(v1, v4)
                    }
                }
                v11 := GenS1{F0: "foo bar", F1: 0}
                _ = fi9
                _ = v11
            }
        }
        for fi12 := 0; fi12 < 4; fi12++ {
            v13 := GenS1{F0: "a", F1: -723}
            v7 = append(v7, fi12)
            fmt.Println(v2, v5)
            _ = fi12
            _ = v13
        }
        _ = fi8
    }
    v4 += -2.25
    var v14 int8 = 0
    v15 := GenS1{F0: ("\n\t" + "\n\t"), F1: -372}
    v1 = v1
    if len(v7) > 0 {
        v7[0] = 1
    }
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    if false {
        var v16 float64 = ((-2.25 + ((v4 * (v4 * v4)) * (0.0 + ((v4 * 0.1) / v4)))) * (float64((v1 / v1)) * v4))
        v14 = (v14 - -81)
        v1 = v1
        _ = v16
    }
    v4 -= (v4 + (-2.25 + (v4 / v4)))
    var v17 bool = (!false)
    var v18 uint64 = v2
    fmt.Println(v2, v1)
    var v19 uint16 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v14
    _ = v15
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
