package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64; F2 int16 }
type GenS2 struct { F0 uint64; F1 uint8 }
func genFn0(p0 bool) int {
    p0 = (340 >= 42)
    v4 := []int{}
    fmt.Println(p0)
    v5 := GenS1{F0: 0, F1: 1, F2: 42}
    v6 := GenS1{F0: 42, F1: 635, F2: 32767}
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    return 32767
}

func main() {
    var v1 int16 = -953
    var v2 int = 0
    var v3 int64 = 760
    fmt.Println(v2, v3)
    v7 := []int{v2, v2}
    if true {
        var v8 float32 = 1.5
        _ = v8
    }
    var v9 uint16 = 982
    for fi10 := 0; fi10 < 15; fi10++ {
        v2 = (v2 ^ (fi10 ^ genFn0(false)))
        _ = fi10
    }
    if len(v7) > 0 {
        v7[0] = genFn0((v9 >= (v9 / ((255) | 1))))
    }
    v11 := GenS1{F0: 753, F1: -297, F2: (v1 << ((629) & 63))}
    for fi12 := 0; fi12 < 8; fi12++ {
        v13 := GenS1{F0: 2147483647, F1: (v3 ^ v3), F2: v1}
        _ = fi12
        _ = v13
    }
    switch "foo bar" {
    case "hello":
        var v14 uint = 255
        _ = v14
    case "foo bar":
        v15 := GenS2{F0: 0, F1: 0}
        _ = v15
    }
    var v16 string = ""
    switch v2 {
    case -348:
        v9 -= v9
    case -305:
        var v17 int16 = (v1 ^ 1)
        _ = v17
    case 32768:
        var v18 bool = ((-1 < 79) && (42 <= 255))
        _ = v18
    }
    var v19 uint = uint(v9)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v16
    _ = v19
    _ = v2
    _ = v3
    _ = v7
    _ = v9
}
