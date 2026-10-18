package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8; F2 int32 }
type GenS2 struct { F0 int16; F1 string }
func genFn0(p0 bool, p1 bool, p2 string) uint8 {
    v5 := []int{-9223372036854775808, -2147483648, 32767}
    fmt.Println(p2, p2, p1)
    var v6 float64 = 3.14159
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 42
}

func main() {
    v1 := []uint32{}
    var v2 int16 = 42
    var v3 int64 = 42
    var v4 uint16 = 670
    v4 += (v4 & v4)
    fmt.Println(v4, v3)
    v7 := GenS1{F0: false, F1: 167, F2: -1}
    var v8 int8 = 42
    var v9 bool = (42 != 42)
    if v9 {
        if (false || true) {
            var v10 uint64 = 4294967295
            var v11 string = "a"
            _ = v10
            _ = v11
        }
        var v12 int16 = (v2 ^ v2)
        for fi13 := 0; fi13 < 9; fi13++ {
            if v9 {
                switch fi13 {
                case 4294967295:
                    v4 = 164
                }
            } else {
                if (!v9) {
                    v14 := GenS2{F0: 42, F1: "foo bar"}
                    fmt.Println(v9)
                    _ = v14
                }
            }
            _ = fi13
        }
        _ = v12
    } else {
        switch 0 {
        case -9223372036854775808:
            fmt.Println(v3, v2)
        case 219:
            v2 -= (v2 << ((990) & 63))
        }
        if len(v1) > 0 {
            v1[0] = 2147483648
        }
        if len(v1) > 0 {
            v1[0] = 1
        }
    }
    v4 += (v4 * v4)
    fmt.Println(v9, v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
