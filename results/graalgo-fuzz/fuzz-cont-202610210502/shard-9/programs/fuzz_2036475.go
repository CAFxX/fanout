package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16; F2 uint64 }
type GenS2 struct { F0 uint64; F1 uint }
type GenS3 struct { F0 uint32; F1 bool; F2 bool }
func genFn0() string {
    v7 := GenS1{F0: true, F1: 65535, F2: 0}
    v8 := GenS2{F0: 4294967295, F1: 563}
    var v9 uint8 = 128
    v10 := GenS1{F0: true, F1: 2, F2: 18446744073709551615}
    if (42 >= 42) {
        v11 := GenS1{F0: (1.5 > -2.25), F1: 2, F2: 294}
        for fi12 := 0; fi12 < 8; fi12++ {
            v13 := []int{9223372036854775807, fi12, fi12, fi12}
            if len(v13) > 0 {
                v13[0] = -1
            }
            _ = fi12
            _ = v13
        }
        v14 := []int{42, 42}
        _ = v11
        _ = v14
    }
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return ("" + "a")
}

func main() {
    v1 := []string{(("\n\t" + "hello") + "a")}
    v2 := []uint32{}
    var v3 float32 = float32(float64(0.1))
    var v4 int64 = -535
    v5 := make(map[int]bool)
    var v6 int32 = 2147483647
    v15 := GenS3{F0: 543, F1: ((float32(1e10) != v3) && false), F2: true}
    for fi16 := 0; fi16 < 17; fi16++ {
        var v17 uint = 255
        _ = fi16
        _ = v17
    }
    var v18 int8 = 42
    delete(v5, -1)
    var v19 int8 = 89
    v20 := GenS2{F0: 336, F1: 255}
    v21 := GenS3{F0: 95, F1: (18446744073709551615 != 373), F2: (1 != 4294967295)}
    for fi22 := 0; fi22 < 10; fi22++ {
        fmt.Println(fi22)
        v23 := GenS3{F0: 2147483647, F1: (!false), F2: (true && false)}
        v3 -= (-2.25 + ((v3 + 0.1) + (3.14159 + v3)))
        _ = fi22
        _ = v23
    }
    var v24 int64 = (v4 - 0)
    v6 += (v6 + v6)
    _ = len(v5)
    _ = len(v5)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
