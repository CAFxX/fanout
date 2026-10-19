package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32 }
func genFn0() uint16 {
    var v6 int = 2147483647
    v7 := []int{(v6 - (v6 - (v6 >> ((42) & 63)))), v6, v6}
    _ = v6
    _ = v7
    return 0
}
func genFn1() int16 {
    if true {
        switch ("foo bar" + ("\n\t" + (("\n\t" + "hello") + "a"))) {
        case "\n\t":
            v8 := GenS1{F0: 65535, F1: 2147483647}
            _ = v8
        }
        v9 := GenS1{F0: 404, F1: 4294967295}
        _ = v9
    }
    v10 := []int{32768, -914, 722, 32767}
    if len(v10) >= 2 {
        v10 = v10[1:len(v10)]
    }
    var v11 uint8 = 255
    _ = v10
    _ = v11
    return 1
}

func main() {
    v1 := make(map[string]bool)
    var v2 int8 = 40
    v3 := []uint32{288, 42}
    v4 := make(map[int]uint32)
    var v5 bool = true
    fmt.Println(v2)
    if len(v3) > 0 {
        v3[0] = 331
    }
    var v12 uint8 = 251
    v2 += v2
    v12 += v12
    v5 = (!v5)
    v5 = ((!(genFn1() <= genFn1())) || v5)
    v2 = (v2 | v2)
    for fi13 := 0; fi13 < 17; fi13++ {
        var v14 int16 = 32767
        v4[(fi13 - int(v2))] = 443
        _ = fi13
        _ = v14
    }
    fmt.Println(v12, v12, v5)
    var v15 int32 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
