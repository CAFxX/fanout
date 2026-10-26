package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16; F2 int16 }
func genFn0() int16 {
    var v4 uint8 = 1
    var v5 int32 = 42
    _ = v4
    _ = v5
    return 32767
}
func genFn1(p0 uint, p1 uint64) int {
    var v6 uint = p0
    p0 = p0
    _ = p0
    _ = p1
    _ = v6
    return -2147483648
}

func main() {
    v1 := []uint32{42, 4294967295, 65535}
    var v2 bool = false
    var v3 float32 = 0.0
    v3 -= (v3 / (-2.25 / 3.14159))
    v7 := GenS1{F0: 42, F1: 368, F2: 1}
    v3 -= 1e10
    v8 := GenS1{F0: 16, F1: 73, F2: 152}
    v2 = false
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v9 int = 6
    for fi10 := 0; fi10 < 12; fi10++ {
        v11 := GenS1{F0: 580, F1: 1, F2: 1}
        if (!(!v2)) {
            if len(v1) > 0 {
                v1[0] = 2147483647
            }
            v3 = 1.5
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        } else {
            fmt.Println(v2, fi10, fi10)
        }
        _ = fi10
        _ = v11
    }
    var v12 uint64 = 2
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v7
    _ = v8
    _ = v9
}
