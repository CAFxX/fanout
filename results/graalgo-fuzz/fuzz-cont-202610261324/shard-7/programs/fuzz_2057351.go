package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8; F2 uint32 }
func genFn0(p0 string, p1 int, p2 uint64) int32 {
    var v4 int8 = 100
    var v5 int64 = 42
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return -1
}

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[uint32]uint32)
    v3 := make(map[uint32]bool)
    if false {
        v6 := GenS1{F0: 1, F1: 230, F2: 42}
        if true {
            v7 := GenS1{F0: -22, F1: 204, F2: 477}
            _ = v7
        } else {
            var v8 string = (("hello" + ("" + (("hello" + "foo bar") + ("" + ("" + ("\n\t" + ((("a" + "foo bar") + "hello") + "a"))))))) + ("a" + "\n\t"))
            v9 := GenS1{F0: -69, F1: 199, F2: 1}
            v8 = (v8 + (v8 + "a"))
            _ = v8
            _ = v9
        }
        var v10 bool = false
        _ = v10
        _ = v6
    } else {
        v11 := []int{-686, 4294967295}
        _ = v11
    }
    var v12 int8 = -1
    v12 -= v12
    for fi13 := 0; fi13 < 19; fi13++ {
        for fi14 := 0; fi14 < 16; fi14++ {
            v15 := GenS1{F0: 126, F1: 246, F2: 2147483648}
            _ = fi14
            _ = v15
        }
        var v16 int64 = -9223372036854775808
        _ = fi13
        _ = v16
    }
    switch "a" {
    case "":
        v12 = -23
    case "foo bar":
        var v17 int8 = (v12 - v12)
        _ = v17
    }
    v12 -= v12
    _ = len(v3)
    if ((true && (!true)) || (false || true)) {
        var v18 int = 177
        var v19 int8 = 127
        _ = v18
        _ = v19
    }
    fmt.Println(v12)
    if true {
        v20 := GenS1{F0: -86, F1: 2, F2: 754}
        _ = v20
    } else {
        v21 := []int{127, 6, 32767, 2147483647}
        var v22 int16 = 1
        fmt.Println(v22, v12)
        _ = v21
        _ = v22
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v12
    _ = v2
    _ = v3
}
