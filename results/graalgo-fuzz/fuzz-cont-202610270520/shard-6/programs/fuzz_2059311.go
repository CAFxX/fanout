package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint32; F1 int16; F2 uint32 }

func main() {
    var v1 uint8 = 2
    var v2 string = ""
    v3 := make(map[string]uint32)
    var v4 uint64 = 373
    var v5 bool = (false || (!((1 == -478) || false)))
    if v5 {
        v6 := GenS1{F0: "hello"}
        _ = v6
    }
    var v7 uint64 = v4
    for fi8 := 0; fi8 < 13; fi8++ {
        v2 = "foo bar"
        _ = fi8
    }
    v9 := v3[v2]
    var v10 int64 = 9223372036854775807
    for fi11 := 0; fi11 < 19; fi11++ {
        v4 += 2
        _ = fi11
    }
    fmt.Println(v1, v7, v1)
    if v5 {
        v1 -= 2
        v10 += 42
    }
    v12 := v3[v2]
    switch 42 {
    case 0:
        v13 := []int{-1, 128, 2147483647}
        _ = v13
    case 42:
        var v14 int8 = -1
        _ = v14
    case 9223372036854775807:
        for fi15 := 0; fi15 < 17; fi15++ {
            v16 := []int{fi15, -9223372036854775808}
            v4 = 1
            fmt.Println(v7)
            _ = fi15
            _ = v16
        }
    }
    v17 := GenS2{F0: v12, F1: 1000, F2: 2147483647}
    v18 := v3["\n\t"]
    var v19 float64 = 0.0
    var v20 string = ("hello" + v2)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v12
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v9
}
