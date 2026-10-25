package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32 }
type GenS2 struct { F0 int8; F1 uint; F2 int8 }
func genFn0(p0 bool, p1 uint, p2 int8) uint16 {
    var v5 string = ("a" + "foo bar")
    p0 = (v5 == v5)
    for fi6 := 0; fi6 < 16; fi6++ {
        fmt.Println(fi6, fi6)
        _ = fi6
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 42
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint16 = 300
    var v3 string = "foo bar"
    var v4 int64 = 422
    for fi7 := 0; fi7 < 14; fi7++ {
        var v8 uint64 = 263
        _ = fi7
        _ = v8
    }
    fmt.Println(v2)
    for fi9 := 0; fi9 < 8; fi9++ {
        v3 = "foo bar"
        _ = fi9
    }
    v10 := GenS1{F0: 42, F1: 4294967295}
    v11 := GenS1{F0: 65535, F1: 2}
    for fi12 := 0; fi12 < 3; fi12++ {
        v1[65535] = true
        _ = fi12
    }
    v13 := GenS2{F0: -61, F1: 2, F2: 100}
    v14 := GenS2{F0: 100, F1: 4294967295, F2: 100}
    var v15 uint8 = 128
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
}
