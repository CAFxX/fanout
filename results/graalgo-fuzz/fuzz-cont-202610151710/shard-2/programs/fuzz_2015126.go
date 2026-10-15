package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint; F1 uint8 }
type GenS3 struct { F0 int64; F1 int32 }
func genFn0() int8 {
    var v4 int = 128
    v5 := GenS1{F0: v4}
    _ = v4
    _ = v5
    return 42
}

func main() {
    var v1 float32 = 3.14159
    var v2 uint32 = 2147483648
    var v3 int8 = -1
    fmt.Println(v2, v1)
    if false {
        if (265 > 32767) {
            switch -326 {
            case 32767:
                var v6 uint8 = 221
                _ = v6
            case 9223372036854775807:
                v7 := []int{}
                _ = v7
            }
            v8 := GenS1{F0: -320}
            _ = v8
        }
        fmt.Println(v2, v1)
        fmt.Println(v3, v3)
    } else {
        v3 = genFn0()
        var v9 uint64 = 42
        fmt.Println(v2, v3, v2)
        _ = v9
    }
    var v10 uint64 = uint64(v2)
    v11 := []int{-1, 776, 42, 32768}
    v12 := GenS2{F0: 18446744073709551615, F1: 127}
    v1 = v1
    fmt.Println(v1, v1, v3)
    for fi13 := 0; fi13 < 18; fi13++ {
        v1 = ((0.0 / v1) - float32(-2.25))
        _ = fi13
    }
    var v14 int32 = -2147483648
    v15 := GenS3{F0: 110, F1: v14}
    v16 := GenS3{F0: 42, F1: -1}
    var v17 uint32 = (v2 | (v2 << ((450) & 63)))
    v18 := GenS1{F0: 890}
    v19 := GenS3{F0: 1, F1: (v14 / ((-1) | 1))}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
}
