package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint }
type GenS2 struct { F0 int }
func genFn0(p0 int8, p1 string, p2 uint) uint8 {
    v7 := GenS1{F0: 2147483647, F1: 42}
    fmt.Println(p2)
    p1 = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 128
}
func genFn1(p0 uint32, p1 int32) uint32 {
    v8 := []int{}
    p1 -= (p1 - p1)
    for fi9 := 0; fi9 < 15; fi9++ {
        var v10 uint8 = 42
        if len(v8) > 0 {
            v8[0] = -2147483648
        }
        _ = fi9
        _ = v10
    }
    _ = p0
    _ = p1
    _ = v8
    return (p0 * p0)
}

func main() {
    var v1 bool = false
    var v2 uint8 = 40
    var v3 string = ("" + "hello")
    v4 := []string{((v3 + "foo bar") + v3), v3, "hello", ("hello" + ((v3 + ((((v3 + v3) + (((v3 + "foo bar") + v3) + v3)) + v3) + v3)) + v3))}
    var v5 uint32 = 4294967295
    v6 := []int{2147483647, 1, 0}
    if len(v4) > 0 {
        v4[0] = (v3 + (v3 + (v3 + "\n\t")))
    }
    v11 := GenS1{F0: -1, F1: 965}
    fmt.Println(v2)
    fmt.Println(v2, v3)
    if (!(!v1)) {
        fmt.Println(v5, v1)
    }
    var v12 float32 = 0.1
    v13 := GenS2{F0: 32767}
    var v14 int16 = 0
    fmt.Println(v1, v2)
    v3 = (v3 + v3)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
