package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint; F2 int32 }
type GenS2 struct { F0 bool; F1 uint64 }
type GenS3 struct { F0 uint; F1 int; F2 uint }
type GenS4 struct { F0 uint16; F1 string }
func genFn0(p0 uint32, p1 uint32, p2 int32) uint64 {
    fmt.Println(p1, p2)
    var v5 int8 = 0
    var v6 int = 9223372036854775807
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 1
}

func main() {
    v1 := make(map[uint32]string)
    var v2 string = "hello"
    v3 := make(map[int]int)
    var v4 int32 = -1
    v4 = v4
    v7 := GenS1{F0: 4294967295, F1: 255, F2: (v4 - v4)}
    v4 -= 1
    var v8 float64 = 1.5
    v9 := []int{-29, -9223372036854775808, 0}
    var v10 uint32 = 1
    if (v8 == v8) {
        v10 = (v10 / (((v10 | (v10 + 0))) | 1))
        v11 := v3[32768]
        fmt.Println(v8)
        _ = v11
    }
    v8 = v8
    for fi12 := 0; fi12 < 1; fi12++ {
        v13 := GenS2{F0: (true || (-2.25 >= 1.5)), F1: 717}
        if true {
            var v14 uint8 = 150
            var v15 uint32 = v10
            if false {
                v15 = (v10 % (((v10 + (v15 & (v10 & (v10 ^ (v10 + v10)))))) | 1))
                var v16 uint32 = uint32(v14)
                fmt.Println(v4)
                _ = v16
            } else {
                delete(v3, (fi12 ^ 4294967295))
            }
            _ = v14
            _ = v15
        } else {
            fmt.Println(fi12, v10)
            v17 := v1[(v10 >> ((493) & 63))]
            _ = len(v3)
            _ = v17
        }
        v18 := GenS3{F0: 18446744073709551615, F1: fi12, F2: 4294967295}
        _ = fi12
        _ = v13
        _ = v18
    }
    v19 := GenS4{F0: 1, F1: "hello"}
    fmt.Println(v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
