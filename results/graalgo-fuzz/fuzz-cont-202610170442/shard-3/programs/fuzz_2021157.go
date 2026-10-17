package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 float32 = 1e10
    var v3 int16 = -774
    v1[769] = true
    var v4 float32 = v2
    v5 := []int{934, -9223372036854775808, -316, 32768}
    var v6 int8 = 0
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v7 uint16 = 32767
    if ((v7 & 512) < v7) {
        v6 = int8(v7)
    }
    v7 = v7
    var v8 string = ""
    for fi9 := 0; fi9 < 12; fi9++ {
        if true {
            var v10 uint32 = 1
            _ = len(v5)
            var v11 uint16 = v7
            _ = v10
            _ = v11
        }
        v12 := GenS1{F0: 42}
        for fi13 := 0; fi13 < 14; fi13++ {
            v4 += float32((0.0 * ((float64(0.0) + 1e10) - 1.5)))
            _ = fi13
        }
        _ = fi9
        _ = v12
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
