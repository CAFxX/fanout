package main

import (
    "fmt"
)

func genFn0(p0 int32, p1 uint8, p2 int16) uint {
    if (722 >= 42) {
        p0 += (p0 | (p0 & (p0 >> ((65535) & 63))))
        p2 = -32768
    }
    p1 -= 42
    for fi6 := 0; fi6 < 10; fi6++ {
        var v7 uint8 = (p1 << ((530) & 63))
        _ = fi6
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    return 65535
}

func main() {
    var v1 uint32 = 2
    v2 := []uint32{}
    v3 := make(map[uint32]string)
    var v4 int64 = -1
    v5 := make(map[uint32]uint32)
    for fi8 := 0; fi8 < 9; fi8++ {
        fmt.Println(fi8, fi8, v4)
        _ = fi8
    }
    var v9 int32 = 615
    var v10 int32 = v9
    v1 += (v1 * 4294967295)
    var v11 int16 = 32767
    var v12 float32 = -2.25
    var v13 int = 32768
    v9 = (v10 / ((v10) | 1))
    var v14 bool = (4294967295 != 4294967295)
    fmt.Println(v9, v4, v13)
    v2 = append(v2, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
