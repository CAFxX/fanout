package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8 }
type GenS2 struct { F0 uint64 }

func main() {
    v1 := make(map[string]uint32)
    var v2 uint64 = 42
    var v3 uint16 = 255
    v4 := make(map[int]int)
    var v5 int8 = 42
    var v6 uint64 = v2
    fmt.Println(v2, v6)
    v7 := GenS1{F0: 801, F1: 1}
    v8 := GenS2{F0: 0}
    for fi9 := 0; fi9 < 7; fi9++ {
        fmt.Println(v2)
        _ = fi9
    }
    v4[0] = 128
    if (!(!(v2 <= 42))) {
        v10 := []int{127, int(v6), -9223372036854775808}
        _ = v10
    } else {
        v11 := []int{1, -415, -502, 2147483647}
        v1[""] = 201
        _ = v11
    }
    var v12 int = 32768
    v6 += v2
    fmt.Println(v12, v6)
    if (v5 <= v5) {
        v13 := GenS1{F0: (v3 + v3), F1: (v5 & (v5 % (((v5 - (v5 * (v5 << ((482) & 63))))) | 1)))}
        var v14 int32 = 905
        v2 -= (v6 << ((18446744073709551615) & 63))
        _ = v13
        _ = v14
    } else {
        delete(v1, "foo bar")
        var v15 bool = false
        var v16 uint32 = 2147483648
        _ = v15
        _ = v16
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
