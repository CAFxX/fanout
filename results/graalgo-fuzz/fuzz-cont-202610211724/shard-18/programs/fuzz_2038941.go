package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16 }
type GenS2 struct { F0 int64; F1 int16; F2 uint64 }
func genFn0(p0 int16, p1 uint16, p2 int16) uint16 {
    fmt.Println(p1, p2)
    for fi4 := 0; fi4 < 9; fi4++ {
        v5 := []int{(-(fi4)), (fi4 | 925), (fi4 & 2147483647), (-(fi4))}
        _ = fi4
        _ = v5
    }
    v6 := []int{32767}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return p1
}

func main() {
    var v1 int8 = -31
    v2 := []uint32{489}
    var v3 float32 = 0.1
    v7 := GenS1{F0: 42, F1: 2}
    v8 := GenS2{F0: 1, F1: 200, F2: 2}
    var v9 uint8 = 233
    fmt.Println(v9)
    var v10 int64 = 0
    var v11 float32 = v3
    for fi12 := 0; fi12 < 13; fi12++ {
        var v13 uint8 = 127
        v10 -= v10
        _ = fi12
        _ = v13
    }
    var v14 string = ("foo bar" + (("foo bar" + "hello") + ((("" + "") + (("" + ("" + ("" + "hello"))) + "hello")) + "foo bar")))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v2
    _ = v3
    _ = v7
    _ = v8
    _ = v9
}
