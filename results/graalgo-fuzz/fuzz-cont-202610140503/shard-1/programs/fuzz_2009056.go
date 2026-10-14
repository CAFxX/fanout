package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint8 }
func genFn0(p0 int16) uint {
    p0 -= (p0 * (p0 | 42))
    switch "\n\t" {
    case "hello":
        fmt.Println(p0)
    case "foo bar":
        for fi6 := 0; fi6 < 3; fi6++ {
            var v7 uint = 2
            v8 := []int{fi6, 1, fi6}
            _ = fi6
            _ = v7
            _ = v8
        }
    default:
        p0 += p0
    }
    var v9 bool = false
    _ = p0
    _ = v9
    return 143
}

func main() {
    v1 := make(map[int]int)
    v2 := []int{}
    var v3 uint32 = 2147483647
    var v4 uint32 = 65535
    v5 := []int{32768, 367, 1}
    v4 = (v3 ^ 0)
    var v10 bool = (-128 < 42)
    fmt.Println(v10, v4, v4)
    var v11 float64 = 1e10
    for fi12 := 0; fi12 < 5; fi12++ {
        var v13 int32 = -825
        var v14 uint16 = 0
        _ = fi12
        _ = v13
        _ = v14
    }
    for fi15 := 0; fi15 < 15; fi15++ {
        v3 -= v3
        v4 = (v4 & v3)
        _ = fi15
    }
    for fi16 := 0; fi16 < 9; fi16++ {
        if v10 {
            var v17 int64 = 1
            v4 += v3
            v17 += v17
            _ = v17
        }
        fmt.Println(v4)
        _ = fi16
    }
    v5 = append(v5, -570)
    delete(v1, 95)
    if (32767 >= 32767) {
        v18 := GenS1{F0: (("hello" + ("\n\t" + "hello")) + ("" + "foo bar"))}
        _ = v18
    }
    if (true && v10) {
        var v19 int16 = -32768
        v20 := GenS2{F0: 255}
        _ = v19
        _ = v20
    }
    var v21 int64 = 42
    v1[-2147483648] = 128
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
