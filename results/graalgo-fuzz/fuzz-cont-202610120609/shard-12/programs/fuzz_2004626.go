package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64; F2 uint16 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint16; F1 int8; F2 uint64 }
func genFn0(p0 int32, p1 uint16, p2 int32) uint {
    fmt.Println(p0)
    for fi4 := 0; fi4 < 5; fi4++ {
        var v5 int64 = 9223372036854775807
        _ = fi4
        _ = v5
    }
    _ = p0
    _ = p1
    _ = p2
    return uint(p1)
}
func genFn1(p0 int8) string {
    fmt.Println(p0)
    v6 := GenS1{F0: 0, F1: 9223372036854775807, F2: 2}
    _ = p0
    _ = v6
    return "\n\t"
}

func main() {
    v1 := make(map[int]int)
    var v2 uint = 2
    var v3 int64 = 0
    fmt.Println(v3, v2)
    v7 := GenS1{F0: 32767, F1: (v3 << ((v2) & 63)), F2: 1}
    for fi8 := 0; fi8 < 6; fi8++ {
        var v9 uint64 = 99
        v3 = (v3 / (((v3 | (v3 << ((v2) & 63)))) | 1))
        if true {
            var v10 uint64 = 346
            _ = v10
        } else {
            var v11 uint32 = 0
            fmt.Println(v2)
            v12 := GenS2{F0: 2}
            _ = v11
            _ = v12
        }
        _ = fi8
        _ = v9
    }
    for fi13 := 0; fi13 < 13; fi13++ {
        var v14 bool = true
        _ = fi13
        _ = v14
    }
    v15 := GenS3{F0: -116}
    var v16 int64 = v3
    v3 = -1
    v17 := GenS4{F0: 255, F1: 0, F2: 0}
    for fi18 := 0; fi18 < 12; fi18++ {
        for fi19 := 0; fi19 < 19; fi19++ {
            v3 = 0
            var v20 uint32 = 148
            _ = fi19
            _ = v20
        }
        fmt.Println(v2)
        _ = fi18
    }
    var v21 uint32 = 877
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v21
    _ = v3
    _ = v7
}
