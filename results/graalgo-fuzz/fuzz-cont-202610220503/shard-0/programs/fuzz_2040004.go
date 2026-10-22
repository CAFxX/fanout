package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string }
type GenS2 struct { F0 uint16 }
func genFn0(p0 int16) uint {
    p0 = 1000
    p0 = 790
    _ = p0
    return 18446744073709551615
}

func main() {
    var v1 uint = 4294967295
    var v2 float64 = 0.0
    var v3 float32 = (1.5 * 0.0)
    var v4 uint64 = 725
    v5 := make(map[int]string)
    var v6 int64 = 9223372036854775807
    v7 := GenS1{F0: 1, F1: "hello"}
    delete(v5, 127)
    var v8 int32 = 42
    if (!(true || false)) {
        if true {
            for fi9 := 0; fi9 < 6; fi9++ {
                for fi10 := 0; fi10 < 15; fi10++ {
                    switch fi10 {
                    case 32768:
                        var v11 float64 = v2
                        _ = v11
                    }
                    _ = fi10
                }
                v6 += v6
                var v12 uint16 = 42
                _ = fi9
                _ = v12
            }
            v2 += 0.1
        }
        if false {
            fmt.Println(v6, v8, v4)
            v3 = ((0.1 - 0.0) + v3)
        }
    }
    v13 := []int{0, -1, -9223372036854775808, 543}
    v14 := GenS1{F0: 32767, F1: ""}
    _ = len(v13)
    v1 = (v1 + v1)
    v15 := GenS1{F0: 1000, F1: "a"}
    fmt.Println(v4, v3)
    if false {
        var v16 int64 = (v6 * (v6 << (((v1 / (((v1 / ((v1) | 1))) | 1))) & 63)))
        _ = v16
    } else {
        fmt.Println(v2, v4, v2)
    }
    switch 4294967295 {
    case 200:
        for fi17 := 0; fi17 < 19; fi17++ {
            _ = len(v13)
            _ = fi17
        }
    case 32768:
        v3 = (v3 * (float32(float64((v3 / v3))) + v3))
    case 42:
        fmt.Println(v4, v6)
    }
    v1 -= 0
    v18 := GenS2{F0: 2}
    v8 -= v8
    fmt.Println(v2, v6, v6)
    v19 := v5[0]
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
