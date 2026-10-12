package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint }
type GenS2 struct { F0 uint; F1 uint32; F2 uint32 }
func genFn0(p0 uint8, p1 uint64) uint {
    var v7 string = "a"
    v8 := []int{}
    var v9 uint16 = 409
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return 42
}
func genFn1(p0 bool, p1 uint16) uint {
    var v10 uint64 = 2
    var v11 bool = true
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    return 42
}

func main() {
    var v1 int64 = 1
    var v2 uint32 = 65535
    v3 := make(map[uint32]string)
    v4 := make(map[string]int)
    var v5 string = "foo bar"
    var v6 int = 696
    if (!false) {
        for fi12 := 0; fi12 < 18; fi12++ {
            v6 += v6
            _ = len(v3)
            var v13 string = v5
            _ = fi12
            _ = v13
        }
        var v14 int = (v6 ^ v6)
        _ = v14
    } else {
        for fi15 := 0; fi15 < 6; fi15++ {
            v6 = int(v1)
            _ = fi15
        }
        v16 := []int{(-(v6)), v6, v6}
        _ = v16
    }
    var v17 int = (v6 * (v6 << ((4294967295) & 63)))
    delete(v4, "hello")
    v18 := GenS1{F0: -902, F1: 42}
    v5 = ("hello" + v5)
    for fi19 := 0; fi19 < 14; fi19++ {
        var v20 uint32 = v2
        v21 := GenS1{F0: v1, F1: 556}
        _ = fi19
        _ = v20
        _ = v21
    }
    var v22 int8 = 42
    v1 = -576
    if (int32(v22) >= -1) {
        v23 := GenS2{F0: 4294967295, F1: v2, F2: (v2 | 2)}
        _ = v23
    } else {
        for fi24 := 0; fi24 < 12; fi24++ {
            fmt.Println(v2, v2, v6)
            _ = fi24
        }
        v25 := GenS1{F0: v1, F1: 2}
        _ = v25
    }
    fmt.Println(v22, v1, v22)
    var v26 int32 = int32(v1)
    v17 = (v17 & v17)
    v1 += v1
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v22
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
