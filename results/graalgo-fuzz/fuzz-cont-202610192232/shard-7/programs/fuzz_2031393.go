package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int8; F1 int64 }
func genFn0(p0 uint16) bool {
    p0 = p0
    var v5 float32 = 3.14159
    fmt.Println(p0, v5)
    _ = p0
    _ = v5
    return false
}
func genFn1(p0 uint8, p1 int64) int16 {
    var v6 uint32 = 65535
    var v7 uint8 = p0
    fmt.Println(p0, v6, p0)
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return -495
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[uint32]bool)
    var v3 int = -788
    var v4 uint16 = 255
    var v8 uint64 = 1
    v2[42] = false
    v1[42] = 865
    _ = len(v1)
    for fi9 := 0; fi9 < 17; fi9++ {
        v10 := GenS1{F0: -1}
        if true {
            for fi11 := 0; fi11 < 3; fi11++ {
                fmt.Println(v4)
                v3 -= v3
                _ = fi11
            }
            var v12 uint = uint(v4)
            _ = v12
        } else {
            for fi13 := 0; fi13 < 17; fi13++ {
                delete(v2, 996)
                var v14 float32 = (1.5 + 3.14159)
                var v15 int = fi13
                _ = fi13
                _ = v14
                _ = v15
            }
            var v16 int64 = 170
            _ = v16
        }
        v17 := GenS2{F0: -128, F1: 1}
        _ = fi9
        _ = v10
        _ = v17
    }
    v8 = 1
    v4 -= (v4 | v4)
    for fi18 := 0; fi18 < 17; fi18++ {
        var v19 uint16 = v4
        _ = fi18
        _ = v19
    }
    var v20 int32 = 1
    v4 += v4
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v8
}
