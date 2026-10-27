package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8; F2 int8 }
func genFn0(p0 uint8, p1 uint64, p2 uint64) int {
    fmt.Println(p2, p1)
    p1 = (p1 - uint64(p0))
    _ = p0
    _ = p1
    _ = p2
    return 128
}

func main() {
    v1 := []uint32{541, 65535, 4294967295, 65535}
    v2 := make(map[string]int)
    v3 := make(map[string]string)
    _ = len(v2)
    var v4 float32 = 1e10
    fmt.Println(v4)
    v5 := v2[("" + "hello")]
    v4 = 1e10
    v2[""] = 9223372036854775807
    if true {
        v5 = (v5 % (((v5 % (((v5 & v5)) | 1))) | 1))
    } else {
        var v6 uint64 = 18446744073709551615
        var v7 int64 = 9223372036854775807
        var v8 int8 = -128
        _ = v6
        _ = v7
        _ = v8
    }
    for fi9 := 0; fi9 < 6; fi9++ {
        var v10 int64 = -1
        v10 -= v10
        _ = fi9
        _ = v10
    }
    v5 += (v5 % (((-(v5))) | 1))
    v11 := GenS1{F0: -1, F1: 255, F2: -128}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
