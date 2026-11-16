package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
func genFn0(p0 uint64, p1 string, p2 bool) int16 {
    v7 := GenS1{F0: 1}
    var v8 int16 = 0
    var v9 int64 = -1
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return v8
}
func genFn1() uint64 {
    var v10 int8 = 42
    v11 := []int{2147483647, 339}
    if (42 != 107) {
        v10 = 100
        if (!(!false)) {
            fmt.Println(v10)
            var v12 uint = 1
            _ = v12
        } else {
            for fi13 := 0; fi13 < 9; fi13++ {
                v11 = append(v11, (fi13 + 2147483647))
                v10 += (v10 / ((1) | 1))
                _ = fi13
            }
            fmt.Println(v10)
            v10 -= (v10 | 127)
        }
        var v14 uint16 = 266
        _ = v14
    } else {
        v15 := GenS1{F0: 323}
        switch -560 {
        case 32767:
            v10 = v10
        default:
            var v16 uint64 = 2
            _ = v16
        }
        _ = v15
    }
    _ = v10
    _ = v11
    return 4294967295
}

func main() {
    var v1 float32 = 3.14159
    v2 := make(map[string]int)
    var v3 uint = 65535
    var v4 uint8 = 1
    var v5 int8 = 42
    v6 := make(map[string]uint32)
    var v17 int8 = (v5 | (-(v5)))
    var v18 bool = false
    if v18 {
        if (32767 > 42) {
            if (2147483647 != -1) {
                var v19 float32 = ((0.1 - 0.1) - 3.14159)
                var v20 int64 = -1
                v20 += 1
                _ = v19
                _ = v20
            }
        }
        var v21 int32 = 1
        fmt.Println(v4, v21, v17)
        _ = v21
    }
    var v22 string = "hello"
    v3 -= v3
    if (v18 || v18) {
        if (314 <= 4294967295) {
            v17 += v5
            v23 := []int{-9223372036854775808}
            _ = v23
        }
    }
    for fi24 := 0; fi24 < 11; fi24++ {
        v25 := []int{}
        var v26 int = fi24
        _ = fi24
        _ = v25
        _ = v26
    }
    v27 := GenS1{F0: 121}
    v18 = ((v3 << ((v3) & 63)) > (v3 + v3))
    if v18 {
        var v28 uint8 = 255
        _ = v28
    } else {
        fmt.Println(v4, v3, v4)
        v18 = (0 == -1)
        var v29 int64 = -1
        _ = v29
    }
    v30 := []int{-2147483648, 32768}
    _ = len(v2)
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v22
    _ = v27
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
