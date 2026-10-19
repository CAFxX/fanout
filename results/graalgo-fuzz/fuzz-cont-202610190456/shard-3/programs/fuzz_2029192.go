package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint }
func genFn0() bool {
    v6 := GenS1{F0: 929}
    var v7 uint16 = 32767
    fmt.Println(v7)
    v7 += v7
    _ = v6
    _ = v7
    return ((1.5 - -2.25) == float32(1e10))
}

func main() {
    var v1 int64 = -1
    var v2 int16 = 42
    var v3 float64 = ((0.1 * (float64(1.5) * 0.0)) + 1e10)
    var v4 int64 = 258
    v5 := make(map[uint32]int)
    fmt.Println(v1, v3)
    v3 = v3
    if false {
        switch 522 {
        case 0:
            v2 = int16(v4)
        case 127:
            var v8 int32 = -1
            _ = v8
        case -210:
            v3 -= (((0.0 - 3.14159) - v3) - float64(1e10))
        }
    }
    v9 := []int{127, 32768, 518}
    v10 := GenS2{F0: 255}
    var v11 int8 = 53
    var v12 uint32 = 4294967295
    _ = len(v9)
    if true {
        if (false || false) {
            delete(v5, v12)
            var v13 float64 = v3
            _ = v13
        } else {
            v4 -= v1
        }
    }
    var v14 uint32 = (v12 + (v12 >> ((2) & 63)))
    v9 = append(v9, -9223372036854775808)
    var v15 int64 = (v1 ^ -864)
    var v16 uint8 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
