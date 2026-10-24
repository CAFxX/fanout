package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64 }
type GenS2 struct { F0 bool; F1 uint; F2 string }

func main() {
    var v1 float64 = -2.25
    var v2 uint16 = 835
    v3 := []int{128, int(v2), 392, 32768}
    v4 := GenS1{F0: 0, F1: 9223372036854775807}
    v2 += 0
    v5 := GenS2{F0: (42 == 42), F1: 0, F2: ("\n\t" + "foo bar")}
    v3 = append(v3, 530)
    v1 -= -2.25
    v1 += v1
    fmt.Println(v2, v1)
    for fi6 := 0; fi6 < 15; fi6++ {
        v7 := GenS1{F0: -128, F1: 0}
        _ = fi6
        _ = v7
    }
    for fi8 := 0; fi8 < 18; fi8++ {
        v3 = append(v3, (fi8 / ((fi8) | 1)))
        _ = fi8
    }
    var v9 float64 = 0.1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
