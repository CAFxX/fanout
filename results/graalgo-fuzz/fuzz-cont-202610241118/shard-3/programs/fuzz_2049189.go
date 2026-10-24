package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint64; F1 int64 }
func genFn0(p0 int64, p1 uint16, p2 uint16) uint64 {
    v7 := GenS1{F0: 42}
    v8 := []int{9223372036854775807}
    p2 = p2
    var v9 int32 = 1
    for fi10 := 0; fi10 < 8; fi10++ {
        if (float32((1e10 + (-2.25 + (0.0 - (float64((0.1 - (0.0 / (-2.25 + 1.5)))) - 1e10))))) > (1.5 + -2.25)) {
            p0 = 1
            v11 := GenS1{F0: 1}
            v12 := GenS1{F0: 1}
            _ = v11
            _ = v12
        }
        _ = fi10
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return 18446744073709551615
}

func main() {
    var v1 int32 = -26
    var v2 uint8 = 74
    var v3 uint16 = 65535
    v4 := make(map[uint32]bool)
    v5 := []int{0, 32767, 2147483647, 669}
    v6 := []uint32{622, uint32(v2), 393, 2147483648}
    for fi13 := 0; fi13 < 1; fi13++ {
        v2 = uint8(v3)
        if (!false) {
            for fi14 := 0; fi14 < 15; fi14++ {
                var v15 int8 = 0
                _ = fi14
                _ = v15
            }
            var v16 int = fi13
            if (!true) {
                v16 -= (v16 << ((0) & 63))
                delete(v4, 981)
            } else {
                v16 = (^(fi13))
            }
            _ = v16
        }
        _ = fi13
    }
    v1 -= v1
    v2 -= 128
    fmt.Println(v3, v3)
    v1 = v1
    for fi17 := 0; fi17 < 11; fi17++ {
        delete(v4, 42)
        _ = fi17
    }
    _ = len(v4)
    v18 := GenS1{F0: 1}
    v19 := GenS1{F0: 18446744073709551615}
    v20 := GenS2{F0: 1, F1: 42}
    v3 = 255
    var v21 int32 = v1
    fmt.Println(v2, v1, v2)
    fmt.Println(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
