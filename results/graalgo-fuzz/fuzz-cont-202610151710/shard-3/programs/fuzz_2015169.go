package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
func genFn0() int32 {
    v7 := []int{32767}
    v8 := GenS1{F0: -32768}
    _ = v7
    _ = v8
    return -468
}
func genFn1(p0 uint16, p1 uint32) int16 {
    var v9 int64 = 1
    v10 := []int{0, -437, -9223372036854775808}
    _ = p0
    _ = p1
    _ = v10
    _ = v9
    return 595
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 string = ""
    v3 := []uint32{2147483647, 65535}
    v4 := make(map[string]bool)
    var v5 int8 = 42
    var v6 uint8 = 2
    fmt.Println(v6, v5, v5)
    var v11 int8 = -128
    var v12 uint16 = 42
    var v13 bool = (0 < 1)
    v12 = v12
    v11 = v5
    for fi14 := 0; fi14 < 7; fi14++ {
        var v15 int16 = genFn1((v12 | v12), 1)
        _ = fi14
        _ = v15
    }
    var v16 uint32 = 932
    v13 = (v2 <= v2)
    var v17 uint64 = 18446744073709551615
    var v18 uint64 = (^(v17))
    _ = len(v1)
    fmt.Println(v11)
    v13 = (!(1 < 2147483647))
    v19 := GenS1{F0: -1}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
