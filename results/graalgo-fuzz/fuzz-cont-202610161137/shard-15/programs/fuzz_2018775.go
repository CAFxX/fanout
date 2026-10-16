package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32; F2 int32 }
type GenS2 struct { F0 uint8; F1 string; F2 string }
func genFn0(p0 int8) bool {
    p0 -= p0
    var v4 int16 = 1000
    p0 += (p0 | p0)
    _ = p0
    _ = v4
    return true
}

func main() {
    v1 := make(map[int]string)
    var v2 int8 = -36
    var v3 string = ("hello" + "")
    v5 := GenS1{F0: "\n\t", F1: -625, F2: 2147483647}
    var v6 uint64 = 214
    for fi7 := 0; fi7 < 11; fi7++ {
        v8 := []int{fi7, fi7}
        _ = fi7
        _ = v8
    }
    fmt.Println(v3, v2)
    var v9 uint32 = 0
    var v10 bool = (234 < 1)
    v11 := GenS2{F0: 42, F1: ("hello" + v3), F2: "\n\t"}
    var v12 uint32 = v9
    v12 += (v12 * 503)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v5
    _ = v6
    _ = v9
}
