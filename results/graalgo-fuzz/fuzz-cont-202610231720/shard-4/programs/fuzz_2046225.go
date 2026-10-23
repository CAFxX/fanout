package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32 }
func genFn0(p0 int8, p1 uint8) uint64 {
    var v7 float64 = float64((float32(1.5) * 1.5))
    var v8 float32 = 0.0
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 0
}

func main() {
    v1 := []int{0, 128, -2147483648, 127}
    v2 := make(map[uint32]string)
    v3 := []int{0, -1, 261}
    v4 := []uint32{0, 537}
    var v5 int8 = -128
    var v6 uint32 = 42
    if (!(!true)) {
        if len(v4) > 0 {
            v4[0] = v6
        }
    }
    var v9 float64 = 3.14159
    var v10 float32 = 3.14159
    if false {
        var v11 int32 = 783
        _ = v11
    } else {
        var v12 uint32 = (v6 % ((v6) | 1))
        delete(v2, (v6 + (v6 * v12)))
        var v13 uint32 = v12
        _ = v12
        _ = v13
    }
    for fi14 := 0; fi14 < 6; fi14++ {
        for fi15 := 0; fi15 < 16; fi15++ {
            var v16 bool = (((-(fi15)) != -2147483648) || (fi15 != fi14))
            _ = fi15
            _ = v16
        }
        _ = fi14
    }
    var v17 uint8 = 2
    switch "" {
    case "\n\t":
        for fi18 := 0; fi18 < 4; fi18++ {
            v5 = (v5 | v5)
            if ((!((uint(v17) == 4294967295) && (1 == (^(fi18))))) && false) {
                var v19 uint16 = 255
                var v20 uint8 = v17
                _ = v19
                _ = v20
            }
            _ = fi18
        }
    case "foo bar":
        v9 += v9
    case "a":
        fmt.Println(v6, v5, v9)
    default:
        v21 := GenS1{F0: v5, F1: 0}
        _ = v21
    }
    for fi22 := 0; fi22 < 8; fi22++ {
        var v23 float32 = (0.1 + 1e10)
        _ = fi22
        _ = v23
    }
    switch 32767 {
    case 127:
        if false {
            _ = len(v3)
        }
    case 32768:
        v6 += 65535
    default:
        var v24 uint64 = 4294967295
        _ = v24
    }
    v6 -= 660
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
