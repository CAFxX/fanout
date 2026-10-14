package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int8; F1 uint32; F2 bool }
func genFn0(p0 int32, p1 int, p2 int) int64 {
    switch (p1 * p2) {
    case 32767:
        fmt.Println(p2, p0, p2)
    default:
        v6 := GenS1{F0: p2, F1: 706}
        _ = v6
    }
    var v7 uint8 = 127
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 42
}

func main() {
    v1 := make(map[int]int)
    v2 := []uint32{65535, 1, 65535, 882}
    v3 := []int{2147483647, 32767, 1}
    var v4 int32 = 551
    v5 := []string{}
    v8 := GenS2{F0: 268}
    var v9 int32 = 250
    fmt.Println(v4, v9)
    var v10 uint8 = 48
    var v11 uint16 = 47
    v9 -= (v9 & (v4 - v9))
    for fi12 := 0; fi12 < 19; fi12++ {
        if len(v5) > 0 {
            v5[0] = "hello"
        }
        v13 := GenS3{F0: 127, F1: 725, F2: true}
        v9 += v4
        _ = fi12
        _ = v13
    }
    v4 += (v9 | v4)
    var v14 uint32 = 4294967295
    var v15 int = 127
    for fi16 := 0; fi16 < 13; fi16++ {
        fmt.Println(v11, v11)
        v17 := GenS1{F0: 42, F1: v14}
        _ = fi16
        _ = v17
    }
    var v18 bool = (uint32(v11) != uint32(v10))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v15
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
