package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32; F2 int64 }
type GenS2 struct { F0 uint64; F1 int8; F2 bool }
type GenS3 struct { F0 bool; F1 uint8 }
func genFn0(p0 uint32, p1 uint32) int {
    v6 := GenS1{F0: -9223372036854775808, F1: -1, F2: -1}
    p1 = p0
    for fi7 := 0; fi7 < 15; fi7++ {
        if (365 <= 26) {
            var v8 bool = (127 <= 42)
            v9 := GenS1{F0: -9223372036854775808, F1: 1, F2: 42}
            fmt.Println(fi7, p0, fi7)
            _ = v8
            _ = v9
        }
        v10 := GenS1{F0: int64(fi7), F1: -288, F2: 42}
        _ = fi7
        _ = v10
    }
    _ = p0
    _ = p1
    _ = v6
    return 9223372036854775807
}
func genFn1() uint32 {
    for fi11 := 0; fi11 < 9; fi11++ {
        v12 := GenS2{F0: 4294967295, F1: 33, F2: (485 >= 32767)}
        _ = fi11
        _ = v12
    }
    for fi13 := 0; fi13 < 7; fi13++ {
        if (133 <= 526) {
            var v14 uint8 = 192
            v15 := []int{(-(fi13)), genFn0(2147483648, 2147483647), (^(fi13)), 888}
            _ = v14
            _ = v15
        } else {
            var v16 int = 127
            _ = v16
        }
        v17 := GenS2{F0: 268, F1: 100, F2: true}
        _ = fi13
        _ = v17
    }
    var v18 int64 = 9223372036854775807
    v19 := GenS3{F0: false, F1: 0}
    v20 := []int{-9223372036854775808, -2147483648, 1}
    _ = v18
    _ = v19
    _ = v20
    return 65535
}

func main() {
    var v1 uint32 = 2
    var v2 string = ""
    v3 := make(map[uint32]int)
    var v4 int16 = -32768
    var v5 float64 = 3.14159
    fmt.Println(v4)
    var v21 uint32 = 96
    v21 -= 4294967295
    var v22 bool = false
    v22 = (-1 <= 94)
    var v23 int64 = -1
    v22 = false
    var v24 int64 = v23
    v25 := []int{597, 265, -354}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
