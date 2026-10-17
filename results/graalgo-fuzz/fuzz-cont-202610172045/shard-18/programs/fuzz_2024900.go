package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string }
type GenS2 struct { F0 bool }
func genFn0(p0 bool) uint16 {
    fmt.Println(p0)
    p0 = (1 <= -507)
    _ = p0
    return 32767
}
func genFn1(p0 uint32) uint64 {
    p0 += (p0 << ((330) & 63))
    p0 -= p0
    _ = p0
    return 4294967295
}

func main() {
    var v1 int32 = 219
    v2 := []string{}
    var v3 float32 = 0.1
    v1 = (v1 | v1)
    if len(v2) > 0 {
        v2[0] = ("\n\t" + "")
    }
    if true {
        var v4 int64 = 821
        v5 := GenS1{F0: "\n\t", F1: ("foo bar" + "\n\t")}
        _ = v4
        _ = v5
    } else {
        var v6 int64 = 604
        v1 += v1
        _ = v6
    }
    var v7 uint8 = 0
    v2 = append(v2, (((("hello" + (("\n\t" + "") + "")) + "\n\t") + "hello") + "hello"))
    v3 = v3
    v8 := GenS2{F0: false}
    v1 += (v1 * (v1 / (((v1 & (-(v1)))) | 1)))
    for fi9 := 0; fi9 < 20; fi9++ {
        v10 := GenS2{F0: (1 <= 128)}
        _ = fi9
        _ = v10
    }
    fmt.Println(v7, v3, v7)
    if true {
        v1 = (^(v1))
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v7
    _ = v8
}
