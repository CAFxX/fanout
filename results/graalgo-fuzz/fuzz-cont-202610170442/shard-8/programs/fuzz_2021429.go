package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 string; F1 int; F2 uint64 }

func main() {
    var v1 uint8 = 2
    var v2 uint64 = 18446744073709551615
    var v3 int = 32767
    v4 := make(map[string]int)
    fmt.Println(v3)
    v5 := GenS1{F0: 18446744073709551615}
    for fi6 := 0; fi6 < 17; fi6++ {
        v1 -= (v1 >> ((0) & 63))
        var v7 int16 = 1000
        _ = fi6
        _ = v7
    }
    v1 = v1
    switch v3 {
    case 42:
        v8 := []int{649, v3, (^(v3))}
        _ = v8
    case -607:
        var v9 string = "\n\t"
        _ = v9
    default:
        v10 := GenS2{F0: ("foo bar" + "\n\t"), F1: (-(v3)), F2: (^(v2))}
        _ = v10
    }
    var v11 int = (v3 >> ((188) & 63))
    var v12 int64 = 91
    for fi13 := 0; fi13 < 12; fi13++ {
        var v14 bool = ((1.5 + (3.14159 + (float32(((1.5 * -2.25) / 0.1)) - (float32(1e10) - (float32(3.14159) * -2.25))))) < 0.1)
        var v15 uint64 = 1
        _ = fi13
        _ = v14
        _ = v15
    }
    v1 -= (v1 & (v1 - (v1 * (v1 + 2))))
    if (v3 < v11) {
        if ((-32768 <= 289) || true) {
            var v16 int16 = -32768
            v11 += (v3 + v11)
            _ = v16
        }
        if (!(255 == 42)) {
            v17 := []int{-813, (v11 | 593)}
            _ = v17
        }
    } else {
        for fi18 := 0; fi18 < 9; fi18++ {
            for fi19 := 0; fi19 < 1; fi19++ {
                var v20 int = (v3 * v11)
                _ = fi19
                _ = v20
            }
            for fi21 := 0; fi21 < 10; fi21++ {
                v1 += v1
                _ = fi21
            }
            var v22 int64 = (v12 >> ((2) & 63))
            _ = fi18
            _ = v22
        }
        if ((float32(float64(((1e10 * 0.1) + float32(0.0)))) + 0.1) >= (1.5 / -2.25)) {
            v3 = 32768
            fmt.Println(v2, v1)
            v23 := GenS1{F0: (v2 ^ (v2 << ((4294967295) & 63)))}
            _ = v23
        }
        var v24 uint8 = v1
        _ = v24
    }
    var v25 string = "\n\t"
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
