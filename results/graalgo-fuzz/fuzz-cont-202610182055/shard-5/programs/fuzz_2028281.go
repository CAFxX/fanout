package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16 }
type GenS2 struct { F0 int32; F1 int16; F2 int8 }
func genFn0(p0 uint64, p1 bool, p2 string) uint64 {
    var v4 float64 = (3.14159 / float64(float32(0.1)))
    var v5 int16 = -32768
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return p0
}

func main() {
    var v1 uint32 = 0
    v2 := []string{"a", "a"}
    var v3 int16 = int16(v1)
    var v6 int8 = -45
    v6 -= 127
    var v7 int32 = 0
    if len(v2) > 0 {
        v2[0] = "\n\t"
    }
    fmt.Println(v6, v6)
    var v8 int = 32768
    _ = len(v2)
    v9 := GenS1{F0: false, F1: 154}
    v7 += v7
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v10 := GenS2{F0: v7, F1: -32768, F2: (-(v6))}
    v11 := GenS2{F0: (v7 | (v7 - v7)), F1: (v3 << ((18446744073709551615) & 63)), F2: 42}
    if (!true) {
        v1 = 2147483648
    }
    switch 1 {
    case 32767:
        v1 += (^(v1))
    case 32768:
        v6 += v6
    case 1:
        _ = len(v2)
    default:
        v8 += v8
    }
    for fi12 := 0; fi12 < 11; fi12++ {
        fmt.Println(v8, v7)
        v8 += (fi12 / (((-(v8))) | 1))
        _ = fi12
    }
    switch (v8 - (-(v8))) {
    case 748:
        if false {
            fmt.Println(v6)
            v8 -= (v8 << ((174) & 63))
            if len(v2) > 0 {
                v2[0] = ("\n\t" + "a")
            }
        }
    case 127:
        v8 += v8
    case 32767:
        var v13 uint32 = v1
        _ = v13
    default:
        v3 -= (v3 & (v3 >> ((255) & 63)))
    }
    if false {
        v8 += (v8 * (v8 << ((255) & 63)))
        v3 = (v3 ^ (-(v3)))
        for fi14 := 0; fi14 < 20; fi14++ {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = fi14
        }
    } else {
        v3 = (v3 % ((v3) | 1))
    }
    var v15 float64 = 0.0
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v15
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
