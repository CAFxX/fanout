package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint64 }
func genFn0(p0 int16) uint8 {
    p0 = 1000
    for fi4 := 0; fi4 < 13; fi4++ {
        var v5 int16 = 930
        v5 = (p0 * (p0 ^ p0))
        v6 := GenS1{F0: 4294967295}
        _ = fi4
        _ = v5
        _ = v6
    }
    fmt.Println(p0)
    _ = p0
    return 2
}

func main() {
    var v1 int16 = 1
    var v2 int64 = 0
    v3 := make(map[string]bool)
    v7 := GenS1{F0: 457}
    var v8 bool = false
    for fi9 := 0; fi9 < 19; fi9++ {
        v10 := []int{(fi9 & fi9), -225, fi9, fi9}
        v11 := GenS2{F0: 18446744073709551615}
        var v12 uint32 = 4294967295
        _ = fi9
        _ = v10
        _ = v11
        _ = v12
    }
    var v13 int16 = (v1 & v1)
    var v14 uint = 786
    v2 -= v2
    v15 := GenS1{F0: 2147483648}
    if ((2 < 1) && (1.5 <= 1.5)) {
        fmt.Println(v13, v8, v13)
        if v8 {
            v2 = v2
        }
        v8 = (0.1 > (1e10 - (1e10 - 0.0)))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v7
    _ = v8
}
