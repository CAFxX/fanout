package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16 }
type GenS2 struct { F0 string }
func genFn0(p0 uint8, p1 int16, p2 uint16) uint8 {
    p1 += p1
    fmt.Println(p1, p0)
    _ = p0
    _ = p1
    _ = p2
    return 0
}
func genFn1(p0 bool, p1 uint32) uint32 {
    v4 := GenS1{F0: -1, F1: 32767}
    var v5 int8 = -1
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return p1
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[int]string)
    v3 := make(map[int]int)
    var v6 int32 = -39
    if (true || (1 >= 42)) {
        var v7 uint16 = 165
        fmt.Println(v6)
        var v8 int16 = -657
        _ = v7
        _ = v8
    } else {
        v6 -= v6
        if true {
            v9 := GenS1{F0: -9223372036854775808, F1: 1}
            _ = v9
        }
        v6 = (v6 | v6)
    }
    v10 := GenS2{F0: (("a" + "\n\t") + ("foo bar" + ("\n\t" + "a")))}
    v6 -= (v6 + 981)
    v6 = v6
    var v11 uint16 = 230
    for fi12 := 0; fi12 < 1; fi12++ {
        var v13 int8 = -34
        var v14 string = ""
        _ = fi12
        _ = v13
        _ = v14
    }
    v6 -= (v6 ^ 2147483647)
    if (!false) {
        v6 += -1
    }
    v15 := GenS1{F0: -1, F1: v11}
    v6 = 2147483647
    v11 = (v11 << ((0) & 63))
    var v16 float32 = -2.25
    v11 += (v11 * 424)
    v6 -= v6
    var v17 int64 = 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v6
}
