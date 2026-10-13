package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16; F2 uint }
func genFn0() bool {
    if (true && false) {
        for fi5 := 0; fi5 < 3; fi5++ {
            v6 := GenS1{F0: -2147483648, F1: 42, F2: 18446744073709551615}
            if (965 != 4294967295) {
                var v7 int16 = int16(fi5)
                v8 := []int{(fi5 & int(v7)), fi5}
                var v9 uint64 = 4294967295
                _ = v7
                _ = v8
                _ = v9
            }
            _ = fi5
            _ = v6
        }
    } else {
        v10 := []int{-300, 128}
        _ = v10
    }
    var v11 uint8 = 168
    _ = v11
    return false
}

func main() {
    v1 := make(map[int]uint32)
    v2 := []uint32{}
    var v3 int8 = -16
    v4 := make(map[string]string)
    if (2147483647 >= -685) {
        for fi12 := 0; fi12 < 8; fi12++ {
            if true {
                fmt.Println(fi12, v3)
            } else {
                var v13 uint64 = 1
                _ = len(v2)
                var v14 int32 = 0
                _ = v13
                _ = v14
            }
            _ = fi12
        }
        var v15 int = 32767
        var v16 uint8 = 253
        _ = v15
        _ = v16
    }
    if (false || false) {
        v3 += (v3 / ((v3) | 1))
        var v17 float32 = 0.1
        _ = v17
    }
    fmt.Println(v3)
    var v18 int32 = 2147483647
    v18 = (v18 | v18)
    if (-2.25 != 3.14159) {
        v19 := GenS1{F0: 2147483647, F1: 32767, F2: 112}
        v20 := v1[4294967295]
        var v21 int = -9223372036854775808
        _ = v19
        _ = v20
        _ = v21
    }
    v3 = v3
    if true {
        v2 = append(v2, 2)
        if false {
            var v22 int = -1
            _ = len(v2)
            fmt.Println(v18)
            _ = v22
        } else {
            v2 = append(v2, 42)
        }
    }
    var v23 int = 4294967295
    v24 := v4["foo bar"]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v18
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
}
