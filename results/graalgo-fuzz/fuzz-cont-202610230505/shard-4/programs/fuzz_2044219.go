package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8 }
type GenS2 struct { F0 int8; F1 int8; F2 uint64 }
type GenS3 struct { F0 int64; F1 int; F2 int16 }
func genFn0(p0 bool) string {
    p0 = p0
    if true {
        p0 = (!false)
    } else {
        p0 = (1 == 0)
    }
    fmt.Println(p0)
    p0 = p0
    _ = p0
    return ("hello" + ("a" + ""))
}

func main() {
    v1 := make(map[string]uint32)
    var v2 int32 = 2147483647
    var v3 int64 = 91
    var v4 bool = false
    var v5 uint8 = 1
    v6 := GenS1{F0: 32767, F1: 127}
    fmt.Println(v2, v3)
    v5 += v5
    v7 := []int{163, 32767}
    v3 -= (v3 << ((367) & 63))
    v8 := GenS2{F0: 42, F1: 0, F2: 496}
    v9 := GenS3{F0: v3, F1: 32768, F2: 1}
    for fi10 := 0; fi10 < 7; fi10++ {
        var v11 int8 = 127
        v12 := GenS1{F0: fi10, F1: v11}
        fmt.Println(v5, v5, fi10)
        _ = fi10
        _ = v11
        _ = v12
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
