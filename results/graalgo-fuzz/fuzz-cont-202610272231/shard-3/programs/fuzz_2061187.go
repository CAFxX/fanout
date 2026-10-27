package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32; F2 string }

func main() {
    v1 := make(map[int]int)
    var v2 uint8 = 251
    var v3 uint16 = 105
    var v4 uint16 = (v3 & v3)
    var v5 string = "\n\t"
    _ = len(v1)
    fmt.Println(v4, v2)
    var v6 int64 = -1
    for fi7 := 0; fi7 < 17; fi7++ {
        v5 = ((v5 + v5) + (v5 + ((v5 + v5) + (("a" + v5) + v5))))
        v8 := []int{}
        _ = fi7
        _ = v8
    }
    v9 := GenS1{F0: 2147483647, F1: 740, F2: v5}
    switch 42 {
    case 4294967295:
        var v10 int64 = (v6 & 34)
        _ = v10
    case -587:
        v3 = (v3 - (v3 / ((v3) | 1)))
    }
    for fi11 := 0; fi11 < 8; fi11++ {
        delete(v1, (fi11 % (((-(fi11))) | 1)))
        var v12 int8 = 24
        _ = fi11
        _ = v12
    }
    v4 += v3
    v6 += -9223372036854775808
    v13 := []int{380, 743, 128, 127}
    v6 -= (v6 / (((-(v6))) | 1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
