package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32 }
type GenS2 struct { F0 int }
func genFn0() uint8 {
    v4 := []int{-9223372036854775808, 127}
    var v5 bool = false
    _ = v4
    _ = v5
    return 255
}

func main() {
    v1 := []string{"\n\t"}
    v2 := make(map[string]int)
    var v3 float32 = (float32(float64(((1e10 / ((0.1 * (1.5 + 0.0)) * (0.1 / 1.0))) / 1.0))) + (0.0 - -2.25))
    _ = len(v1)
    v3 = -2.25
    v6 := GenS1{F0: 878, F1: 247}
    var v7 float64 = float64(3.14159)
    var v8 bool = (v7 == v7)
    for fi9 := 0; fi9 < 6; fi9++ {
        if (v8 || (4294967295 < 65535)) {
            fmt.Println(v7, fi9, v8)
            for fi10 := 0; fi10 < 8; fi10++ {
                fmt.Println(v7, v7)
                if v8 {
                    v7 += v7
                } else {
                    var v11 int = fi9
                    if true {
                        v3 = (v3 / 1e10)
                    }
                    _ = v11
                }
                var v12 uint = 203
                _ = fi10
                _ = v12
            }
            v13 := GenS1{F0: 2, F1: 2}
            _ = v13
        } else {
            v3 -= v3
            v7 = 0.0
        }
        var v14 uint8 = 203
        fmt.Println(v8)
        _ = fi9
        _ = v14
    }
    delete(v2, ("a" + "foo bar"))
    fmt.Println(v8, v7)
    var v15 uint64 = 18446744073709551615
    var v16 uint8 = genFn0()
    v17 := GenS1{F0: 31, F1: 830}
    fmt.Println(v7)
    v18 := GenS2{F0: 32767}
    _ = len(v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v8
}
