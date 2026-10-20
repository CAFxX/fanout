package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
func genFn0() uint {
    if (576 <= 735) {
        var v6 int8 = 42
        _ = v6
    } else {
        var v7 float64 = 0.0
        _ = v7
    }
    var v8 int = 127
    _ = v8
    return 27
}
func genFn1(p0 uint16, p1 uint32, p2 string) uint16 {
    p1 -= p1
    var v9 int64 = -913
    for fi10 := 0; fi10 < 19; fi10++ {
        fmt.Println(p0, v9)
        _ = fi10
    }
    v9 += v9
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return (p0 / ((65535) | 1))
}

func main() {
    v1 := make(map[string]uint32)
    var v2 uint = 4294967295
    v3 := make(map[string]uint32)
    var v4 uint8 = 128
    var v5 uint8 = (v4 * 26)
    fmt.Println(v4, v2)
    var v11 float64 = 0.1
    fmt.Println(v2, v11, v4)
    if false {
        if false {
            v12 := GenS1{F0: 575}
            _ = v12
        } else {
            v13 := []int{32767, 2147483647}
            v13 = append(v13, 1)
            _ = v13
        }
    }
    v14 := GenS1{F0: 0}
    v15 := GenS1{F0: genFn1(uint16(v5), 0, "\n\t")}
    for fi16 := 0; fi16 < 4; fi16++ {
        var v17 int = fi16
        _ = fi16
        _ = v17
    }
    var v18 int32 = 2147483647
    v2 -= 419
    v3[("a" + "\n\t")] = 313
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v14
    _ = v15
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
