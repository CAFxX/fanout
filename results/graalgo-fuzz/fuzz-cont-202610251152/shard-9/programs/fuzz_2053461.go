package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }

func main() {
    var v1 int = -103
    v2 := []int{v1}
    v3 := make(map[int]string)
    v4 := []uint32{2, 0}
    var v5 uint16 = 0
    v6 := []string{}
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    var v7 int64 = 943
    v8 := v3[int(v7)]
    if (0 >= 87) {
        var v9 uint32 = 1
        _ = v9
    }
    v5 = v5
    fmt.Println(v5, v5)
    var v10 uint32 = 888
    var v11 bool = false
    if v11 {
        var v12 int8 = -96
        var v13 string = (v8 + "\n\t")
        var v14 string = v8
        _ = v12
        _ = v13
        _ = v14
    } else {
        var v15 int32 = -2147483648
        _ = v15
    }
    v16 := GenS1{F0: int16(v10)}
    if v11 {
        switch v1 {
        case 32767:
            v17 := GenS1{F0: 32767}
            _ = v17
        case -9223372036854775808:
            v18 := GenS1{F0: 165}
            _ = v18
        case 584:
            switch (v1 - v1) {
            case -717:
                var v19 uint32 = (v10 ^ 0)
                _ = v19
            case 32767:
                _ = len(v4)
            case 1:
                var v20 int64 = (v7 * (v7 & -583))
                _ = v20
            default:
                var v21 uint16 = v5
                _ = v21
            }
        }
    } else {
        var v22 uint16 = (^(v5))
        var v23 int16 = -32768
        _ = v22
        _ = v23
    }
    var v24 float64 = (0.1 * (-2.25 - 0.0))
    v25 := GenS1{F0: 1000}
    v1 += (v1 << ((uint(v10)) & 63))
    v26 := v3[(^(v1))]
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v16
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
