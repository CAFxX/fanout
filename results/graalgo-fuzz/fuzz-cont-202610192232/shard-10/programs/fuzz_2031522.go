package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8; F2 uint }
type GenS2 struct { F0 int64; F1 uint16 }
func genFn0(p0 bool, p1 int8) uint8 {
    v6 := GenS1{F0: 4294967295, F1: 1, F2: 769}
    if p0 {
        v7 := GenS2{F0: -818, F1: 65535}
        _ = v7
    }
    _ = p0
    _ = p1
    _ = v6
    return 160
}

func main() {
    var v1 int64 = -9223372036854775808
    v2 := make(map[int]uint32)
    var v3 string = ("" + "\n\t")
    var v4 int = 2147483647
    var v5 int = -599
    v3 = v3
    if (2147483647 < 413) {
        v8 := GenS1{F0: 50, F1: 26, F2: 18446744073709551615}
        _ = v8
    }
    for fi9 := 0; fi9 < 8; fi9++ {
        fmt.Println(v4, v5)
        if true {
            v1 += v1
            for fi10 := 0; fi10 < 9; fi10++ {
                var v11 float32 = 1.5
                delete(v2, 9223372036854775807)
                _ = fi10
                _ = v11
            }
        }
        _ = fi9
    }
    v12 := []int{v5}
    v13 := GenS1{F0: 195, F1: 63, F2: 2}
    var v14 uint64 = 868
    v5 -= (v5 * -9223372036854775808)
    v15 := GenS2{F0: (^(v1)), F1: 360}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
