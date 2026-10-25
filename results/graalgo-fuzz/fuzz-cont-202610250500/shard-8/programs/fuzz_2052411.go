package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
func genFn0(p0 int16, p1 int8, p2 uint) string {
    v6 := GenS1{F0: 2}
    fmt.Println(p1, p1)
    v7 := []int{-1}
    p0 = (p0 ^ p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return "a"
}
func genFn1(p0 int32, p1 int) uint32 {
    var v8 string = ("\n\t" + "hello")
    v9 := []int{}
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return 211
}

func main() {
    var v1 uint32 = 217
    v2 := []string{}
    var v3 int = -1
    var v4 uint64 = 18446744073709551615
    v5 := make(map[uint32]int)
    var v10 int16 = -32768
    var v11 int8 = 100
    v10 = v10
    v12 := GenS1{F0: 4294967295}
    var v13 float32 = (float32(-2.25) / 1.0)
    v11 -= (v11 + v11)
    var v14 int = -1
    var v15 int32 = -2147483648
    var v16 int64 = -9223372036854775808
    var v17 string = "\n\t"
    for fi18 := 0; fi18 < 20; fi18++ {
        var v19 int = (fi18 >> ((444) & 63))
        _ = fi18
        _ = v19
    }
    if len(v2) > 0 {
        v2[0] = v17
    }
    fmt.Println(v1, v17)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
