package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint32; F1 int8; F2 uint64 }
type GenS4 struct { F0 uint; F1 int64; F2 int16 }
func genFn0() uint {
    for fi6 := 0; fi6 < 8; fi6++ {
        var v7 uint64 = 0
        v8 := GenS1{F0: -1}
        _ = fi6
        _ = v7
        _ = v8
    }
    v9 := GenS2{F0: 1}
    _ = v9
    return 805
}
func genFn1(p0 uint, p1 int32, p2 uint) int {
    fmt.Println(p1)
    p1 -= p1
    if ((!(((0.1 + ((3.14159 * (((float32(-2.25) + -2.25) + 0.0) + 0.1)) / 1.0)) >= float32(3.14159)) && false)) || (-2.25 >= (3.14159 - float32((3.14159 - 0.0))))) {
        if false {
            fmt.Println(p2, p2, p1)
        }
        var v10 float32 = (0.0 + 0.1)
        v11 := GenS3{F0: 42, F1: 106, F2: 554}
        _ = v10
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    return -1
}

func main() {
    var v1 int8 = 100
    var v2 int = 128
    var v3 uint64 = 515
    v4 := make(map[string]bool)
    var v5 uint64 = v3
    var v12 uint8 = 197
    v13 := []int{}
    v2 -= v2
    v5 -= (v5 - (v5 << ((641) & 63)))
    var v14 uint64 = (^(v3))
    var v15 int = 32767
    if len(v13) >= 2 {
        v13 = v13[1:len(v13)]
    }
    v16 := GenS4{F0: 18446744073709551615, F1: 499, F2: 91}
    delete(v4, "\n\t")
    v2 = -9223372036854775808
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
