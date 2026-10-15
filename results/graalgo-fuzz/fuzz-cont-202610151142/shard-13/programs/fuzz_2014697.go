package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32; F2 uint }
type GenS2 struct { F0 int8; F1 int64 }
type GenS3 struct { F0 string; F1 uint16; F2 uint32 }
type GenS4 struct { F0 uint; F1 int }
func genFn0() uint8 {
    var v4 float64 = 3.14159
    v5 := GenS1{F0: 1, F1: -2147483648, F2: 0}
    _ = v4
    _ = v5
    return 116
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[string]uint32)
    var v3 int64 = -935
    var v6 float32 = 1.5
    v7 := GenS2{F0: -19, F1: 9223372036854775807}
    if (!(false || (0 < 4294967295))) {
        for fi8 := 0; fi8 < 4; fi8++ {
            v3 = (-(v3))
            _ = fi8
        }
    }
    var v9 int16 = -32768
    fmt.Println(v6)
    if false {
        var v10 int32 = 0
        if true {
            v11 := GenS3{F0: "a", F1: 0, F2: 2147483648}
            _ = v11
        } else {
            v12 := GenS4{F0: 767, F1: -9223372036854775808}
            _ = len(v1)
            _ = v12
        }
        _ = v10
    } else {
        v13 := []int{584, 32768}
        if true {
            var v14 bool = (!true)
            _ = v14
        }
        _ = v13
    }
    v3 = v3
    var v15 float32 = (((v6 - (0.1 - v6)) * (0.1 / v6)) * 0.1)
    fmt.Println(v15)
    v16 := GenS3{F0: "hello", F1: 2, F2: 4294967295}
    fmt.Println(v9, v6, v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v9
}
