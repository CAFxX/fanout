package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 bool; F1 int }
func genFn0(p0 bool) int32 {
    p0 = p0
    for fi6 := 0; fi6 < 20; fi6++ {
        var v7 int16 = -1
        v8 := GenS1{F0: 832}
        _ = fi6
        _ = v7
        _ = v8
    }
    _ = p0
    return 42
}
func genFn1(p0 string, p1 uint64) int {
    p1 -= (p1 % (((^(p1))) | 1))
    v9 := []int{-322}
    v10 := GenS1{F0: 496}
    fmt.Println(p1, p0)
    var v11 int = 354
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v9
    return v11
}

func main() {
    var v1 int32 = 0
    var v2 float64 = 3.14159
    var v3 int16 = 573
    v4 := make(map[string]string)
    var v5 uint32 = 0
    v12 := GenS1{F0: 42}
    var v13 int8 = -128
    fmt.Println(v1, v13)
    v14 := []int{}
    v15 := GenS1{F0: 334}
    v16 := GenS2{F0: ((4294967295 == 42) || false), F1: 127}
    fmt.Println(v3, v13)
    var v17 uint = 18446744073709551615
    fmt.Println(v5, v5)
    v17 = (^(v17))
    if true {
        v1 = (^(v1))
        var v18 float32 = 0.1
        _ = v18
    }
    var v19 int8 = -60
    v2 = (((v2 / float64(3.14159)) - (v2 / 3.14159)) - v2)
    for fi20 := 0; fi20 < 1; fi20++ {
        v5 += (v5 << ((41) & 63))
        v21 := v4[("\n\t" + (("hello" + ("a" + "")) + ""))]
        fmt.Println(fi20, v13, v2)
        _ = fi20
        _ = v21
    }
    v2 = (v2 * v2)
    switch "" {
    case "":
        var v22 uint32 = (v5 * v5)
        _ = v22
    case "hello":
        delete(v4, (("hello" + "hello") + (("hello" + "foo bar") + ("" + "hello"))))
    }
    v2 = (v2 + v2)
    v3 = (v3 / (((v3 + v3)) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
