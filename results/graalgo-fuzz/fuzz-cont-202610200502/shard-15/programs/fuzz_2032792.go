package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64; F2 int }
type GenS2 struct { F0 uint16 }
func genFn0(p0 bool, p1 int8) uint8 {
    if ((872 < 4294967295) || p0) {
        p1 += -120
        v7 := []int{4294967295, 218, -255}
        _ = v7
    } else {
        p0 = p0
        v8 := []int{32767, 9223372036854775807, 32767}
        if (!(!(!true))) {
            p0 = p0
            v8 = append(v8, 32767)
            var v9 uint32 = 4294967295
            _ = v9
        } else {
            p0 = (42 < 128)
        }
        _ = v8
    }
    p0 = ((p1 << ((18446744073709551615) & 63)) == -120)
    v10 := GenS1{F0: 0, F1: -206, F2: -1}
    var v11 int32 = 2147483647
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    return 1
}

func main() {
    v1 := make(map[string]uint32)
    var v2 int16 = 1
    var v3 uint64 = 244
    var v4 uint32 = 4294967295
    v5 := make(map[string]int)
    v6 := make(map[int]bool)
    var v12 int64 = -914
    v6[9223372036854775807] = true
    var v13 uint32 = v4
    var v14 uint64 = (^(v3))
    _ = len(v6)
    var v15 uint8 = uint8(v3)
    var v16 int16 = v2
    var v17 uint16 = 32767
    var v18 int = 1
    _ = len(v1)
    v19 := GenS2{F0: v17}
    if true {
        fmt.Println(v2, v15)
    }
    fmt.Println(v4, v15)
    fmt.Println(v18, v17, v13)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
