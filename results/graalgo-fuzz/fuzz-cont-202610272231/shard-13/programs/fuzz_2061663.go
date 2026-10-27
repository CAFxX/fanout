package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64; F2 uint8 }
type GenS2 struct { F0 string }
func genFn0() int16 {
    var v4 bool = false
    var v5 uint64 = 225
    for fi6 := 0; fi6 < 12; fi6++ {
        var v7 uint64 = 1
        var v8 uint64 = v7
        _ = fi6
        _ = v7
        _ = v8
    }
    var v9 int32 = 421
    _ = v4
    _ = v5
    _ = v9
    return int16(v5)
}

func main() {
    v1 := []uint32{65535, 924}
    v2 := make(map[uint32]bool)
    var v3 uint8 = 1
    var v10 string = "foo bar"
    var v11 uint64 = 1
    v10 = "hello"
    v1 = append(v1, 222)
    _ = len(v1)
    v3 = (v3 >> ((42) & 63))
    var v12 uint64 = (v11 | (v11 & (v11 - v11)))
    v3 += v3
    var v13 int = 32768
    var v14 float64 = (1e10 - 0.1)
    fmt.Println(v3)
    v14 = ((3.14159 / 1e10) + 3.14159)
    switch (v13 << ((42) & 63)) {
    case 4294967295:
        v15 := GenS1{F0: (false && true), F1: 0, F2: v3}
        _ = v15
    case -823:
        for fi16 := 0; fi16 < 20; fi16++ {
            for fi17 := 0; fi17 < 19; fi17++ {
                v18 := GenS2{F0: v10}
                var v19 uint32 = uint32(v3)
                _ = fi17
                _ = v18
                _ = v19
            }
            _ = fi16
        }
    }
    _ = len(v1)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
}
