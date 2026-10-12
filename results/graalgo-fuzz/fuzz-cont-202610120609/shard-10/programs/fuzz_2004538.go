package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 string; F1 uint64; F2 int8 }
func genFn0(p0 int, p1 int32) uint16 {
    p1 = 792
    v6 := []int{(p0 & (-(p0))), p0, 326}
    var v7 uint64 = 2
    if ((!true) && (true && true)) {
        v7 = (v7 >> ((0) & 63))
        v8 := GenS1{F0: (p1 - (p1 % ((p1) | 1)))}
        v7 += v7
        _ = v8
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 2
}

func main() {
    var v1 int16 = 42
    v2 := make(map[int]uint32)
    v3 := []string{("hello" + "hello")}
    v4 := make(map[uint32]uint32)
    var v5 bool = (true || false)
    v3 = append(v3, "")
    switch "hello" {
    case "\n\t":
        var v9 float32 = -2.25
        _ = v9
    }
    v10 := GenS2{F0: "a", F1: 1, F2: 100}
    var v11 string = "foo bar"
    v2[-517] = 4294967295
    if len(v3) > 0 {
        v3[0] = v11
    }
    v5 = false
    v2[509] = 817
    for fi12 := 0; fi12 < 18; fi12++ {
        if true {
            for fi13 := 0; fi13 < 15; fi13++ {
                v1 += (v1 % ((v1) | 1))
                _ = fi13
            }
        }
        for fi14 := 0; fi14 < 18; fi14++ {
            v15 := GenS2{F0: v11, F1: 4294967295, F2: -1}
            _ = fi14
            _ = v15
        }
        _ = fi12
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
