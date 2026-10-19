package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint16; F1 uint }
type GenS3 struct { F0 int64; F1 int64; F2 int32 }
type GenS4 struct { F0 int8 }
func genFn0() int8 {
    v4 := GenS1{F0: 197}
    var v5 uint32 = 42
    _ = v4
    _ = v5
    return -22
}
func genFn1(p0 int64, p1 uint, p2 uint64) int16 {
    var v6 int8 = 81
    var v7 int32 = 0
    fmt.Println(v6, p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return int16(p2)
}

func main() {
    var v1 uint64 = 0
    v2 := make(map[string]int)
    v3 := make(map[uint32]uint32)
    v1 = (v1 | v1)
    if true {
        if false {
            v8 := GenS2{F0: 42, F1: 1}
            _ = v8
        }
        var v9 float64 = ((0.1 - 0.1) / float64(3.14159))
        _ = v9
    }
    if true {
        switch 2147483647 {
        case 32768:
            var v10 int32 = 929
            _ = v10
        case 4294967295:
            v1 -= 2
        case 0:
            v11 := GenS3{F0: 282, F1: 9223372036854775807, F2: -11}
            _ = v11
        default:
            fmt.Println(v1)
        }
        fmt.Println(v1)
        v12 := GenS3{F0: 745, F1: 1, F2: -2147483648}
        _ = v12
    }
    if ((0.1 - 1.5) >= 3.14159) {
        v1 = v1
        fmt.Println(v1)
    } else {
        var v13 uint64 = (v1 | v1)
        _ = v13
    }
    switch "" {
    case "foo bar":
        fmt.Println(v1)
    case "":
        v14 := GenS4{F0: genFn0()}
        _ = v14
    case "hello":
        var v15 float64 = -2.25
        _ = v15
    default:
        v16 := []int{-916}
        _ = v16
    }
    var v17 uint16 = 2
    for fi18 := 0; fi18 < 11; fi18++ {
        fmt.Println(v1, v1)
        v17 += v17
        _ = fi18
    }
    var v19 string = ""
    v19 = v19
    v1 += (v1 + v1)
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v17
    _ = v19
    _ = v2
    _ = v3
}
