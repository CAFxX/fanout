package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
func genFn0(p0 uint8) int16 {
    v7 := []int{42, 404, 127, 219}
    fmt.Println(p0)
    _ = p0
    _ = v7
    return 234
}

func main() {
    var v1 string = "hello"
    var v2 float32 = (0.1 - 0.0)
    var v3 float32 = v2
    var v4 uint8 = 69
    v5 := []uint32{0}
    v6 := []uint32{2147483647}
    var v8 uint = 882
    for fi9 := 0; fi9 < 5; fi9++ {
        var v10 int = fi9
        v10 -= 286
        _ = fi9
        _ = v10
    }
    v1 = v1
    v11 := GenS1{F0: (v8 % (((v8 ^ (v8 >> ((v8) & 63)))) | 1))}
    var v12 int = 0
    for fi13 := 0; fi13 < 18; fi13++ {
        fmt.Println(fi13, v3, v8)
        v4 = v4
        _ = fi13
    }
    v6 = append(v6, 65535)
    fmt.Println(v2)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
