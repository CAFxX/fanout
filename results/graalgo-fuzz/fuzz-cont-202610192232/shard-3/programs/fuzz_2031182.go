package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64; F2 uint16 }
type GenS2 struct { F0 int32; F1 int64; F2 uint8 }
type GenS3 struct { F0 uint16; F1 int16 }
type GenS4 struct { F0 uint32; F1 uint8 }
func genFn0(p0 int8, p1 string, p2 uint8) uint32 {
    v6 := []int{-9223372036854775808, -636}
    _ = len(v6)
    v7 := GenS1{F0: p1, F1: 636, F2: 0}
    v8 := GenS1{F0: p1, F1: 565, F2: 545}
    v9 := GenS2{F0: 1, F1: -580, F2: p2}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 0
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[uint32]uint32)
    v3 := make(map[uint32]int)
    v4 := []string{"a", ("" + "a"), "hello"}
    var v5 uint64 = 42
    for fi10 := 0; fi10 < 2; fi10++ {
        var v11 int = 42
        fmt.Println(v11, fi10)
        var v12 int16 = 498
        _ = fi10
        _ = v11
        _ = v12
    }
    if (120 > 42) {
        for fi13 := 0; fi13 < 13; fi13++ {
            v5 = (v5 / ((v5) | 1))
            v5 += 4294967295
            v5 = 706
            _ = fi13
        }
    }
    for fi14 := 0; fi14 < 13; fi14++ {
        var v15 string = ("\n\t" + "a")
        var v16 uint64 = v5
        _ = fi14
        _ = v15
        _ = v16
    }
    fmt.Println(v5)
    v4 = append(v4, "a")
    if (892 <= 1) {
        switch -154 {
        case 9223372036854775807:
            var v17 int8 = 1
            _ = v17
        default:
            v18 := GenS3{F0: 1, F1: -32768}
            _ = v18
        }
        v19 := v1[42]
        _ = v19
    }
    v5 = 42
    for fi20 := 0; fi20 < 5; fi20++ {
        var v21 uint16 = 1
        _ = fi20
        _ = v21
    }
    var v22 uint64 = 18446744073709551615
    v22 += v5
    for fi23 := 0; fi23 < 7; fi23++ {
        v24 := GenS4{F0: 85, F1: 42}
        _ = fi23
        _ = v24
    }
    var v25 uint32 = 362
    v26 := v1[2147483647]
    for fi27 := 0; fi27 < 13; fi27++ {
        var v28 int8 = 46
        var v29 string = v26
        _ = fi27
        _ = v28
        _ = v29
    }
    var v30 bool = (!true)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v25
    _ = v26
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
