package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16 }
func genFn0(p0 int32) uint64 {
    var v5 float32 = (float32((1.5 * ((0.0 - (0.0 + -2.25)) / float64((1e10 / float32(float64((3.14159 / 1.5)))))))) * 1.5)
    for fi6 := 0; fi6 < 4; fi6++ {
        v5 += (-2.25 + float32(0.1))
        _ = fi6
    }
    v5 -= (((0.0 * v5) + 0.0) / ((((v5 / -2.25) * 1e10) + -2.25) + float32((1.5 * (0.0 - 1e10)))))
    _ = p0
    _ = v5
    return 235
}

func main() {
    var v1 int16 = 838
    v2 := make(map[string]bool)
    var v3 uint8 = 127
    var v4 uint32 = 42
    v4 -= v4
    var v7 uint16 = 32767
    v8 := []int{32767}
    if (128 >= (v3 >> ((2) & 63))) {
        v4 = v4
        var v9 int = -748
        v7 += (v7 % ((v7) | 1))
        _ = v9
    }
    v10 := GenS1{F0: (v1 / (((v1 | v1)) | 1)), F1: v7}
    _ = len(v8)
    if (("a" >= (("foo bar" + "\n\t") + "hello")) && false) {
        v4 = (v4 & (v4 % (((^(v4))) | 1)))
    } else {
        var v11 int8 = 1
        _ = len(v8)
        _ = v11
    }
    v12 := GenS1{F0: v1, F1: v7}
    fmt.Println(v4)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v10
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
