package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 bool }
func genFn0(p0 int16) int {
    for fi7 := 0; fi7 < 3; fi7++ {
        var v8 uint64 = 1
        p0 = p0
        _ = fi7
        _ = v8
    }
    var v9 int16 = 993
    p0 = p0
    var v10 int16 = v9
    p0 -= v10
    _ = p0
    _ = v10
    _ = v9
    return 127
}
func genFn1() int16 {
    v11 := GenS1{F0: 366}
    if true {
        switch -487 {
        case -490:
            var v12 uint8 = 2
            _ = v12
        case -134:
            var v13 uint = 65535
            _ = v13
        default:
            var v14 uint64 = 4294967295
            _ = v14
        }
        v15 := GenS1{F0: 65535}
        var v16 uint16 = 42
        _ = v15
        _ = v16
    }
    var v17 uint8 = 188
    v17 = v17
    var v18 int64 = -1
    _ = v11
    _ = v17
    _ = v18
    return 1000
}

func main() {
    var v1 uint16 = 255
    v2 := make(map[uint32]string)
    v3 := []uint32{284, 2147483648, 42}
    var v4 uint64 = 4294967295
    v5 := make(map[uint32]string)
    var v6 uint64 = 42
    var v19 bool = false
    v20 := GenS2{F0: false}
    v19 = (!false)
    var v21 float64 = 3.14159
    v22 := GenS2{F0: (!(!(!(!true))))}
    for fi23 := 0; fi23 < 9; fi23++ {
        _ = len(v3)
        if len(v3) > 0 {
            v3[0] = 2
        }
        _ = fi23
    }
    fmt.Println(v21, v4, v21)
    var v24 float32 = (((1.5 * 1e10) + 3.14159) - 0.1)
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
