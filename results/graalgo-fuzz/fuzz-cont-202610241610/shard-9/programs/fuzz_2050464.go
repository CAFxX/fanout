package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16 }
func genFn0() uint32 {
    var v6 string = "hello"
    var v7 uint16 = 32767
    fmt.Println(v7)
    var v8 int32 = -745
    _ = v6
    _ = v7
    _ = v8
    return 982
}
func genFn1(p0 uint32) bool {
    v9 := GenS1{F0: 412, F1: 747}
    v10 := []int{int(p0), 127, -885}
    var v11 uint64 = 1
    _ = p0
    _ = v10
    _ = v11
    _ = v9
    return (true || (100 > 6))
}

func main() {
    var v1 uint32 = 607
    var v2 uint16 = 65535
    var v3 int32 = 110
    v4 := []uint32{v1}
    v5 := make(map[string]int)
    if false {
        fmt.Println(v2, v1)
        fmt.Println(v3, v2, v1)
    }
    var v12 int = 853
    for fi13 := 0; fi13 < 15; fi13++ {
        switch "hello" {
        case "":
            v1 -= (v1 % (((^(v1))) | 1))
        }
        _ = fi13
    }
    v3 += (-(v3))
    fmt.Println(v3)
    var v14 uint8 = 17
    v1 -= v1
    fmt.Println(v1, v3)
    fmt.Println(v12)
    v14 += 2
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v12
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
