package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64 }
type GenS2 struct { F0 uint8; F1 int32; F2 int8 }
func genFn0(p0 uint32) int64 {
    p0 += p0
    var v5 uint32 = p0
    var v6 int16 = 1
    _ = p0
    _ = v5
    _ = v6
    return 42
}

func main() {
    var v1 int8 = 83
    v2 := make(map[uint32]int)
    v3 := []string{}
    v4 := make(map[uint32]string)
    var v7 uint32 = 42
    v4[v7] = "hello"
    v7 += 2147483648
    v8 := GenS1{F0: -1, F1: 368}
    var v9 uint8 = 1
    v1 = v1
    for fi10 := 0; fi10 < 14; fi10++ {
        v11 := GenS1{F0: -1, F1: 4294967295}
        _ = fi10
        _ = v11
    }
    v1 -= (v1 ^ (v1 & (v1 & v1)))
    for fi12 := 0; fi12 < 18; fi12++ {
        v13 := GenS2{F0: v9, F1: 1, F2: int8(fi12)}
        _ = fi12
        _ = v13
    }
    v7 -= 2
    v1 = 65
    switch 32767 {
    case 127:
        var v14 float32 = 1e10
        _ = v14
    case 4294967295:
        var v15 uint64 = 1
        _ = v15
    default:
        if (1 != 1) {
            var v16 int8 = v1
            _ = v16
        } else {
            v9 -= (^(v9))
            v7 = v7
        }
    }
    var v17 uint64 = 18446744073709551615
    fmt.Println(v7, v1, v9)
    var v18 int64 = int64(v1)
    fmt.Println(v17, v1)
    var v19 uint16 = 65535
    var v20 int32 = -1
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
