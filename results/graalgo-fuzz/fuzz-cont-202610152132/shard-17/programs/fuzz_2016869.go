package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32 }
type GenS2 struct { F0 int64; F1 string; F2 uint64 }
func genFn0(p0 int64, p1 int64, p2 uint64) uint16 {
    fmt.Println(p1, p1)
    var v5 uint = 77
    v6 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 547
}

func main() {
    v1 := make(map[string]int)
    v2 := make(map[uint32]uint32)
    v3 := make(map[int]int)
    var v4 int8 = 1
    for fi7 := 0; fi7 < 13; fi7++ {
        v4 -= 100
        if true {
            v8 := []int{(fi7 ^ (fi7 | (-(fi7))))}
            for fi9 := 0; fi9 < 8; fi9++ {
                delete(v3, (fi9 << ((65535) & 63)))
                _ = fi9
            }
            _ = v8
        }
        if false {
            for fi10 := 0; fi10 < 12; fi10++ {
                if (42 == 2147483648) {
                    v4 = v4
                }
                _ = fi10
            }
            v4 = (v4 >> ((1000) & 63))
            v4 = v4
        }
        _ = fi7
    }
    v4 -= v4
    var v11 int16 = 0
    for fi12 := 0; fi12 < 13; fi12++ {
        _ = len(v2)
        _ = fi12
    }
    v13 := []int{92, 9223372036854775807, -2147483648, 865}
    for fi14 := 0; fi14 < 1; fi14++ {
        v15 := GenS1{F0: (v4 ^ -1), F1: 42}
        _ = len(v13)
        _ = fi14
        _ = v15
    }
    v3[int(v11)] = 215
    v4 -= (v4 % ((v4) | 1))
    v4 -= 127
    if len(v13) > 0 {
        v13[0] = 981
    }
    var v16 uint64 = 42
    v17 := GenS2{F0: 42, F1: ("foo bar" + ("" + "a")), F2: (v16 + 620)}
    var v18 uint16 = 276
    var v19 float32 = (((float32(3.14159) - float32(1.5)) / (3.14159 * 1e10)) * 0.0)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v13
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
