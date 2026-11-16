package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
func genFn0(p0 uint32) int8 {
    var v7 int8 = 1
    v8 := []int{}
    _ = p0
    _ = v7
    _ = v8
    return (v7 + v7)
}
func genFn1(p0 int8, p1 string) uint {
    v9 := []int{-18, -946}
    var v10 bool = false
    _ = p0
    _ = p1
    _ = v10
    _ = v9
    return 255
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint64 = 18446744073709551615
    var v3 string = ""
    var v4 int32 = 2147483647
    v5 := make(map[string]bool)
    v6 := []uint32{2147483647, 2147483648, 0}
    fmt.Println(v4, v2)
    v11 := v1[1]
    var v12 uint16 = 42
    var v13 string = v3
    if false {
        for fi14 := 0; fi14 < 12; fi14++ {
            v12 -= (^(v12))
            var v15 uint32 = 0
            _ = fi14
            _ = v15
        }
    }
    var v16 uint = 65535
    v17 := GenS1{F0: (v12 >> ((v16) & 63))}
    if false {
        fmt.Println(v2, v16)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
