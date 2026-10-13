package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16; F2 uint64 }

func main() {
    var v1 float64 = 0.1
    v2 := []string{("" + "\n\t")}
    v3 := make(map[uint32]bool)
    var v4 int = 32767
    var v5 int16 = 1
    v6 := make(map[uint32]int)
    v1 = float64(0.0)
    var v7 int8 = -128
    var v8 int32 = 1
    var v9 string = ""
    var v10 uint = 0
    if (1 < 0) {
        v8 += v8
        var v11 float32 = 1.5
        _ = v11
    } else {
        if true {
            if (!true) {
                delete(v3, 2147483647)
            } else {
                _ = len(v3)
                delete(v3, 2147483647)
            }
            var v12 int64 = -9223372036854775808
            var v13 int32 = v8
            _ = v12
            _ = v13
        }
    }
    var v14 int8 = (v7 & (v7 & -54))
    for fi15 := 0; fi15 < 4; fi15++ {
        v16 := GenS1{F0: ("\n\t" + v9), F1: -32768, F2: 20}
        _ = fi15
        _ = v16
    }
    var v17 int64 = -1
    var v18 uint64 = 1
    for fi19 := 0; fi19 < 16; fi19++ {
        var v20 int16 = (v5 + (v5 ^ v5))
        v4 = fi19
        var v21 float32 = (float32(v1) / 1.0)
        _ = fi19
        _ = v20
        _ = v21
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v22 int8 = -26
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v14
    _ = v17
    _ = v18
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
