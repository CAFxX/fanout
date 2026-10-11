package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool; F2 uint }
func genFn0() uint8 {
    if (("a" + ("hello" + ("hello" + "foo bar"))) < "a") {
        for fi4 := 0; fi4 < 15; fi4++ {
            var v5 uint64 = 18446744073709551615
            var v6 uint = 4294967295
            _ = fi4
            _ = v5
            _ = v6
        }
        var v7 bool = true
        _ = v7
    }
    var v8 int32 = 640
    _ = v8
    return 2
}
func genFn1(p0 uint64, p1 int) string {
    var v9 uint32 = 578
    var v10 int32 = 239
    var v11 string = "hello"
    v11 = ""
    v12 := GenS1{F0: (v9 | v9), F1: false, F2: 4294967295}
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v12
    _ = v9
    return v11
}

func main() {
    v1 := make(map[int]int)
    v2 := []uint32{}
    v3 := make(map[string]uint32)
    var v13 uint32 = 1
    v13 = (^(v13))
    v14 := GenS1{F0: (v13 >> ((339) & 63)), F1: true, F2: 778}
    if len(v2) > 0 {
        v2[0] = v13
    }
    var v15 int32 = -1
    var v16 float32 = 0.1
    v16 = (0.1 / (1e10 / (v16 * v16)))
    v13 = v13
    v17 := GenS1{F0: v13, F1: true, F2: 726}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
}
