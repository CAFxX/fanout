package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint }
type GenS2 struct { F0 int32; F1 int16; F2 bool }
type GenS3 struct { F0 string; F1 int16 }
func genFn0(p0 int16, p1 uint8) int8 {
    v7 := []int{-1, 42, 693}
    p1 -= p1
    v8 := GenS1{F0: (p1 << ((1) & 63)), F1: 664}
    p1 = 1
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 1
}

func main() {
    var v1 uint32 = 936
    v2 := make(map[string]string)
    var v3 string = ("a" + "a")
    v4 := make(map[uint32]int)
    v5 := make(map[string]bool)
    var v6 float32 = -2.25
    for fi9 := 0; fi9 < 13; fi9++ {
        v10 := GenS2{F0: -844, F1: -32768, F2: false}
        _ = fi9
        _ = v10
    }
    v11 := GenS1{F0: 128, F1: 0}
    var v12 int64 = 1
    v3 = (("a" + "") + ("hello" + v3))
    for fi13 := 0; fi13 < 16; fi13++ {
        for fi14 := 0; fi14 < 13; fi14++ {
            v15 := []int{fi13, (^(fi13)), fi13}
            var v16 uint32 = v1
            switch fi14 {
            case -970:
                v12 += v12
            }
            fmt.Println(fi13)
            var v17 uint = 231
            v12 = v12
            _ = fi14
            _ = v15
            _ = v16
            _ = v17
        }
        _ = fi13
    }
    var v18 float32 = (v6 * v6)
    v6 += (v6 + v6)
    for fi19 := 0; fi19 < 14; fi19++ {
        v20 := GenS3{F0: v3, F1: 42}
        _ = fi19
        _ = v20
    }
    var v21 uint8 = 247
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v18
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
