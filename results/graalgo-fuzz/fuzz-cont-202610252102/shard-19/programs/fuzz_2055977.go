package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint64; F1 int8 }
func genFn0(p0 uint8, p1 uint64) int {
    v5 := GenS1{F0: 255}
    switch -347 {
    case 238:
        p1 = 4294967295
    case 704:
        v6 := []int{-581, -132, 42, 4294967295}
        _ = v6
    default:
        for fi7 := 0; fi7 < 16; fi7++ {
            var v8 uint8 = (^(p0))
            var v9 uint16 = 32767
            _ = fi7
            _ = v8
            _ = v9
        }
    }
    v10 := GenS2{F0: 32767}
    var v11 int64 = -1
    var v12 int8 = 1
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v12
    _ = v5
    return 9223372036854775807
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[int]uint32)
    var v3 int32 = 0
    v4 := make(map[int]string)
    v13 := []int{-1, -882}
    for fi14 := 0; fi14 < 19; fi14++ {
        var v15 uint64 = 18446744073709551615
        _ = fi14
        _ = v15
    }
    v3 -= 2147483647
    v4[32768] = (("" + "") + "a")
    for fi16 := 0; fi16 < 6; fi16++ {
        v17 := GenS3{F0: 2147483648}
        _ = fi16
        _ = v17
    }
    for fi18 := 0; fi18 < 10; fi18++ {
        v3 = -181
        var v19 uint32 = 1
        _ = fi18
        _ = v19
    }
    v20 := GenS3{F0: 181}
    v3 = 2147483647
    v21 := GenS4{F0: 861, F1: 52}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v13
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
}
