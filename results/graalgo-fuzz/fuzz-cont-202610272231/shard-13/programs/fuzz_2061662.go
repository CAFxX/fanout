package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64; F2 uint32 }
type GenS2 struct { F0 bool }
func genFn0(p0 uint64) int64 {
    v6 := GenS1{F0: 203, F1: p0, F2: 1}
    for fi7 := 0; fi7 < 5; fi7++ {
        var v8 float32 = 0.0
        for fi9 := 0; fi9 < 19; fi9++ {
            if (0 < 101) {
                if (1 > 349) {
                    var v10 int16 = -32768
                    _ = v10
                }
                v8 += v8
            }
            v11 := GenS1{F0: 243, F1: 42, F2: 2}
            var v12 uint8 = 255
            _ = fi9
            _ = v11
            _ = v12
        }
        _ = fi7
        _ = v8
    }
    fmt.Println(p0)
    for fi13 := 0; fi13 < 2; fi13++ {
        v14 := GenS2{F0: true}
        var v15 float64 = (-2.25 / (((float64(1.5) / 0.1) * 0.1) - 3.14159))
        _ = fi13
        _ = v14
        _ = v15
    }
    _ = p0
    _ = v6
    return 42
}

func main() {
    var v1 int32 = -2147483648
    var v2 int16 = -1
    v3 := make(map[uint32]int)
    v4 := []string{"hello", "\n\t", "hello"}
    v5 := []string{("hello" + (("a" + (("" + "hello") + "hello")) + "a")), "", "hello", "foo bar"}
    var v16 uint32 = 362
    v1 += (^(v1))
    fmt.Println(v1, v2, v1)
    v17 := GenS1{F0: 255, F1: 2, F2: 0}
    var v18 uint64 = 4294967295
    v19 := GenS2{F0: (255 == 37)}
    if false {
        var v20 int8 = 127
        _ = v20
    }
    var v21 int8 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
