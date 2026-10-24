package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
func genFn0(p0 int16, p1 int32) uint {
    v5 := GenS1{F0: ("a" + ("foo bar" + "a"))}
    v6 := []int{32768, -9223372036854775808, -266, -1}
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 348
}

func main() {
    v1 := []int{}
    var v2 int32 = -2147483648
    v3 := make(map[uint32]uint32)
    var v4 float64 = 1.5
    delete(v3, 533)
    fmt.Println(v4, v4)
    var v7 int8 = -1
    var v8 string = (((("" + "hello") + "foo bar") + "\n\t") + ("foo bar" + "a"))
    for fi9 := 0; fi9 < 11; fi9++ {
        v7 = (v7 | (v7 | (v7 << ((2) & 63))))
        _ = fi9
    }
    var v10 int32 = v2
    if (!true) {
        fmt.Println(v10, v2, v2)
    }
    var v11 uint64 = 4294967295
    v10 += v2
    fmt.Println(v4)
    v4 -= v4
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
