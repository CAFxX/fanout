package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int }
func genFn0(p0 uint32) uint8 {
    var v5 float32 = 1e10
    fmt.Println(v5, v5)
    var v6 int = -9223372036854775808
    _ = p0
    _ = v5
    _ = v6
    return 144
}

func main() {
    var v1 uint = 565
    var v2 int16 = 32767
    v3 := []int{42, 37}
    v4 := []uint32{2147483647, 638, 2147483648, 798}
    if len(v3) > 0 {
        v3[0] = 0
    }
    var v7 int32 = 42
    v4 = append(v4, 4294967295)
    var v8 bool = (128 < 32767)
    for fi9 := 0; fi9 < 6; fi9++ {
        var v10 int32 = (-(v7))
        _ = fi9
        _ = v10
    }
    v7 += (v7 * v7)
    for fi11 := 0; fi11 < 9; fi11++ {
        if (false && v8) {
            if v8 {
                var v12 int8 = 1
                _ = v12
            } else {
                for fi13 := 0; fi13 < 3; fi13++ {
                    var v14 int8 = 127
                    var v15 float64 = 0.1
                    _ = fi13
                    _ = v14
                    _ = v15
                }
                fmt.Println(fi11, v8)
            }
            v7 = v7
        }
        _ = fi11
    }
    var v16 uint8 = 127
    v17 := GenS1{F0: v16, F1: 131}
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
