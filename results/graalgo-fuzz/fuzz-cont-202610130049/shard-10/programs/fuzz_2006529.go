package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16 }
func genFn0(p0 int64) bool {
    p0 = p0
    var v6 uint64 = 186
    fmt.Println(v6, v6)
    fmt.Println(v6, v6)
    _ = p0
    _ = v6
    return false
}
func genFn1(p0 int8, p1 uint) uint {
    for fi7 := 0; fi7 < 3; fi7++ {
        var v8 uint8 = 1
        v9 := []int{fi7, (^(fi7)), (fi7 + fi7), -1}
        _ = fi7
        _ = v8
        _ = v9
    }
    var v10 uint32 = 65535
    if (true || true) {
        switch -1 {
        case -349:
            p0 -= (p0 >> (((p1 & 42)) & 63))
        case -2147483648:
            v10 = (v10 | 2147483648)
        }
    }
    p0 = 100
    v10 -= v10
    _ = p0
    _ = p1
    _ = v10
    return (^(p1))
}

func main() {
    var v1 int8 = -128
    var v2 float32 = (float32(0.1) + (1.5 + (((3.14159 * float32(((0.0 - 1e10) / ((((1.5 - 3.14159) * 0.0) + -2.25) / 1.0)))) / -2.25) - 0.1)))
    var v3 int8 = -1
    v4 := make(map[uint32]uint32)
    v5 := make(map[uint32]int)
    for fi11 := 0; fi11 < 6; fi11++ {
        v12 := GenS1{F0: int16(v1), F1: 1}
        for fi13 := 0; fi13 < 13; fi13++ {
            delete(v5, 4294967295)
            _ = len(v5)
            fmt.Println(v3)
            switch (fi13 / ((fi11) | 1)) {
            case -1:
                v14 := GenS1{F0: 250, F1: 42}
                _ = v14
            case 4294967295:
                v15 := GenS1{F0: -1, F1: 0}
                _ = v15
            }
            _ = fi13
        }
        _ = len(v5)
        _ = fi11
        _ = v12
    }
    v1 = 23
    var v16 int8 = -91
    var v17 int = 389
    v3 = (v16 / ((v16) | 1))
    var v18 float64 = 0.0
    v19 := GenS1{F0: 159, F1: 324}
    var v20 int16 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
