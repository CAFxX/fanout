package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32; F2 int32 }
func genFn0(p0 int16) bool {
    p0 += 32767
    var v4 uint64 = 1
    var v5 int16 = -1
    _ = p0
    _ = v4
    _ = v5
    return true
}
func genFn1(p0 string, p1 bool) bool {
    p1 = p1
    var v6 uint32 = 1
    if p1 {
        p0 = (p0 + "hello")
        p0 = ((p0 + p0) + "")
        p1 = p1
    } else {
        v6 -= 65535
    }
    var v7 int32 = -643
    var v8 int64 = 9223372036854775807
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    return ((!p1) && p1)
}

func main() {
    var v1 int64 = 0
    v2 := []string{(((("\n\t" + "\n\t") + "") + (("hello" + ("" + "hello")) + "")) + ""), ("" + ""), "foo bar"}
    v3 := make(map[uint32]int)
    fmt.Println(v1)
    var v9 int8 = -75
    v3[940] = 128
    v10 := GenS1{F0: 32767, F1: 304, F2: 42}
    var v11 bool = ((-224 <= 32768) && true)
    v3[2] = 32768
    v1 -= (v1 - (v1 & (v1 ^ 9223372036854775807)))
    fmt.Println(v11, v11)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v9
}
