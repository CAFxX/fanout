package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
func genFn0(p0 int8, p1 uint) uint64 {
    v4 := []int{-722}
    for fi5 := 0; fi5 < 14; fi5++ {
        var v6 float32 = 1e10
        _ = fi5
        _ = v6
    }
    if true {
        fmt.Println(p1, p0)
        if len(v4) > 0 {
            v4[0] = -9223372036854775808
        }
        p0 = (p0 * (p0 - 100))
    }
    _ = p0
    _ = p1
    _ = v4
    return 498
}
func genFn1() int16 {
    if true {
        v7 := []int{32767, 0}
        _ = v7
    }
    var v8 uint64 = 18446744073709551615
    v9 := GenS1{F0: 65535}
    _ = v8
    _ = v9
    return 1000
}

func main() {
    v1 := make(map[int]bool)
    var v2 uint32 = 4294967295
    var v3 uint = 851
    v2 = 145
    var v10 int64 = 1
    var v11 int64 = v10
    delete(v1, 158)
    v3 = 330
    if ((v3 - (v3 - v3)) >= v3) {
        var v12 uint = v3
        fmt.Println(v11, v12, v2)
        if true {
            var v13 bool = false
            v2 = v2
            _ = v13
        }
        _ = v12
    } else {
        var v14 int32 = 1
        _ = v14
    }
    v3 = (v3 / ((42) | 1))
    v15 := []int{32768, 511}
    for fi16 := 0; fi16 < 7; fi16++ {
        v17 := GenS1{F0: 1}
        _ = fi16
        _ = v17
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v15
    _ = v2
    _ = v3
}
