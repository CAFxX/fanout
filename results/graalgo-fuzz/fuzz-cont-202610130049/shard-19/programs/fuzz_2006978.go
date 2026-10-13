package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32; F2 bool }
type GenS2 struct { F0 uint; F1 int }
func genFn0(p0 string, p1 uint16, p2 int) uint8 {
    var v6 int64 = -229
    switch (^(p2)) {
    case 42:
        v7 := GenS1{F0: 775, F1: -2147483648, F2: false}
        _ = v7
    case 179:
        fmt.Println(p0, p2)
    default:
        p2 += p2
    }
    var v8 uint32 = 671
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v8
    return 42
}

func main() {
    v1 := make(map[string]string)
    var v2 string = "\n\t"
    v3 := []uint32{4294967295}
    v4 := []int{1, -1, -904}
    var v5 int16 = -54
    fmt.Println(v5, v2)
    switch -167 {
    case 1:
        var v9 int64 = int64(v5)
        _ = v9
    case 9223372036854775807:
        var v10 string = ((("" + "a") + (v2 + v2)) + v2)
        _ = v10
    case -668:
        var v11 uint32 = 1
        _ = v11
    default:
        var v12 int8 = -1
        _ = v12
    }
    v13 := GenS1{F0: 504, F1: 0, F2: (false && (!false))}
    delete(v1, (v2 + ("foo bar" + ("a" + v2))))
    var v14 uint8 = 255
    for fi15 := 0; fi15 < 12; fi15++ {
        var v16 int64 = 241
        _ = fi15
        _ = v16
    }
    v17 := GenS1{F0: 4294967295, F1: -775, F2: false}
    if (false || (2 != 1)) {
        v2 = v2
    } else {
        v18 := GenS2{F0: 858, F1: 32768}
        v14 += v14
        v1[v2] = (v2 + "\n\t")
        _ = v18
    }
    v19 := GenS1{F0: 659, F1: 2147483647, F2: (127 == 127)}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v17
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
