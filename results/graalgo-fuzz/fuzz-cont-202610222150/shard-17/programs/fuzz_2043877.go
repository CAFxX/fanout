package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64 }
func genFn0() uint32 {
    v4 := GenS1{F0: 1, F1: -9223372036854775808}
    var v5 string = (("" + ("\n\t" + "hello")) + "hello")
    v6 := GenS1{F0: -753, F1: -137}
    v7 := []int{277}
    var v8 int32 = -898
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return 145
}

func main() {
    v1 := []uint32{1, 2147483648}
    v2 := make(map[int]string)
    var v3 float32 = 0.0
    for fi9 := 0; fi9 < 18; fi9++ {
        v3 -= float32((3.14159 / ((-2.25 + -2.25) * 1.5)))
        _ = fi9
    }
    _ = len(v1)
    v3 += 3.14159
    for fi10 := 0; fi10 < 9; fi10++ {
        var v11 int = fi10
        if len(v1) > 0 {
            v1[0] = 1
        }
        _ = fi10
        _ = v11
    }
    var v12 int = -2147483648
    v13 := GenS1{F0: -896, F1: 1}
    v3 = v3
    v14 := GenS1{F0: -7, F1: 42}
    if len(v1) > 0 {
        v1[0] = 65535
    }
    v15 := GenS1{F0: 0, F1: -894}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
}
