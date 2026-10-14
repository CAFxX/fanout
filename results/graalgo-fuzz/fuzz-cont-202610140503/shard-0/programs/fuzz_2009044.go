package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int16 }
func genFn0(p0 int64, p1 int8, p2 uint16) uint32 {
    if (!true) {
        p0 = p0
    } else {
        switch 128 {
        case 0:
            p0 = p0
        case -1:
            var v6 int32 = 503
            _ = v6
        case -352:
            var v7 int = -2147483648
            _ = v7
        }
    }
    var v8 int64 = p0
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 2147483647
}
func genFn1(p0 bool, p1 uint8, p2 bool) uint16 {
    var v9 uint32 = 2147483648
    var v10 float64 = 1.5
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v9
    return uint16(v9)
}

func main() {
    v1 := make(map[int]bool)
    var v2 string = "hello"
    var v3 float32 = -2.25
    v4 := []uint32{65535, 2147483648}
    var v5 int = 42
    v4 = append(v4, 2147483647)
    fmt.Println(v5, v2)
    v11 := GenS1{F0: 32767}
    var v12 float32 = v3
    for fi13 := 0; fi13 < 16; fi13++ {
        v14 := GenS2{F0: 114}
        _ = fi13
        _ = v14
    }
    _ = len(v1)
    for fi15 := 0; fi15 < 2; fi15++ {
        var v16 float32 = 1.5
        _ = fi15
        _ = v16
    }
    if len(v4) > 0 {
        v4[0] = 2147483648
    }
    v17 := v1[v5]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
