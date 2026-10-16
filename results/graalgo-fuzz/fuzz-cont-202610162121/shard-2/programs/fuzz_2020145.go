package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64; F2 int32 }
type GenS2 struct { F0 int8; F1 int; F2 int64 }
func genFn0(p0 int64, p1 int8, p2 int8) int16 {
    p0 = p0
    for fi4 := 0; fi4 < 2; fi4++ {
        v5 := GenS1{F0: 412, F1: 504, F2: -50}
        p2 -= -128
        for fi6 := 0; fi6 < 10; fi6++ {
            if false {
                if (!true) {
                    var v7 int8 = p2
                    p2 -= 100
                    v8 := GenS2{F0: v7, F1: (fi4 ^ (^(fi6))), F2: p0}
                    _ = v7
                    _ = v8
                }
                v9 := []int{(fi4 & (fi6 >> ((42) & 63))), (fi6 * (-(fi6))), (fi4 % ((fi6) | 1)), fi4}
                _ = v9
            }
            _ = fi6
        }
        _ = fi4
        _ = v5
    }
    fmt.Println(p1)
    var v10 string = ("" + "")
    var v11 uint8 = 249
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    return -32768
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[string]int)
    v3 := make(map[uint32]bool)
    var v12 uint8 = 1
    var v13 uint = 18446744073709551615
    for fi14 := 0; fi14 < 9; fi14++ {
        var v15 int8 = -38
        _ = fi14
        _ = v15
    }
    fmt.Println(v12, v12)
    v12 -= (^(v12))
    var v16 uint16 = 0
    if (!true) {
        switch (("" + "a") + ("a" + "\n\t")) {
        case "\n\t":
            v16 = (v16 % (((v16 - (v16 - v16))) | 1))
        default:
            v17 := []int{127, -1, 42, 32767}
            _ = v17
        }
    }
    var v18 uint = v13
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v12
    _ = v13
    _ = v16
    _ = v18
    _ = v2
    _ = v3
}
