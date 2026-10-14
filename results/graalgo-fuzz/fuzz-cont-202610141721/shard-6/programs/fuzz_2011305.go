package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint64 }
func genFn0(p0 uint16) uint8 {
    p0 = (p0 ^ p0)
    for fi6 := 0; fi6 < 11; fi6++ {
        p0 = 2
        _ = fi6
    }
    for fi7 := 0; fi7 < 4; fi7++ {
        var v8 int = (fi7 << ((1) & 63))
        v9 := GenS1{F0: fi7}
        _ = fi7
        _ = v8
        _ = v9
    }
    p0 += p0
    _ = p0
    return 1
}

func main() {
    var v1 float32 = 3.14159
    v2 := make(map[uint32]string)
    var v3 int32 = 1
    v4 := make(map[uint32]bool)
    v5 := make(map[int]string)
    var v10 uint32 = 2
    v11 := []int{}
    v11 = append(v11, 32768)
    v11 = append(v11, 1)
    for fi12 := 0; fi12 < 3; fi12++ {
        v1 = v1
        _ = fi12
    }
    fmt.Println(v3, v1)
    if len(v11) > 0 {
        v11[0] = 4294967295
    }
    var v13 uint32 = 32
    var v14 int = -990
    var v15 int16 = 0
    switch v14 {
    case 0:
        var v16 string = "a"
        _ = v16
    case -33:
        delete(v2, v10)
    }
    if (true || (int8(v15) > -1)) {
        var v17 float64 = 1e10
        v10 = 65535
        fmt.Println(v3)
        _ = v17
    } else {
        var v18 uint64 = 490
        v15 = (v15 + -894)
        var v19 int8 = -57
        _ = v18
        _ = v19
    }
    v20 := GenS2{F0: 342}
    v10 += (v13 + 42)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
