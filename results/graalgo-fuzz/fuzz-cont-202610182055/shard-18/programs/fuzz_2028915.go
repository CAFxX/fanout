package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16; F2 int }
type GenS2 struct { F0 uint16; F1 uint8 }

func main() {
    v1 := []int{}
    var v2 uint64 = 1
    v3 := make(map[int]string)
    var v4 float32 = (float32(0.0) / 1.0)
    var v5 int = 771
    var v6 uint8 = uint8(v2)
    var v7 int32 = 1
    v2 += (v2 ^ (v2 + (v2 >> ((18446744073709551615) & 63))))
    v7 += 42
    var v8 uint = 0
    var v9 uint64 = (^(v2))
    var v10 uint16 = 2
    for fi11 := 0; fi11 < 1; fi11++ {
        var v12 uint = uint(v2)
        _ = fi11
        _ = v12
    }
    if true {
        if ((v6 >= (v6 | 42)) || (true || false)) {
            for fi13 := 0; fi13 < 19; fi13++ {
                for fi14 := 0; fi14 < 3; fi14++ {
                    var v15 string = "foo bar"
                    _ = fi14
                    _ = v15
                }
                _ = fi13
            }
            var v16 uint8 = 74
            v17 := GenS1{F0: 210, F1: (v10 << ((uint(v6)) & 63)), F2: v5}
            _ = v16
            _ = v17
        }
    }
    v18 := GenS2{F0: (v10 ^ 992), F1: (v6 / ((uint8(v8)) | 1))}
    v3[v5] = ("" + "foo bar")
    v9 = 2
    var v19 int = v5
    if true {
        var v20 int32 = v7
        var v21 int32 = 42
        _ = v20
        _ = v21
    }
    v8 += (v8 - (v8 % (((^(v8))) | 1)))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
