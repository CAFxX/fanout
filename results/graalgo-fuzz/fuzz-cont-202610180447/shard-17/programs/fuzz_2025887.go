package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64; F2 uint16 }
type GenS2 struct { F0 bool; F1 uint32; F2 int32 }
func genFn0(p0 int32, p1 uint) bool {
    var v5 int64 = 1
    v6 := GenS1{F0: "\n\t", F1: v5, F2: 255}
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return (1.5 != 0.0)
}

func main() {
    var v1 int64 = 1
    var v2 uint64 = 1
    var v3 int8 = 42
    var v4 float32 = 3.14159
    if true {
        if ((false || (!false)) || false) {
            var v7 int = 1
            _ = v7
        } else {
            v8 := GenS2{F0: (false || true), F1: 42, F2: -1}
            for fi9 := 0; fi9 < 12; fi9++ {
                fmt.Println(v2, v4, v3)
                _ = fi9
            }
            _ = v8
        }
        var v10 string = "\n\t"
        _ = v10
    }
    var v11 float32 = ((1.5 + (float32(1e10) - v4)) / -2.25)
    v3 = v3
    v12 := GenS1{F0: ("\n\t" + ""), F1: v1, F2: 32767}
    fmt.Println(v3)
    v4 += ((v4 / v11) * 0.0)
    for fi13 := 0; fi13 < 11; fi13++ {
        v14 := []int{-813, (fi13 ^ (fi13 ^ 4294967295)), -2147483648, fi13}
        _ = fi13
        _ = v14
    }
    v15 := []int{0, 32768, 517, -31}
    for fi16 := 0; fi16 < 11; fi16++ {
        v3 += v3
        _ = fi16
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v15
    _ = v2
    _ = v3
    _ = v4
}
