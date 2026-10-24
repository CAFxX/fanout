package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
func genFn0() int32 {
    if (127 < 1) {
        if (!false) {
            var v6 uint32 = 42
            switch -2147483648 {
            case 0:
                fmt.Println(v6)
            case 4294967295:
                v7 := GenS1{F0: 146}
                _ = v7
            case 889:
                var v8 float64 = (float64(1.5) / 1e10)
                _ = v8
            }
            v9 := []int{274, -1}
            _ = v6
            _ = v9
        }
    }
    var v10 int64 = -980
    if (1e10 != 0.0) {
        if (-2147483648 <= -136) {
            v11 := GenS1{F0: 73}
            v12 := []int{4294967295, 42, -809, 2147483647}
            var v13 uint32 = 2147483647
            _ = v11
            _ = v12
            _ = v13
        }
    } else {
        var v14 int = 2147483647
        _ = v14
    }
    var v15 uint32 = 2
    _ = v10
    _ = v15
    return 1
}
func genFn1(p0 bool, p1 int, p2 uint8) bool {
    p2 = (p2 | p2)
    switch p1 {
    case 42:
        p2 = p2
    default:
        fmt.Println(p1)
    }
    fmt.Println(p0)
    if ((p1 / ((p1) | 1)) < (p1 * 9223372036854775807)) {
        fmt.Println(p1)
    }
    var v16 int8 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    return (!true)
}

func main() {
    v1 := []string{"a", (("a" + ("foo bar" + "foo bar")) + "foo bar"), "hello", "hello"}
    var v2 uint = 0
    var v3 uint = v2
    v4 := make(map[uint32]int)
    v5 := make(map[string]uint32)
    var v17 uint64 = 42
    var v18 uint = (v2 << ((v2) & 63))
    v17 -= (^(v17))
    var v19 float64 = 1e10
    v18 += uint(v17)
    var v20 int8 = 127
    if false {
        v21 := GenS1{F0: 222}
        v19 -= v19
        var v22 uint64 = 999
        _ = v21
        _ = v22
    } else {
        fmt.Println(v17)
        for fi23 := 0; fi23 < 19; fi23++ {
            v24 := GenS1{F0: 2}
            var v25 float64 = v19
            _ = fi23
            _ = v24
            _ = v25
        }
        v26 := v4[2]
        _ = v26
    }
    for fi27 := 0; fi27 < 17; fi27++ {
        fmt.Println(v17, v17)
        _ = fi27
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
