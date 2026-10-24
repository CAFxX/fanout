package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint32; F1 int8 }
func genFn0() uint8 {
    var v5 int32 = 42
    v6 := GenS1{F0: 241}
    v7 := GenS1{F0: -32768}
    _ = v5
    _ = v6
    _ = v7
    return 181
}
func genFn1() uint32 {
    var v8 uint8 = genFn0()
    var v9 string = (((("\n\t" + "\n\t") + "\n\t") + "hello") + ("hello" + ("hello" + (("a" + "") + "a"))))
    var v10 int = -947
    var v11 uint64 = 0
    _ = v10
    _ = v11
    _ = v8
    _ = v9
    return 240
}

func main() {
    var v1 uint64 = 18446744073709551615
    v2 := make(map[string]string)
    var v3 int64 = -952
    v4 := make(map[uint32]bool)
    v12 := GenS2{F0: 2147483648, F1: -128}
    v13 := GenS2{F0: 0, F1: -1}
    fmt.Println(v3)
    fmt.Println(v3, v3)
    v14 := []int{-1, 128}
    v1 -= (v1 | (v1 * v1))
    _ = len(v2)
    fmt.Println(v1, v1)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
}
