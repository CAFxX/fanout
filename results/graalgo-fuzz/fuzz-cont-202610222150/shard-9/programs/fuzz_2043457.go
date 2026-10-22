package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32; F2 uint16 }
type GenS2 struct { F0 string; F1 uint16; F2 string }
func genFn0(p0 int64, p1 uint32, p2 int16) int8 {
    var v7 int64 = (p0 + p0)
    p0 = (p0 >> ((42) & 63))
    p1 = p1
    var v8 uint = 4294967295
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 42
}

func main() {
    var v1 uint = 2
    v2 := make(map[string]uint32)
    var v3 string = "foo bar"
    var v4 bool = (!((^(v1)) <= 1))
    v5 := []uint32{582, 42, 4294967295}
    v6 := make(map[int]bool)
    var v9 uint32 = 2147483647
    _ = len(v6)
    v2[""] = (v9 >> (((v1 * v1)) & 63))
    v3 = "hello"
    var v10 uint = (^(v1))
    var v11 int = 42
    var v12 int32 = -1
    for fi13 := 0; fi13 < 18; fi13++ {
        var v14 float32 = 1.5
        v15 := GenS1{F0: (v9 + (v9 | v9)), F1: (v12 / ((v12) | 1)), F2: 2}
        _ = fi13
        _ = v14
        _ = v15
    }
    var v16 uint64 = 663
    fmt.Println(v1, v16, v1)
    for fi17 := 0; fi17 < 12; fi17++ {
        fmt.Println(v16, v1, v12)
        var v18 int = 1
        _ = fi17
        _ = v18
    }
    var v19 uint64 = v16
    v20 := GenS1{F0: (^(v9)), F1: (v12 / (((v12 * (v12 - (v12 - v12)))) | 1)), F2: 32767}
    v21 := GenS2{F0: v3, F1: 2, F2: (v3 + "a")}
    fmt.Println(v11)
    var v22 uint32 = (^(v9))
    var v23 uint64 = v16
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v16
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
