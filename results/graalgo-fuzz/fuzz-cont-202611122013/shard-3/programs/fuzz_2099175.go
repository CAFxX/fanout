package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8 }
func genFn0() uint64 {
    if (0.1 == 1e10) {
        v5 := GenS1{F0: 128, F1: 1}
        if true {
            for fi6 := 0; fi6 < 11; fi6++ {
                fmt.Println(fi6)
                v7 := []int{(fi6 ^ (^(fi6))), fi6, fi6, fi6}
                _ = fi6
                _ = v7
            }
            var v8 int32 = 0
            _ = v8
        }
        _ = v5
    }
    v9 := []int{4294967295, 32768, 0}
    var v10 bool = true
    v10 = (0.1 == (-2.25 - 1.5))
    _ = v10
    _ = v9
    return 18446744073709551615
}
func genFn1(p0 uint8, p1 string, p2 int) string {
    var v11 uint16 = 42
    switch (p2 << ((854) & 63)) {
    case 42:
        var v12 int = (p2 << ((2) & 63))
        _ = v12
    case 2147483647:
        var v13 int = p2
        _ = v13
    default:
        p0 = p0
    }
    var v14 int64 = 9223372036854775807
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v14
    return p1
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[int]uint32)
    v3 := make(map[uint32]bool)
    var v4 uint = 375
    var v15 uint32 = 2
    v16 := GenS1{F0: 128, F1: 1}
    v15 += 251
    v17 := v3[(v15 * (v15 >> ((2) & 63)))]
    v18 := GenS1{F0: 42, F1: -128}
    if v17 {
        v4 += (v4 * v4)
        v4 -= (v4 << (((v4 >> ((0) & 63))) & 63))
        var v19 uint64 = 735
        _ = v19
    }
    v4 = (v4 % ((v4) | 1))
    _ = len(v2)
    v17 = v17
    v20 := GenS1{F0: 255, F1: -128}
    var v21 int8 = 100
    v1[(v15 >> ((v4) & 63))] = ("\n\t" + ("a" + ("foo bar" + ((("" + "foo bar") + (("a" + "a") + "a")) + "a"))))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
}
