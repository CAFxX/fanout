package main

import (
    "fmt"
)


func main() {
    var v1 int16 = -172
    var v2 float64 = 1.5
    var v3 int64 = -965
    v4 := []uint32{2147483648}
    v5 := make(map[int]uint32)
    v6 := []int{32768, 32767, 2147483647, -376}
    _ = len(v5)
    v1 = (v1 << ((868) & 63))
    fmt.Println(v1, v2)
    if (!false) {
        var v7 float64 = v2
        _ = v7
    }
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    var v8 int = 9223372036854775807
    _ = len(v5)
    fmt.Println(v3, v1, v3)
    fmt.Println(v1)
    for fi9 := 0; fi9 < 20; fi9++ {
        v10 := v5[fi9]
        _ = fi9
        _ = v10
    }
    v8 = (^(v8))
    if (!false) {
        var v11 bool = false
        var v12 int = (-(v8))
        _ = v11
        _ = v12
    }
    v2 = v2
    for fi13 := 0; fi13 < 9; fi13++ {
        var v14 int32 = -1
        for fi15 := 0; fi15 < 13; fi15++ {
            var v16 int = 711
            v8 += (v8 & 1)
            if true {
                v14 = (-(v14))
            }
            _ = fi15
            _ = v16
        }
        _ = fi13
        _ = v14
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
