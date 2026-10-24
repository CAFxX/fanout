package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16 }
func genFn0(p0 bool, p1 int32) int8 {
    var v6 int = -2147483648
    fmt.Println(p0, v6, p0)
    var v7 bool = p0
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return -87
}

func main() {
    v1 := make(map[uint32]int)
    v2 := []int{}
    v3 := make(map[int]bool)
    v4 := make(map[string]bool)
    v5 := make(map[string]bool)
    var v8 uint8 = 6
    if false {
        _ = len(v2)
        for fi9 := 0; fi9 < 14; fi9++ {
            v8 += 220
            _ = fi9
        }
    }
    var v10 uint8 = 1
    v8 = v8
    for fi11 := 0; fi11 < 2; fi11++ {
        v8 -= 128
        fmt.Println(fi11, v8)
        v10 = 128
        _ = fi11
    }
    var v12 uint16 = 339
    if true {
        v10 -= (v10 << ((255) & 63))
    }
    var v13 int16 = 42
    v12 = v12
    v10 -= v10
    switch ("\n\t" + "foo bar") {
    case "":
        v14 := GenS1{F0: 62, F1: v12}
        _ = v14
    case "foo bar":
        switch 565 {
        case 9223372036854775807:
            v13 = v13
        }
    case "a":
        fmt.Println(v12)
    }
    v15 := GenS1{F0: genFn0((0.1 == 3.14159), 1), F1: 2}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
