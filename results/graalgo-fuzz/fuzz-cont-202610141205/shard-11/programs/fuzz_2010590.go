package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int8; F1 uint32 }
func genFn0(p0 string, p1 uint32, p2 int) uint64 {
    var v5 int32 = -1
    var v6 uint64 = 2
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 462
}

func main() {
    var v1 int32 = 0
    var v2 uint32 = 158
    var v3 int = -495
    v4 := make(map[string]bool)
    v2 += 42
    v7 := GenS1{F0: 32767}
    var v8 int = -573
    if (!(238 <= 18446744073709551615)) {
        v9 := GenS2{F0: -37, F1: (^(v2))}
        var v10 uint8 = 255
        v11 := GenS2{F0: 0, F1: (v2 ^ v2)}
        _ = v10
        _ = v11
        _ = v9
    } else {
        var v12 float32 = 1.5
        fmt.Println(v8)
        v3 += v8
        _ = v12
    }
    v13 := []int{v8, v3, (v8 & (v3 | 4294967295))}
    fmt.Println(v1)
    v1 -= 0
    v1 = (v1 >> ((1) & 63))
    v3 -= int(v1)
    if (v2 != 576) {
        for fi14 := 0; fi14 < 15; fi14++ {
            v15 := GenS1{F0: 1}
            var v16 uint32 = (v2 % ((v2) | 1))
            _ = fi14
            _ = v15
            _ = v16
        }
    } else {
        if len(v13) >= 2 {
            v13 = v13[1:len(v13)]
        }
    }
    if (true || ((122 <= 193) && ((!(true || (65535 < 460))) || false))) {
        v2 += v2
        if (0 < 42) {
            v2 = v2
        } else {
            var v17 uint8 = 1
            var v18 int16 = 42
            v18 = v18
            _ = v17
            _ = v18
        }
    } else {
        delete(v4, (("\n\t" + "foo bar") + "\n\t"))
        var v19 float64 = 0.0
        _ = v19
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
