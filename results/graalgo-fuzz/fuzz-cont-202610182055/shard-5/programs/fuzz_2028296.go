package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8 }
func genFn0(p0 int, p1 int8, p2 bool) int {
    var v6 int64 = -624
    v7 := []int{}
    for fi8 := 0; fi8 < 2; fi8++ {
        var v9 uint8 = 146
        p2 = true
        fmt.Println(v6, v9, v6)
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return p0
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 bool = true
    var v3 uint = 65535
    var v4 float32 = float32(-2.25)
    var v5 uint32 = 65535
    var v10 bool = (false && v2)
    v11 := GenS1{F0: (("hello" + "") + ""), F1: -128}
    for fi12 := 0; fi12 < 1; fi12++ {
        v13 := []int{fi12, fi12, fi12, (fi12 - (^(fi12)))}
        fmt.Println(fi12, v10)
        v2 = false
        _ = fi12
        _ = v13
    }
    if true {
        v1[(^(v5))] = 4294967295
        var v14 bool = true
        _ = v14
    }
    for fi15 := 0; fi15 < 5; fi15++ {
        v5 = (v5 - (v5 ^ 4294967295))
        _ = fi15
    }
    var v16 int16 = 0
    var v17 uint8 = 1
    var v18 int = genFn0(-35, -122, (v17 <= v17))
    var v19 uint8 = uint8(v5)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
