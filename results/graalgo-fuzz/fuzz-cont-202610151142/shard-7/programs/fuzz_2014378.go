package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32; F2 uint }
type GenS2 struct { F0 bool; F1 bool }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint; F1 uint32 }
type GenS5 struct { F0 int32 }
func genFn0() uint16 {
    if (!false) {
        var v6 uint16 = 1
        v7 := GenS1{F0: 42, F1: 2147483647, F2: 0}
        _ = v6
        _ = v7
    }
    var v8 int64 = 1
    v9 := GenS2{F0: false, F1: (!true)}
    _ = v8
    _ = v9
    return 2
}

func main() {
    var v1 uint64 = 4294967295
    v2 := make(map[string]bool)
    var v3 float32 = 1.5
    v4 := []int{640, 32767, 2147483647, 32768}
    var v5 float32 = 0.1
    v1 = v1
    if false {
        fmt.Println(v3, v3)
        fmt.Println(v1, v1, v3)
        v10 := GenS2{F0: false, F1: ((!false) && (699 <= -1))}
        _ = v10
    }
    var v11 int64 = 9223372036854775807
    var v12 uint64 = (^(v1))
    v13 := GenS3{F0: (true || (42 <= 75))}
    var v14 uint32 = 2
    v15 := GenS4{F0: 105, F1: (v14 - (v14 + (v14 * (v14 * v14))))}
    v16 := GenS2{F0: true, F1: (!true)}
    if true {
        if (true || false) {
            v4 = append(v4, 719)
            v11 = v11
            v17 := GenS2{F0: (128 == 128), F1: (false && (false || (42 <= 531)))}
            _ = v17
        } else {
            fmt.Println(v5, v12, v5)
        }
        fmt.Println(v11)
    } else {
        v4 = append(v4, 1)
    }
    if (2147483647 >= (v14 | (v14 * 268))) {
        v12 += v12
        var v18 int64 = (v11 ^ (v11 % (((v11 | (^(v11)))) | 1)))
        _ = v18
    }
    if true {
        v5 = -2.25
        for fi19 := 0; fi19 < 8; fi19++ {
            if len(v4) > 0 {
                v4[0] = fi19
            }
            v11 += v11
            v4 = append(v4, fi19)
            _ = fi19
        }
    }
    v20 := GenS4{F0: 4294967295, F1: (v14 + (v14 >> ((706) & 63)))}
    v3 -= v3
    v21 := GenS5{F0: -333}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
