package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32; F2 uint16 }
func genFn0(p0 bool, p1 uint32, p2 uint64) string {
    for fi7 := 0; fi7 < 16; fi7++ {
        var v8 uint8 = 0
        var v9 uint8 = (v8 % (((v8 - (v8 >> ((2) & 63)))) | 1))
        v10 := []int{fi7}
        _ = fi7
        _ = v10
        _ = v8
        _ = v9
    }
    var v11 float64 = (1e10 - 1.5)
    v12 := GenS1{F0: -1, F1: p1, F2: 255}
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    return "foo bar"
}
func genFn1(p0 uint64) int32 {
    v13 := []int{-9223372036854775808}
    var v14 int16 = 42
    v14 = (v14 >> ((0) & 63))
    _ = p0
    _ = v13
    _ = v14
    return 67
}

func main() {
    v1 := []uint32{2147483648}
    var v2 float32 = (1.5 * 3.14159)
    v3 := []string{"", ("foo bar" + "\n\t")}
    var v4 uint32 = 2
    var v5 uint32 = v4
    v6 := []int{-1}
    v4 -= (^(v4))
    var v15 uint32 = v5
    var v16 uint32 = (v4 | (v15 | (v15 / (((v15 % (((v5 / ((v15) | 1))) | 1))) | 1))))
    fmt.Println(v4, v15, v5)
    var v17 int64 = 1
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println(v15, v17)
    fmt.Println(v4)
    fmt.Println(v16, v17)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
