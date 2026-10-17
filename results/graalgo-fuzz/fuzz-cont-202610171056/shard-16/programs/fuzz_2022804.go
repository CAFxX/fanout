package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }

func main() {
    var v1 float64 = (float64(-2.25) + 3.14159)
    v2 := make(map[int]int)
    v3 := []int{-1, -2147483648, 128, 9223372036854775807}
    v1 += 1e10
    fmt.Println(v1)
    fmt.Println(v1)
    if false {
        var v4 bool = (-193 == 42)
        var v5 uint32 = 65535
        _ = v4
        _ = v5
    }
    v1 = v1
    if false {
        var v6 uint8 = 127
        _ = v6
    } else {
        v1 = v1
    }
    v7 := GenS1{F0: 138}
    v1 = v1
    var v8 uint16 = 2
    for fi9 := 0; fi9 < 12; fi9++ {
        v8 -= (v8 ^ (v8 >> ((4294967295) & 63)))
        _ = fi9
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v7
    _ = v8
}
