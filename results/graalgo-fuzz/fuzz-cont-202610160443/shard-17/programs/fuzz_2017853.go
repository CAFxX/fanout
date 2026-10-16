package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8 }
type GenS2 struct { F0 uint64 }
func genFn0() uint16 {
    var v7 bool = (1 > -1)
    v8 := []int{4294967295, -600, 451, -1}
    _ = v7
    _ = v8
    return 255
}
func genFn1(p0 string, p1 uint) int32 {
    fmt.Println(p0, p0)
    v9 := []int{127}
    _ = p0
    _ = p1
    _ = v9
    return 2147483647
}

func main() {
    var v1 uint64 = 4294967295
    v2 := make(map[uint32]uint32)
    var v3 uint64 = (^(v1))
    var v4 uint64 = 272
    var v5 float32 = -2.25
    v6 := make(map[string]bool)
    var v10 string = ("foo bar" + (("foo bar" + "foo bar") + "\n\t"))
    v11 := []int{}
    fmt.Println(v1, v10, v10)
    v12 := GenS1{F0: false, F1: 127}
    v11 = append(v11, -111)
    v13 := v2[672]
    v14 := GenS1{F0: (!false), F1: 1}
    var v15 int32 = 315
    fmt.Println(v10, v4)
    delete(v2, 2)
    fmt.Println(v3, v1)
    for fi16 := 0; fi16 < 11; fi16++ {
        var v17 uint64 = (^(v4))
        v5 += 1.5
        _ = fi16
        _ = v17
    }
    switch (v10 + "hello") {
    case "a":
        v2[v13] = v13
    case "hello":
        v5 += (v5 / float32(0.1))
    }
    v18 := GenS2{F0: 4294967295}
    _ = len(v6)
    v19 := GenS1{F0: (-1 <= 42), F1: -1}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
