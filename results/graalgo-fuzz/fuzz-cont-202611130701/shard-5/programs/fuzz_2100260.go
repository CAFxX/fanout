package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8; F2 int16 }
type GenS2 struct { F0 int8; F1 uint32 }
type GenS3 struct { F0 int8; F1 int64; F2 int64 }
func genFn0(p0 uint64) uint {
    if true {
        v7 := GenS1{F0: 9223372036854775807, F1: 42, F2: 954}
        _ = v7
    }
    v8 := GenS2{F0: -128, F1: 42}
    v9 := GenS1{F0: 9223372036854775807, F1: 0, F2: 0}
    _ = p0
    _ = v8
    _ = v9
    return 42
}
func genFn1(p0 int32) uint {
    v10 := GenS1{F0: 0, F1: 100, F2: 42}
    p0 += (p0 | p0)
    _ = p0
    _ = v10
    return genFn0(42)
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint = 65535
    v3 := []string{"foo bar", "", "foo bar"}
    v4 := []uint32{}
    var v5 int64 = -1
    var v6 uint8 = 167
    v11 := v1["\n\t"]
    var v12 int8 = 100
    delete(v1, ("\n\t" + "\n\t"))
    v13 := GenS2{F0: v12, F1: 4294967295}
    v2 += (v2 % ((genFn0(2)) | 1))
    v14 := GenS3{F0: -1, F1: (v5 % ((v5) | 1)), F2: 0}
    v2 = (v2 << ((v2) & 63))
    var v15 bool = false
    v16 := GenS1{F0: (v5 - v5), F1: v12, F2: -838}
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v17 := GenS3{F0: (v12 + (v12 * 42)), F1: 9223372036854775807, F2: v5}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
