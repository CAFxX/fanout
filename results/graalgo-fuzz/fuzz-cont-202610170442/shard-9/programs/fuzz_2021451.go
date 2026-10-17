package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
func genFn0() bool {
    var v5 uint64 = 4294967295
    v5 += (v5 % ((v5) | 1))
    _ = v5
    return true
}
func genFn1(p0 int32) uint32 {
    var v6 uint32 = 65535
    var v7 int8 = -1
    _ = p0
    _ = v6
    _ = v7
    return 441
}

func main() {
    v1 := make(map[string]uint32)
    var v2 uint8 = 42
    v3 := make(map[int]int)
    var v4 float64 = (0.1 / 0.1)
    v8 := v1[""]
    fmt.Println(v2, v8, v2)
    v4 = v4
    v4 = -2.25
    if false {
        v2 += 255
    } else {
        for fi9 := 0; fi9 < 11; fi9++ {
            v1["a"] = (v8 - (v8 ^ (v8 ^ v8)))
            switch fi9 {
            case 26:
                var v10 string = "a"
                _ = v10
            case 736:
                var v11 uint32 = (v8 * (v8 + 4294967295))
                _ = v11
            }
            _ = fi9
        }
        fmt.Println(v4)
    }
    v8 += (v8 * (v8 - v8))
    v12 := GenS1{F0: -9223372036854775808}
    v13 := GenS1{F0: 113}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}
