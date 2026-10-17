package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint64 }
func genFn0(p0 uint16, p1 bool, p2 int32) uint64 {
    p0 = 431
    var v7 float32 = -2.25
    var v8 string = (("a" + "") + "")
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 0
}
func genFn1(p0 uint32, p1 int16) uint {
    p0 = p0
    var v9 uint8 = 42
    p0 += (p0 | (p0 | (^(p0))))
    var v10 int16 = p1
    var v11 int64 = 0
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v9
    return 2
}

func main() {
    var v1 uint64 = 18446744073709551615
    var v2 uint8 = 2
    var v3 int64 = -1
    var v4 int8 = 0
    var v5 bool = ((^(v4)) > int8(v2))
    var v6 int64 = (v3 + (v3 % ((v3) | 1)))
    var v12 uint64 = genFn0(32767, (!(160 <= 20)), 2147483647)
    v13 := GenS1{F0: 375}
    for fi14 := 0; fi14 < 20; fi14++ {
        var v15 float64 = 0.1
        var v16 uint8 = v2
        _ = fi14
        _ = v15
        _ = v16
    }
    var v17 uint64 = (v12 * (^(v1)))
    v18 := GenS2{F0: (v1 * (v12 % ((547) | 1)))}
    var v19 bool = (true || v5)
    var v20 float64 = (1.5 * 1.5)
    v21 := []int{2147483647, 9223372036854775807, 128}
    if (true && v5) {
        v3 -= v6
    }
    v3 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
