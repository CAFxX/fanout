package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64 }
type GenS2 struct { F0 uint; F1 int32; F2 uint8 }
type GenS3 struct { F0 int16; F1 int8 }
func genFn0(p0 int32, p1 int32) uint32 {
    p1 -= (p0 / (((p0 - (^(p0)))) | 1))
    v7 := GenS1{F0: -9223372036854775808, F1: 0}
    v8 := []int{128}
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 2147483647
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int32 = 1
    var v3 int64 = 0
    v4 := make(map[uint32]string)
    var v5 uint16 = 2
    var v6 uint16 = (v5 + 65535)
    if ("" >= "\n\t") {
        v3 = 1
        v9 := []int{42, 128, 1, 1}
        _ = v9
    }
    v10 := GenS1{F0: v3, F1: (v3 % ((42) | 1))}
    if true {
        v11 := GenS1{F0: v3, F1: v3}
        var v12 uint16 = 32767
        switch "" {
        case "\n\t":
            v13 := []int{128, 528, 2147483647}
            _ = v13
        case "foo bar":
            fmt.Println(v5, v2, v5)
        case "hello":
            var v14 uint8 = 1
            _ = v14
        }
        _ = v11
        _ = v12
    } else {
        v3 -= (-(v3))
        v15 := GenS2{F0: 1, F1: -907, F2: 1}
        _ = v15
    }
    v2 = v2
    v16 := v4[0]
    delete(v1, 0)
    for fi17 := 0; fi17 < 13; fi17++ {
        var v18 int16 = -32768
        _ = fi17
        _ = v18
    }
    switch 851 {
    case -203:
        v19 := []int{4294967295}
        _ = v19
    }
    for fi20 := 0; fi20 < 9; fi20++ {
        v21 := GenS3{F0: -32768, F1: 0}
        v22 := GenS3{F0: int16(fi20), F1: 92}
        _ = fi20
        _ = v21
        _ = v22
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
