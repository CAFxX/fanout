package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8; F2 int8 }
func genFn0(p0 int32, p1 int16, p2 int8) int16 {
    v7 := []int{32768, 530, 259}
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 1
}

func main() {
    v1 := []uint32{442, 374, 2147483647}
    var v2 bool = (((0 <= 255) || false) && false)
    var v3 uint64 = 4294967295
    var v4 uint16 = 401
    var v5 int64 = 986
    v6 := make(map[string]uint32)
    delete(v6, ((("" + "foo bar") + "foo bar") + "hello"))
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println(v2)
    var v8 int16 = 32767
    v5 -= v5
    var v9 float64 = float64(1.5)
    var v10 uint32 = 4294967295
    if true {
        if v2 {
            if false {
                _ = len(v6)
                fmt.Println(v9, v3, v10)
            } else {
                var v11 float64 = v9
                _ = v11
            }
            v4 -= 32767
            if ((v8 << ((1) & 63)) < genFn0(int32(v8), (v8 % ((v8) | 1)), -77)) {
                fmt.Println(v9, v5, v2)
            } else {
                fmt.Println(v9, v2, v2)
                v12 := GenS1{F0: 4294967295, F1: 42, F2: -63}
                _ = v12
            }
        }
    }
    for fi13 := 0; fi13 < 7; fi13++ {
        v14 := GenS1{F0: fi13, F1: 0, F2: 100}
        v10 = v10
        _ = fi13
        _ = v14
    }
    var v15 float32 = (0.0 + 1.5)
    for fi16 := 0; fi16 < 5; fi16++ {
        for fi17 := 0; fi17 < 17; fi17++ {
            v3 = 2
            _ = fi17
        }
        v18 := GenS1{F0: fi16, F1: 255, F2: 100}
        _ = fi16
        _ = v18
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
