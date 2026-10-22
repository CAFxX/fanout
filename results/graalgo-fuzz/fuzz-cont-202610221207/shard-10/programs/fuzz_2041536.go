package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int }

func main() {
    v1 := []string{}
    v2 := []string{"hello"}
    v3 := make(map[int]string)
    v4 := make(map[string]string)
    v5 := []int{771, -741}
    v6 := make(map[uint32]string)
    v7 := GenS1{F0: 0, F1: -858}
    delete(v4, (("hello" + "\n\t") + "a"))
    var v8 bool = false
    fmt.Println(v8)
    v8 = (!false)
    if len(v2) > 0 {
        v2[0] = "\n\t"
    }
    v8 = (0 != 0)
    v8 = (!(818 <= 0))
    if (("" + ((("hello" + "hello") + "") + ("" + "a"))) <= "foo bar") {
        var v9 int64 = 42
        v8 = v8
        _ = v9
    }
    if v8 {
        v10 := GenS1{F0: 255, F1: 42}
        v8 = v8
        _ = v10
    }
    v8 = (false && (9223372036854775807 <= 128))
    var v11 float64 = 0.0
    if ("a" >= "\n\t") {
        var v12 uint16 = 32767
        var v13 uint16 = (v12 & (v12 - v12))
        _ = v12
        _ = v13
    }
    var v14 int16 = 42
    v15 := GenS1{F0: 255, F1: 32768}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
