package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint; F2 uint32 }
type GenS2 struct { F0 uint }
func genFn0(p0 int16) int64 {
    var v6 uint32 = 4294967295
    v7 := []int{32768}
    var v8 int8 = -128
    v8 += 100
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    return -9223372036854775808
}

func main() {
    v1 := []string{""}
    v2 := []int{32768, 32767, -9223372036854775808, -2147483648}
    var v3 uint32 = 65535
    v4 := make(map[string]int)
    v5 := []uint32{188, v3, v3}
    var v9 float64 = ((((1.5 * 1e10) + -2.25) + ((-2.25 - 0.0) + 0.1)) / 1e10)
    v10 := GenS1{F0: (^(v3)), F1: 241, F2: 65535}
    v11 := GenS2{F0: 226}
    var v12 int = 281
    var v13 uint8 = 2
    if (!true) {
        if len(v1) > 0 {
            v1[0] = ""
        }
        v14 := GenS2{F0: 18446744073709551615}
        _ = v14
    } else {
        var v15 float64 = 1.5
        var v16 int16 = -32768
        var v17 string = ("" + "foo bar")
        _ = v15
        _ = v16
        _ = v17
    }
    var v18 uint16 = 252
    v3 -= 360
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
