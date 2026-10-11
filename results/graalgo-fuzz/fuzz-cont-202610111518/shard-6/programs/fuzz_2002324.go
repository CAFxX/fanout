package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16 }
type GenS2 struct { F0 int16; F1 int16; F2 int16 }
func genFn0(p0 int, p1 uint32, p2 string) int16 {
    fmt.Println(p1, p0, p2)
    if (("" + p2) <= "") {
        p0 = p0
    }
    _ = p0
    _ = p1
    _ = p2
    return -303
}

func main() {
    var v1 int16 = 32767
    var v2 uint8 = 127
    var v3 uint16 = 32767
    v4 := []uint32{419, 1, 42}
    v5 := make(map[uint32]uint32)
    var v6 bool = (!(v2 < (v2 + 42)))
    var v7 float32 = 3.14159
    for fi8 := 0; fi8 < 14; fi8++ {
        v9 := GenS1{F0: v3, F1: (v1 | 959)}
        v6 = (false && ((9223372036854775807 == -9223372036854775808) || ((v6 || true) || v6)))
        v7 -= (v7 + (3.14159 + 1.5))
        _ = fi8
        _ = v9
    }
    v4 = append(v4, 2)
    var v10 int8 = 0
    v1 -= int16(v10)
    var v11 int32 = 1
    v12 := v5[2147483648]
    v5[(v12 * 0)] = (v12 >> ((1) & 63))
    switch 2147483647 {
    case 32768:
        for fi13 := 0; fi13 < 15; fi13++ {
            if v6 {
                v3 += v3
            }
            v1 += v1
            _ = fi13
        }
    case 4294967295:
        fmt.Println(v3, v1, v3)
    default:
        v14 := GenS2{F0: v1, F1: (v1 << ((737) & 63)), F2: v1}
        _ = v14
    }
    for fi15 := 0; fi15 < 14; fi15++ {
        if (!("hello" <= "a")) {
            var v16 uint64 = 18446744073709551615
            _ = v16
        }
        _ = fi15
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
