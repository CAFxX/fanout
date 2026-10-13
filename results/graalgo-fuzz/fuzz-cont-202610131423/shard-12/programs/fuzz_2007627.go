package main

import (
    "fmt"
)


func main() {
    var v1 bool = (42 != 2)
    var v2 bool = v1
    v3 := make(map[string]int)
    var v4 uint8 = 255
    var v5 string = ""
    v4 = 255
    v2 = (!v2)
    v2 = v1
    v1 = true
    for fi6 := 0; fi6 < 8; fi6++ {
        var v7 uint16 = 967
        var v8 uint = 0
        _ = fi6
        _ = v7
        _ = v8
    }
    for fi9 := 0; fi9 < 10; fi9++ {
        v10 := v3[v5]
        var v11 float64 = (1e10 * 1e10)
        _ = fi9
        _ = v10
        _ = v11
    }
    fmt.Println(v5, v1, v1)
    var v12 int = -420
    v4 -= (^(v4))
    var v13 int64 = 1
    if v1 {
        for fi14 := 0; fi14 < 16; fi14++ {
            if v1 {
                v2 = v1
                v12 = 172
                var v15 int32 = -542
                _ = v15
            }
            _ = fi14
        }
    }
    for fi16 := 0; fi16 < 8; fi16++ {
        v17 := []int{(fi16 << ((740) & 63)), int(v4), (fi16 % (((fi16 ^ 39)) | 1)), int(v4)}
        var v18 uint32 = 65535
        var v19 int8 = 117
        _ = fi16
        _ = v17
        _ = v18
        _ = v19
    }
    v20 := []int{32768}
    v3[v5] = (^(v12))
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
