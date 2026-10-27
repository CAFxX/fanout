package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool; F2 int8 }
type GenS2 struct { F0 uint8 }
func genFn0(p0 uint8) uint {
    p0 -= p0
    v4 := GenS1{F0: 9223372036854775807, F1: ((0 < 816) || (!(p0 > 42))), F2: -42}
    _ = p0
    _ = v4
    return 1
}
func genFn1(p0 int) uint32 {
    fmt.Println(p0)
    for fi5 := 0; fi5 < 17; fi5++ {
        if true {
            p0 = fi5
            var v6 uint16 = 711
            _ = v6
        }
        _ = fi5
    }
    var v7 float32 = 1e10
    v8 := []int{(^(p0)), p0, p0}
    _ = p0
    _ = v7
    _ = v8
    return 821
}

func main() {
    v1 := []int{325}
    v2 := make(map[int]bool)
    var v3 int64 = -1
    v3 -= (-(v3))
    switch "hello" {
    case "hello":
        v3 += -9223372036854775808
    default:
        v9 := GenS2{F0: 20}
        _ = v9
    }
    v10 := v2[-938]
    v1 = append(v1, 128)
    for fi11 := 0; fi11 < 3; fi11++ {
        v10 = ((-1 > 32767) && v10)
        var v12 int16 = -32768
        _ = fi11
        _ = v12
    }
    v2[-1] = (!(!v10))
    switch 128 {
    case 4294967295:
        fmt.Println(v3, v10)
    }
    fmt.Println(v3, v10)
    var v13 uint = 65535
    var v14 int32 = 0
    v15 := GenS1{F0: -2147483648, F1: v10, F2: 100}
    var v16 int32 = -1
    if (!true) {
        fmt.Println(v10, v16, v14)
        v1 = append(v1, 42)
    }
    fmt.Println(v10, v3, v13)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
}
