package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32; F2 uint8 }

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[uint32]uint32)
    var v3 int = -992
    var v4 uint16 = 335
    var v5 uint64 = 1
    v6 := make(map[string]int)
    v7 := GenS1{F0: (v5 - 544), F1: 2147483647, F2: 1}
    if ((false || true) && false) {
        v5 -= (v5 & uint64(v4))
        delete(v6, ("\n\t" + "foo bar"))
    }
    v8 := []int{(v3 << ((42) & 63)), 42, v3}
    v5 = v5
    if ((!true) || (1.5 > 0.0)) {
        v3 -= (v3 * (-(v3)))
    } else {
        var v9 float32 = 0.0
        _ = v9
    }
    for fi10 := 0; fi10 < 1; fi10++ {
        if true {
            v5 = 18446744073709551615
        }
        _ = fi10
    }
    v3 += (v3 % (((v3 + (v3 & (v3 % ((42) | 1))))) | 1))
    switch (v3 << ((0) & 63)) {
    case 127:
        v3 = (v3 << ((4294967295) & 63))
    case -845:
        switch 570 {
        case -1:
            v11 := v6["foo bar"]
            _ = v11
        case 4294967295:
            v3 = v3
        case -2147483648:
            v4 += v4
        }
    default:
        for fi12 := 0; fi12 < 8; fi12++ {
            v5 += 18446744073709551615
            v4 = (v4 | v4)
            _ = fi12
        }
    }
    var v13 int8 = -1
    var v14 float64 = -2.25
    var v15 uint16 = v4
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
