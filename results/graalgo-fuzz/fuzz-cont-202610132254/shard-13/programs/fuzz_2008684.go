package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64; F2 bool }
func genFn0() int16 {
    var v4 int = 32767
    var v5 int16 = -273
    v4 += v4
    if (!false) {
        v5 -= v5
        v4 = (v4 & (v4 + (v4 ^ (v4 * v4))))
    }
    _ = v4
    _ = v5
    return v5
}
func genFn1(p0 uint16) uint16 {
    v6 := []int{}
    var v7 float32 = -2.25
    _ = p0
    _ = v6
    _ = v7
    return 544
}

func main() {
    v1 := make(map[string]uint32)
    var v2 int8 = 0
    var v3 string = "foo bar"
    fmt.Println(v2, v3)
    v3 = ("" + ((v3 + (("" + v3) + (v3 + ("foo bar" + v3)))) + v3))
    v3 = v3
    v8 := GenS1{F0: (!false), F1: 42, F2: (float64((1e10 / 1e10)) == 1e10)}
    var v9 uint64 = 1
    var v10 float32 = 0.0
    var v11 float32 = 1.5
    v12 := GenS1{F0: true, F1: (^(v9)), F2: (9223372036854775807 >= 9223372036854775807)}
    v10 = v11
    var v13 uint8 = 0
    v14 := GenS1{F0: ((32767 >= 1000) && false), F1: (v9 * (v9 * 2)), F2: (-572 == 1)}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
