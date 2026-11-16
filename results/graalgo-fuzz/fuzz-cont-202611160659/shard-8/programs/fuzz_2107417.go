package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int }
func genFn0(p0 bool, p1 int64, p2 int16) uint32 {
    for fi6 := 0; fi6 < 1; fi6++ {
        var v7 int32 = -981
        var v8 int8 = 100
        v9 := []int{(fi6 + (fi6 << ((0) & 63))), (fi6 & -2147483648), fi6, (fi6 >> ((0) & 63))}
        _ = fi6
        _ = v7
        _ = v8
        _ = v9
    }
    var v10 uint32 = 600
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return 1
}

func main() {
    var v1 uint64 = 999
    var v2 bool = true
    var v3 uint64 = (^(v1))
    v4 := make(map[int]uint32)
    v5 := make(map[int]int)
    var v11 int32 = 2147483647
    for fi12 := 0; fi12 < 9; fi12++ {
        var v13 uint16 = 255
        _ = fi12
        _ = v13
    }
    if ("hello" < "foo bar") {
        v14 := []int{128, 50, 32768}
        fmt.Println(v3, v3)
        fmt.Println(v2, v2, v11)
        _ = v14
    } else {
        var v15 int64 = 0
        _ = v15
    }
    _ = len(v5)
    fmt.Println(v2)
    var v16 int = 4294967295
    for fi17 := 0; fi17 < 17; fi17++ {
        for fi18 := 0; fi18 < 11; fi18++ {
            v19 := []int{(v16 - -669), 1, fi17, 4294967295}
            var v20 uint64 = 42
            _ = fi18
            _ = v19
            _ = v20
        }
        v21 := GenS1{F0: 896, F1: (fi17 | 0)}
        if v2 {
            fmt.Println(v16, v1)
        }
        _ = fi17
        _ = v21
    }
    if true {
        var v22 uint8 = 128
        _ = v22
    } else {
        v23 := []int{}
        _ = v23
    }
    switch v16 {
    case 877:
        fmt.Println(v11, v3, v3)
    default:
        v24 := GenS1{F0: 309, F1: (v16 - (v16 * (v16 >> ((516) & 63))))}
        _ = v24
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
