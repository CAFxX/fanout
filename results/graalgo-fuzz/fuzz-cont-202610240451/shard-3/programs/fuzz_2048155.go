package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int8; F1 int16 }
func genFn0(p0 uint64, p1 int, p2 int8) uint64 {
    switch (p1 & (p1 + (p1 * p1))) {
    case 2147483647:
        v4 := GenS1{F0: p0}
        _ = v4
    }
    v5 := []int{9223372036854775807, (-(p1))}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return (p0 & (p0 / ((1) | 1)))
}
func genFn1() string {
    for fi6 := 0; fi6 < 10; fi6++ {
        var v7 uint16 = 32767
        for fi8 := 0; fi8 < 12; fi8++ {
            v7 = v7
            fmt.Println(fi8, v7, fi6)
            _ = fi8
        }
        if (true || true) {
            v9 := []int{}
            _ = v9
        }
        _ = fi6
        _ = v7
    }
    var v10 int8 = 1
    var v11 int8 = v10
    _ = v10
    _ = v11
    return ""
}

func main() {
    var v1 bool = (128 <= 42)
    var v2 int8 = 32
    var v3 int8 = (v2 + (v2 << ((255) & 63)))
    fmt.Println(v2, v1, v2)
    for fi12 := 0; fi12 < 20; fi12++ {
        switch "" {
        case "foo bar":
            v1 = (!true)
        }
        v13 := []int{}
        v1 = v1
        _ = fi12
        _ = v13
    }
    v14 := GenS2{F0: v2, F1: 900}
    var v15 int64 = -1
    v16 := GenS2{F0: v2, F1: -72}
    v17 := []int{}
    v17 = append(v17, 32768)
    fmt.Println(v3)
    v15 -= v15
    var v18 float64 = 1.5
    fmt.Println(v2)
    var v19 int = 42
    var v20 int16 = -187
    v15 = (v15 * -9223372036854775808)
    v1 = v1
    var v21 int32 = 298
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
}
