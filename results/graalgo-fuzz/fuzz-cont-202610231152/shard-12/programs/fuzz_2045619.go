package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
func genFn0(p0 uint) bool {
    switch 84 {
    case 46:
        v6 := []int{32768, 275, 127}
        _ = v6
    }
    var v7 uint16 = 763
    v8 := []int{128}
    if (!("foo bar" >= "a")) {
        var v9 int32 = 42
        _ = v9
    } else {
        fmt.Println(v7)
    }
    if len(v8) > 0 {
        v8[0] = 4294967295
    }
    _ = p0
    _ = v7
    _ = v8
    return true
}
func genFn1(p0 bool, p1 int16, p2 int16) uint8 {
    var v10 uint8 = 128
    fmt.Println(p0, p0)
    v11 := []int{9223372036854775807, 142, 32768}
    fmt.Println(p2, p0, p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    return 128
}

func main() {
    var v1 float64 = (1e10 / 1.5)
    var v2 uint32 = 4294967295
    v3 := []string{}
    var v4 bool = (!true)
    v5 := make(map[int]string)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v3 = append(v3, (("a" + "hello") + "a"))
    var v12 uint32 = (^(v2))
    for fi13 := 0; fi13 < 16; fi13++ {
        v12 -= 0
        _ = len(v3)
        _ = fi13
    }
    fmt.Println(v2, v12, v2)
    v14 := GenS1{F0: 4294967295}
    v15 := GenS1{F0: 4294967295}
    if (float32(1.5) < 3.14159) {
        var v16 int = 0
        _ = v16
    }
    v1 = v1
    var v17 uint16 = 310
    v18 := GenS1{F0: 0}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v12
    _ = v14
    _ = v15
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
