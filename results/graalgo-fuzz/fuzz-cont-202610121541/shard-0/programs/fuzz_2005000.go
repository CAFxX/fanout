package main

import (
    "fmt"
)

func genFn0(p0 uint32, p1 uint8) uint64 {
    var v6 bool = false
    v6 = (2147483647 != 128)
    for fi7 := 0; fi7 < 2; fi7++ {
        fmt.Println(fi7, fi7)
        var v8 uint16 = 32767
        _ = fi7
        _ = v8
    }
    for fi9 := 0; fi9 < 10; fi9++ {
        var v10 int64 = -243
        _ = fi9
        _ = v10
    }
    fmt.Println(v6)
    _ = p0
    _ = p1
    _ = v6
    return 117
}
func genFn1(p0 string, p1 int8, p2 int16) string {
    fmt.Println(p1)
    var v11 uint16 = 32767
    var v12 int32 = 42
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    return "foo bar"
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int = -649
    v3 := []uint32{0}
    var v4 int8 = 42
    v5 := make(map[int]uint32)
    var v13 int16 = 1
    v4 = v4
    var v14 int16 = v13
    fmt.Println(v14, v14)
    for fi15 := 0; fi15 < 2; fi15++ {
        var v16 int32 = 1
        _ = fi15
        _ = v16
    }
    _ = len(v5)
    fmt.Println(v2, v13, v2)
    v4 = (v4 - (v4 / (((v4 % ((v4) | 1))) | 1)))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
