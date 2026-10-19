package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64 }
func genFn0() int16 {
    var v6 uint16 = 1
    v6 += (v6 - (v6 & (v6 << ((255) & 63))))
    v7 := []int{32768, -518}
    _ = v6
    _ = v7
    return 1
}

func main() {
    var v1 int = -240
    var v2 uint8 = 128
    v3 := make(map[string]bool)
    var v4 int16 = -32768
    v5 := make(map[int]string)
    v8 := GenS1{F0: 489, F1: 42}
    var v9 float32 = -2.25
    if false {
        v1 = 128
    }
    fmt.Println(v2, v4)
    if (100 != 1) {
        v2 = v2
    }
    switch 1 {
    case 557:
        var v10 int64 = 0
        _ = v10
    default:
        v11 := GenS1{F0: 18446744073709551615, F1: -278}
        _ = v11
    }
    v1 -= v1
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
