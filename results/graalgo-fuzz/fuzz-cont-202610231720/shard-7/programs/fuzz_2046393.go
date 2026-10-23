package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64 }
type GenS2 struct { F0 int8; F1 int64; F2 int32 }

func main() {
    var v1 uint = 0
    var v2 int64 = 1
    var v3 int = 32768
    v4 := make(map[uint32]int)
    v2 -= (v2 - -1)
    v5 := []int{(v3 % ((658) | 1)), (^(v3)), (^(v3)), v3}
    var v6 string = "\n\t"
    for fi7 := 0; fi7 < 12; fi7++ {
        for fi8 := 0; fi8 < 12; fi8++ {
            var v9 uint32 = 2
            v10 := GenS1{F0: 726, F1: v2}
            _ = fi8
            _ = v10
            _ = v9
        }
        _ = fi7
    }
    fmt.Println(v3, v6)
    switch ((v6 + v6) + (v6 + v6)) {
    case "":
        v11 := GenS1{F0: 9223372036854775807, F1: 767}
        _ = v11
    case "\n\t":
        for fi12 := 0; fi12 < 15; fi12++ {
            var v13 int32 = -591
            var v14 uint64 = 217
            _ = fi12
            _ = v13
            _ = v14
        }
    default:
        v15 := GenS1{F0: v2, F1: v2}
        _ = v15
    }
    for fi16 := 0; fi16 < 17; fi16++ {
        var v17 uint32 = 2147483647
        var v18 uint32 = 2
        _ = fi16
        _ = v17
        _ = v18
    }
    v6 = ("hello" + v6)
    v2 = (v2 | (v2 | v2))
    if true {
        var v19 int8 = 79
        _ = len(v5)
        _ = v19
    }
    for fi20 := 0; fi20 < 7; fi20++ {
        v6 = v6
        var v21 uint = v1
        for fi22 := 0; fi22 < 9; fi22++ {
            v21 -= v21
            v23 := GenS1{F0: 0, F1: v2}
            v24 := GenS1{F0: (^(v2)), F1: (v2 ^ v2)}
            v25 := GenS2{F0: 96, F1: v2, F2: 0}
            _ = fi22
            _ = v23
            _ = v24
            _ = v25
        }
        _ = fi20
        _ = v21
    }
    if false {
        var v26 uint16 = 1
        v26 = 42
        _ = len(v5)
        _ = v26
    }
    fmt.Println(v1, v2, v3)
    switch (v3 * v3) {
    case -410:
        var v27 uint32 = 0
        _ = v27
    case -9223372036854775808:
        var v28 string = ("foo bar" + "foo bar")
        _ = v28
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
