package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16 }
type GenS2 struct { F0 uint }
func genFn0() uint64 {
    var v5 float32 = (((3.14159 * ((((((1.5 / -2.25) - 0.0) * 0.0) / 0.1) - float32(float64(0.0))) - 1.5)) * 0.1) * (float32(0.1) - (-2.25 - 3.14159)))
    var v6 int = 494
    fmt.Println(v6)
    _ = v5
    _ = v6
    return 852
}
func genFn1(p0 uint8) int32 {
    p0 -= (^(p0))
    v7 := GenS1{F0: 793, F1: 855}
    _ = p0
    _ = v7
    return -1
}

func main() {
    v1 := []int{128, 4294967295, 66}
    v2 := make(map[int]int)
    v3 := make(map[int]int)
    v4 := []string{(("\n\t" + ("hello" + "a")) + "")}
    var v8 bool = true
    if true {
        v8 = true
    } else {
        v9 := GenS1{F0: -32768, F1: 32767}
        v8 = (v8 || (v8 && (((("hello" + ("" + "a")) < ("foo bar" + "")) && false) || true)))
        _ = v9
    }
    _ = len(v3)
    var v10 uint32 = 0
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v11 := GenS2{F0: 160}
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v12 int32 = genFn1(173)
    v10 += (v10 ^ 230)
    var v13 int32 = 603
    v14 := GenS2{F0: 581}
    v12 += v12
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}
