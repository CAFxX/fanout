package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32; F2 int64 }

func main() {
    var v1 int16 = 621
    v2 := make(map[uint32]bool)
    var v3 uint8 = 0
    for fi4 := 0; fi4 < 9; fi4++ {
        var v5 int32 = -580
        _ = fi4
        _ = v5
    }
    var v6 float64 = -2.25
    fmt.Println(v3, v1)
    v3 = v3
    v7 := GenS1{F0: 100, F1: 2147483647, F2: 16}
    v1 = v1
    v3 -= 255
    var v8 bool = (!false)
    v9 := GenS1{F0: 127, F1: -104, F2: -9223372036854775808}
    switch 2147483647 {
    case -909:
        v10 := []int{-9223372036854775808}
        _ = v10
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
