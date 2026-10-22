package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16 }
func genFn0(p0 uint16, p1 int8) string {
    var v4 uint16 = p0
    if (0 < 42) {
        p0 = v4
    }
    _ = p0
    _ = p1
    _ = v4
    return "\n\t"
}
func genFn1(p0 int16, p1 int8, p2 bool) uint16 {
    var v5 int8 = 0
    v6 := GenS1{F0: false, F1: p0}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 42
}

func main() {
    var v1 int32 = 1
    v2 := []string{}
    v3 := make(map[string]uint32)
    v1 += 957
    fmt.Println(v1)
    var v7 int64 = 9223372036854775807
    v8 := GenS1{F0: (1 != int16(v1)), F1: 361}
    var v9 float64 = 1.5
    var v10 uint32 = 42
    fmt.Println(v9)
    v10 = 575
    var v11 int32 = -1
    fmt.Println(v11)
    for fi12 := 0; fi12 < 9; fi12++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        for fi13 := 0; fi13 < 4; fi13++ {
            fmt.Println(fi13, fi13)
            v7 = (v7 % ((1) | 1))
            _ = len(v2)
            _ = fi13
        }
        for fi14 := 0; fi14 < 10; fi14++ {
            if false {
                if (!(65535 >= 1)) {
                    _ = len(v3)
                }
            }
            v11 -= v1
            _ = fi14
        }
        _ = fi12
    }
    v10 -= (v10 << ((128) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v7
    _ = v8
    _ = v9
}
