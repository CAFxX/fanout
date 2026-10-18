package main

import (
    "fmt"
)

func genFn0(p0 uint8) int {
    for fi6 := 0; fi6 < 18; fi6++ {
        p0 = p0
        _ = fi6
    }
    var v7 uint8 = 2
    _ = p0
    _ = v7
    return 466
}
func genFn1(p0 bool, p1 uint8, p2 uint32) uint64 {
    fmt.Println(p1)
    var v8 uint8 = 42
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 1
}

func main() {
    var v1 uint16 = 552
    v2 := make(map[int]uint32)
    var v3 int16 = 1000
    v4 := []uint32{}
    v5 := []string{"", "\n\t", "", "a"}
    v1 -= (v1 - 2)
    var v9 int8 = 1
    v1 -= (v1 << ((uint(v1)) & 63))
    var v10 int64 = -646
    var v11 bool = (false && (153 < 1))
    var v12 uint32 = 65535
    var v13 uint16 = v1
    delete(v2, 32767)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
