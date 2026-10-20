package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
func genFn0(p0 int8, p1 uint16, p2 uint16) uint16 {
    p1 = p1
    v5 := []int{-1}
    var v6 uint8 = 136
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return (p1 << ((398) & 63))
}

func main() {
    var v1 int64 = -9223372036854775808
    var v2 uint = 673
    v3 := make(map[int]int)
    var v4 uint16 = 65535
    v1 = -558
    if (!(!(false || (false && (4294967295 > 890))))) {
        v7 := GenS1{F0: v4}
        _ = v7
    } else {
        v4 = 65535
        delete(v3, 1)
        if (v1 > v1) {
            _ = len(v3)
            var v8 float32 = 3.14159
            _ = v8
        }
    }
    v9 := []int{-9223372036854775808, 2147483647, 1, -206}
    if len(v9) > 0 {
        v9[0] = 128
    }
    if len(v9) > 0 {
        v9[0] = -2147483648
    }
    if ((("foo bar" + "hello") + "foo bar") > ("" + "hello")) {
        var v10 int8 = 91
        _ = v10
    } else {
        v11 := GenS1{F0: v4}
        if (false && (!true)) {
            v1 += v1
        } else {
            if false {
                var v12 uint16 = 670
                _ = v12
            } else {
                fmt.Println(v4, v4)
                if true {
                    var v13 uint = v2
                    var v14 float32 = 1.5
                    _ = v13
                    _ = v14
                } else {
                    var v15 float32 = (float32(0.0) + -2.25)
                    fmt.Println(v1, v15)
                    _ = v15
                }
            }
            var v16 uint64 = 18446744073709551615
            var v17 float32 = (1e10 * 1.5)
            _ = v16
            _ = v17
        }
        var v18 uint8 = 135
        _ = v11
        _ = v18
    }
    var v19 int16 = 1
    var v20 float64 = 1e10
    var v21 int8 = 100
    var v22 int32 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v9
}
