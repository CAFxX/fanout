package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
func genFn0() uint64 {
    v4 := []int{343, 454, -854}
    var v5 int8 = 1
    var v6 uint64 = 18446744073709551615
    v5 += (v5 * v5)
    if len(v4) > 0 {
        v4[0] = 9223372036854775807
    }
    _ = v4
    _ = v5
    _ = v6
    return (v6 ^ 454)
}
func genFn1() string {
    var v7 uint8 = 0
    var v8 int8 = 100
    v8 = v8
    fmt.Println(v7)
    _ = v7
    _ = v8
    return (("hello" + "") + "hello")
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[uint32]bool)
    v3 := []uint32{1}
    var v9 int64 = -694
    delete(v1, 42)
    v9 = v9
    fmt.Println(v9)
    v9 += 9223372036854775807
    var v10 uint = 4294967295
    var v11 int8 = 42
    v12 := GenS1{F0: 158}
    for fi13 := 0; fi13 < 11; fi13++ {
        if false {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            fmt.Println(v11, fi13)
        } else {
            var v14 uint16 = 204
            for fi15 := 0; fi15 < 18; fi15++ {
                _ = len(v3)
                _ = fi15
            }
            _ = v14
        }
        var v16 uint64 = 42
        _ = fi13
        _ = v16
    }
    v9 = (v9 % ((1) | 1))
    var v17 float64 = 3.14159
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v17
    _ = v2
    _ = v3
    _ = v9
}
