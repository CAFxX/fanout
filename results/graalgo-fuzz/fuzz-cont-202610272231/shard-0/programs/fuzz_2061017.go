package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64 }
func genFn0(p0 int, p1 uint8) int16 {
    v7 := GenS1{F0: 0, F1: 42}
    v8 := GenS1{F0: 4294967295, F1: 9223372036854775807}
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return -952
}

func main() {
    var v1 bool = false
    v2 := []string{("a" + "hello"), "a", "a", ("hello" + ("" + "hello"))}
    var v3 int64 = 42
    v4 := make(map[uint32]uint32)
    var v5 float32 = -2.25
    v6 := make(map[uint32]uint32)
    if true {
        var v9 string = ("a" + ("foo bar" + "hello"))
        _ = v9
    }
    var v10 uint32 = 749
    fmt.Println(v3, v1)
    if v1 {
        v3 -= v3
    } else {
        v3 = 0
        v10 = v10
        _ = len(v2)
    }
    for fi11 := 0; fi11 < 5; fi11++ {
        var v12 uint16 = 167
        _ = fi11
        _ = v12
    }
    _ = len(v4)
    fmt.Println(v10)
    var v13 int32 = 0
    for fi14 := 0; fi14 < 16; fi14++ {
        v10 -= (v10 % ((4294967295) | 1))
        v10 -= (^(v10))
        var v15 int16 = genFn0((fi14 / (((fi14 | fi14)) | 1)), 251)
        _ = fi14
        _ = v15
    }
    v13 -= -746
    var v16 int32 = v13
    for fi17 := 0; fi17 < 19; fi17++ {
        var v18 uint32 = v10
        _ = fi17
        _ = v18
    }
    v1 = (128 == 220)
    fmt.Println(v1, v16, v5)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v13
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
