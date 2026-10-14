package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16 }
func genFn0() uint64 {
    v7 := []int{9223372036854775807, 312}
    var v8 uint16 = 114
    v9 := GenS1{F0: 2147483647, F1: (v8 >> ((658) & 63))}
    _ = v7
    _ = v8
    _ = v9
    return 42
}
func genFn1(p0 string, p1 uint64) int {
    v10 := GenS1{F0: -1, F1: 1}
    fmt.Println(p1, p1)
    p1 += p1
    _ = p0
    _ = p1
    _ = v10
    return 127
}

func main() {
    var v1 int = 4294967295
    var v2 int16 = -32768
    var v3 int8 = 100
    var v4 uint = 563
    var v5 uint = (v4 & 2)
    var v6 float32 = 0.1
    v11 := []int{188, (v1 - (v1 ^ (^(v1)))), v1, -49}
    fmt.Println(v5, v5, v1)
    for fi12 := 0; fi12 < 17; fi12++ {
        var v13 uint8 = 128
        var v14 uint16 = 2
        _ = fi12
        _ = v13
        _ = v14
    }
    v4 += v4
    if (2 <= genFn0()) {
        var v15 uint8 = 42
        var v16 int32 = 1
        fmt.Println(v3)
        _ = v15
        _ = v16
    }
    var v17 bool = true
    for fi18 := 0; fi18 < 4; fi18++ {
        if true {
            if v17 {
                switch (v1 | genFn1(("a" + "\n\t"), 1)) {
                case 4294967295:
                    var v19 uint32 = 21
                    _ = v19
                case -1:
                    fmt.Println(v6)
                case 42:
                    v11 = append(v11, fi18)
                default:
                    if len(v11) > 0 {
                        v11[0] = v1
                    }
                }
            }
        } else {
            var v20 uint16 = 1
            fmt.Println(v4, v17)
            _ = v20
        }
        _ = fi18
    }
    var v21 int8 = v3
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v17
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
