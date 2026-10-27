package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint }
type GenS2 struct { F0 int32 }

func main() {
    var v1 uint8 = 255
    var v2 int8 = -1
    var v3 uint16 = 255
    v4 := make(map[string]bool)
    if false {
        v3 -= (v3 & (^(v3)))
    } else {
        var v5 int32 = -261
        _ = v5
    }
    v6 := []int{32767, -1, 9223372036854775807}
    v2 -= v2
    for fi7 := 0; fi7 < 2; fi7++ {
        var v8 bool = false
        var v9 int32 = 0
        _ = fi7
        _ = v8
        _ = v9
    }
    var v10 string = ""
    var v11 int8 = v2
    for fi12 := 0; fi12 < 4; fi12++ {
        v13 := GenS1{F0: (-2147483648 <= 538), F1: 415}
        switch (fi12 - fi12) {
        case -1:
            if len(v6) > 0 {
                v6[0] = int(v3)
            }
        default:
            v14 := GenS2{F0: 0}
            _ = v14
        }
        _ = fi12
        _ = v13
    }
    v1 += (v1 * 1)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v6
}
