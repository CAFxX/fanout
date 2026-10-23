package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int; F1 uint16 }
type GenS3 struct { F0 uint; F1 int16 }
func genFn0(p0 uint8, p1 uint) uint8 {
    if false {
        p0 += p0
        p0 = p0
        var v6 uint32 = 2
        _ = v6
    }
    p0 = p0
    _ = p0
    _ = p1
    return 1
}

func main() {
    var v1 float64 = float64((1e10 / 0.1))
    var v2 uint16 = 809
    var v3 uint64 = 734
    var v4 bool = (88 == 255)
    var v5 int64 = 931
    for fi7 := 0; fi7 < 9; fi7++ {
        v8 := GenS1{F0: 318}
        _ = fi7
        _ = v8
    }
    v9 := []int{2147483647, -1, 2147483647}
    fmt.Println(v5, v4, v2)
    var v10 int = 32767
    v9 = append(v9, v10)
    v11 := GenS2{F0: int(v5), F1: 42}
    v2 -= v2
    v3 = 1
    fmt.Println(v3, v1, v3)
    v4 = v4
    var v12 uint64 = (v3 % ((v3) | 1))
    fmt.Println(v10)
    for fi13 := 0; fi13 < 19; fi13++ {
        v5 -= v5
        fmt.Println(v2, v5, fi13)
        _ = fi13
    }
    v12 = (v3 + v3)
    switch ("a" + "a") {
    case "":
        v14 := GenS3{F0: uint(v3), F1: 1000}
        _ = v14
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
