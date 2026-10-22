package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64; F2 uint32 }
func genFn0(p0 bool, p1 int) uint32 {
    var v6 float32 = (float32(0.0) / 3.14159)
    var v7 int16 = -32768
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 757
}

func main() {
    v1 := make(map[string]bool)
    v2 := []uint32{2}
    var v3 uint8 = 128
    v4 := make(map[int]int)
    v5 := make(map[int]int)
    if len(v2) > 0 {
        v2[0] = 4294967295
    }
    v3 -= v3
    switch 192 {
    case 128:
        v1[("foo bar" + "")] = (false && false)
    case -1:
        var v8 int8 = -13
        _ = v8
    case 510:
        if (false && (false || ("a" <= "hello"))) {
            v9 := GenS1{F0: 121, F1: 42, F2: 1}
            _ = v9
        } else {
            delete(v4, 42)
            if len(v2) > 0 {
                v2[0] = 2147483647
            }
        }
    }
    v10 := GenS1{F0: 2147483647, F1: -774, F2: 4294967295}
    _ = len(v2)
    var v11 bool = true
    if len(v2) > 0 {
        v2[0] = 566
    }
    v3 = v3
    fmt.Println(v11)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
