package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16; F2 int }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int16; F1 int64 }
func genFn0(p0 int16, p1 int64, p2 uint8) int8 {
    var v4 int = 127
    for fi5 := 0; fi5 < 1; fi5++ {
        v6 := GenS1{F0: 42, F1: p0, F2: fi5}
        var v7 int8 = 127
        _ = fi5
        _ = v6
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 100
}

func main() {
    var v1 int8 = 1
    v2 := make(map[string]string)
    v3 := []string{}
    fmt.Println(v1)
    for fi8 := 0; fi8 < 12; fi8++ {
        var v9 uint = 42
        for fi10 := 0; fi10 < 9; fi10++ {
            _ = len(v3)
            _ = fi10
        }
        _ = fi8
        _ = v9
    }
    var v11 int16 = 408
    v12 := GenS2{F0: 0}
    v11 += (v11 << ((532) & 63))
    v13 := GenS3{F0: (v11 * 0), F1: -9223372036854775808}
    v11 += (-(v11))
    v1 -= (v1 % (((v1 | 42)) | 1))
    fmt.Println(v11, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
}
