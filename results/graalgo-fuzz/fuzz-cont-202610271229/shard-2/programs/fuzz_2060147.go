package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int }
type GenS2 struct { F0 uint }

func main() {
    var v1 uint32 = 747
    var v2 float64 = (0.0 * (((-2.25 - 1.5) / float64(1e10)) - (0.0 + ((1.5 + ((-2.25 + 3.14159) * float64(-2.25))) / 1.0))))
    var v3 int16 = -1
    v4 := []int{}
    v5 := make(map[string]uint32)
    if false {
        v6 := GenS1{F0: 2, F1: -1}
        v1 = v1
        _ = v6
    }
    switch "\n\t" {
    case "foo bar":
        v1 += v1
    case "":
        v1 = (v1 % ((65535) | 1))
    }
    for fi7 := 0; fi7 < 6; fi7++ {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = fi7
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    for fi8 := 0; fi8 < 13; fi8++ {
        v9 := GenS1{F0: 139, F1: (fi8 + fi8)}
        _ = fi8
        _ = v9
    }
    fmt.Println(v1)
    v10 := GenS2{F0: 18446744073709551615}
    var v11 int32 = 633
    if len(v4) > 0 {
        v4[0] = 465
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
