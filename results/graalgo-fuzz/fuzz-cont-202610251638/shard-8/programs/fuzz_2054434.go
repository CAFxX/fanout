package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64 }
type GenS2 struct { F0 int32; F1 uint; F2 int64 }
type GenS3 struct { F0 uint32 }
func genFn0(p0 uint, p1 int) int16 {
    v7 := GenS1{F0: (^(p1)), F1: 462}
    v8 := GenS1{F0: (p1 - 959), F1: 9223372036854775807}
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 839
}

func main() {
    var v1 int16 = -967
    v2 := make(map[uint32]uint32)
    var v3 uint8 = 2
    v4 := make(map[string]uint32)
    var v5 uint = 255
    v6 := []uint32{990}
    if len(v6) > 0 {
        v6[0] = 107
    }
    if (true || false) {
        var v9 uint16 = 956
        var v10 bool = false
        _ = v10
        _ = v9
    }
    var v11 uint32 = 65535
    v6 = append(v6, (^(v11)))
    for fi12 := 0; fi12 < 3; fi12++ {
        v5 += v5
        _ = fi12
    }
    v13 := GenS2{F0: -2147483648, F1: 1, F2: 0}
    v14 := GenS3{F0: (v11 % ((uint32(v5)) | 1))}
    var v15 uint = 392
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v11
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
