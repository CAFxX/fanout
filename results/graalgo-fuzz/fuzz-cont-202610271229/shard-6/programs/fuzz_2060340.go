package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
func genFn0(p0 int, p1 bool, p2 uint32) bool {
    p1 = p1
    fmt.Println(p1, p2, p0)
    var v4 float64 = 1e10
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return (-2147483648 < -1)
}
func genFn1(p0 int32) uint64 {
    v5 := []int{32768, int(p0), 32767}
    if (true || false) {
        var v6 int64 = 9223372036854775807
        _ = v6
    } else {
        if len(v5) > 0 {
            v5[0] = 934
        }
        p0 += (p0 ^ p0)
        var v7 bool = (false || (p0 >= 0))
        _ = v7
    }
    var v8 uint64 = 0
    v9 := GenS1{F0: (1.5 == (float32(0.0) / 3.14159))}
    for fi10 := 0; fi10 < 2; fi10++ {
        switch (fi10 % ((42) | 1)) {
        case 0:
            var v11 uint64 = (v8 + (v8 * v8))
            _ = v11
        }
        v8 -= (v8 & v8)
        v8 = 0
        _ = fi10
    }
    _ = p0
    _ = v5
    _ = v8
    _ = v9
    return 896
}

func main() {
    v1 := make(map[int]uint32)
    var v2 int32 = 2147483647
    v3 := make(map[uint32]int)
    var v12 int = -424
    for fi13 := 0; fi13 < 17; fi13++ {
        v14 := []int{fi13}
        _ = fi13
        _ = v14
    }
    for fi15 := 0; fi15 < 16; fi15++ {
        v12 -= fi15
        _ = fi15
    }
    v2 += -1
    if (!false) {
        v12 = -500
    } else {
        fmt.Println(v2, v2)
        v2 += -72
        fmt.Println(v12, v12)
    }
    var v16 uint32 = 2147483647
    fmt.Println(v12)
    v17 := GenS1{F0: ((100 != 4) && ((false && true) || true))}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v12
    _ = v16
    _ = v17
    _ = v2
    _ = v3
}
