package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool; F2 uint32 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int16; F1 int8; F2 int64 }
func genFn0(p0 uint64) uint64 {
    v5 := GenS1{F0: 508, F1: false, F2: 4294967295}
    for fi6 := 0; fi6 < 2; fi6++ {
        for fi7 := 0; fi7 < 4; fi7++ {
            v8 := GenS2{F0: 100}
            fmt.Println(fi7)
            _ = fi7
            _ = v8
        }
        if (!false) {
            fmt.Println(fi6)
        }
        _ = fi6
    }
    _ = p0
    _ = v5
    return p0
}

func main() {
    v1 := []string{"a", "a", ("\n\t" + "a"), "foo bar"}
    v2 := make(map[int]int)
    v3 := make(map[string]string)
    var v4 uint = 817
    v9 := GenS3{F0: 1000, F1: 1, F2: -9223372036854775808}
    v10 := GenS2{F0: 42}
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if (-638 > 440) {
        var v11 uint = (^(v4))
        _ = v11
    }
    v4 -= (v4 * v4)
    fmt.Println(v4)
    var v12 uint32 = 42
    v12 = (v12 >> (((^(v4))) & 63))
    for fi13 := 0; fi13 < 18; fi13++ {
        v4 = v4
        var v14 uint32 = v12
        _ = fi13
        _ = v14
    }
    var v15 uint16 = 862
    var v16 int32 = -32
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v12
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
