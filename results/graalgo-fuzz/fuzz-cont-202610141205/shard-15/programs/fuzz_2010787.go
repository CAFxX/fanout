package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16; F2 int16 }

func main() {
    var v1 uint64 = 1
    v2 := []uint32{701}
    v3 := []uint32{2, 42, 2147483648, 42}
    var v4 int64 = 42
    var v5 float32 = 1.5
    v2 = append(v2, 65535)
    var v6 float32 = -2.25
    v4 = (v4 + 1)
    v7 := GenS1{F0: -2147483648, F1: 65535, F2: 1}
    var v8 float32 = (v6 - (1e10 * v6))
    v1 = v1
    if (255 <= 681) {
        for fi9 := 0; fi9 < 2; fi9++ {
            v8 -= 1.5
            fmt.Println(v8)
            v10 := GenS1{F0: 2147483647, F1: 0, F2: 42}
            _ = fi9
            _ = v10
        }
    }
    v11 := GenS1{F0: 206, F1: 959, F2: 32767}
    var v12 uint = 695
    if true {
        switch -9223372036854775808 {
        case -912:
            if len(v2) > 0 {
                v2[0] = 65535
            }
        case -871:
            for fi13 := 0; fi13 < 11; fi13++ {
                for fi14 := 0; fi14 < 5; fi14++ {
                    v1 = v1
                    _ = fi14
                }
                _ = fi13
            }
        case 128:
            fmt.Println(v12, v4, v12)
        }
        v12 += v12
    }
    if true {
        for fi15 := 0; fi15 < 2; fi15++ {
            v8 = float32((1e10 * 1e10))
            _ = fi15
        }
    }
    v16 := GenS1{F0: 1, F1: 0, F2: 32767}
    if ((v4 % ((v4) | 1)) >= (-(v4))) {
        v5 -= 0.0
        v1 = (v1 >> ((v12) & 63))
        _ = len(v3)
    }
    if (2 == 42) {
        v4 += (-(v4))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
