package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32 }
type GenS2 struct { F0 int }
func genFn0(p0 uint16) int8 {
    p0 = (^(p0))
    v5 := []int{-1, 32767, -292, 2147483647}
    for fi6 := 0; fi6 < 13; fi6++ {
        p0 = 42
        _ = fi6
    }
    var v7 int = -1
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    _ = p0
    _ = v5
    _ = v7
    return 42
}

func main() {
    var v1 int8 = 100
    v2 := []uint32{0, 2147483647, 65535}
    var v3 int64 = -312
    var v4 int64 = v3
    switch "\n\t" {
    case "":
        v8 := GenS1{F0: "", F1: -632}
        _ = v8
    case "\n\t":
        v9 := GenS2{F0: 9223372036854775807}
        _ = v9
    }
    fmt.Println(v4, v1)
    v3 = v3
    var v10 uint16 = 0
    _ = len(v2)
    var v11 uint16 = v10
    var v12 uint8 = 255
    v13 := GenS1{F0: ("hello" + "a"), F1: 0}
    v10 -= v11
    fmt.Println(v12, v3, v12)
    v4 += (v3 | v3)
    v11 -= v10
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
}
