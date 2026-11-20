package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16; F2 int32 }
type GenS2 struct { F0 uint }
func genFn0(p0 uint8) int16 {
    switch 2147483647 {
    case 32767:
        v5 := GenS1{F0: (0 <= 71), F1: 0, F2: 125}
        _ = v5
    case 42:
        var v6 int8 = 1
        _ = v6
    case 583:
        p0 -= 1
    }
    var v7 int16 = -266
    _ = p0
    _ = v7
    return 1000
}

func main() {
    v1 := make(map[string]bool)
    var v2 int64 = -675
    var v3 uint64 = 1
    v4 := make(map[uint32]bool)
    v8 := GenS2{F0: 1}
    v9 := []int{32767, -2147483648, -2147483648}
    fmt.Println(v3)
    if len(v9) >= 2 {
        v9 = v9[1:len(v9)]
    }
    switch -686 {
    case 574:
        v3 += 0
    case 128:
        v3 += (^(v3))
    }
    var v10 uint32 = 4294967295
    v11 := GenS1{F0: ("" <= "hello"), F1: 461, F2: 42}
    if false {
        v12 := GenS2{F0: 717}
        var v13 string = "a"
        _ = v12
        _ = v13
    }
    fmt.Println(v10)
    _ = len(v9)
    var v14 float32 = 3.14159
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
