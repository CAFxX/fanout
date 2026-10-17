package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8; F2 uint64 }
func genFn0(p0 bool, p1 uint8) string {
    var v7 uint = 7
    p0 = true
    v8 := GenS1{F0: 1, F1: p1, F2: 0}
    if p0 {
        for fi9 := 0; fi9 < 1; fi9++ {
            var v10 float32 = 0.1
            v11 := GenS1{F0: 842, F1: (p1 - (p1 << (((v7 % ((v7) | 1))) & 63))), F2: 2}
            _ = fi9
            _ = v10
            _ = v11
        }
        v12 := GenS1{F0: 137, F1: 79, F2: 4294967295}
        _ = v12
    } else {
        var v13 int8 = -1
        fmt.Println(v13, p0, v7)
        v14 := GenS1{F0: 1, F1: (^(p1)), F2: 42}
        _ = v13
        _ = v14
    }
    var v15 uint64 = 1
    _ = p0
    _ = p1
    _ = v15
    _ = v7
    _ = v8
    return ""
}

func main() {
    v1 := make(map[int]int)
    var v2 uint16 = 255
    v3 := make(map[uint32]int)
    var v4 uint = 4294967295
    v5 := []string{("" + "a"), "foo bar"}
    var v6 int16 = -1
    v16 := GenS1{F0: 0, F1: 127, F2: 67}
    fmt.Println(v6)
    if (-2147483648 >= 742) {
        var v17 float32 = ((1e10 + 0.1) / 0.1)
        fmt.Println(v2, v2, v6)
        _ = v17
    }
    fmt.Println(v4, v2)
    for fi18 := 0; fi18 < 3; fi18++ {
        v4 = (v4 << ((v4) & 63))
        _ = fi18
    }
    fmt.Println(v6, v6, v4)
    var v19 int8 = -1
    var v20 uint8 = 8
    var v21 int16 = v6
    var v22 int = 730
    v20 = v20
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
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
