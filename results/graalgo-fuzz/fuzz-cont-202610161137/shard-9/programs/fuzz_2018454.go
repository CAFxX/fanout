package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8 }
type GenS2 struct { F0 uint16 }

func main() {
    var v1 int32 = -2147483648
    v2 := make(map[int]bool)
    v3 := make(map[int]string)
    var v4 string = ((("" + "foo bar") + "") + "a")
    var v5 int32 = v1
    fmt.Println(v4, v4, v4)
    for fi6 := 0; fi6 < 17; fi6++ {
        v7 := GenS1{F0: -380, F1: 0}
        _ = fi6
        _ = v7
    }
    v4 = v4
    v5 -= 42
    v1 += v1
    v4 = v4
    v5 = (v5 / (((^(v1))) | 1))
    v1 -= (v1 << ((0) & 63))
    v8 := GenS2{F0: 2}
    for fi9 := 0; fi9 < 17; fi9++ {
        var v10 int16 = 0
        _ = fi9
        _ = v10
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
