package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16; F2 uint }
func genFn0(p0 int64, p1 int64, p2 uint8) int32 {
    var v7 uint8 = p2
    v8 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return -1
}
func genFn1(p0 int16, p1 uint8) uint64 {
    p0 -= (p0 & p0)
    var v9 uint64 = 12
    _ = p0
    _ = p1
    _ = v9
    return v9
}

func main() {
    v1 := make(map[uint32]int)
    var v2 uint8 = 42
    var v3 int32 = 649
    v4 := make(map[string]bool)
    v5 := make(map[string]int)
    var v6 uint64 = 1
    v2 -= (v2 | 42)
    v6 -= (v6 & (v6 & (^(v6))))
    v10 := GenS1{F0: (1 >= 9223372036854775807), F1: 139, F2: 2}
    for fi11 := 0; fi11 < 12; fi11++ {
        v3 = (^(v3))
        var v12 uint32 = 1
        _ = fi11
        _ = v12
    }
    v2 -= (v2 & v2)
    var v13 uint16 = 0
    v14 := GenS1{F0: true, F1: 0, F2: 42}
    var v15 float32 = 1e10
    for fi16 := 0; fi16 < 3; fi16++ {
        var v17 uint16 = 65535
        v3 = v3
        _ = fi16
        _ = v17
    }
    v18 := GenS1{F0: (9223372036854775807 < 42), F1: int16(v2), F2: 262}
    for fi19 := 0; fi19 < 20; fi19++ {
        fmt.Println(v3, v2)
        v20 := v5["\n\t"]
        _ = fi19
        _ = v20
    }
    var v21 uint8 = 37
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v15
    _ = v18
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
