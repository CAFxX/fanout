package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16 }
func genFn0() int64 {
    var v5 int = 1
    v5 -= v5
    v6 := []int{(v5 % ((v5) | 1)), v5, (v5 | v5), v5}
    _ = v5
    _ = v6
    return -1
}
func genFn1(p0 int16, p1 int64, p2 int64) uint8 {
    v7 := GenS1{F0: -73, F1: 32767}
    p1 = p1
    v8 := GenS1{F0: 127, F1: 42}
    v9 := GenS1{F0: 100, F1: 32767}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return 47
}

func main() {
    v1 := make(map[string]bool)
    var v2 int8 = 127
    var v3 uint16 = 32767
    var v4 uint8 = 129
    for fi10 := 0; fi10 < 17; fi10++ {
        v4 = (v4 ^ (v4 << ((0) & 63)))
        switch -1 {
        case 1:
            var v11 int8 = -128
            _ = v11
        case -480:
            if (42 <= 484) {
                var v12 uint32 = 511
                _ = v12
            }
        }
        _ = len(v1)
        _ = fi10
    }
    v13 := GenS1{F0: -54, F1: (v3 % (((v3 % ((255) | 1))) | 1))}
    fmt.Println(v4)
    v2 += -107
    v14 := []int{32768, 1}
    v15 := v1[("a" + "a")]
    var v16 uint16 = v3
    v17 := v1[""]
    var v18 int8 = -1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
