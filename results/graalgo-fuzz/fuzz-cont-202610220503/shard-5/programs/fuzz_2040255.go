package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32; F2 uint8 }
type GenS2 struct { F0 uint64 }
func genFn0(p0 string, p1 uint32) int64 {
    var v7 bool = false
    var v8 uint8 = 42
    var v9 float32 = 1e10
    v8 -= (v8 >> ((42) & 63))
    for fi10 := 0; fi10 < 16; fi10++ {
        var v11 int = 424
        v12 := []int{v11, (fi10 + v11)}
        fmt.Println(v7, p0, v8)
        _ = fi10
        _ = v11
        _ = v12
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return 1
}
func genFn1() int {
    var v13 float64 = 1e10
    v14 := []int{0, -330}
    _ = v13
    _ = v14
    return 9223372036854775807
}

func main() {
    v1 := []string{"foo bar", "\n\t", "", "foo bar"}
    var v2 uint32 = 4294967295
    var v3 float32 = (float32(float64(float32(-2.25))) / (3.14159 + 1e10))
    v4 := make(map[string]bool)
    var v5 int8 = 42
    var v6 int = 378
    delete(v4, ("" + ("hello" + "")))
    _ = len(v4)
    v15 := GenS1{F0: "foo bar", F1: v2, F2: 2}
    switch (v6 ^ (v6 & v6)) {
    case -674:
        var v16 int16 = 0
        _ = v16
    default:
        v5 -= v5
    }
    v6 -= v6
    var v17 uint64 = 42
    v6 = genFn1()
    v18 := GenS2{F0: (v17 >> ((713) & 63))}
    if true {
        var v19 uint16 = 32767
        fmt.Println(v6)
        _ = v19
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
