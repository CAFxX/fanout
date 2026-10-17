package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint }
type GenS2 struct { F0 int32 }
func genFn0(p0 uint64, p1 uint32) uint16 {
    var v5 uint32 = p1
    var v6 float32 = 1.5
    var v7 float64 = float64(v6)
    var v8 float64 = 1e10
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return 255
}
func genFn1() int32 {
    v9 := GenS1{F0: "", F1: 1000}
    v10 := GenS2{F0: 2147483647}
    var v11 string = ""
    fmt.Println(v11)
    v11 = ("" + "hello")
    _ = v10
    _ = v11
    _ = v9
    return -864
}

func main() {
    v1 := []int{42, 1}
    var v2 uint32 = 4294967295
    var v3 string = (("a" + ("\n\t" + (((("hello" + "\n\t") + "hello") + "\n\t") + ""))) + "hello")
    v4 := make(map[string]bool)
    var v12 string = v3
    v12 = v3
    var v13 float64 = -2.25
    switch 32768 {
    case 1:
        var v14 bool = false
        _ = v14
    }
    v15 := GenS2{F0: 2147483647}
    var v16 float64 = (v13 - v13)
    var v17 uint64 = 42
    v4[v12] = true
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v12
    _ = v13
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
