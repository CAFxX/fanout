package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16 }
func genFn0(p0 string) uint {
    v6 := []int{-9223372036854775808}
    var v7 uint16 = 856
    v7 += v7
    _ = p0
    _ = v6
    _ = v7
    return 255
}

func main() {
    v1 := make(map[string]string)
    var v2 uint32 = 65535
    var v3 int16 = -680
    var v4 float64 = float64(1e10)
    v5 := make(map[int]string)
    v8 := GenS1{F0: 2, F1: (v3 | v3)}
    v2 = v2
    v2 = v2
    v9 := GenS1{F0: 32767, F1: (v3 + (-(v3)))}
    var v10 int32 = -2147483648
    for fi11 := 0; fi11 < 6; fi11++ {
        var v12 float32 = 1e10
        v10 += (v10 + v10)
        _ = fi11
        _ = v12
    }
    v10 += v10
    var v13 string = "a"
    if true {
        var v14 int8 = 0
        for fi15 := 0; fi15 < 9; fi15++ {
            v2 -= v2
            _ = fi15
        }
        _ = v14
    }
    var v16 uint = uint(v2)
    fmt.Println(v16)
    for fi17 := 0; fi17 < 18; fi17++ {
        v13 = v13
        _ = fi17
    }
    v18 := []int{-1, 32767, 4294967295, 9223372036854775807}
    v4 -= (v4 / float64(1e10))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v13
    _ = v16
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
