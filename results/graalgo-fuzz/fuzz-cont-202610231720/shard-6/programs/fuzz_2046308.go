package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
func genFn0(p0 int32, p1 int32, p2 uint16) string {
    var v6 uint = 1
    var v7 int16 = 42
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return "foo bar"
}

func main() {
    var v1 string = "hello"
    var v2 uint32 = 2147483648
    var v3 uint64 = 4294967295
    v4 := make(map[uint32]string)
    var v5 int16 = -32768
    for fi8 := 0; fi8 < 13; fi8++ {
        var v9 int8 = 1
        v2 = v2
        var v10 int32 = 636
        _ = fi8
        _ = v10
        _ = v9
    }
    if (173 < 0) {
        var v11 int32 = int32(v2)
        var v12 float32 = float32((0.0 / 1.0))
        fmt.Println(v5, v1, v1)
        _ = v11
        _ = v12
    }
    fmt.Println(v3, v5, v3)
    fmt.Println(v3, v5, v1)
    v5 -= (v5 << ((1) & 63))
    var v13 int16 = 1000
    for fi14 := 0; fi14 < 8; fi14++ {
        switch fi14 {
        case 128:
            v13 -= (v13 >> ((0) & 63))
        case 346:
            v15 := []int{(fi14 ^ 32768), (fi14 << ((234) & 63)), fi14, (-(fi14))}
            _ = v15
        }
        _ = fi14
    }
    if (!(!(680 > 0))) {
        v16 := []int{-833}
        var v17 uint32 = (v2 ^ v2)
        for fi18 := 0; fi18 < 1; fi18++ {
            _ = len(v4)
            v2 = uint32(v3)
            _ = fi18
        }
        _ = v16
        _ = v17
    }
    var v19 int8 = 127
    v20 := GenS1{F0: 272}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
