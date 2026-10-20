package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool }
type GenS2 struct { F0 uint; F1 int32 }
func genFn0(p0 uint) int8 {
    if (127 >= -1) {
        v4 := GenS1{F0: -46, F1: ((32767 >= 903) || true)}
        _ = v4
    } else {
        p0 -= (p0 ^ (p0 + (p0 >> (((p0 | p0)) & 63))))
    }
    fmt.Println(p0)
    _ = p0
    return 127
}
func genFn1(p0 bool) string {
    var v5 int64 = -9223372036854775808
    v6 := []int{522, 766, 2147483647, 0}
    var v7 uint16 = 438
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return "\n\t"
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[uint32]string)
    v3 := make(map[string]uint32)
    if false {
        var v8 uint16 = 32767
        delete(v3, ("foo bar" + "foo bar"))
        _ = v8
    } else {
        v9 := GenS2{F0: 211, F1: 625}
        var v10 uint32 = 4294967295
        v11 := []int{981, 218, 42}
        _ = v10
        _ = v11
        _ = v9
    }
    switch 469 {
    case -1:
        var v12 float64 = 1e10
        _ = v12
    default:
        var v13 string = ""
        _ = v13
    }
    switch 2147483647 {
    case 128:
        switch 572 {
        case 127:
            var v14 int8 = 127
            _ = v14
        case 4294967295:
            for fi15 := 0; fi15 < 6; fi15++ {
                v16 := v2[1]
                _ = fi15
                _ = v16
            }
        }
    case -968:
        var v17 int64 = -1
        _ = v17
    case 4294967295:
        var v18 uint16 = 132
        _ = v18
    default:
        var v19 uint32 = 65535
        _ = v19
    }
    var v20 int = -9223372036854775808
    for fi21 := 0; fi21 < 1; fi21++ {
        if false {
            var v22 uint64 = 0
            if true {
                v22 += v22
            }
            _ = v22
        }
        v23 := GenS2{F0: 255, F1: 0}
        _ = fi21
        _ = v23
    }
    fmt.Println(v20)
    v20 -= v20
    var v24 uint64 = 858
    if true {
        v25 := v2[893]
        _ = v25
    } else {
        switch "hello" {
        case "foo bar":
            var v26 uint32 = 4294967295
            _ = v26
        case "a":
            v24 += (v24 << ((448) & 63))
        default:
            var v27 uint32 = 1
            _ = v27
        }
    }
    v24 += v24
    for fi28 := 0; fi28 < 4; fi28++ {
        v29 := GenS2{F0: 1, F1: -2147483648}
        var v30 uint8 = 27
        var v31 string = ""
        _ = fi28
        _ = v29
        _ = v30
        _ = v31
    }
    v24 = v24
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v20
    _ = v24
    _ = v3
}
