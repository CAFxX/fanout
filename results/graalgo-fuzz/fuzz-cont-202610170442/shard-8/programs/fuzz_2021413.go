package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint32; F1 uint64 }
func genFn0(p0 int16) int {
    var v7 string = "a"
    for fi8 := 0; fi8 < 20; fi8++ {
        var v9 int8 = -29
        _ = fi8
        _ = v9
    }
    v7 = (v7 + v7)
    p0 = -425
    v7 = v7
    _ = p0
    _ = v7
    return -2147483648
}

func main() {
    var v1 int32 = -889
    v2 := []string{}
    v3 := make(map[int]bool)
    var v4 float32 = 1.5
    var v5 uint8 = 42
    v6 := []string{"hello"}
    for fi10 := 0; fi10 < 9; fi10++ {
        v11 := GenS1{F0: 1}
        _ = fi10
        _ = v11
    }
    if false {
        _ = len(v2)
        var v12 bool = (1 > 2)
        fmt.Println(v12, v12, v5)
        _ = v12
    } else {
        v5 += 64
    }
    v13 := GenS2{F0: 2147483647, F1: 18446744073709551615}
    v5 += 128
    var v14 float32 = v4
    delete(v3, 803)
    _ = len(v3)
    var v15 uint16 = 228
    v3[42] = false
    if true {
        var v16 int32 = (v1 + 42)
        _ = v16
    }
    var v17 string = "foo bar"
    var v18 int16 = 1000
    v4 -= v14
    if len(v6) > 0 {
        v6[0] = "foo bar"
    }
    v19 := GenS1{F0: -369}
    var v20 string = ""
    var v21 uint16 = (v15 ^ v15)
    var v22 bool = (v1 == v1)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
