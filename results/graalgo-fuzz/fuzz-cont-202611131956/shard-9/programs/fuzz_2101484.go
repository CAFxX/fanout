package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint }
type GenS2 struct { F0 uint64; F1 string }
func genFn0(p0 int8, p1 uint8) bool {
    var v5 string = ("foo bar" + "\n\t")
    var v6 int = 42
    v7 := GenS1{F0: (uint32(p1) < 648), F1: 65535}
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    return (v6 == v6)
}

func main() {
    v1 := make(map[uint32]int)
    v2 := []int{}
    var v3 string = (("" + "\n\t") + "")
    var v4 int16 = -1
    v3 = v3
    v3 = ""
    var v8 uint16 = 65535
    var v9 int32 = -2147483648
    if (v4 > v4) {
        if false {
            v9 += v9
            fmt.Println(v9, v8, v8)
        }
    }
    if (true && false) {
        v10 := GenS2{F0: 924, F1: (v3 + v3)}
        v8 -= v8
        v3 = "\n\t"
        _ = v10
    }
    var v11 uint16 = (v8 / ((65535) | 1))
    var v12 uint8 = 206
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
